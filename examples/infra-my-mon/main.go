package main

import (
	"github.com/bwplotka/gocodeit"
	"github.com/bwplotka/gocodeit/encoding"
	"github.com/bwplotka/gocodeit/providers/dockercompose"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var secretFile *string
	gci := gocodeit.New(
		func(cmd *kingpin.CmdClause) {
			secretFile = cmd.Flag("secret-file", "YAML file with secrets").Required().String()
		},
	)
	defer gci.Generate()

	var secrets Secrets
	gocodeit.UnmarshalSecretFile(*secretFile, &secrets)
	genMyMonAll(gci, secrets)
}

func genMyMonAll(gci *gocodeit.Gen, secrets Secrets) {
	for _, env := range Environments {
		gci := gci.With(env.Name)
		for _, cl := range ClustersByEnv[env] {
			gci := gci.With(cl.Name)

			genMyMon(gci, secrets)
		}
	}

}

func genMyMon(gci *gocodeit.Gen, secrets Secrets) {
	dpl := dockercompose.Config{
		//Services: []*dockercompose.Service{
		//	{
		//		Image:
		//	},
		//},
	}

	gci.Add("mon-compose.yaml", encoding.YAML(dpl))
}
