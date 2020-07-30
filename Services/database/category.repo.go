package database

import (
	"database/sql"
	"webservice/handlers/logger"
	"webservice/models"
)

// CategoryRepo represents the DTL for category object
type CategoryRepo struct {
	db *sql.DB
}

// NewCategoryRepo creates new CategoryRepo
func NewCategoryRepo() *CategoryRepo {
	return &CategoryRepo{
		db: Db,
	}
}

// GetAll returns all the categories from database
func (cr *CategoryRepo) GetAll() []models.Category {
	var categoryList []models.Category
	row, err := cr.db.Query(`SELECT id, name, description FROM category`)
	if err != nil {
		logger.LogErr(err)
	}
	defer row.Close()

	for row.Next() {
		var category models.Category
		row.Scan(&category.ID, &category.Name, &category.Description)
		categoryList = append(categoryList, category)
	}

	return categoryList
}

// GetCategory returns all the product of specified category
func (cr *CategoryRepo) GetCategory(catID int) []models.Product {
	var productList []models.Product
	row, err := cr.db.Query(`SELECT id, name, price, categoryId FROM Product WHERE categoryId = ?`, catID)
	if err != nil {
		logger.LogErr(err)
	}
	defer row.Close()

	for row.Next() {
		var product models.Product
		row.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)
		productList = append(productList, product)
	}

	return productList
}

// Add inserts the prodcut to database
func (cr *CategoryRepo) Add(c models.Category) models.Category {
	_, err := cr.db.Exec(`INSERT into TheGreatWall.category (name, description) 
		VALUES (?, ?)`, c.Name, c.Description)
	if err != nil {
		logger.LogErr(err)
	}
	logger.LogInfo("Category successfully added to database")

	var category models.Category
	row := cr.db.QueryRow(`SELECT id, name, description FROM category WHERE id = last_insert_id()`)
	row.Scan(&category.ID, &category.Name, &category.Description)

	return category
}
