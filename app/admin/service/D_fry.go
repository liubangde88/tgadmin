package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type DFry struct {
	service.Service
}

// GetPage 获取DFry列表
func (e *DFry) GetPage(c *dto.DFryGetPageReq, p *actions.DataPermission, list *[]models.DFry, count *int64) error {
	var err error
	var data models.DFry

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("DFryService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取DFry对象
func (e *DFry) Get(d *dto.DFryGetReq, p *actions.DataPermission, model *models.DFry) error {
	var data models.DFry

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetDFry error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建DFry对象
func (e *DFry) Insert(c *dto.DFryInsertReq) error {
    var err error
    var data models.DFry
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("DFryService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改DFry对象
func (e *DFry) Update(c *dto.DFryUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.DFry{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("DFryService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除DFry
func (e *DFry) Remove(d *dto.DFryDeleteReq, p *actions.DataPermission) error {
	var data models.DFry

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveDFry error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
