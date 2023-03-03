// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package mimic

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Generator manages a pool of generated files.
type Generator struct {
	FilePool

	out       string
	generated bool
}

// New returns a new Generator that parses os.Args as command line arguments.
// It allows passing closure BEFORE parsing the flags to allow defining additional flags.
//
// NOTE: Read README.md before using. This is intentionally NOT following Go library patterns like:
// * It uses panics as the main error handling way.
// * It creates CLI command inside constructor.
// * It does not allow custom loggers etc
func New(injs ...func(cmd *kingpin.CmdClause)) *Generator {
	app := kingpin.New("mimic", "mimic: https://github.com/bwplotka/mimic")
	app.HelpFlag.Short('h')

	gen := app.Command("generate", "generates output files from all registered files via Add method.")
	out := gen.Flag("output", "output directory for generated files.").Short('o').Default("gen").String()

	for _, inj := range injs {
		inj(gen)
	}

	logLevel := app.Flag("log.level", "Log filtering level.").
		Default("info").Enum("error", "warn", "info", "debug")

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("error parsing commandline arguments: %v", err))
		app.Usage(os.Args[1:])
		os.Exit(2)
	}

	var logger log.Logger
	{
		var lvl level.Option
		switch *logLevel {
		case "error":
			lvl = level.AllowError()
		case "warn":
			lvl = level.AllowWarn()
		case "info":
			lvl = level.AllowInfo()
		case "debug":
			lvl = level.AllowDebug()
		default:
			panic("unexpected log level")
		}
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		logger = level.NewFilter(logger, lvl)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	a := &Generator{out: *out}
	switch cmd {
	case gen.FullCommand():
		a.FilePool = FilePool{Logger: logger, m: map[string]string{}}
		return a
	}

	_ = level.Error(logger).Log("err", "command not found", "command", cmd)
	os.Exit(2)

	return nil
}

// With behaves like linux `cd` command. It allows to "walk" & organize output files in a desired way for ease of use.
// Example:
//
// ```
//
//	gen := gen.With("mycompany.com", "production", "eu1", "kubernetes", "thanos")
//
// ```
// Giving the path `mycompany.com/production/eu1/kubernetes/thanos`.
//
// With return a Generator pointing at the specified path which can be specified even further:
// Example:
// ```
//
//	gen := mimic.New()
//	// gen/
//	...
//	gen = gen.With('foo')
//	// gen/foo
//	...
//	{
//	  gen := gen.With('bar')
//	  // gen/foo/bar
//	}
//	// gen/foo
//
// ```
func (g *Generator) With(parts ...string) *Generator {
	// TODO(bwplotka): Support "..", to get back?

	return &Generator{
		out: g.out,
		FilePool: FilePool{
			Logger:           g.Logger,
			path:             append(g.path, parts...),
			m:                g.m,
			topLevelComments: g.topLevelComments,
		},
	}
}

// WithTopLevelComment enables mimic to add any string as a header comment for a genenrated
// file. Follows same usage semantics as With().
//
// Example:
//
//	gen := mimic.New()
//	defer gen.Generate()
//
//	gen.With("config").WithTopLevelComment(mimic.GeneratedComment).Add(name+".yaml", encoding.GhodssYAML(config))
//
// Like With, you can also chain multiple WithTopLevelComment(), which will be added to the top of a file in that order.
//
// Example:
//
//	gen := mimic.New()
//
//	defer gen.Generate()
//	gen = gen.WithTopLevelComment("Foo.")
//
//	gen.With("config").WithTopLevelComment("Bar.").Add(name+".yaml", encoding.GhodssYAML(config))
//
// This will result in a YAML file like,
//
//	# Foo.
//	# Bar.
//	config:
//	- prometheus:
//
// NOTE: This option will be noop for encodings that does not support comment (e.g. encoding.JSON).
func (g *Generator) WithTopLevelComment(content ...string) *Generator {
	return &Generator{
		out: g.out,
		FilePool: FilePool{
			Logger:           g.Logger,
			path:             g.path,
			m:                g.m,
			topLevelComments: append(g.topLevelComments, content...),
		},
	}
}

// Generate generates the configuration files that have been defined and added to a generator.
func (g *Generator) Generate() {
	if g.generated {
		PanicErr(errors.New("generate method already invoked once"))
	}
	defer func() { g.generated = true }()

	_ = level.Info(g.Logger).Log("msg", "generated output", "dir", g.out)
	g.write(g.out)
}
