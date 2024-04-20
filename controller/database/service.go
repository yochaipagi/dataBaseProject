package database

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// dropTables contains SQL commands to drop all the relevant tables in the database.
const dropTables = "DROP TABLE article_words; DROP TABLE article_lines; DROP TABLE article_pages; DROP TABLE articles; DROP TABLE words; DROP TABLE word_groups; DROP TABLE linguistic_exprs;"

// ListArticles retrieves all articles from the database.
func ListArticles() (any, error) {
	var articles []Article
	// Execute the predefined SQL query to fetch articles and scan the results into the articles slice.
	res := DB.Raw(getArticles).Scan(&articles)
	// Check the query result and return the articles or an error.
	return handleQueryResult(articles, res)
}

// ListWordGroups fetches all word groups from the database.
func ListWordGroups() (any, error) {
	var wordGroups []WordGroup
	// Execute the predefined SQL query to fetch word groups and scan the results into the wordGroups slice.
	res := DB.Raw(getWordGroups).Scan(&wordGroups)
	// Check the query result and return the word groups or an error.
	return handleQueryResult(wordGroups, res)
}

// ListLinguisticExpr retrieves all linguistic expressions from the database.
func ListLinguisticExpr() (any, error) {
	var lingExprs []LinguisticExpr
	// Execute the predefined SQL query to fetch linguistic expressions and scan the results into the lingExprs slice.
	res := DB.Raw(getLinguisticExprs).Scan(&lingExprs)
	// Check the query result and return the linguistic expressions or an error.
	return handleQueryResult(lingExprs, res)
}

// GetArticle fetches a specific article by its ID, including aggregated data like average words per line.
func GetArticle(id string) (any, error) {
	var article struct {
		Article     string `json:"content"`
		WordsInLine string `json:"avg_words_in_line"`
		CharsInWord string `json:"avg_chars_in_word"`
		PagesNum    string `json:"avg_pages_num"`
	}
	// Execute the predefined SQL query with the article ID to fetch its details.
	res := DB.Raw(getRawArticleByID, id).Scan(&article)
	// Check the query result and return the article details or an error.
	return handleQueryResult(article, res)
}

// GetWordsIndex returns an index of words, optionally filtered by article ID and word group ID.
func GetWordsIndex(articleID string, wordGroupId string) (any, error) {
	var articleWords []struct {
		Word  string `json:"word"`
		Count int    `json:"count"`
		Index string `json:"index"`
	}
	// Dynamically set filters based on the provided articleID and wordGroupId.
	articleFilter := "1=1"
	if articleID != "" {
		articleFilter = fmt.Sprintf("a.id = %s", articleID)
	}
	wordGroupFilter := "1=1"
	if wordGroupId != "" {
		wordGroupFilter = fmt.Sprintf(wordsIndexWithWordGroup, wordGroupId)
	}
	// Execute the query with the filters and scan the results into articleWords.
	res := DB.Raw(getWordsIndex, gorm.Expr(articleFilter), gorm.Expr(wordGroupFilter)).Scan(&articleWords)
	// Check the query result and return the word index or an error.
	return handleQueryResult(articleWords, res)
}

// GetWordByPosition retrieves a specific word based on its article ID, line number, and word number.
func GetWordByPosition(articleID string, lineNum string, wordNum string) (any, error) {
	// Convert the line number and word number from string to int.
	lineNumInt, err := strconv.Atoi(lineNum)
	if err != nil {
		return nil, errors.Wrap(err, "convert line number to int")
	}

	// Fetch context lines for the specified line number.
	lines, err := getWordContext(articleID, lineNumInt)
	if err != nil {
		return nil, err
	}

	wordNumInt, err := strconv.Atoi(wordNum)
	if err != nil {
		return nil, errors.Wrap(err, "convert word number to int")
	}

	// Find and return the specific word along with its context lines.
	var word string
	for _, line := range lines {
		if line.LineNumber == lineNumInt {
			words := strings.Split(line.Content, " ")
			word = words[wordNumInt-1]
		}
	}

	// Ensure the lines are sorted by their line numbers.
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].LineNumber < lines[j].LineNumber
	})

	// Return the word and its context lines.
	return wordByPositionRes{
		Lines: lines,
		Word:  word,
	}, nil
}

// getWordContext fetches lines surrounding a specific line number to provide context for the word.
func getWordContext(articleID string, lineNumInt int) (textLines, error) {
	// Define which lines to fetch based on the provided line number.
	linesToGet := fmt.Sprintf("(%d,%d,%d)", lineNumInt-1, lineNumInt, lineNumInt+1)
	lines := textLines{}
	// Execute the query to fetch the context lines and scan the results into lines.
	tx := DB.Raw(getContextByPosition, articleID, gorm.Expr(linesToGet)).Scan(&lines)
	// Check the query result and return the lines or an error.
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return lines, nil
}

// CreateArticle inserts a new article into the database.
func CreateArticle(newArticle NewArticle) (any, error) {
	// Parse the newArticle to conform to the database schema.
	articleToInsert, err := parseArticle(newArticle)
	if err != nil {
		return nil, err
	}

	// Insert the new article into the database.
	res := DB.Create(&articleToInsert)
	// Check the insertion result and return the new article or an error.
	if res.Error != nil {
		return nil, res.Error
	}

	return articleToInsert, nil
}

