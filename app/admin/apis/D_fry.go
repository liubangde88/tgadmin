package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type DFry struct {
	api.Api
}

// GetPage 获取DFry列表
// @Summary 获取DFry列表
// @Description 获取DFry列表
// @Tags DFry
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.DFry}} "{"code": 200, "data": [...]}"
// @Router /api/v1/D-fry [get]
// @Security Bearer
func (e DFry) GetPage(c *gin.Context) {
    req := dto.DFryGetPageReq{}
    s := service.DFry{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.DFry, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取DFry失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取DFry
// @Summary 获取DFry
// @Description 获取DFry
// @Tags DFry
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.DFry} "{"code": 200, "data": [...]}"
// @Router /api/v1/D-fry/{id} [get]
// @Security Bearer
func (e DFry) Get(c *gin.Context) {
	req := dto.DFryGetReq{}
	s := service.DFry{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.DFry

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取DFry失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建DFry
// @Summary 创建DFry
// @Description 创建DFry
// @Tags DFry
// @Accept application/json
// @Product application/json
// @Param data body dto.DFryInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/D-fry [post]
// @Security Bearer
func (e DFry) Insert(c *gin.Context) {
    req := dto.DFryInsertReq{}
    s := service.DFry{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建DFry失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改DFry
// @Summary 修改DFry
// @Description 修改DFry
// @Tags DFry
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.DFryUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/D-fry/{id} [put]
// @Security Bearer
func (e DFry) Update(c *gin.Context) {
    req := dto.DFryUpdateReq{}
    s := service.DFry{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改DFry失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除DFry
// @Summary 删除DFry
// @Description 删除DFry
// @Tags DFry
// @Param data body dto.DFryDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/D-fry [delete]
// @Security Bearer
func (e DFry) Delete(c *gin.Context) {
    s := service.DFry{}
    req := dto.DFryDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除DFry失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
