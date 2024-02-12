package backend

import (
	"log/slog"
	"net/http"
	"os"

	"utf-moodle-scraper/internal/scraper"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ApiKey   string `json:"api_key" binding:"required"`
}

var API_KEY string

func getApiKey() string {
	slog.Info("Loading API_KEY from .env file...")
	err := godotenv.Load()
	envApiKey := os.Getenv("API_KEY")
	if err != nil {
		slog.Warn("Error loading .env file", "error", err)
	}

	if envApiKey == "" {
		panic("API_KEY not found")
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
	slog.Info("Service started on port 8080")

	engine.Run(":8080")
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
	var authRequest AuthRequest
	err := c.BindJSON(&authRequest)

	if err != nil {
		slog.Error("Invalid JSON.", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON.",
		})
		return
	}

	username := authRequest.Username
	password := authRequest.Password
	apiKey := authRequest.ApiKey

	slog.Info("Received login request", "user", username)

	if username == "" || password == "" || apiKey == "" {
		slog.Warn("Missing parameters.")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing parameters.",
		})
		return
	}

	if apiKey != API_KEY {
		slog.Warn("Invalid API key.")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid API key.",
		})
		return
	}

	valid, err := scraper.ScrapeLogin("https://moodle.utfpr.edu.br/login/index.php", username, password)

	if err != nil {
		slog.Error("Error scraping login.", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !valid {
		slog.Warn("Invalid credentials.")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials.",
		})
		return
	}
	slog.Info("Login successful.", "user", username)
	c.JSON(http.StatusOK, gin.H{
		"allow_login": true,
	})

}
