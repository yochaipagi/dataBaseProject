package database

import (
	"strings"
	"time"
)

// dateLayout specifies the format expected for article publication dates.
const dateLayout = time.RFC1123

// parseArticle converts a NewArticle struct into an Article struct, structuring its content into pages and lines.
func parseArticle(newArticle NewArticle) (Article, error) {
	// Split the raw article content into lines.
	rawArticleLines := strings.Split(newArticle.RawContent, "\n")

	// Organize the lines into pages.
	pages := getArticlePages(rawArticleLines)

	// Parse the publication date using the specified layout.
	publishedAt, err := time.Parse(dateLayout, newArticle.PublishedAt)
	if err != nil {
		return Article{}, err // Return an error if the date parsing fails.
	}

	// Construct and return an Article struct with the parsed data.
	return Article{
		Title:        newArticle.Title,
		Author:       newArticle.Author,
		PublishedAt:  publishedAt,
		Source:       newArticle.Source,
		PagesCount:   len(pages),
		ArticlePages: pages,
	}, nil
}

// getArticlePages divides the article's lines into pages based on a predefined number of lines per page.
func getArticlePages(rawLines []string) []ArticlePage {
	// Parse the raw lines into ArticleLine structs.
	articleLines := parseLines(rawLines)
	var articlePages []ArticlePage

	// Distribute the lines across pages.
	for i := 1; len(articleLines) > 0; i++ {
		// Determine the number of lines on the current page.
		numOfLinesInPage := linesPerPage
		if len(articleLines) < linesPerPage {
			numOfLinesInPage = len(articleLines)
		}

		// Create a new ArticlePage and append it to the slice.
		articlePages = append(articlePages, ArticlePage{
			ArticleLines: articleLines[:numOfLinesInPage],
			PageNumber:   i,
		})
		// Move to the next set of lines for the next page.
		articleLines = articleLines[numOfLinesInPage:]
	}
	return articlePages
}

// parseLines converts raw text lines into ArticleLine structs, also parsing out individual words.
func parseLines(rawLines []string) (articleLines []ArticleLine) {
	for i, line := range rawLines {
		var lineWords []ArticleWord
		words := strings.Split(line, " ")
		// Parse each word in the line.
		for j, word := range words {
			lineWords = append(lineWords, ArticleWord{
				WordNumber: j + 1,
				Word:       word,
				CharCount:  len(word),
			})
		}
		// Add the parsed line to the articleLines slice.
		articleLines = append(articleLines, ArticleLine{
			LineNumber:   i + 1,
			ArticleWords: lineWords,
			WordCount:    len(lineWords),
		})
	}
	return
}

// trimPrefixOrSuffix removes a specified prefix and suffix from a word.
func trimPrefixOrSuffix(word string, prefix string, suffix string) string {
	word = strings.TrimPrefix(word, prefix)
	return strings.TrimSuffix(word, suffix)
}
