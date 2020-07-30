package models

// Product defines the product object
type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID int    `json:"categoryId"`
}
