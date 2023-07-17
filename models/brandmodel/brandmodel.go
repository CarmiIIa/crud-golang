package brandmodel

import (
	"golang-crud/config"
	"golang-crud/entities"
)

func GetAll() []entities.Brand {
	rows, err := config.DB.Query(`SELECT * FROM brands`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var brands []entities.Brand

	for rows.Next() {
		var brand entities.Brand
		if err := rows.Scan(&brand.Id, &brand.Name, &brand.CreatedAt, &brand.UpdatedAt); err != nil {
			panic(err)
		}

		brands = append(brands, brand)
	}

	return brands
}

func Create(brand entities.Brand) bool {
	result, err := config.DB.Exec(`
		INSERT INTO brands (name, created_at, updated_at)
		VALUE (?, ?, ?)`,
		brand.Name, brand.CreatedAt, brand.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsertId > 0
}

func Detail(id int) entities.Brand {
	row := config.DB.QueryRow(`
	SELECT 
		id,
		name
	FROM brands
	WHERE id = ?
	`, id)

	var brand entities.Brand
	if err := row.Scan(&brand.Id, &brand.Name); err != nil {
		panic(err.Error())
	}

	return brand
}

func Update(id int, brand entities.Brand) (bool, error) {
	query, err := config.DB.Exec(`
		UPDATE brands SET
			name = ?,
			updated_at = ?
		WHERE id = ?
	`,
		brand.Name,	
		brand.UpdatedAt,
		id,
	)

	if err != nil {
		return false, err
	}

	result, err := query.RowsAffected()
	if err != nil {
		return false, err
	}

	return result > 0, nil
}

func Delete(id int) error {
	_, err := config.DB.Exec(`
		Delete FROM brands
		WHERE id = ?
	`, id)

	return err
}