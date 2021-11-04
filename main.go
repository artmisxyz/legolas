package main

import "github.com/artmisxyz/blockinspector/syncer"

func main() {
	conf := syncer.Config{}
	conf.General.StartBlockNumber = 13550914
	conf.General.BlockLag = 10
	conf.General.PosFileLocation = "."
	conf.General.PosFileName = "block_inspector.pos"
	conf.Logger.Level = "debug"
	conf.Node.Websocket = "wss://mainnet.infura.io/ws/v3/83af4d404170428f866ad492288eafac"
	conf.Node.RPC = "https://mainnet.infura.io/v3/83af4d404170428f866ad492288eafac"

	s := &syncer.Syncer{}
	s.Init(conf)

	for {
		s.Sync()
	}
}