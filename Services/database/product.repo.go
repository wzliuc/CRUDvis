package database

import (
	"database/sql"
	"fmt"
	"webservice/handlers/logger"
	"webservice/models"
)

// ProductRepo represents the DTL for product object
type ProductRepo struct {
	db *sql.DB
}

// NewProductRepo creates new ProductRepo
func NewProductRepo() *ProductRepo {
	return &ProductRepo{
		db: Db,
	}
}

// Get returns the product of specified id from database
func (pr *ProductRepo) Get(id int) models.Product {
	var product models.Product
	row := pr.db.QueryRow(`SELECT id, name, price, categoryid FROM product WHERE id = ?`, id)
	row.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)

	return product
}

// GetAll returns all the products from database
func (pr *ProductRepo) GetAll() []models.Product {
	var productList []models.Product
	row, err := pr.db.Query(`SELECT id, name, price FROM product`)
	if err != nil {
		logger.LogErr(err)
	}
	defer row.Close()

	for row.Next() {
		var product models.Product
		row.Scan(&product.ID, &product.Name, &product.Price)
		productList = append(productList, product)
	}

	return productList
}

// Add inserts the prodcut to database
func (pr *ProductRepo) Add(p models.Product) models.Product {
	_, err := pr.db.Exec(`INSERT into TheGreatWall.product (name, price, categoryId) 
		VALUES (?, ?, ?)`, p.Name, p.Price, p.CategoryID)
	if err != nil {
		logger.LogErr(err)
	} else {
		logger.LogInfo("Product successfully added to database")
	}

	var product models.Product
	row := pr.db.QueryRow(`SELECT id, name, price, categoryId FROM product WHERE id = last_insert_id()`)
	row.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)

	return product
}

// Update updates the prodcut to database
func (pr *ProductRepo) Update(p models.Product) models.Product {
	_, err := pr.db.Exec(`
		Update TheGreatWall.product 
		SET name=?, price=?, categoryId=? 
		WHERE id = ?`,
		p.Name, p.Price, p.CategoryID, p.ID)
	if err != nil {
		logger.LogErr(err)
	} else {
		logger.LogInfo(fmt.Sprintf("Product with ID %v successfully updated", p.ID))
	}

	var product models.Product
	row := pr.db.QueryRow(`SELECT id, name, price, categoryId FROM product WHERE id = ?`, p.ID)
	row.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)

	return product
}

// Delete deletes the prodcut to database
func (pr *ProductRepo) Delete(ID int) {
	_, err := pr.db.Exec(`DELETE FROM TheGreatWall.product WHERE id = ?`, ID)
	if err != nil {
		logger.LogErr(err)
	} else {
		logger.LogInfo(fmt.Sprintf("Product with ID %v successfully deleted", ID))
	}
}
