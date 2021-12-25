package model

import (
	"gin-use/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 新增分类
func CreateCate(data *Category) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户是否存在
func CheckCategory(name string) int {
	var cate Category
	db.Select("ID").Where("name = ?", name).First(&cate)

	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

func GetCategories(pageSize int, pageNum int) []Category {
	var categories []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categories
}

func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeletedCate(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
