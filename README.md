# UTF Moodle Scraper

This is a simple login forwarder for different applications by web scraping the Moodle website of the Universidade Tecnológica Federal do Paraná.

> [!IMPORTANT]  
> This project was made with the explicit permission of the administrators of the targeted website, also requires explicit **permission form the user** to forward the credentials.

> [!WARNING]  
> You may not use this software without **EXPLICIT PERMISSION FROM ALL THE PARTS INVOLVED**.

## Routes

### GET `/`

Returns a string with "hello world"

### POST `/auth`

Auth path requires

```json
{
	"username": "abcdef",
	"password": "123123",
	"api_key": "a1b2c3"
}
```

If succeeded responds

```json
{
	"allow_login": true
}
```

## Using docker

This app's docker image is available at [zrafaf/utf-moodle-scraper](https://hub.docker.com/r/zrafaf/utf-moodle-scraper)

### Building a docker image

```sh
docker build -t utf-moodle-scraper .
```

### Running docker image

```sh
docker run -p 8080:8080 -e API_KEY=yourapikey utf-moodle-scraper

```

> Pay attention to the `-e` flag, it sets the `API_KEY` environment variable, without it the service will not start
