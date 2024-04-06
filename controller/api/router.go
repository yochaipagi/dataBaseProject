package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the Gin router, setting up routes and their corresponding handlers.
func SetupRouter() *gin.Engine {
	// Initialize the default Gin router with default middleware (logger and recovery middleware).
	r := gin.Default()

	// Configure CORS (Cross-Origin Resource Sharing) settings for the API to allow interactions from different origins.
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},                                                                      // Allow all origins
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},                                         // Allow specific HTTP methods
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"}, // Allow specific headers
	}))

	// Define API endpoints and associate them with handler functions.

	// Article-related routes:
	r.GET("/articles", listArticles)                     // Fetches a list of articles.
	r.POST("/articles", createArticle)                   // Creates a new article.
	r.GET("/articles/:id", getArticle)                   // Fetches a specific article by ID.
	r.GET("/articles/:id/ling_expr_pos", getLingExprPos) // Fetches linguistic expression positions within an article.

	// Word-related routes:
	r.GET("/article_words", getWordByPosition)   // Fetches a word by its position in an article.
	r.GET("/article_words/index", getWordsIndex) // Fetches an index of words within articles.

	// Word group-related routes:
	r.GET("/word_groups", listWordGroups)          // Lists all word groups.
	r.POST("/word_groups", createWordGroup)        // Creates a new word group.
	r.POST("/word_groups/:id", addWordToWordGroup) // Adds a new word to a word group.

	// Linguistic expression-related routes:
	r.GET("/ling_exprs", listLinguisticExpr)    // Lists all linguistic expressions.
	r.POST("/ling_exprs", createLinguisticExpr) // Creates a new linguistic expression.

	// Benchmark-related route:
	r.POST("/benchmark", benchmark) // Triggers a benchmarking process. Warning: this can reset the database.

	// Return the configured router.
	return r
}
