package menus

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

func (s *Store) GetAllMenuByRestaurantID(rID int) (*[]types.Menu, error) {
	var menus []types.Menu

	stmt, err := s.db.Prepare("SELECT * FROM menus WHERE restaurant_id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	result, err := stmt.Query(rID)

	for result.Next() {
		var menu types.Menu
		err := result.Scan(&menu.ID, &menu.Name, &menu.Description, &menu.Available, &menu.CategoryID, &menu.RestaurantID, &menu.Price, &menu.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("internal server error when reading from menus table")
		}
		menus = append(menus, menu)
	}

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no menu found")
	}

	return &menus, nil
}

func (s *Store) GetMenuByID(mID, rID int) (*types.Menu, error) {
	var menu types.Menu

	stmt, err := s.db.Prepare("SELECT * FROM menus WHERE id = ? AND restaurant_id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = stmt.QueryRow(mID, rID).Scan(&menu.ID, &menu.Name, &menu.Description, &menu.Available, &menu.CategoryID, &menu.RestaurantID, &menu.Price, &menu.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no menu found")
	}
	return &menu, nil
}

func (s *Store) CreateMenu(m types.CreateMenu, rID int) (*types.Menu, error) {
	stmt, err := s.db.Prepare("INSERT INTO menus (name, description, available, category_id, price, restaurant_id) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	defer stmt.Close()
	result, err := stmt.Exec(m.Name, m.Description, m.Available, m.CategoryID, m.Price, rID)
	if err != nil {
		fmt.Println(err, m.CategoryID)
		return nil, fmt.Errorf("internal server error")
	}

	mID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	createdRTable, err := s.GetMenuByID(int(mID), rID)
	if err != nil {
		return nil, err
	}

	return createdRTable, nil
}
