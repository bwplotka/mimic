package main

import (
	"github.com/bwplotka/gocodeit"
	enc "github.com/bwplotka/gocodeit/encoding"
	"github.com/bwplotka/gocodeit/providers/dockercompose"
)

func main() {
	gci := gocodeit.New()
	defer gci.Generate()

	genMyMonAll(gci)
}

func genMyMonAll(gci *gocodeit.Gen) {
	for _, env := range Environments {
		gci := gci.With(env.Name)
		for _, cl := range ClustersByEnv[env] {
			gci := gci.With(cl.Name)

			genMyMon(gci)
		}
	}

}

func genMyMon(gci *gocodeit.Gen) {
	composeCfg := dockercompose.Config{
		Name: "LOL",
	}

	gci.Output("mon-compose.yaml", enc.YAML(composeCfg))
}
