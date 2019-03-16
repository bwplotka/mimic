package gocodeit

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path/filepath"
)

func App() {
	app := kingpin.New(filepath.Base(os.Args[0]), "A block storage based long-term storage for Prometheus")
	app.Version(version.Print("thanos"))
	app.HelpFlag.Short('h')
}
