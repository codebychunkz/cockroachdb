package main

import ()

type MovieRepository struct {
	storage func() MovieStorage
}

func NewStore(storage func() MovieStorage) *MovieRepository {
	return &MovieRepository{storage: storage}
}

func (store *MovieRepository) Initialize() error {
	/*store := store.storage()
	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS MovieDB"); err != nil {
		return err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS movies (id SERIAL PRIMARY KEY, title VARCHAR(128), summery VARCHAR(255))"); err != nil {
		return err
	}


		stmt, err := db.Prepare("INSERT INTO movies (title, summery) VALUES ($1, $2)")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmt.Exec("Cloud Atlas", "Great movie!")
		if err != nil {
			log.Fatal(err)
		}


	return err
	*/
	return nil
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
