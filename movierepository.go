package main


type MovieRepository struct {
	storage func() MovieStorage
}

func NewStore(storage func() MovieStorage) *MovieRepository {
	return &MovieRepository{storage: storage}
}

func (store *MovieRepository) Add(movie Movie) error {
	storage := store.storage()
	return storage.Add(movie)
}

func (store *MovieRepository) All() ([]Movie, error) {
	storage := store.storage()
	movies, err := storage.All()
	return movies, err
}

func (store *MovieRepository) Remove(movie Movie) error {
	storage := store.storage()
	err := storage.Remove(movie)
	return err
}
