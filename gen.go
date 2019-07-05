package gocodeit

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

// Generator is a module that hep
type Generator struct {
	FilePool

	out       string
	generated bool
}

// New returns new Generator that parses os.Args as command line arguments.
// It allows passing closure BEFORE parsing the flags to allow defining additional flags.
//
// NOTE: Read README.md before using. This is intentionally NOT following Go library patterns like:
// * It uses panics as the main error handling way.
// * It creates CLI command inside constructor.
// * It does not allow custom loggers etc
func New(injs ...func(cmd *kingpin.CmdClause)) *Generator {
	app := kingpin.New("gocodeit", "GoCodeIt: https://github.com/bwplotka/gocodeit")
	app.HelpFlag.Short('h')

	gen := app.Command("generate", "generates output files from all registered files via Add method.")
	out := gen.Flag("output", "output directory for generated files.").Short('o').Default("gcigen").String()

	for _, inj := range injs {
		inj(gen)
	}

	logLevel := app.Flag("log.level", "Log filtering level.").
		Default("info").Enum("error", "warn", "info", "debug")

	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing commandline arguments"))
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
//  gen := With("mycompany.com", "production", "eu1", "kubernetes" "thanos")
// ```
func (g *Generator) With(parts ...string) *Generator {
	// TODO(bwplotka): Support "..", to get back?

	return &Generator{
		out: g.out,
		FilePool: FilePool{
			Logger: g.Logger,
			path:   append(g.path, parts...),
			m:      g.m,
		},
	}
}

// Generate generates all the files that were registered before.
func (g *Generator) Generate() {
	if g.generated {
		PanicErr(errors.New("generate method already invoked once."))
	}
	defer func() { g.generated = true }()

	_ = level.Info(g.Logger).Log("msg", "generated output", "dir", g.out)
	g.write(g.out)
}

// UnmarshalSecretFile allows to easily manage your secrets passed to Go defined configuration via custom file.
func UnmarshalSecretFile(file string, in interface{}) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		Panicf("read file from: %v", err)
	}

	if err := yaml.Unmarshal(b, in); err != nil {
		Panicf("failed to unmarshal file from: %v", err)
	}
}
