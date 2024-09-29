package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	model := db.Model(&product)
	model.Update("Price", 200)
	// Update - update multiple fields
	model.Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	model.Updates(map[string]interface{}{           // use map to update fields
		"Price": 200,
		"Code":  "F42",
	})

	// Delete - delete product
	db.Delete(&product, 1)
}
