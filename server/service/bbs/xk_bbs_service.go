package bbsService

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	bbsModel "github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type XkBbsService struct {
}

// 保存
func (bbsService *XkBbsService) CreateXkBbs(xkBbs *bbsModel.XkBbs) (err error) {
	//1.获取数据库连接对象
	err = global.GVA_DB.Create(xkBbs).Error
	return err
}

// 更新
func (bbsService *XkBbsService) UpdateXkBbs(xkBbs *bbsModel.XkBbs) (err error) {
	err = global.GVA_DB.Save(xkBbs).Error
	return err
}

// 删除
func (bbsService *XkBbsService) DeleteXkBbs(xkBbs *bbsModel.XkBbs) (err error) {
	err = global.GVA_DB.Delete(&xkBbs).Error
	return err
}

// 根据id查询
func (bbsService *XkBbsService) GetXkBbs(id uint) (xkBbs bbsModel.XkBbs, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xkBbs).Error
	return
}

// 分页
func (bbsService *XkBbsService) GetXkBbsPage(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bbsModel.XkBbs{})

	var XkBbsList []bbsModel.XkBbs
	err = db.Count(&total).Error
	if err != nil {
		return XkBbsList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&XkBbsList).Error
	}
	return XkBbsList, total, err
}
