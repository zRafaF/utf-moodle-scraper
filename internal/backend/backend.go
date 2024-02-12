package backend

import (
	"log/slog"
	"net/http"
	"os"

	"utf-moodle-scraper/internal/scraper"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var API_KEY string

func getApiKey() string {
	slog.Info("Loading API_KEY from .env file...")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	envApiKey := os.Getenv("API_KEY")

	if envApiKey == "" {
		panic("API_KEY not found in .env file")
	}

	slog.Info("API_KEY loaded successfully.")

	return envApiKey
}

func init() {
	API_KEY = getApiKey()
}

// Run starts the backend service.
func Run(debug bool) {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	var engine = gin.Default()

	engine.GET("/", getRoot)
	engine.GET("/auth", getAuth)
	engine.POST("/auth", postAuth)
	engine.Run(":8080")

	slog.Info("Service started on port 8080")
}

func getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, world!")
	slog.Info("Received request to root.")
}

func getAuth(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"error":   "Method not allowed.",
		"message": "Use POST method to authenticate. Must contain 'username', 'password' and 'api_key' parameters.",
	})
}

func postAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	apiKey := c.PostForm("api_key")

	slog.Info("Received login request", "user", username)

	if username == "" || password == "" || apiKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters.",
		})
		return
	}

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
		"allow_login": true,
	})

}
