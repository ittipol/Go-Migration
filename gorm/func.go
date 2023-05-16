package gorm

import "gorm.io/gorm"

type GormFunc interface {
	CreateTable(dst ...interface{}) error
	Create(value *interface{}) error
	Update() error
	Delete() error
}

type gormFunc struct {
	db *gorm.DB
}

func NewGormFunc(db *gorm.DB) GormFunc {
	return &gormFunc{db}
}

func (obj gormFunc) CreateTable(dst ...interface{}) error {
	err := obj.db.Migrator().CreateTable(dst)

	if err != nil {
		return err
	}

	return err
}

func (obj gormFunc) Create(value *interface{}) error {
	tx := obj.db.Create(value)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (obj gormFunc) Update() error {

	// // db.Save(&User{ID: 1, Name: "jinzhu", Age: 100})
	// UPDATE `users` SET `name`="jinzhu",`age`=100,`birthday`="0000-00-00 00:00:00",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1

	// Update with conditions
	// // db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// Update with conditions and model value
	// // db.Model(&user).Where("active = ?", true).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// Update attributes with `struct`, will only update non-zero fields
	// // db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	return nil

}

func (obj gormFunc) Delete() error {

	// Email's ID is `10`
	// db.Delete(&email)
	// DELETE from emails where id = 10;

	// Delete with additional conditions
	// db.Where("name = ?", "jinzhu").Delete(&email)
	// DELETE from emails where id = 10 AND name = "jinzhu";

	// db.Delete(&User{}, 10)
	// DELETE FROM users WHERE id = 10;

	// db.Delete(&User{}, "10")
	// DELETE FROM users WHERE id = 10;

	// db.Delete(&users, []int{1,2,3})
	// DELETE FROM users WHERE id IN (1,2,3);

	return nil
}
