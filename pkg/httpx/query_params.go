package httpx

type QueryParams struct {
	Page    *int    `json:"page"`
	Limit   *int    `json:"limit"`
	Search  *string `json:"search"`
	OrderBy *string `json:"order_by"`
	Sort    *string `json:"sort" `
}

func (q *QueryParams) GetSearch() string {
	if q.Search != nil {
		return *q.Search
	}
	return ""
}

func (q *QueryParams) GetSort() string {
	if q.Sort != nil {
		return *q.Sort
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
