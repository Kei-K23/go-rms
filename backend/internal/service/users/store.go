package users

import "github.com/Kei-K23/go-rms/backend/internal/types"

type Store struct {
	store types.UserStore
}

func (s *Store) NewStore(store types.UserStore) *Store {
	return &Store{store: store}
}
