package restaurantTables

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

func (s *Store) GetRestaurantTableByID(rTID int) (*types.RestaurantTable, error) {
	var restaurantTable types.RestaurantTable

	stmt, err := s.db.Prepare("SELECT * FROM restaurant_tables WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = stmt.QueryRow(rTID).Scan(&restaurantTable.ID, &restaurantTable.TableNumber, &restaurantTable.Capacity, &restaurantTable.Status, &restaurantTable.RestaurantID, &restaurantTable.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no restaurant table found")
	}
	return &restaurantTable, nil
}

func (s *Store) CreateRestaurantTable(rT types.CreateRestaurantTable) (*types.RestaurantTable, error) {
	stmt, err := s.db.Prepare("INSERT INTO restaurant_tables (table_number, status, capacity, restaurant_id) VALUES (?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	defer stmt.Close()

	result, err := stmt.Exec(rT.TableNumber, rT.Status, rT.Capacity, rT.RestaurantID)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	rTId, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	createdRTable, err := s.GetRestaurantTableByID(int(rTId))
	if err != nil {
		return nil, err
	}

	return createdRTable, nil
}
