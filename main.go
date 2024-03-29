package main

import (
	"github.com/artmisxyz/legolas/api"
	"github.com/artmisxyz/legolas/syncer"
)

func main() {
	conf := syncer.Config{}
	conf.General.StartBlockNumber = 13639719
	conf.General.BlockLag = 10
	conf.General.PosFileLocation = "."
	conf.General.PosFileName = "block_inspector.pos"
	conf.Logger.Level = "error"
	conf.Node.Websocket = "wss://mainnet.infura.io/ws/v3/83af4d404170428f866ad492288eafac"
	conf.Node.RPC = "https://mainnet.infura.io/v3/83af4d404170428f866ad492288eafac"

	s := &syncer.Syncer{}
	s.Init(conf)
	go func() {
		for {
			s.Sync()
		}
	}()
	server := &api.Server{}
	server.Init(&api.Config{Port: 8080})
	server.Listen()
}
