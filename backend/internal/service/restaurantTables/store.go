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

func (s *Store) GetRestaurantTableByID(rTID, rID int) (*types.RestaurantTable, error) {
	var restaurantTable types.RestaurantTable

	stmt, err := s.db.Prepare("SELECT * FROM restaurant_tables WHERE id = ? AND restaurant_id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = stmt.QueryRow(rTID, rID).Scan(&restaurantTable.ID, &restaurantTable.TableNumber, &restaurantTable.Capacity, &restaurantTable.Status, &restaurantTable.RestaurantID, &restaurantTable.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no restaurant table found")
	}
	return &restaurantTable, nil
}

func (s *Store) GetRestaurantTables(rID int) (*[]types.RestaurantTable, error) {
	var restaurantTables []types.RestaurantTable

	stmt, err := s.db.Prepare("SELECT * FROM restaurant_tables WHERE restaurant_id = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	result, err := stmt.Query(rID)

	for result.Next() {
		var restaurantTable types.RestaurantTable
		err := result.Scan(&restaurantTable.ID, &restaurantTable.TableNumber, &restaurantTable.Capacity, &restaurantTable.Status, &restaurantTable.RestaurantID, &restaurantTable.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("internal server error when reading from restaurantTable")
		}
		restaurantTables = append(restaurantTables, restaurantTable)
	}

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("no restaurant table found")
	}

	return &restaurantTables, nil
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

	createdRTable, err := s.GetRestaurantTableByID(int(rTId), rT.RestaurantID)
	if err != nil {
		return nil, err
	}

	return createdRTable, nil
}

func (s *Store) UpdateRestaurantTable(rT types.UpdateRestaurantTable, rID, rTID int) (*types.RestaurantTable, error) {
	stmt, err := s.db.Prepare("UPDATE restaurant_tables SET status = ?, capacity = ? WHERE id = ? AND restaurant_id = ?")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(rT.Status, rT.Capacity, rTID, rID)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("internal server error")
	}

	updatedRT, err := s.GetRestaurantTableByID(rTID, rID)
	if err != nil {
		return nil, err
	}

	return updatedRT, nil
}
