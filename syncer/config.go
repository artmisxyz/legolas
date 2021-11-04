package syncer

type Config struct {
	General struct {
		PosFileLocation  string `yaml:"PosFileLocation"`
		PosFileName      string `yaml:"PosFileName"`
		BlockLag         int    `yaml:"BlockLag"`
		StartBlockNumber uint64 `yaml:"StartBlockNumber"`
	}
	Node struct {
		RPC       string `yaml:"RPC"`
		Websocket string `yaml:"Websocket"`
	} `yaml:"Node"`
	Logger struct {
		Level string `yaml:"Level"`
	}
}
