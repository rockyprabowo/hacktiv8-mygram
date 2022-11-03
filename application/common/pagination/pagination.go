package pagination

import (
	"github.com/jellydator/validation"
	"math"
)

const defaultPaginationLimit = 25
const maxPaginationLimit = 25

type Pagination[T any] struct {
	Data []T `json:"data"`
	State
}

type State struct {
	Limit      int   `json:"limit,omitempty" query:"limit"`
	Page       int   `json:"page,omitempty" query:"page"`
	Count      int64 `json:"count"`
	TotalPages int   `json:"total_pages"`
}

func (s *State) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.Limit, validation.Required),
		validation.Field(&s.Page, validation.Required),
	)
}

func (s *State) SetPaginateTotalCount(totalCount int64) {
	s.Count = totalCount
	totalPages := int(math.Ceil(float64(totalCount) / float64(s.GetLimit())))
	s.TotalPages = totalPages
}

func (s *State) GetOffset() int {
	return (s.GetPage() - 1) * s.GetLimit()
}

func (s *State) GetLimit() int {
	if s.Limit == 0 {
		s.Limit = defaultPaginationLimit
	}
	if s.Limit > maxPaginationLimit {
		s.Limit = maxPaginationLimit
	}
	return s.Limit
}

func (s *State) GetPage() int {
	if s.Page == 0 {
		s.Page = 1
	}
	return s.Page
}
