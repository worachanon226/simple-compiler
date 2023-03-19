package config

type Config struct {
	App Fiber
}

type Fiber struct {
	Host string
	Port string
}
