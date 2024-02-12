run:
	@go run cmd/utf-moodle-scraper/utf-moodle-scraper.go --debug

ifeq ($(OS),Windows_NT)
build:
	@echo Building for Windows
	@go build -o utf-moodle-scraper.exe cmd/utf-moodle-scraper/utf-moodle-scraper.go
else
build:
	@echo Building for Linux
	@go build -o utf-moodle-scraper cmd/utf-moodle-scraper/utf-moodle-scraper.go
endif
