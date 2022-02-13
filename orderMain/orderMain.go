package orderMain

import "time"

type ORDER_MAIN struct {
	ORDER_ID       uint      `gorm:"primary_key" json:"ORDER_ID"`
	OIS_RUNNING_NO string    `json:"oisRunningNo"`
	AFS_USER       string    `json:"afsUser"`
	STATUS         string    `json:"status"`
	STATUS_MESSAGE string    `json:"statusMessage"`
	COMPANY        string    `json:"company"`
	PRODUCT        string    `json:"product"`
	BRANCH         string    `json:"branch"`
	PARTNER_CODE   string    `json:"partnerCode"`
	CREATE_DATE    time.Time `json:"createDate"`
	CREATE_BY      string    `json:"createBy"`
	CREATE_BY_NAME string    `json:"createByName"`
	UPDATE_DATE    time.Time `json:"updateDate"`
	UPDATE_BY      string    `json:"updateBy"`
	UPDATE_BY_NAME string    `json:"updateByName"`
	COMPLETE_DATE  time.Time `json:"completeDate"`
	DEL_NO         bool      `json:"delNo"`
}
