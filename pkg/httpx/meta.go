package httpx

type Meta struct {
	Search     *Search     `json:"search"`
	Sort       *Sort       `json:"sort"`
	Pagination *Pagination `json:"pagination"`
}

type Search struct {
	Search string `json:"search"`
}

type Sort struct {
	OrderBy string `json:"order_by"`
	Sort    string `json:"sort"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func BuildMeta(meta *QueryParams, total int64) *Meta {
	page := 1
	if meta.Page != nil {
		page = *meta.Page
	}

	limit := meta.GetLimit()
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &Meta{
		Search: &Search{
			Search: meta.GetSearch(),
		},
		Sort: &Sort{
			OrderBy: meta.GetOrderBy(),
			Sort:    meta.GetSort(),
		},
		Pagination: &Pagination{
			Page:       page,
			Limit:      limit,
			Total:      int(total),
			TotalPages: totalPages,
		},
	}
}
