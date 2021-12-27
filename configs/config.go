package configs

type Config struct {
	General struct {
		BlockLag         int    `yaml:"BlockLag"`
	} `yaml:"General"`
	Node struct {
		RPC       string `yaml:"RPC"`
		Websocket string `yaml:"Websocket"`
	} `yaml:"Node"`
	Logger struct {
		Level string `yaml:"Level"`
	}
	Contracts struct {
	} `yaml:"Contracts"`
	Database struct {
		Driver       string `yaml:"Driver"`
		Host         string `yaml:"Host"`
		Port         int    `yaml:"Port"`
		User         string `yaml:"User"`
		Password     string `yaml:"Password"`
		DatabaseName string `yaml:"DatabaseName"`
		SSLMode      string `yaml:"SSLMode"`
	}
	Server struct {
		Port int `yaml:"Port"`
	}
	Redis struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		Password string `yaml:"Password"`
	}
}
