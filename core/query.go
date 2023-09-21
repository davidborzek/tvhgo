package core

import (
	"errors"

	"github.com/davidborzek/tvhgo/tvheadend"
)

type (
	// PaginationQueryParams defines query params to paginate the result.
	PaginationQueryParams struct {
		// (Optional) Limit the result.
		Limit int64 `schema:"limit"`
		// (Optional) Offset the result.
		Offset int64 `schema:"offset"`
	}

	// SortQueryParams defines query params to sort the result.
	SortQueryParams struct {
		// (Optional) Sort key.
		SortKey string `schema:"sort_key"`
		// (Optional) Sort direction.
		SortDirection string `schema:"sort_dir"`
	}

	// PaginationSortQueryParams defines query params to paginate and sort the result.
	PaginationSortQueryParams struct {
		PaginationQueryParams
		SortQueryParams
	}
)

var (
	ErrPaginationInvalidLimit  = errors.New("pagination limit invalid")
	ErrPaginationInvalidOffset = errors.New("pagination offset invalid")
	ErrSortInvalidDirection    = errors.New("sort direction invalid")
)

func (p *PaginationQueryParams) Validate() error {
	switch {
	case p.Limit < 0:
		return ErrPaginationInvalidLimit
	case p.Offset < 0:
		return ErrPaginationInvalidOffset
	}
	return nil
}

func (p *SortQueryParams) Validate() error {
	if p.SortDirection != "" && p.SortDirection != "desc" && p.SortDirection != "asc" {
		return ErrSortInvalidDirection
	}

	return nil
}

func (p *PaginationSortQueryParams) Validate() error {
	if err := p.PaginationQueryParams.Validate(); err != nil {
		return err
	}
	if err := p.SortQueryParams.Validate(); err != nil {
		return err
	}

	return nil
}

func (p *PaginationQueryParams) MapToTvheadendQuery() tvheadend.Query {
	t := tvheadend.NewQuery()

	if p.Limit > 0 {
		t.Limit(p.Limit)
	}

	if p.Offset > 0 {
		t.Start(p.Offset)
	}

	return t
}

func (p *SortQueryParams) applyTvheadendQueryMapping(
	sortKeyMapping map[string]string,
	t *tvheadend.Query,
) {
	mappedKey, ok := sortKeyMapping[p.SortKey]
	if ok {
		t.SortKey(mappedKey)
	}

	if p.SortDirection != "" {
		t.SortDir(p.SortDirection)
	}
}

func (p *SortQueryParams) MapToTvheadendQuery(sortKeyMapping map[string]string) tvheadend.Query {
	t := tvheadend.NewQuery()
	p.applyTvheadendQueryMapping(sortKeyMapping, &t)
	return t
}

func (p *PaginationSortQueryParams) MapToTvheadendQuery(
	sortKeyMapping map[string]string,
) tvheadend.Query {
	t := p.PaginationQueryParams.MapToTvheadendQuery()
	p.applyTvheadendQueryMapping(sortKeyMapping, &t)
	return t
}
