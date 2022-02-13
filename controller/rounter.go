package controller

import (
	"github.com/gin-gonic/gin"
	"loc-system-order/orderMain"
)

func InitRounter() *gin.Engine {
	r := gin.Default()
	repo := orderMain.OrderMainRepository{}
	repo.RepositoryCustomerinit()

	non := orderMain.RepositoryNonORM{}
	non.RepositoryCustomerinitNonORM()
	r.POST("/save/orderMain", repo.SaveOrderTransaction)
	r.POST("/saveNon/orderMain", non.InsertMain)
	return r
}
