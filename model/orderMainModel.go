package model

type RequestOrderMain struct {
	OrderMain []OrderMain `json:"OrderMain"`
}

type OrderMain struct {
	OisRunningNo  string          `json:"oisRunningNo"`
	AfsUser       string          `json:"afsUser"`
	Status        string          `json:"status"`
	StatusMessage string          `json:"statusMessage"`
	Company       string          `json:"company"`
	Product       string          `json:"product"`
	Branch        string          `json:"branch"`
	PartnerCode   string          `json:"partnerCode"`
	ContractNo    string          `json:"contractNo"`
	OrderItemList []OrderItemList `json:"orderItemList"`
	CreateDate    string          `json:"createDate"`
	CreateBy      string          `json:"createBy"`
	CreateByName  string          `json:"createByName"`
	UpdateDate    string          `json:"updateDate"`
	UpdateBy      string          `json:"updateBy"`
	UpdateByName  string          `json:"updateByName"`
	CompletedDate string          `json:"completedDate"`
}
type OrderItemList struct {
	OrderId             int                   `json:"orderId"`
	ParentItemId        int                   `json:"parentItemId"`
	ItemName            string                `json:"itemName"`
	Status              string                `json:"status"`
	StatusMessage       string                `json:"statusMessage"`
	ContractNo          string                `json:"contractNo"`
	ReferenceNo         string                `json:"referenceNo"`
	Channel             string                `json:"channel"`
	OrderProductKeyList []OrderProductKeyList `json:"orderProductKeyList"`
}

type OrderProductKeyList struct {
	ProductKeyName string `json:"productKeyName"`
	ProductKeyId   int    `json:"productKeyId"`
}
