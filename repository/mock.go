package repository

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func MockData(db *gorm.DB) error {

	var count int64
	tx := db.Model(&Product{}).Count(&count)

	if tx.Error != nil {
		return tx.Error
	}

	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	products := []Product{}
	for i := 0; i < 500; i++ {
		products = append(products, Product{
			Name:     fmt.Sprintf("Product %v", i+1),
			Quantity: random.Intn(100),
		})
	}

	return db.Create(&products).Error
}
