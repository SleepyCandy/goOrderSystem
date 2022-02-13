package orderMain

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"loc-system-order/model"
	"net/http"
	"sync"
)

type RepositoryNonORM struct {
	db *sql.DB
}

func (r *RepositoryNonORM) RepositoryCustomerinitNonORM() {
	db, err := sql.Open("mysql", "puppysql:1234@tcp(68.183.178.30:3306)/ordersystem?charset=utf8&parseTime=True")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	r.db = db
}

func (r *RepositoryNonORM) InsertMain(c *gin.Context) {
	requestOrderMain := model.RequestOrderMain{}
	response := model.ResponseOrder{}
	if err := c.ShouldBindJSON(&requestOrderMain); err != nil {
		c.Status(http.StatusBadRequest)
		panic(err.Error())
		return
	}
	tx, err := r.db.Begin()
	orderInsert, err := tx.Prepare("INSERT INTO ordersystem.order_mains\n(ois_running_no, afs_user, status, status_message, company, product, branch, partner_code, create_date, create_by, create_by_name, update_date, update_by, update_by_name, complete_date, del_no)\nVALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);\n")
	itemInsert, err := tx.Prepare("INSERT INTO ordersystem.order_items\n(order_id, parent_item_id, item_name, status, status_message, contract_no, reference_no, channel)\nVALUES(?, ?, ?, ?, ?, ?, ?, ?);\n")
	orderMain := requestOrderMain.OrderMain[0]
	result, err := orderInsert.Exec(
		orderMain.OisRunningNo,
		orderMain.AfsUser,
		orderMain.Status,
		orderMain.StatusMessage,
		orderMain.Company,
		orderMain.Product,
		orderMain.Branch,
		orderMain.PartnerCode,
		orderMain.CreateDate,
		orderMain.CreateBy,
		orderMain.CreateByName,
		orderMain.UpdateDate,
		orderMain.UpdateBy,
		orderMain.UpdateByName,
		orderMain.CompletedDate,
		0,
	)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}
	lastID, err := result.LastInsertId()
	response.OrderMain = int(lastID)
	orderId := int(lastID)
	fmt.Printf("orderId ID=%d\n", lastID)
	orderInsert.Close()
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for _, item := range requestOrderMain.OrderMain[0].OrderItemList {
		wg.Add(1)
		go func(item model.OrderItemList) {
			//defer wg.Done()
			result, err = itemInsert.Exec(
				orderId,
				item.ParentItemId,
				item.ItemName,
				item.Status,
				item.StatusMessage,
				item.ContractNo,
				item.ReferenceNo,
				item.Channel,
			)
			lastID, err := result.LastInsertId()
			fmt.Printf("item ID=%d\n", lastID)
			if err != nil {
				fmt.Printf("item ID error ")
				tx.Rollback()
				panic(err.Error())
			}
			fmt.Printf("done..")
			wg.Done()
			ch <- int(lastID)
		}(item)
	}
	fmt.Printf("wait..")
	wg.Wait()
	tx.Commit()
	itemInsert.Close()
	for range requestOrderMain.OrderMain[0].OrderItemList {
		response.OrderItem = append(response.OrderItem, <-ch)
	}
	c.JSON(http.StatusOK, response)
}
