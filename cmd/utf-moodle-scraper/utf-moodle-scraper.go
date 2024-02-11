package main

import (
	"utf-moodle-scraper/internal/backend"

	"log"
	// "utf-moodle-scrape/internal/scraper"
)

func main() {
	//backend.Run()
	log.Println("Starting service...")

	backend.Run()

}
