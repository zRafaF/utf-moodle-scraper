package backend

import (
	"log"
	"net/http"
	"os"

	"utf-moodle-scraper/internal/scraper"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var API_KEY string

func getApiKey() string {
	log.Println("Loading API_KEY from .env file...")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	envApiKey := os.Getenv("API_KEY")

	if envApiKey == "" {
		panic("API_KEY not found in .env file")
	}

	log.Println("API_KEY loaded successfully.")

	log.Println("API_KEY:", envApiKey)

	return envApiKey
}

func init() {
	API_KEY = getApiKey()
}

func Run() {
	var engine = gin.Default()

	engine.GET("/ping", getPing)
	engine.POST("/login-moodle", postLoginMoodle)
	engine.Run(":8080")
}

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func postLoginMoodle(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	apiKey := c.PostForm("api_key")

	log.Println("Received request with username:", username)

	if apiKey != API_KEY {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid API key.",
		})
		return
	}

	valid, err := scraper.ScrapeLogin("https://moodle.utfpr.edu.br/login/index.php", username, password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "authorized",
	})

}
