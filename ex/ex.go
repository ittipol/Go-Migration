package ex

import (
	"database/sql"
	"fmt"
	"migration/repository"

	"gorm.io/gorm"
)

func innerJoin(db *gorm.DB) {
	result := []repository.ItemToCategory{}

	// Join 3 tables
	// ItemToCategory > Item
	// ItemToCategory > Category
	db.Joins("Item").Joins("Category").Where("item_id=@itemId", sql.Named("itemId", 2)).Find(&result)

	fmt.Println(result)

	for _, v := range result {
		fmt.Printf("%v: %v | %v: %v\n", v.Item.ID, v.Item.Name, v.CategoryID, v.Category.Name)
	}
}

func joinWithConditionEx(db *gorm.DB) {
	result := []repository.ItemToCategory{}

	// Join 3 tables
	// ItemToCategory > Item
	// ItemToCategory > Category
	db.Joins("Item", db.Where(&repository.Item{Model: gorm.Model{ID: 2}})).Find(&result)
}

func join3Table(db *gorm.DB) {

	user := repository.User{}

	db.Joins("Company").Joins("Company.Product").Where("users.id = 3").Find(&user)
}
