package models

import (
	"go-admin/common/models"
)

type DFry struct {
	models.Model

	InviteCode   string `json:"inviteCode" gorm:"type:varchar(255);comment:InviteCode"`
	Phone        string `json:"phone" gorm:"type:varchar(255);comment:Phone"`
	Pwd          string `json:"pwd" gorm:"type:varchar(255);comment:Pwd"`
	Remark       string `json:"remark" gorm:"type:varchar(255);comment:Remark"`
	StateId      string `json:"stateId" gorm:"type:varchar(255);comment:StateId"`
	DcServerSalt string `json:"dcServerSalt" gorm:"type:varchar(255);comment:DcServerSalt"`
	DcAuthKey    string `json:"dcAuthKey" gorm:"type:text;comment:DcAuthKey"`
	Url          string `json:"url" gorm:"type:varchar(255);comment:Url"`
	UserAuthDcId string `json:"userAuthDcId" gorm:"type:int(11);comment:UserAuthDcId"`
	UserAuthDate string `json:"userAuthDate" gorm:"type:bigint(20);comment:UserAuthDate"`
	UserAuthId   string `json:"userAuthId" gorm:"type:bigint(20);comment:UserAuthId"`
	CreatedAt    string `json:"CreatedAt" gorm:"type:bigint(20);comment:CreatedAt"`
	UpdatedAt    string `json:"UpdatedAt" gorm:"type:bigint(20);comment:UpdatedAt"`
	models.ControlBy
}

func (DFry) TableName() string {
	return "D_fry"
}

func (e *DFry) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *DFry) GetId() interface{} {
	return e.Id
}
