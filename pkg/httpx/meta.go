package httpx

import (
	"github.com/faisallbhr/light-pos-be/database"
	"gorm.io/gorm"
)

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

func ApplyMetaQuery(db *database.DB, model any, params *QueryParams, searchFields []string) (*gorm.DB, int64, error) {
	query := db.Model(model)
	search := params.GetSearch()

	if search != "" && len(searchFields) > 0 {
		for i, field := range searchFields {
			if i == 0 {
				query = query.Where(field+" LIKE ?", "%"+search+"%")
			} else {
				query = query.Or(field+" LIKE ?", "%"+search+"%")
			}
		}
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query = query.Order(params.GetOrderBy() + " " + params.GetSort())
	query = query.Offset(params.Offset()).Limit(params.GetLimit())

	return query, total, nil
}
