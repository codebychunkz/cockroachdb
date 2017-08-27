package main

import (
	"database/sql"
)

type SqlMovieStorage struct {
	db *sql.DB
}

func NewSqlMovieStorage(db *sql.DB) MovieStorage {
	return &SqlMovieStorage{db: db}
}

func (store *SqlMovieStorage) Add(movie Movie) error {
	db := store.db

	stmt, err := db.Prepare("INSERT INTO movies (title, summery) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(movie.Title, movie.Summery)
	if err != nil {
		return err
	}
	return nil
}

func (store *SqlMovieStorage) All() ([]Movie, error) {

	db := store.db

	rows, err := db.Query("SELECT id, title, summery FROM movies")
	if err != nil {
		return nil, err
	}
	defer LoggingClose(rows.Close)

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Summery); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

func (store *SqlMovieStorage) Remove(Movie) error {
	return nil
}

func connectDatabase(url string) (*sql.DB, error) {

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(num_workers)

	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS MovieDB"); err != nil {
		return nil, err
	}

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS movies (id SERIAL PRIMARY KEY, title VARCHAR(128), summery VARCHAR(255))"); err != nil {
		return nil, err
	}

	return db, err
}
