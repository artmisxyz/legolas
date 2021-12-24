package api

import (
	"fmt"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector/uniswapv3"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db   *ent.Client
	port int
}

func (s *Server) Init(conf *Config) {
	s.port = conf.Port
	db, err := ent.Open("sqlite3", "file:db?cache=shared&_fk=1")
	if err != nil {
		panic("failed opening connection to sqlite.")
	}
	s.db = db
}

func (s *Server) Listen() {
	r := gin.Default()
	univ3api := uniswapv3.NewApi(s.db)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "uniswapv3 events",
		})
	})

	r.GET("/events", univ3api.Events)
	r.GET("/swaps", univ3api.Swaps)
	r.Run(fmt.Sprintf(":%d", s.port))
}
