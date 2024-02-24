package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ozhey/concordance/controller/database"
	"net/http"
	"strconv"
)

func listArticles(c *gin.Context) {
	articles, err := db.ListArticles()
	handleResponse(c, articles, err)
}

func listWordGroups(c *gin.Context) {
	wordGroups, err := db.ListWordGroups()
	handleResponse(c, wordGroups, err)
}

func listLinguisticExpr(c *gin.Context) {
	exprs, err := db.ListLinguisticExpr()
	handleResponse(c, exprs, err)
}

func getArticle(c *gin.Context) {
	article, err := db.GetArticle(c.Params.ByName("id"))
	handleResponse(c, article, err)
}

func getWordsIndex(c *gin.Context) {
	articleWords, err := db.GetWordsIndex(c.Query("article_id"), c.Query("word_group_id"))
	handleResponse(c, articleWords, err)
}

func getWordByPosition(c *gin.Context) {
	word, err := db.GetWordByPosition(c.Query("article_id"), c.Query("line_num"), c.Query("word_num"))
	handleResponse(c, word, err)
}

func createArticle(c *gin.Context) {
	var body db.NewArticle
	err := c.BindJSON(&body)
	if err != nil {
		handleResponse(c, nil, err)
	}

	article, err := db.CreateArticle(body)
	handleResponse(c, article, err)
}

func createWordGroup(c *gin.Context) {
	var body db.WordGroup
	err := c.BindJSON(&body)
	if err != nil {
		handleResponse(c, nil, err)
	}

	wg, err := db.CreateWordGroup(body)
	handleResponse(c, wg, err)
}

func createLinguisticExpr(c *gin.Context) {
	var body db.LinguisticExpr
	err := c.BindJSON(&body)
	if err != nil {
		handleResponse(c, nil, err)
	}

	article, err := db.CreateLinguisticExpr(body)
	handleResponse(c, article, err)
}

func addWordToWordGroup(c *gin.Context) {
	var body db.Word
	err := c.BindJSON(&body)
	if err != nil {
		handleResponse(c, nil, err)
	}

	wordGroupId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		handleResponse(c, nil, err)
	}

	body.WordGroupID = uint(wordGroupId)
	article, err := db.AddWordToWordGroup(body)
	handleResponse(c, article, err)
}

func getLingExprPos(c *gin.Context) {
	exprPos, err := db.GetLingExprPos(c.Params.ByName("id"), c.Query("expr"))
	handleResponse(c, exprPos, err)
}

func benchmark(c *gin.Context) {
	replicates, err := strconv.Atoi(c.Query("replicates"))
	if err != nil {
		handleResponse(c, nil, err)
	}

	dbSize, err := strconv.Atoi(c.Query("db_size"))
	if err != nil {
		handleResponse(c, nil, err)
	}

	res, err := db.BenchmarkQuery(replicates, dbSize)
	handleResponse(c, res, err)
}

func handleResponse(c *gin.Context, res any, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}
