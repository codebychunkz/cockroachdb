package main

type MovieRepository struct {
	storage func() MovieStorage
}

func NewStore(storage func() MovieStorage) *MovieRepository {
	return &MovieRepository{storage: storage}
}

func (store *MovieRepository) Add(movie Movie) error {

	var err error
	defer SimpleErrorRecovery(&err)

	storage := store.storage()
	err = storage.Add(movie)
	return err
}

func (store *MovieRepository) All() ([]Movie, error) {

	var err error
	defer SimpleErrorRecovery(&err)

	storage := store.storage()
	movies, err := storage.All()
	return movies, err
}

func (store *MovieRepository) Remove(movie Movie) error {

	var err error
	defer SimpleErrorRecovery(&err)

	storage := store.storage()
	err = storage.Remove(movie)
	return err
}
