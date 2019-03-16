package gocodeit

import (
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
	"honnef.co/go/tools/version"
)

func App() {
	app := kingpin.New(filepath.Base(os.Args[0]), "A block storage based long-term storage for Prometheus")
	app.Version(version.Print("thanos"))
	app.HelpFlag.Short('h')
}
