package orderMain

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loc-system-order/model"
	"log"
	"net/http"
	"sync"
	"time"
)

type OrderMainRepository struct {
	DB *gorm.DB
}

type ResponseSaveOrder struct {
	OrderMain int
	OrderItem []int
}

func (h *OrderMainRepository) RepositoryCustomerinit() {
	dsn := "puppysql:1234@tcp(68.183.178.30:3306)/ordersystem?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true, SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(ORDER_MAIN{})
	db.AutoMigrate(ORDER_ITEM{})
	h.DB = db
}

func (h *OrderMainRepository) SaveOrderTransaction(c *gin.Context) {
	orderMain := ORDER_MAIN{}
	response := ResponseSaveOrder{}
	requestOrderMain := model.RequestOrderMain{}
	if err := c.ShouldBindJSON(&requestOrderMain); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	pOrderMain := &orderMain
	orderMain = OrderMapping(&requestOrderMain)
	//tx := h.DB.Begin()
	tx := h.DB
	if err := tx.Save(&orderMain).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		tx.Rollback()
		return
	}
	id := pOrderMain.ORDER_ID
	response.OrderMain = int(id)
	//orderid := *pOrderMain.ORDER_ID
	orderitemlist := requestOrderMain.OrderMain[0].OrderItemList
	ch := make(chan int)
	fmt.Println(len(orderitemlist))
	var wg sync.WaitGroup
	for _, s := range orderitemlist {
		wg.Add(1)
		go h.mapOrderItemAndSave(tx, s, int(id), ch, &wg)
	}
	wg.Wait()
	tx.Commit()
	for _, _ = range orderitemlist {
		response.OrderItem = append(response.OrderItem, <-ch)
	}

	//loopselect(ch,exit ,&response)
	c.JSON(http.StatusOK, response)
}

func loopselect(ch, exit chan int, rs *ResponseSaveOrder) {
	for {
		select {
		case <-ch:
			rs.OrderItem = append(rs.OrderItem, <-ch)
		case <-exit:
			return
		}
	}
}

func OrderMapping(requestOrderMain *model.RequestOrderMain) (om ORDER_MAIN) {

	om.OIS_RUNNING_NO = requestOrderMain.OrderMain[0].OisRunningNo
	om.AFS_USER = requestOrderMain.OrderMain[0].AfsUser
	om.STATUS = requestOrderMain.OrderMain[0].Status
	om.STATUS_MESSAGE = requestOrderMain.OrderMain[0].StatusMessage
	om.COMPANY = requestOrderMain.OrderMain[0].Company
	om.PRODUCT = requestOrderMain.OrderMain[0].Product
	om.BRANCH = requestOrderMain.OrderMain[0].Branch
	om.PARTNER_CODE = requestOrderMain.OrderMain[0].PartnerCode
	om.CREATE_DATE = time.Now()
	om.CREATE_BY = requestOrderMain.OrderMain[0].CreateBy
	om.CREATE_BY_NAME = requestOrderMain.OrderMain[0].CreateByName
	om.UPDATE_DATE = time.Now()
	om.UPDATE_BY = requestOrderMain.OrderMain[0].UpdateBy
	om.UPDATE_BY_NAME = requestOrderMain.OrderMain[0].UpdateByName
	om.COMPLETE_DATE = time.Now()
	om.DEL_NO = false
	return om

}

func (h *OrderMainRepository) mapOrderItemAndSave(tx *gorm.DB, s model.OrderItemList, orderId int, ch chan<- int, wg *sync.WaitGroup) (ot ORDER_ITEM) {
	ot.OrderId = orderId
	ot.ParentItemId = s.ParentItemId
	ot.ItemName = s.ItemName
	ot.Status = s.Status
	ot.StatusMessage = s.StatusMessage
	ot.ContractNo = s.ContractNo
	ot.ReferenceNo = s.ReferenceNo
	ot.Channel = s.Channel
	if err := tx.Save(&ot).Error; err != nil {
		tx.Rollback()
		panic("error")
	}
	wg.Done()
	ch <- ot.OrderItem
	return ot
}
