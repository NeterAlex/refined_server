package sql

import (
	"gorm.io/gorm"
)

func Create[T any](elements []*T) error {
	return DB.Create(elements).Error
}

func Delete[T any](id int64) error {
	var t T
	return DB.Where("id = ?", id).Delete(&t).Error
}

func Update[T any](id int64, element *T) error {
	var t T
	return DB.Model(t).Where("id = ?", id).Updates(element).Error
}

func Query[T any](where string, value string) ([]*T, int64, error) {
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

func Count[T any]() (int64, error) {
	var t T
	var total int64
	if err := DB.Model(&t).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func QueryExclude[T any](where, value, exclude string) ([]*T, int64, error) {
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

func QueryByOrder[T any](key, order string, limit int) ([]*T, error) {
	var t T
	db := DB.Model(t)
	var res []*T
	if err := db.Order(key + " " + order).Limit(limit).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryPreloadAll[T any](page, pageSize int64, preload string) ([]*T, int64, error) {
	var t T
	db := DB.Model(t)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*T
	if err := db.Preload(preload).Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func QueryAllExclude[T any](exclude string, page, pageSize int64) ([]*T, int64, error) {
	var t T
	db := DB.Model(t)
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*T
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Omit(exclude).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func CountPlus[T any](where, whereValue, item string, count int64) error {
	var t T
	return DB.Model(t).Where(where, whereValue).UpdateColumn(item, gorm.Expr(item+" + ?", count)).Error
}

func SumColumn[T any](column string) (int64, error) {
	var t T
	var result []int64
	var sum int64 = 0
	if err := DB.Model(t).Pluck(column, &result).Error; err != nil {
		return 0, err
	}
	for _, v := range result {
		sum += v
	}
	return sum, nil
}
