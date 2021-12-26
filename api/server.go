package api

import (
	"fmt"
	"github.com/artmisxyz/legolas/configs"
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector/uniswapv3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	db     *ent.Client
	port   int
	logger *zap.Logger
}

func (s *Server) Init(conf configs.Config) {
	s.port = conf.Server.Port
	db, err := database.NewPostgresClient(conf)
	if err != nil {
		panic(err)
	}
	s.db = db

}

func (s *Server) Listen() {
	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

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
