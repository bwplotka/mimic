package main

// Mini typed cluster & environment setup for inspiration.

type Environment struct {
	Name string
}

func (e *Environment) String() string { return e.Name }

type Cluster struct {
	Name        string
	Desc        string
	Environment *Environment
}

func (c *Cluster) String() string { return c.Name }

var (
	Environments = []*Environment{
		Production, Testing,
	}

	Production = &Environment{
		Name: "production",
	}
	Testing = &Environment{
		Name: "testing",
	}

	Clusters = []*Cluster{
		{
			Name:        "prod-par1-mon0",
			Environment: Production,
			Desc:        "bwplotka.dev monitoring docker based cluster",
		},
	}

	ClustersByName = map[string]*Cluster{}
	ClustersByEnv  = map[*Environment][]*Cluster{}
)

func init() {
	// Fill the clusters into useful helpers like ClustersByName and ClustersByEnv in compile time etc.
	for _, cl := range Clusters {
		ClustersByName[cl.Name] = cl
	}

	for _, cl := range Clusters {
		ClustersByEnv[cl.Environment] = append(ClustersByEnv[cl.Environment], cl)
	}
}
