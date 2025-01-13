package store

import (
	"log/slog"

	"trill/db"
	"trill/types"
)

func GetCategories() (types.APICategories, error) {

	var categories types.APICategories
	var category types.APICategory

	query := db.Connection.Table("categories")
	query.Select("categories.id, categories.name")
	query.Where("categories.deleted_at IS NULL")
	query.Order("categories.id")
	rows, err := query.Rows()
	if err != nil {
		slog.Error(err.Error())
		return categories, err
	}
	defer rows.Close()

	count := 0
	all := types.APICategory{
		ID:   0,
		Name: "すべて",
	}
	categories.Categories = append(categories.Categories, all)
	for rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			slog.Error(err.Error())
			return categories, err
		}
		categories.Categories = append(categories.Categories, category)
		count++
	}
	categories.Quantity = count

	return categories, nil
}
