package restaurants

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

func (s *Store) GetRestaurantByID(rID int, accessToken string) (*types.Restaurant, error) {
	var restaurant types.Restaurant

	stmt, err := s.db.Prepare("SELECT * FROM restaurants WHERE id = ? AND access_token = ?")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	err = stmt.QueryRow(rID, accessToken).Scan(&restaurant.ID,
		&restaurant.Name, &restaurant.Address,
		&restaurant.Phone, &restaurant.OpenHours,
		&restaurant.CloseHours, &restaurant.CuisineType,
		&restaurant.AccessToken, &restaurant.UserID, &restaurant.Capacity, &restaurant.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("no restaurant found")
	}
	return &restaurant, nil
}

func (s *Store) CreateRestaurant(r types.CreateRestaurant) (*types.Restaurant, error) {
	stmt, err := s.db.Prepare("INSERT INTO restaurants (name, address, phone, open_hours, close_hours, cuisine_type, access_token, user_id, capacity) VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	defer stmt.Close()

	result, err := stmt.Exec(r.Name, r.Address, r.Phone, r.OpenHours, r.CloseHours, r.CuisineType, r.AccessToken, r.UserID, r.Capacity)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	rId, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	createdR, err := s.GetRestaurantByID(int(rId), r.AccessToken)
	if err != nil {
		return nil, err
	}

	return createdR, nil
}

func (s *Store) UpdateRestaurant(r types.UpdateRestaurant, rAccessToken string, rID int) (*types.Restaurant, error) {
	stmt, err := s.db.Prepare("UPDATE restaurants SET name = ?, address = ?, phone = ?, open_hours = ?, close_hours = ?, cuisine_type = ?, capacity = ? WHERE id = ? AND access_token = ? ")
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	defer stmt.Close()

	_, err = stmt.Exec(r.Name, r.Address, r.Phone, r.OpenHours, r.CloseHours, r.CuisineType, r.Capacity, rID, rAccessToken)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	updatedR, err := s.GetRestaurantByID(rID, rAccessToken)
	if err != nil {
		return nil, err
	}

	return updatedR, nil
}

func (s *Store) DeleteRestaurant(rID int, accessToken string) (*types.HTTPGeneralRes, error) {
	_, err := s.GetRestaurantByID(rID, accessToken)
	if err != nil {
		return nil, err
	}

	stmt, err := s.db.Prepare("DELETE FROM restaurants WHERE id = ? AND access_token = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(rID, accessToken)
	if err != nil {
		return nil, err
	}

	return &types.HTTPGeneralRes{Success: true, Message: "Deleted restaurant with ID: " + fmt.Sprintf("%d", rID)}, nil
}
