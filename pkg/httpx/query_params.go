package httpx

import (
	"slices"
	"strings"
)

type QueryParams struct {
	Page    *int    `form:"page"`
	Limit   *int    `form:"limit"`
	Search  *string `form:"search"`
	OrderBy *string `form:"order_by"`
	Sort    *string `form:"sort" `
}

func (q *QueryParams) GetSearch() string {
	if q.Search != nil {
		return *q.Search
	}
	return ""
}

func (q *QueryParams) GetSort() string {
	if q.Sort != nil && (strings.ToUpper(*q.Sort) == "ASC" || strings.ToUpper(*q.Sort) == "DESC") {
		return strings.ToUpper(*q.Sort)
	}
	return "DESC"
}

func (q *QueryParams) GetOrderBy() string {
	if q.OrderBy != nil {
		return *q.OrderBy
	}
	return "id"
}

func (q *QueryParams) GetPage() int {
	if q.Page != nil && *q.Page > 0 {
		return *q.Page
	}
	return 1
}

func (q *QueryParams) GetLimit() int {
	if q.Limit != nil && *q.Limit > 0 {
		return *q.Limit
	}
	return 10
}

func (q *QueryParams) Offset() int {
	if q.Page != nil && q.Limit != nil && *q.Page > 1 {
		return (*q.Page - 1) * *q.Limit
	}
	return 0
}

func (q *QueryParams) IsValidOrderField(validFields []string) bool {
	if q.OrderBy == nil || *q.OrderBy == "" {
		return true
	}
	return slices.Contains(validFields, *q.OrderBy)
}
