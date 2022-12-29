package core_test

import (
	"testing"

	"github.com/davidborzek/tvhgo/core"
	"github.com/stretchr/testify/assert"
)

func TestSortQueryParamsValidate(t *testing.T) {
	p := core.SortQueryParams{
		SortKey:       "someKey",
		SortDirection: "asc",
	}

	err := p.Validate()
	assert.Nil(t, err)
}

func TestSortQueryParamsValidateEmpty(t *testing.T) {
	p := core.SortQueryParams{}

	err := p.Validate()
	assert.Nil(t, err)
}

func TestSortQueryParamsValidateInvalidDirection(t *testing.T) {
	p := core.SortQueryParams{
		SortDirection: "garbage",
	}

	err := p.Validate()
	assert.Equal(t, core.ErrSortInvalidDirection, err)
}

func TestPaginationQueryParamsValidate(t *testing.T) {
	p := core.PaginationQueryParams{
		Limit:  1,
		Offset: 1,
	}

	err := p.Validate()
	assert.Nil(t, err)
}

func TestPaginationQueryParamsValidateEmpty(t *testing.T) {
	p := core.PaginationQueryParams{}

	err := p.Validate()
	assert.Nil(t, err)
}

func TestPaginationQueryParamsValidateInvalidLimit(t *testing.T) {
	p := core.PaginationQueryParams{
		Limit: -1,
	}

	err := p.Validate()
	assert.Equal(t, core.ErrPaginationInvalidLimit, err)
}

func TestPaginationQueryParamsValidateInvalidOffset(t *testing.T) {
	p := core.PaginationQueryParams{
		Offset: -1,
	}

	err := p.Validate()
	assert.Equal(t, core.ErrPaginationInvalidOffset, err)
}

func TestPaginationSortQueryParamsValidate(t *testing.T) {
	p := core.PaginationSortQueryParams{
		PaginationQueryParams: core.PaginationQueryParams{
			Limit:  1,
			Offset: 1,
		},
		SortQueryParams: core.SortQueryParams{
			SortKey:       "someKey",
			SortDirection: "desc",
		},
	}

	err := p.Validate()
	assert.Nil(t, err)
}

func TestPaginationSortQueryParamsValidateInvalidPagination(t *testing.T) {
	p := core.PaginationSortQueryParams{
		PaginationQueryParams: core.PaginationQueryParams{
			Limit: -1,
		},
		SortQueryParams: core.SortQueryParams{
			SortKey:       "someKey",
			SortDirection: "desc",
		},
	}

	err := p.Validate()
	assert.Equal(t, core.ErrPaginationInvalidLimit, err)
}

func TestPaginationSortQueryParamsValidateInvalidSort(t *testing.T) {
	p := core.PaginationSortQueryParams{
		PaginationQueryParams: core.PaginationQueryParams{
			Limit: 1,
		},
		SortQueryParams: core.SortQueryParams{
			SortDirection: "garbage",
		},
	}

	err := p.Validate()
	assert.Equal(t, core.ErrSortInvalidDirection, err)
}

func TestPaginationSortQueryParamsMapToTvheadendQuery(t *testing.T) {
	p := core.PaginationSortQueryParams{
		PaginationQueryParams: core.PaginationQueryParams{
			Limit:  1,
			Offset: 2,
		},
		SortQueryParams: core.SortQueryParams{
			SortDirection: "asc",
			SortKey:       "field",
		},
	}

	q := p.MapToTvheadendQuery(map[string]string{
		"field": "mappedField",
	})

	assert.Equal(t, "1", q.Get("limit"))
	assert.Equal(t, "2", q.Get("start"))
	assert.Equal(t, "asc", q.Get("dir"))
	assert.Equal(t, "mappedField", q.Get("sort"))
}
