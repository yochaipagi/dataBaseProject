package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

// Set the path to the articles and prefixes/suffixes used in article parsing.
const (
	articlesPath      = "../articles"
	authorPrefix      = "By "
	authorSuf         = ", CNN"
	publishedAtPrefix = "Updated: "
	sourcePrefix      = "Source: "
	linesPerPage      = 10
)

// Index constants to identify parts of an article.
const (
	titleIndex = iota
	authorIndex
	publishedAtIndex
	sourceIndex
	contentIndex
)

// populateDB reads and inserts articles and other entities into the database.
func populateDB() error {
	var articlesToInsert []Article

	// Loop through article files.
	for i := 1; ; i++ {
		articlePath := fmt.Sprintf("%s/%d.txt", articlesPath, i)
		if _, err := os.Stat(articlePath); errors.Is(err, os.ErrNotExist) {
			break // Exit loop if file doesn't exist.
		}

		rawArticle, err := os.ReadFile(articlePath)
		if err != nil {
			return err // Stop and return error if file reading fails.
		}

		// Convert the article text to an Article struct.
		article, err := parseRawArticle(string(rawArticle))
		if err != nil {
			return err // Stop and return error if parsing fails.
		}
		articlesToInsert = append(articlesToInsert, article)
	}

	// Prepare a WordGroup with common personal pronouns.
	wordGroupToInsert := WordGroup{
		Name: "Personal Pronouns",
		Words: []Word{
			{Word: "i"}, {Word: "he"}, {Word: "she"}, {Word: "it"}, {Word: "we"}, {Word: "you"},
		},
	}

	// Prepare a linguistic expression to insert.
	lingToInsert := LinguisticExpr{
		Expression: "Even with",
	}

	// Insert data into the database within a transaction.
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(&articlesToInsert, 5).Error; err != nil {
			return err
		}
		if err := tx.Create(&wordGroupToInsert).Error; err != nil {
			return err
		}
		if err := tx.Create(&lingToInsert).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err // Return error if the transaction fails.
	}

	return nil // Return nil if everything succeeds.
}

// parseRawArticle creates an Article struct from raw article text.
func parseRawArticle(rawArticle string) (Article, error) {
	// Normalize and split the article into lines.
	rawArticle = strings.ReplaceAll(rawArticle, "\n\n", "\n")
	rawArticleLines := strings.Split(rawArticle, "\n")

	// Build the Article struct from the lines.
	newArticle := NewArticle{
		Title:       rawArticleLines[titleIndex],
		Author:      trimPrefixOrSuffix(rawArticleLines[authorIndex], authorPrefix, authorSuf),
		PublishedAt: strings.TrimPrefix(rawArticleLines[publishedAtIndex], publishedAtPrefix),
		Source:      trimPrefixOrSuffix(rawArticleLines[sourceIndex], sourcePrefix, ""),
		RawContent:  strings.Join(rawArticleLines[contentIndex:], "\n"),
	}
	return parseArticle(newArticle)
}
