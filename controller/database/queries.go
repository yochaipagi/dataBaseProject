package database

const (
	getArticleByID = `
SELECT page_number, line_number, word_number, word
FROM article_words
         JOIN article_lines al ON al.id = article_words.article_line_id
         JOIN article_pages ap ON al.article_page_id = ap.id
         JOIN articles a ON ap.article_id = a.id
WHERE a.id = ?
ORDER BY page_number, line_number, word_number
`

	getRawArticleByID = `
SELECT string_agg(article_lines.line, E'\n') as article,
       AVG(words_in_line)                    AS words_in_line,
       AVG(chars_in_word)                    AS chars_in_word,
       AVG(pages_num)                        AS pages_num
FROM (SELECT string_agg(word, ' ' ORDER BY word_number) AS line,
             AVG(a.pages_count)                         AS pages_num,
             AVG(al.word_count)                         AS words_in_line,
             AVG(char_count)                            AS chars_in_word
      FROM article_words
               JOIN article_lines al ON al.id = article_words.article_line_id
               JOIN article_pages ap ON al.article_page_id = ap.id
               JOIN articles a ON ap.article_id = a.id
      WHERE a.id = ?
      GROUP BY page_number, line_number
      ORDER BY page_number, line_number) AS article_lines`

	getArticles = `
SELECT *
FROM articles
`

	getWordGroups = `
SELECT *
FROM word_groups
`

	getLinguisticExprs = `
SELECT *
FROM linguistic_exprs
`

	getWordsIndex = `
SELECT LOWER(word)                                                      AS word,
       COUNT(word),
       string_agg(CONCAT(a.id, ',', ap.page_number::text, ',', al.line_number::text, ',', word_number::text), E'\n'
                  ORDER BY ap.page_number, al.line_number, word_number) AS index
FROM article_words aw
         JOIN article_lines al ON al.id = aw.article_line_id
         JOIN article_pages ap ON ap.id = al.article_page_id
         JOIN articles a ON a.id = ap.article_id
WHERE ?
  AND ?
GROUP BY LOWER(word)
ORDER BY word`

	wordsIndexWithWordGroup = `
LOWER(word) IN (SELECT w.word
                      FROM word_groups wg
                               JOIN words w ON wg.id = w.word_group_id
                      WHERE wg.id = %s)`

	getContextByPosition = `
SELECT line_number, string_agg(word, ' ') AS content
FROM article_words aw
         JOIN article_lines al ON al.id = aw.article_line_id
         JOIN article_pages ap ON ap.id = al.article_page_id
         JOIN articles a ON a.id = ap.article_id
WHERE a.id = ?
  AND line_number IN ?
GROUP BY line_number`
)