// CreateWordGroup inserts a new word group into the database.
func CreateWordGroup(group WordGroup) (any, error) {
	// Insert the new word group into the database.
	res := DB.Create(&group)
	// Check the insertion result and return the new word group or an error.
	if res.Error != nil {
		return nil, res.Error
	}

	return group, nil
}

// CreateLinguisticExpr inserts a new linguistic expression into the database.
func CreateLinguisticExpr(expr LinguisticExpr) (any, error) {
	// Insert the new linguistic expression into the database.
	res := DB.Create(&expr)
	// Check the insertion result and return the new expression or an error.
	if res.Error != nil {
		return nil, res.Error
	}

	return expr, nil
}

// AddWordToWordGroup adds a new word to an existing word group in the database.
func AddWordToWordGroup(word Word) (any, error) {
	// Insert the new word into the database, associating it with a word group.
	res := DB.Create(&word)
	// Check the insertion result and return the new word or an error.
	if res.Error != nil {
		return nil, res.Error
	}

	return word, nil
}

// GetLingExprPos finds all occurrences of a linguistic expression in an article.
func GetLingExprPos(articleId string, expr string) (any, error) {
	var words wordsRes
	// Fetch all words from the specified article.
	tx := DB.Raw(getArticleByID, articleId).Scan(&words)
	// Check the query result and return an error if necessary.
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	// Get all occurrences of the specified expression within the fetched words.
	return getExprOccurrences(words, expr), nil
}

// getExprOccurrences identifies all positions where a specific expression occurs within the provided words.
func getExprOccurrences(words wordsRes, expr string) wordsRes {
	exprList := strings.Split(expr, " ")
	matches := wordsRes{}

	// Search for the expression in the list of words.
	for i := 0; i <= len(words)-len(exprList); i++ {
		matchFound := true
		for j, wordPart := range exprList {
			if words[i+j].Word != wordPart {
				matchFound = false
				break
			}
		}

		// If the entire expression is found, add the starting word to the matches.
		if matchFound {
			matches = append(matches, words[i])
		}
	}

	return matches
}

// BenchmarkQuery performs a series of database queries to measure performance.
// It replicates queries based on a given number of replicates and increases the database size for the test.
func BenchmarkQuery(replicates int, dbSize int) (any, error) {
	// Attempt to duplicate the database contents to expand the database size for benchmarking purposes.
	// This helps in simulating a more loaded database environment.
	insertionTime, err := duplicateDB(dbSize)
	if err != nil {
		// Return an error if duplicating the database fails
		return nil, err
	}

	// Retrieve a random article from the database to use as the subject of the benchmarking queries.
	randomArticle := &Article{}
	res := DB.First(randomArticle)
	if res.Error != nil {
		// Return an error if fetching the article fails
		return nil, res.Error
	}

	// Log the start of the benchmark using the ID of the article
	log.Printf("Starting benchmark with article ID %d", randomArticle.ID)
	totalTime := 0
	var results []int // This will store the duration of each query replicate for analysis

	// Execute the benchmark query multiple times according to the number of replicates specified
	for i := 0; i < replicates; i++ {
		start := time.Now() // Record the start time of the query
		_, err := GetWordsIndex(strconv.Itoa(int(randomArticle.ID)), "")
		if err != nil {
			// Return an error if any query fails during the benchmarking
			return nil, err
		}
		elapsed := int(time.Since(start).Milliseconds()) // Calculate the elapsed time in milliseconds
		totalTime += elapsed                             // Add the elapsed time to the total time
		results = append(results, elapsed)               // Append the result for this replicate to the results slice
	}

	// Reset the database to its original state after benchmarking to avoid polluting subsequent tests.
	if err := resetDB(); err != nil {
		// Return an error if resetting the database fails
		return nil, err
	}

	// Compile the benchmarking results into a structured response.
	// This includes the total insertion time for duplicating the database, the average query time, and all individual query times.
	return benchmarkRes{
		InsertionTime:      insertionTime,
		AverageQueryTime:   totalTime / replicates,
		DiscreteQueryTimes: results,
	}, nil
}

// resetDB clears the database and reinitializes it.
func resetDB() error {
	log.Println("Resetting database...")
	if res := DB.Exec(dropTables); res.Error != nil {
		return res.Error
	}

	// Recreate the tables and constraints after dropping them.
	err := DB.AutoMigrate(&Article{}, &ArticleLine{}, &ArticleWord{}, &WordGroup{}, &Word{}, &LinguisticExpr{})
	if err != nil {
		return err
	}

	// Repopulate the database after resetting.
	return populateDB()
}

// duplicateDB replicates the database content to increase its size, used for benchmarking.
func duplicateDB(dbSizeDuplicate int) (int, error) {
	start := time.Now() // Record the start time for the duplication process.
	for i := 0; i < dbSizeDuplicate; i++ {
		// Populate the database multiple times to increase its size.
		if err := populateDB(); err != nil {
			return 0, err // Return an error if populating the database fails.
		}
		log.Printf("Database size increased: now %dx larger.", i+1)
	}
	// Calculate and return the total time taken for the duplication process.
	return int(time.Since(start).Milliseconds()), nil
}

// handleQueryResult processes the result of a database query, checking for errors and data presence.
func handleQueryResult(res any, tx *gorm.DB) (any, error) {
	if tx.Error != nil {
		return nil, tx.Error // Return an error if the query encountered an issue.
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("No data found.") // Return an error if no rows were affected.
	}

	// Return the query result and nil error if everything is fine.
	return res, nil
}
