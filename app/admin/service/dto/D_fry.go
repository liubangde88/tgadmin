package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type DFryGetPageReq struct {
	dto.Pagination     `search:"-"`
    DFryOrder
}

type DFryOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:D_fry"`
    InviteCode string `form:"inviteCodeOrder"  search:"type:order;column:invite_code;table:D_fry"`
    Phone string `form:"phoneOrder"  search:"type:order;column:phone;table:D_fry"`
    Pwd string `form:"pwdOrder"  search:"type:order;column:pwd;table:D_fry"`
    Remark string `form:"remarkOrder"  search:"type:order;column:remark;table:D_fry"`
    StateId string `form:"stateIdOrder"  search:"type:order;column:state_id;table:D_fry"`
    DcServerSalt string `form:"dcServerSaltOrder"  search:"type:order;column:dc_server_salt;table:D_fry"`
    DcAuthKey string `form:"dcAuthKeyOrder"  search:"type:order;column:dc_auth_key;table:D_fry"`
    Url string `form:"urlOrder"  search:"type:order;column:url;table:D_fry"`
    UserAuthDcId string `form:"userAuthDcIdOrder"  search:"type:order;column:user_auth_dc_id;table:D_fry"`
    UserAuthDate string `form:"userAuthDateOrder"  search:"type:order;column:user_auth_date;table:D_fry"`
    UserAuthId string `form:"userAuthIdOrder"  search:"type:order;column:user_auth_id;table:D_fry"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:D_fry"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:D_fry"`
    
}

func (m *DFryGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type DFryInsertReq struct {
    Id int `json:"-" comment:""` // 
    InviteCode string `json:"inviteCode" comment:""`
    Phone string `json:"phone" comment:""`
    Pwd string `json:"pwd" comment:""`
    Remark string `json:"remark" comment:""`
    StateId string `json:"stateId" comment:""`
    DcServerSalt string `json:"dcServerSalt" comment:""`
    DcAuthKey string `json:"dcAuthKey" comment:""`
    Url string `json:"url" comment:""`
    UserAuthDcId string `json:"userAuthDcId" comment:""`
    UserAuthDate string `json:"userAuthDate" comment:""`
    UserAuthId string `json:"userAuthId" comment:""`
    common.ControlBy
}

func (s *DFryInsertReq) Generate(model *models.DFry)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.InviteCode = s.InviteCode
    model.Phone = s.Phone
    model.Pwd = s.Pwd
    model.Remark = s.Remark
    model.StateId = s.StateId
    model.DcServerSalt = s.DcServerSalt
    model.DcAuthKey = s.DcAuthKey
    model.Url = s.Url
    model.UserAuthDcId = s.UserAuthDcId
    model.UserAuthDate = s.UserAuthDate
    model.UserAuthId = s.UserAuthId
}

func (s *DFryInsertReq) GetId() interface{} {
	return s.Id
}

type DFryUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    InviteCode string `json:"inviteCode" comment:""`
    Phone string `json:"phone" comment:""`
    Pwd string `json:"pwd" comment:""`
    Remark string `json:"remark" comment:""`
    StateId string `json:"stateId" comment:""`
    DcServerSalt string `json:"dcServerSalt" comment:""`
    DcAuthKey string `json:"dcAuthKey" comment:""`
    Url string `json:"url" comment:""`
    UserAuthDcId string `json:"userAuthDcId" comment:""`
    UserAuthDate string `json:"userAuthDate" comment:""`
    UserAuthId string `json:"userAuthId" comment:""`
    common.ControlBy
}

func (s *DFryUpdateReq) Generate(model *models.DFry)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.InviteCode = s.InviteCode
    model.Phone = s.Phone
    model.Pwd = s.Pwd
    model.Remark = s.Remark
    model.StateId = s.StateId
    model.DcServerSalt = s.DcServerSalt
    model.DcAuthKey = s.DcAuthKey
    model.Url = s.Url
    model.UserAuthDcId = s.UserAuthDcId
    model.UserAuthDate = s.UserAuthDate
    model.UserAuthId = s.UserAuthId
}

func (s *DFryUpdateReq) GetId() interface{} {
	return s.Id
}

// DFryGetReq 功能获取请求参数
type DFryGetReq struct {
     Id int `uri:"id"`
}
func (s *DFryGetReq) GetId() interface{} {
	return s.Id
}

// DFryDeleteReq 功能删除请求参数
type DFryDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *DFryDeleteReq) GetId() interface{} {
	return s.Ids
}
