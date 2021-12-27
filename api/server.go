package api

import (
	"context"
	"fmt"
	"github.com/artmisxyz/legolas/configs"
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/ent"
	"github.com/artmisxyz/legolas/inspector/uniswapv3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	db     *ent.Client
	redis  redis.UniversalClient
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
		all := s.db.Syncer.Query().AllX(context.Background())
		c.JSON(http.StatusOK, all)
	})

	r.GET("/events", univ3api.Events)
	r.GET("/swaps", univ3api.Swaps)
	r.Run(fmt.Sprintf(":%d", s.port))
}
