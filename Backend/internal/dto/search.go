package dto



type AssetQueryParams struct {
	Limit  	int 	`form:"limit"`
	Page 	int 	`form:"page"`
	Filter string 	`form:"filter"`
	SortBy string	`form:"sortBy"`
	
}
type SearchFilters map[string]map[string]map[string]string
