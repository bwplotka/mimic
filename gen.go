package gocodeit

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Gen struct {
	Files

	out string

	Logger log.Logger
}

func New(injs ...func(cmd *kingpin.CmdClause)) *Gen {
	app := kingpin.New("gocodeit", "GoCodeIt")
	app.HelpFlag.Short('h')

	gen := app.Command("generate", "generate your everything!")
	out := gen.Flag("output", "output directory").Short('o').Default("gcigen").String()

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

	a := &Gen{
		out: *out,

		Logger: logger,
	}
	switch cmd {
	case gen.FullCommand():
		a.Files = Files{m: map[string]string{}}
		return a
	}

	level.Error(logger).Log("err", "command not found", "command", cmd)
	os.Exit(2)

	return nil
}

func (g *Gen) With(parts ...string) *Gen {
	return &Gen{
		Logger: g.Logger,
		out:    g.out,
		Files: Files{
			path: append(g.path, parts...),
			m:    g.m,
		},
	}
}

func (g *Gen) Generate() {
	g.write(g.out)
}
