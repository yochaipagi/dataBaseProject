package database

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

const dropTables = "DROP TABLE article_words; DROP TABLE article_lines; DROP TABLE article_pages; DROP TABLE articles; DROP TABLE words; DROP TABLE word_groups; DROP TABLE linguistic_exprs;"

func ListArticles() (any, error) {
	var articles []Article
	res := DB.Raw(getArticles).Scan(&articles)
	return handleQueryResult(articles, res)
}

func ListWordGroups() (any, error) {
	var wordGroups []WordGroup
	res := DB.Raw(getWordGroups).Scan(&wordGroups)
	return handleQueryResult(wordGroups, res)
}

func ListLinguisticExpr() (any, error) {
	var lingExprs []LinguisticExpr
	res := DB.Raw(getLinguisticExprs).Scan(&lingExprs)
	return handleQueryResult(lingExprs, res)
}

func GetArticle(id string) (any, error) {
	var article struct {
		Article     string `json:"content"`
		WordsInLine string `json:"avg_words_in_line"`
		CharsInWord string `json:"avg_chars_in_word"`
		PagesNum    string `json:"avg_pages_num"`
	}
	res := DB.Raw(getRawArticleByID, id).Scan(&article)
	return handleQueryResult(article, res)

}

func GetWordsIndex(articleID string, wordGroupId string) (any, error) {
	var articleWords []struct {
		Word  string `json:"word"`
		Count int    `json:"count"`
		Index string `json:"index"`
	}
	articleFilter := "1=1"
	if articleID != "" {
		articleFilter = fmt.Sprintf("a.id = %s", articleID)
	}
	wordGroupFilter := "1=1"
	if wordGroupId != "" {
		wordGroupFilter = fmt.Sprintf(wordsIndexWithWordGroup, wordGroupId)
	}
	res := DB.Raw(getWordsIndex, gorm.Expr(articleFilter), gorm.Expr(wordGroupFilter)).Scan(&articleWords)
	return handleQueryResult(articleWords, res)

}

func GetWordByPosition(articleID string, lineNum string, wordNum string) (any, error) {
	lineNumInt, err := strconv.Atoi(lineNum)
	if err != nil {
		return nil, errors.Wrap(err, "convert line number to int")
	}

	lines, err := getWordContext(articleID, lineNumInt)
	if err != nil {
		return nil, err
	}

	wordNumInt, err := strconv.Atoi(wordNum)
	if err != nil {
		return nil, errors.Wrap(err, "convert word number to int")
	}

	var word string
	for _, line := range lines {
		if line.LineNumber == lineNumInt {
			words := strings.Split(line.Content, " ")
			word = words[wordNumInt-1]
		}
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].LineNumber < lines[j].LineNumber
	})

	return wordByPositionRes{
		Lines: lines,
		Word:  word,
	}, nil
}

func getWordContext(articleID string, lineNumInt int) (textLines, error) {
	linesToGet := fmt.Sprintf("(%d,%d,%d)", lineNumInt-1, lineNumInt, lineNumInt+1)
	lines := textLines{}
	tx := DB.Raw(getContextByPosition, articleID, gorm.Expr(linesToGet)).Scan(&lines)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return lines, nil
}

func CreateArticle(newArticle NewArticle) (any, error) {
	articleToInsert, err := parseArticle(newArticle)
	if err != nil {
		return nil, err
	}

	res := DB.Create(&articleToInsert)
	if res.Error != nil {
		return nil, res.Error
	}

	return articleToInsert, nil
}

func CreateWordGroup(group WordGroup) (any, error) {
	res := DB.Create(&group)
	if res.Error != nil {
		return nil, res.Error
	}

	return group, nil
}

func CreateLinguisticExpr(expr LinguisticExpr) (any, error) {
	res := DB.Create(&expr)
	if res.Error != nil {
		return nil, res.Error
	}

	return expr, nil
}

func AddWordToWordGroup(word Word) (any, error) {
	res := DB.Create(&word)
	if res.Error != nil {
		return nil, res.Error
	}

	return word, nil
}

func GetLingExprPos(articleId string, expr string) (any, error) {
	var words wordsRes
	tx := DB.Raw(getArticleByID, articleId).Scan(&words)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return getExprOccurrences(words, expr), nil
}

func getExprOccurrences(words wordsRes, expr string) wordsRes {
	exprList := strings.Split(expr, " ")
	matches := wordsRes{}
	for i := 0; i <= len(words)-len(exprList); i++ {
		j := 0
		for j < len(exprList) {
			if words[i+j].Word != exprList[j] {
				break
			}
			j += 1
		}
		if j == len(exprList) {
			matches = append(matches, words[i])
		}
	}
	return matches
}

func BenchmarkQuery(replicates int, dbSize int) (any, error) {
	insertionTime, err := duplicateDB(dbSize)
	if err != nil {
		return nil, err
	}

	randomArticle := &Article{}
	res := DB.First(randomArticle)
	if res.Error != nil {
		return nil, res.Error
	}

	log.Printf("queries benchmark: starting, fetched random article (id: %d)", randomArticle.ID)
	totalTime := 0
	var results []int
	for i := 0; i < replicates; i++ {
		start := time.Now()
		if _, err := GetWordsIndex(strconv.Itoa(int(randomArticle.ID)), ""); err != nil {
			return nil, err
		}
		elapsed := int(time.Since(start).Milliseconds())
		totalTime += elapsed
		results = append(results, elapsed)
	}
	log.Printf("queries benchmark: finished")

	if err := resetDB(); err != nil {
		return nil, err
	}

	return benchmarkRes{InsertionTime: insertionTime, AverageQueryTime: totalTime / replicates, DiscreteQueryTimes: results}, nil
}

func resetDB() error {
	log.Printf("reset db: dropping tables")
	if res := DB.Exec(dropTables); res.Error != nil {
		return res.Error
	}
	log.Printf("reset db: recreating tables and constraints")
	err := DB.AutoMigrate(&Article{}, &ArticleLine{}, &ArticleWord{}, &WordGroup{}, &Word{}, &LinguisticExpr{})
	if err != nil {
		return err
	}
	return populateDB()
}

func duplicateDB(dbSizeDuplicate int) (int, error) {
	start := time.Now()
	for i := 0; i < dbSizeDuplicate; i++ {
		if err := populateDB(); err != nil {
			return 0, err
		}
		log.Printf("duplicate db: now %dx bigger", i+1)
	}
	return int(time.Since(start).Milliseconds()), nil

}

func handleQueryResult(res any, tx *gorm.DB) (any, error) {
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return res, nil
}
