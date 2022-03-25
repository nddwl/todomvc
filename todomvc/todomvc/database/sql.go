package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `gorm:"size:50;notnull;Index:taskName" json:"name"`
	Condition int    `gorm:"tinyint;notnull;Index:taskCondition" json:"condition"`
}

// Client :数据库todomvc存在的前提
func Client(name string, password string) (db *gorm.DB, err error) {
	dns := name + ":" + password + "@(127.0.0.1:3306)/todomvc?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
func AutoMigrate(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&Task{})
	return err
}
func Select(db *gorm.DB) []Task {
	var l int64
	db.Model(&Task{}).Count(&l)
	task := make([]Task, l)
	db.Model(&Task{}).Find(&task)
	return task
}
func AddTask(db *gorm.DB, name string) {
	db.Model(&Task{}).Create(map[string]interface{}{
		"Condition": 1,
		"Name":      name,
	})
}
func Delete(db *gorm.DB, name string) {
	db.Model(&Task{}).Where("name = ?", name).Delete(&Task{})
}
func SelectAll(db *gorm.DB, condition int) []Task {
	if condition == -1 {
		var l int64
		db.Model(&Task{}).Count(&l)
		var task = make([]Task, l)
		db.Model(&Task{}).Find(&task)
		return task
	}
	var l int64
	db.Model(&Task{}).Where("condition = ?", condition).Count(&l)
	var task = make([]Task, l)
	db.Model(&Task{}).Where("condition = ?", condition).Find(&task)
	return task
}
func Sign(db *gorm.DB, name string, condition int) {
	if name == "1" {
		db.Model(&Task{}).Where("1 = 1").Update("condition", condition)
	}
	db.Model(&Task{}).Where("name = ?", name).Update("condition", condition)
}
func SelectLike(db *gorm.DB, name string, condition int) []Task {
	if condition == -1 {
		var l int64
		db.Model(&Task{}).Where("name LIKE ?", name).Count(&l)
		var task = make([]Task, l)
		db.Model(&task).Where("name LIKE ?", name).Find(&task)
		return task
	}
	var l int64
	db.Model(&Task{}).Where("name LIKE ?", name).Where("condition = ?", condition).Count(&l)
	var task = make([]Task, l)
	db.Model(&task).Where("name LIKE ?", name).Where("condition = ?", condition).Find(&task)
	return task
}
