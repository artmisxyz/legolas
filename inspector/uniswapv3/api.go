package uniswapv3

import (
	"github.com/artmisxyz/legolas/database"
	"github.com/artmisxyz/legolas/ent"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Api struct {
	storage *database.Storage
}

func NewApi(db *ent.Client) *Api {
	return &Api{
		storage: database.NewPostgresStorage(db),
	}
}

type EventsQuery struct {
	Cursor int    `form:"cursor"`
	Limit  int    `form:"limit,default=10"`
	From   uint64 `form:"from"`
	To     uint64 `form:"to"`
}

type Query struct {
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
	var req Query
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

func (a *Api) Burns(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	burns, err := a.storage.GetBurns(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, burns)
}

func (a *Api) Collects(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	collects, err := a.storage.GetCollects(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, collects)
}

func (a *Api) DecreaseLiquidity(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	dls, err := a.storage.GetDecreaseLiquidity(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, dls)
}

func (a *Api) PoolCreated(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	pls, err := a.storage.GetPoolCreated(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, pls)
}

func (a *Api) Flashes(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	flashes, err := a.storage.GetFlash(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, flashes)
}

func (a *Api) IncreaseLiquidity(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	ils, err := a.storage.GetIncreaseLiquidity(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, ils)
}

func (a *Api) PoolInitialize(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	pis, err := a.storage.GetPoolInitialize(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, pis)
}

func (a *Api) Mints(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	mints, err := a.storage.GetMints(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, mints)
}

func (a *Api) Transfers(c *gin.Context) {
	var req Query
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	transfers, err := a.storage.GetTransfers(req.Cursor, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, transfers)
}
