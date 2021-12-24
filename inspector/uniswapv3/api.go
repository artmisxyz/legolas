package uniswapv3

import (
	"github.com/artmisxyz/legolas/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	storage *Postgres
}

func NewApi(db *ent.Client) *Api {
	return &Api{
		storage: NewPostgres(db),
	}
}

type EventsQuery struct {
	Cursor int    `form:"cursor"`
	Limit  int    `form:"limit,default=10"`
	From   uint64 `form:"from"`
	To     uint64 `form:"to"`
}

type SwapsQuery struct {
	Cursor int `form:"cursor"`
	Limit  int `form:"limit,default=10"`
}

func (a *Api) Events(c *gin.Context) {
	var req EventsQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	if (req.From == 0 && req.To != 0) || (req.From != 0 && req.To == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "you have to specify none or both"})
		return
	}
	if req.From == 0 && req.To == 0 {
		events, err := a.storage.GetEvents(req.Cursor, req.Limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, events)
		return
	}
	events, err := a.storage.GetEventsByBlockNumber(req.From, req.To, req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, events)
}

func (a *Api) Swaps(c *gin.Context) {
	var req SwapsQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	swaps, err := a.storage.GetSwaps(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, swaps)
}
