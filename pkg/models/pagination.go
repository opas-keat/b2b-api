package models

type Pagination struct {
	Limit  int `query:"limit" json:"limit" validate:"required,max=100"`
	Offset int `query:"offset" json:"offset" validate:"min=0"`
}

func (p *Pagination) Safety() {
	if p.Limit == 0 {
		p.Limit = 20
	}
}

type ListRequest[T any] struct {
	Pagination
	Criteria *T `json:"criteria" validate:"omitempty,required"`
}

type General struct {
	GeneralSearch *string `json:"general_search"`
}

type ListResponse[T any] struct {
	Rows       []T   `json:"rows"`
	TotalCount int64 `json:"total_count"`
}
