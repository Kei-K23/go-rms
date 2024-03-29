package category

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-rms/backend/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCategoryByID(id int) (*types.Category, error) {
	var category types.Category

	stmt, err := s.db.Prepare("SELECT * FROM categories WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = stmt.QueryRow(id).Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no category found")
	}
	return &category, nil
}

func (s *Store) GetCategories() (*[]types.Category, error) {
	var categories []types.Category

	stmt, err := s.db.Prepare("SELECT * FROM categories")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	result, err := stmt.Query()

	for result.Next() {
		var category types.Category
		err := result.Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("internal server error when reading from restaurantTable")
		}
		categories = append(categories, category)
	}

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no categories found")
	}

	return &categories, nil
}

func (s *Store) CreateCategory(ct types.CreateCategory) (*types.Category, error) {
	stmt, err := s.db.Prepare("INSERT INTO categories(name, description) VALUES(?, ?)")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	defer stmt.Close()

	result, err := stmt.Exec(ct.Name, ct.Description)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	ctID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	createdCt, err := s.GetCategoryByID(int(ctID))
	if err != nil {
		return nil, err
	}

	return createdCt, nil
}

func (s *Store) UpdateCategory(ct types.UpdateCategory, id int) (*types.Category, error) {
	stmt, err := s.db.Prepare("UPDATE categories SET name = ?, description = ? WHERE id = ?")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(ct.Name, ct.Description, id)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}

	updatedCt, err := s.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return updatedCt, nil
}

func (s *Store) DeleteCategory(id int) (*types.HTTPGeneralRes, error) {
	_, err := s.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare("DELETE FROM categories WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return nil, err
	}

	return &types.HTTPGeneralRes{Success: true, Message: "Deleted menu with ID: " + fmt.Sprintf("%d", id)}, nil
}
