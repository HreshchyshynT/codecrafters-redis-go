package store

type Store struct {
	store map[Key]Data
}

type Key string

func NewStore() *Store {
	return &Store{
		store: make(map[Key]Data),
	}
}

func (s *Store) Put(key Key, data Data) {
	s.store[key] = data
}

/*
bool is true when Data exists under specified key and is not expired
*/
func (s *Store) Get(key Key) (Data, bool) {
	data, ok := s.store[key]

	// isExpired can't be true for zero value data
	// so either its true and we must clean it up
	// or we can return as it is in the map
	if data.isExpired() {
		delete(s.store, key)
		return Data{}, false
	}

	return data, ok
}
