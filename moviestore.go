package main

type MovieStorage interface {
	Add(Movie) error
	All() ([]Movie, error)
	Remove(Movie) error
}
