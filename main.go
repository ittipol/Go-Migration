package main

import (
	"database/sql"
	"fmt"
	"migration/orm/db"
	"migration/repository"
	"time"
)

func main() {

	println("start...")

	initTimeZone()

	// PostgreSQL
	// dsn := "host=localhost user=postgres password=1234 dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	// Mysql
	dsn := "root:1234@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"

	db := db.GetConnection(dsn, false)

	// err := db.Migrator().CreateTable(repository.Product{})

	// err := db.AutoMigrate(repository.Product{}, repository.User{}, repository.Company{})

	// if err != nil {
	// 	panic(err)
	// }

	// repository.MockData(db)

	// company := repository.Company{
	// 	Name: "Company A",
	// }

	// db.Create(&company)

	// tx := db.Create(&repository.User{
	// 	Name:      "Test Name 3",
	// 	CompanyID: company.ID,
	// })

	// if tx.Error != nil {
	// 	panic(tx.Error)
	// }

	// company := repository.Company{}

	// db.Find(&company, 1)

	// company.ProductId = 1

	// db.Save(company)

	// db.Model(&company).Updates(repository.Company{ProductId: 1})

	// User ============================================================

	users := []repository.User{}

	// Join 3 tables
	// User > Company > Product
	db.Joins("Company").Joins("Company.Product").Where("users.id IN @ids", sql.Named("ids", []int{2, 3})).Find(&users)

	for _, user := range users {
		fmt.Printf("%#v \n\n", user.Company.Product)
	}
	// ============================================================

	// Product ============================================================

	// var products = []repository.Product{}

	// OR
	// result := db.Where("id = @id1 OR id = @id2", sql.Named("id1", 1), sql.Named("id2", 2)).Find(&products)

	// Not
	// result := db.Not("id IN @ids", sql.Named("ids", []int{1, 2, 3, 4})).Find(&products)

	// println(result.RowsAffected)

	// fmt.Printf("%v\n", products)

	// ============================================================
}

func initTimeZone() {
	// LoadLocation looks for the IANA Time Zone database
	// List of tz database time zones
	// https: //en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	// init system time zone
	time.Local = location

	// timeInUTC := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	// fmt.Println(timeInUTC.In(location))
}
