package main

import (
	"flag"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

const num_workers = 5

var initialize bool
var db_url string

func init() {
	flag.BoolVar(&initialize, "init", false, "Set to true for database initialization")
	flag.StringVar(&db_url, "db_url", "", "postgres connection url, ie [postgres://root@<url>:<port>/MovieDB?sslmode=disable]")
	flag.Parse()
}

func main() {
	log.Println("START")

	db, err := connectDatabase(db_url)

	if err != nil {
		log.Fatal(err)
	}
	defer LoggingClose(db.Close)

	sqlStore := func() MovieStorage {
		return &SqlMovieStorage{db: db}
	}

	repository := MovieRepository{storage: sqlStore}

	if initialize {
		setup(repository)
	}

	for i := 1; i <= num_workers; i++ {

		time.Sleep(time.Second / num_workers)

		go func(respository MovieRepository) {
			for {
				time.Sleep(time.Millisecond * 1000)
				movies, err := respository.All()

				if err != nil {
					log.Printf("Error: %s", err)
				} else {
					log.Printf("Successfully fetched %d results", len(movies))
				}
			}
		}(repository)

	}

	log.Println("CTRL-C to exit")
	waitForSigkill()

	log.Println("END")
}

func setup(repository MovieRepository) {

	movies := []Movie{
		Movie{Title: "Cloud atlas", Summery: "Greate scifi"},
		Movie{Title: "Lord of the rings", Summery: "Epic fantasy"},
		Movie{Title: "Die Hard", Summery: "yippiekayay"},
		Movie{Title: "Lock, Stock and Two Smoking Barrels", Summery: "A Guy Ritchie masterpiece"}}

	for _, movie := range movies {
		err := repository.Add(movie)

		if err != nil {
			log.Printf("Unable to insert movie with title %s due to error %s", movie.Title, err)
		} else {
			log.Printf("Inserted title %s", movie.Title)
		}
	}

}

func waitForSigkill() {
	var end_waiter sync.WaitGroup
	end_waiter.Add(1)
	signal_channel := make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		<-signal_channel
		log.Println("Sigkill called")
		end_waiter.Done()
	}()
	end_waiter.Wait()
}
