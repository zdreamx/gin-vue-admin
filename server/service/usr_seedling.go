package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSeedling
//@description: 创建Seedling记录
//@param: seed model.Seedling
//@return: err error

func CreateSeedling(seed model.Seedling) (err error) {
	err = global.GVA_DB.Create(&seed).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSeedling
//@description: 删除Seedling记录
//@param: seed model.Seedling
//@return: err error

func DeleteSeedling(seed model.Seedling) (err error) {
	err = global.GVA_DB.Delete(&seed).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSeedlingByIds
//@description: 批量删除Seedling记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSeedlingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Seedling{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSeedling
//@description: 更新Seedling记录
//@param: seed *model.Seedling
//@return: err error

func UpdateSeedling(seed model.Seedling) (err error) {
	err = global.GVA_DB.Save(&seed).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSeedling
//@description: 根据id获取Seedling记录
//@param: id uint
//@return: err error, seed model.Seedling

func GetSeedling(id uint) (err error, seed model.Seedling) {
	err = global.GVA_DB.Where("id = ?", id).First(&seed).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSeedlingInfoList
//@description: 分页获取Seedling记录
//@param: info request.SeedlingSearch
//@return: err error, list interface{}, total int64

func GetSeedlingInfoList(info request.SeedlingSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Seedling{})
    var seeds []model.Seedling
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Enterprise  != "" {
        db = db.Where("`enterprise ` LIKE ?","%"+ info.Enterprise +"%")
    }
    if info.Mobile != "" {
        db = db.Where("`mobile` LIKE ?","%"+ info.Mobile+"%")
    }
    if info.Technology  != "" {
        db = db.Where("`technology ` LIKE ?","%"+ info.Technology +"%")
    }
    if info.Process != "" {
        db = db.Where("`process` LIKE ?","%"+ info.Process+"%")
    }
    if info.SeedDate != "" {
        db = db.Where("`seed_date` LIKE ?","%"+ info.SeedDate+"%")
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&seeds).Error
	return err, seeds, total
}