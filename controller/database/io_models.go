package database

// models for possible request body payloads
// handle data exchange between client and server

// NewArticle represents the structure for a new article payload, defining how data is received from the client for article creation.
type NewArticle struct {
	Title       string `json:"title" binding:"required"`        // Title of the article, required in the request payload.
	Author      string `json:"author" binding:"required"`       // Author of the article, required in the request payload.
	PublishedAt string `json:"published_at" binding:"required"` // Publication date of the article, required in the request payload.
	Source      string `json:"source" binding:"required"`       // Source of the article, required in the request payload.
	RawContent  string `json:"content" binding:"required"`      // Raw content of the article, required in the request payload.
}

// textLines is a slice of anonymous structs used to represent multiple lines of text, each with content and a line number.
type textLines []struct {
	Content    string `json:"content"`     // The actual content of the line.
	LineNumber int    `json:"line_number"` // The line number in the article or document.
}

// wordByPositionRes is the response structure that provides a word and its surrounding text lines.
type wordByPositionRes struct {
	Lines textLines `json:"lines"` // Contextual lines surrounding the word.
	Word  string    `json:"word"`  // The specific word that was requested.
}

// wordsRes is a slice of anonymous structs that detail words along with their position in an article or document.
type wordsRes []struct {
	PageNumber int    `json:"page_number"` // The page number where the word is found.
	LineNumber int    `json:"line_number"` // The line number where the word is found.
	WordNumber int    `json:"word_number"` // The position of the word in the line.
	Word       string `json:"word"`        // The word itself.
}

// benchmarkRes provides structured results for database benchmarking, including insertion times and query performance metrics.
type benchmarkRes struct {
	InsertionTime      int   `json:"insertion_time"`       // Time taken to insert test data during benchmarking.
	AverageQueryTime   int   `json:"average_query_time"`   // Average time taken per query during benchmarking.
	DiscreteQueryTimes []int `json:"discrete_query_times"` // Individual times for each query executed during benchmarking.
}
