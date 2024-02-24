package database

import (
	"strings"
	"time"
)

const dateLayout = time.RFC1123

func parseArticle(newArticle NewArticle) (Article, error) {
	rawArticleLines := strings.Split(newArticle.RawContent, "\n")
	pages := getArticlePages(rawArticleLines)
	publishedAt, err := time.Parse(dateLayout, newArticle.PublishedAt)
	if err != nil {
		return Article{}, err
	}
	return Article{
		Title:        newArticle.Title,
		Author:       newArticle.Author,
		PublishedAt:  publishedAt,
		Source:       newArticle.Source,
		PagesCount:   len(pages),
		ArticlePages: pages,
	}, nil
}

func getArticlePages(rawLines []string) []ArticlePage {
	articleLines := parseLines(rawLines)
	var articlePages []ArticlePage
	for i := 1; len(articleLines) > 0; i++ {
		numOfLinesInPage := linesPerPage
		if len(articleLines) < linesPerPage {
			numOfLinesInPage = len(articleLines)
		}
		articlePages = append(articlePages, ArticlePage{
			ArticleLines: articleLines[:numOfLinesInPage],
			PageNumber:   i,
		})
		articleLines = articleLines[numOfLinesInPage:]
	}
	return articlePages
}

func parseLines(rawLines []string) (articleLines []ArticleLine) {
	for i, line := range rawLines {
		var lineWords []ArticleWord
		words := strings.Split(line, " ")
		for j, word := range words {
			lineWords = append(lineWords, ArticleWord{
				WordNumber: j + 1,
				Word:       word,
				CharCount:  len(word),
			})
		}
		articleLines = append(articleLines, ArticleLine{
			LineNumber:   i + 1,
			ArticleWords: lineWords,
			WordCount:    len(lineWords),
		})
	}
	return
}

func trimPrefixOrSuffix(word string, prefix string, suffix string) string {
	word = strings.TrimPrefix(word, prefix)
	return strings.TrimSuffix(word, suffix)
}
