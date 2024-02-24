package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"strings"
)

const (
	articlesPath      = "../articles"
	authorPrefix      = "By "
	authorSuffix      = ", CNN"
	publishedAtPrefix = "Updated: "
	sourcePrefix      = "Source: "
	linesPerPage      = 10
)

const (
	titleIndex       = iota
	authorIndex      = iota
	publishedAtIndex = iota
	sourceIndex      = iota
	contentIndex     = iota
)

// populateDB creates multiple articles, word groups and linguistic expressions
func populateDB() error {
	var articlesToInsert []Article
	for i := 1; ; i++ {
		articlePath := fmt.Sprintf("%s/%d.txt", articlesPath, i)
		if _, err := os.Stat(articlePath); errors.Is(err, os.ErrNotExist) {
			break
		}

		rawArticle, err := os.ReadFile(articlePath)
		if err != nil {
			return err
		}

		article, err := parseRawArticle(string(rawArticle))
		if err != nil {
			return err
		}
		articlesToInsert = append(articlesToInsert, article)
	}

	wordGroupToInsert := WordGroup{
		Name: "Personal Pronouns",
		Words: []Word{
			{Word: "i"},
			{Word: "we"},
			{Word: "you"},
			{Word: "he"},
			{Word: "she"},
			{Word: "it"},
			{Word: "they"},
		},
	}

	lingToInsert := LinguisticExpr{
		Expression: "a letter",
	}

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
		return err
	}

	return nil
}

func parseRawArticle(rawArticle string) (Article, error) {
	rawArticle = strings.ReplaceAll(rawArticle, "\n\n", "\n")
	rawArticleLines := strings.Split(rawArticle, "\n")
	newArticle := NewArticle{
		Title:       rawArticleLines[titleIndex],
		Author:      trimPrefixOrSuffix(rawArticleLines[authorIndex], authorPrefix, authorSuffix),
		PublishedAt: strings.TrimPrefix(rawArticleLines[publishedAtIndex], publishedAtPrefix),
		Source:      trimPrefixOrSuffix(rawArticleLines[sourceIndex], sourcePrefix, ""),
		RawContent:  strings.Join(rawArticleLines[contentIndex:], "\n"),
	}
	return parseArticle(newArticle)
}
