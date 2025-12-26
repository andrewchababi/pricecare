package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/andrewchababi/pricecare/backend/models"
)

func GetTestByID(id string) (*models.Test, error) {
	ctx := context.Background()
	const q = `
		SELECT id, name, reagent_cost, list_price
		FROM tests
		WHERE id = ?
	`
	var t models.Test
	err := database.QueryRowContext(ctx, q, id).Scan(
		&t.ID,
		&t.Name,
		&t.ReagentCost,
		&t.ListPrice,
	)

	if err != nil {
		log.Printf("Error querying this test : %v", id)
	}

	return &t, nil
}

func GetMultipleTestsByID(ids []string) ([]models.Test, error) {
	ctx := context.Background()
	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))

	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT id, name, reagent_cost, list_price
		FROM tests
		WHERE id IN (%s)
	`, strings.Join(placeholders, ","))

	rows, err := database.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tests := make([]models.Test, 0)

	for rows.Next() {
		var t models.Test
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.ReagentCost,
			&t.ListPrice,
		); err != nil {
			return nil, err
		}
		tests = append(tests, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tests, nil
}
