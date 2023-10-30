package constant

const (
	NamespaceHeader string = "saas-namespace"
	CtxUserId       string = "userId"
	CtxUserName     string = "userName"
	CtxUserEmail    string = "userEmail"
	CtxCookie       string = "cookie"

	QueryInclude string = "include"
	QueryFilter  string = "filter"
	QuerySortBy  string = "sortBy"
	QueryLimit   string = "limit"
	QueryOffset  string = "offset"
	SortByDesc   string = "desc"
	SortByAsc    string = "asc"

	OrderByDesc string = "DESC"
	OrderByAsc  string = "ASC"
)

const (
	Root             = "root"
	MinPageSize      = -1
	MaxPageSize      = 1000
	DefaultPageSize  = 100
	DefaultSortOrder = "ASC"
)

const (
	AndOperatorDelimiter    = ";"
	KeyValueDelimiter       = ":"
	SearchOperatorDelimiter = "."
	CommaDelimiter          = ","
	ColumnDelimiter         = "-"
)

type SearchOperator string

const (
	Equal              SearchOperator = "eq"
	NotEqual           SearchOperator = "neq"
	IsNull             SearchOperator = "isNull"
	IsNotNull          SearchOperator = "isNotNull"
	GreaterThan        SearchOperator = "gt"
	GreaterThanEqualTo SearchOperator = "gte"
	LessThan           SearchOperator = "lt"
	LessThanEqualTo    SearchOperator = "lte"
	Like               SearchOperator = "like"
	In                 SearchOperator = "in"
	NotIn              SearchOperator = "nin"
	Distinct           SearchOperator = "distinct"			
)

type QueryOperation string

const (
	QueryOr  QueryOperation = "OR"
	QueryAnd QueryOperation = "AND"
	QueryNot QueryOperation = "NOT"
)
