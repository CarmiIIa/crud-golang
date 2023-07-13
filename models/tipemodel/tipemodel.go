package tipemodel

import (
	"golang-crud/config"
	"golang-crud/entities"
)

func GetAll() []entities.Tipe {
	rows, err := config.DB.Query(`SELECT * FROM tipes`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var tipes []entities.Tipe

	for rows.Next() {
		var tipe entities.Tipe
		if err := rows.Scan(&tipe.Id, &tipe.Name, &tipe.CreatedAt, &tipe.UpdatedAt); err != nil{
			panic(err)
		}

		tipes = append(tipes, tipe)
	}

	return tipes

}

func Create(tipe entities.Tipe) bool {
	result, err := config.DB.Exec(`
		INSERT INTO tipes (name, created_at, updated_at)
		VALUE (? , ?, ?)`,
		tipe.Name, tipe.CreatedAt, tipe.UpdatedAt,
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

func Detail(id int) entities.Tipe {
	row := config.DB.QueryRow(`SELECT id, name FROM tipes WHERE id = ?`, id)

	var tipe entities.Tipe
	if err := row.Scan(&tipe.Id, &tipe.Name); err != nil {
		panic(err.Error())
	}

	return tipe
}

func Update(id int, tipe entities.Tipe) (bool, error) {
	query, err := config.DB.Exec(`
		UPDATE tipes SET 
			name = ?, 
			updated_at = ?
		WHERE id = ?
	`, 
		tipe.Name, 
		tipe.UpdatedAt, 
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
		Delete FROM tipes
		WHERE id = ?
	`, id)

	return err
}
