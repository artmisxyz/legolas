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
	s.redis = database.NewRedis(conf)
}

func (s *Server) Listen() {
	r := gin.New()
	r.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	univ3api := uniswapv3.NewApi(s.db)
	r.GET("/", func(c *gin.Context) {
		positionKeys, err := s.redis.Keys(context.Background(), "*:block:position").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not retrieve info"})
			return
		}
		finishKeys, err := s.redis.Keys(context.Background(), "*:block:finish").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not retrieve info"})
			return
		}
		positions := make(map[string]string)
		for _, k := range positionKeys {
			blockNumber, err := s.redis.Get(context.Background(), k).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not retrieve info"})
				return
			}
			positions[k] = blockNumber
		}
		finishes := make(map[string]string)
		for _, k := range finishKeys {
			blockNumber, err := s.redis.Get(context.Background(), k).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not retrieve info"})
				return
			}
			finishes[k] = blockNumber
		}
		c.JSON(200, gin.H{
			"positions": positions,
			"finishes":  finishes,
		})
	})

	r.GET("/events", univ3api.Events)
	r.GET("/swaps", univ3api.Swaps)
	r.Run(fmt.Sprintf(":%d", s.port))
}
