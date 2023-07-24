package productmodel

import (
	"golang-crud/config"
	"golang-crud/entities"
)

func Getall() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			tipes.name as tipe_name,
			brands.name as brand_name,
			products.stock, 
			products.status, 
			products.description, 
			products.created_at, 
			products.updated_at 
		FROM products
		JOIN categories ON products.category_id = categories.id
		JOIN tipes ON products.tipe_id = tipes.id
		JOIN brands ON products.brand_id = brands.id
	`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Tipe.Name,
			&product.Brand.Name,
			&product.Stock,
			&product.Status,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
		INSERT INTO products(
			name, category_id, tipe_Id, brand_id, status, stock, description, created_at, updated_at
		) VALUES (?,?,?,?,?,?,?,?,?)`,
		product.Name,
		product.Category.Id,
		product.Tipe.Id,
		product.Brand.Id,
		product.Status,
		product.Stock,
		product.Description,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			tipes.name as tipe_name,
			brands.name as brand_name,
			products.stock, 
			products.status, 
			products.description, 
			products.created_at, 
			products.updated_at 
		FROM products
		JOIN categories ON products.category_id = categories.id
		JOIN tipes ON products.tipe_id = tipes.id
		JOIN brands ON products.brand_id = brands.id
		WHERE products.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Tipe.Name,
		&product.Brand.Name,
		&product.Stock,
		&product.Status,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
		UPDATE products SET
			name = ?,
			category_id = ?,
			tipe_id = ?,
			brand_id = ?,
			stock = ?,
			description = ?,
			updated_at = ?
		WHERE id = ?
	`,
		product.Name,
		product.Category.Id,
		product.Tipe.Id,
		product.Brand.Id,
		product.Stock,
		product.Description,
		product.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}

	rowsAffected, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return rowsAffected > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`Delete FROM products WHERE id = ?`, id)
	return err
}
