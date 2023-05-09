package sqlite

import (
	"Refined_service/biz/model/comment"
	"Refined_service/biz/model/post"
	"Refined_service/biz/model/user"
	"gorm.io/gorm"
)

func Create[T comment.Comment | post.Post | user.User](elements []*T) error {
	return DB.Create(elements).Error
}

func Delete[T comment.Comment | post.Post | user.User](id int64) error {
	var t T
	return DB.Where("id = ?", id).Delete(&t).Error
}

func Update[T comment.Comment | post.Post | user.User](id int64, element *T) error {
	var t T
	return DB.Model(&t).Where("id = ?", id).Updates(element).Error
}

func Count[T comment.Comment | post.Post | user.User]() (int64, error) {
	var t T
	var total int64
	if err := DB.Model(&t).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func Query[T comment.Comment | post.Post | user.User](where, value string) ([]*T, int64, error) {
	var t T
	db := DB.Model(t)
	if where != "" && value != "" {
		db = db.Where(where, value)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*T
	if err := db.Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func QueryExclude[T comment.Comment | post.Post | user.User](where, value, exclude string) ([]*T, int64, error) {
	var t T
	db := DB.Model(t)
	if where != "" && value != "" {
		db = db.Where(where, value)
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*T
	if err := db.Omit(exclude).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func QueryAll[T comment.Comment | post.Post | user.User](page, pageSize int64) ([]*T, int64, error) {
	var t T
	db := DB.Model(t)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*T
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func CountPlus[T comment.Comment | post.Post | user.User](where, whereValue, item string, count int64) error {
	var t T
	return DB.Model(t).Where(where, whereValue).UpdateColumn(item, gorm.Expr(item+" + ?", count)).Error
}
