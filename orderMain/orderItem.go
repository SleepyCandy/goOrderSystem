package orderMain

type ORDER_ITEM struct {
	OrderItem     int    `gorm:"primary_key" json:"orderItem"`
	OrderId       int    `json:"orderId"`
	ParentItemId  int    `json:"parentItemId"`
	ItemName      string `json:"itemName"`
	Status        string `json:"status"`
	StatusMessage string `json:"statusMessage"`
	ContractNo    string `json:"contractNo"`
	ReferenceNo   string `json:"referenceNo"`
	Channel       string `json:"channel"`
}
