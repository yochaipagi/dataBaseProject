package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/yochaipagi/dataBaseProject/controller/database"
)

// listArticles handles the API endpoint for listing all articles. It communicates with the database to fetch the data and formats the response.
func listArticles(c *gin.Context) {
	articles, err := db.ListArticles() // Fetch all articles from the database.
	handleResponse(c, articles, err)   // Send the response to the client.
}

// listWordGroups handles the request for listing all word groups. It retrieves data from the database and sends a formatted response.
func listWordGroups(c *gin.Context) {
	wordGroups, err := db.ListWordGroups() // Fetch all word groups from the database.
	handleResponse(c, wordGroups, err)     // Format and send the response.
}

// listLinguisticExpr manages the endpoint for listing all linguistic expressions. It gets the data from the database and returns it to the client.
func listLinguisticExpr(c *gin.Context) {
	exprs, err := db.ListLinguisticExpr() // Retrieve all linguistic expressions.
	handleResponse(c, exprs, err)         // Format and return the response.
}

// getArticle is the handler for fetching a single article based on its ID. It extracts the ID from the URL, queries the database, and formats the response.
func getArticle(c *gin.Context) {
	article, err := db.GetArticle(c.Params.ByName("id")) // Get a specific article by ID.
	handleResponse(c, article, err)                      // Send the fetched article or an error message.
}

// getWordsIndex handles the endpoint to get an index of words, potentially filtered by article and word group. It formats and sends the response.
func getWordsIndex(c *gin.Context) {
	articleWords, err := db.GetWordsIndex(c.Query("article_id"), c.Query("word_group_id")) // Fetch the index of words.
	handleResponse(c, articleWords, err)                                                   // Respond with the data or an error.
}

// getWordByPosition retrieves a word based on its position in an article. It takes the article ID, line number, and word number as parameters.
func getWordByPosition(c *gin.Context) {
	word, err := db.GetWordByPosition(c.Query("article_id"), c.Query("line_num"), c.Query("word_num")) // Fetch the specific word.
	handleResponse(c, word, err)                                                                       // Send the word or an error.
}

// createArticle processes the POST request to create a new article. It parses the request body, creates the article, and returns the result.
func createArticle(c *gin.Context) {
	var body db.NewArticle
	if err := c.BindJSON(&body); err != nil { // Parse the JSON body to the NewArticle struct.
		handleResponse(c, nil, err)
		return
	}

	article, err := db.CreateArticle(body) // Insert the new article into the database.
	handleResponse(c, article, err)        // Respond with the created article or an error.
}

// createWordGroup handles the creation of a new word group from the POST request, inserting it into the database and returning the result.
func createWordGroup(c *gin.Context) {
	var body db.WordGroup
	if err := c.BindJSON(&body); err != nil { // Parse the request body to the WordGroup struct.
		handleResponse(c, nil, err)
		return
	}

	wg, err := db.CreateWordGroup(body) // Create the new word group in the database.
	handleResponse(c, wg, err)          // Respond with the word group or an error.
}

// createLinguisticExpr handles POST requests to create new linguistic expressions, inserting them into the database and formatting the response.
func createLinguisticExpr(c *gin.Context) {
	var body db.LinguisticExpr
	if err := c.BindJSON(&body); err != nil { // Parse the request body to the LinguisticExpr struct.
		handleResponse(c, nil, err)
		return
	}

	expr, err := db.CreateLinguisticExpr(body) // Insert the new expression into the database.
	handleResponse(c, expr, err)               // Send the response or an error.
}

// addWordToWordGroup adds a new word to an existing word group. It parses the request body and the word group ID from the URL.
func addWordToWordGroup(c *gin.Context) {
	var body db.Word
	if err := c.BindJSON(&body); err != nil { // Parse the word data from the request body.
		handleResponse(c, nil, err)
		return
	}

	wordGroupId, err := strconv.Atoi(c.Params.ByName("id")) // Convert the word group ID from string to integer.
	if err != nil {
		handleResponse(c, nil, err)
		return
	}

	body.WordGroupID = uint(wordGroupId)     // Set the WordGroupID in the word struct.
	word, err := db.AddWordToWordGroup(body) // Add the word to the specified word group in the database.
	handleResponse(c, word, err)             // Respond with the result or an error.
}

// getLingExprPos finds the positions of a linguistic expression within an article. It takes the article ID and the expression as parameters.
func getLingExprPos(c *gin.Context) {
	exprPos, err := db.GetLingExprPos(c.Params.ByName("id"), c.Query("expr")) // Fetch the positions of the expression.
	handleResponse(c, exprPos, err)                                           // Format and send the response.
}

// benchmark triggers a benchmarking operation based on the provided parameters: the number of replicates and the database size.
func benchmark(c *gin.Context) {
	replicates, err := strconv.Atoi(c.Query("replicates")) // Convert the replicate count from query parameters.
	if err != nil {
		handleResponse(c, nil, err)
		return
	}

	dbSize, err := strconv.Atoi(c.Query("db_size")) // Convert the database size from query parameters.
	if err != nil {
		handleResponse(c, nil, err)
		return
	}

	result, err := db.BenchmarkQuery(replicates, dbSize) // Execute the benchmarking operation.
	handleResponse(c, result, err)                       // Send back the benchmarking results or an error.
}

// handleResponse is a utility function to standardize API responses, encapsulating the logic for error checking and response formatting.
func handleResponse(c *gin.Context, res any, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Send an error response if there's an error.
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res}) // Send a success response with the data.
}
