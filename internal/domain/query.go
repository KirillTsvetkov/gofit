package domain

import "go.mongodb.org/mongo-driver/mongo/options"

type PaginationQuery struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

type SearchQuery struct {
	Search string `form:"search"`
}
type GetWorkoutListQuery struct {
	PaginationQuery
	SearchQuery
}

func (p *PaginationQuery) GetSkip() *int64 {
	var limit *int64 = new(int64)

	limit = p.GetLimit()
	skip := *p.GetPage()**limit - *limit
	return &skip
}

func (p *PaginationQuery) GetPage() *int64 {
	if p.Page == 0 {
		p.Page = 1
	}

	return &p.Page
}

func (p *PaginationQuery) GetLimit() *int64 {
	if p.Limit == 0 {
		p.Limit = 15
	}

	return &p.Limit
}

func (p *PaginationQuery) GetPaginationOpts() *options.FindOptions {
	var opts *options.FindOptions
	if p != nil {
		opts = &options.FindOptions{
			Skip:  p.GetSkip(),
			Limit: p.GetLimit(),
		}
	}

	return opts
}
