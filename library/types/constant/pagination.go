package constant

var (
	WhereClauseName   = "name"
	WhereClauseAuthor = "author"
	WhereClauseType   = "type"
)

type FilterBy string

var (
	FilterByName   FilterBy = "name = ?"
	FilterByAuthor FilterBy = "author = ?"
	FilterByType   FilterBy = "type = ?"
)
