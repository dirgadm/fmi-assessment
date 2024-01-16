package ehttp

import (
	"errors"
	"math"
	"strconv"
)

type (
	// Pagination
	Paginator struct {
		Page     int  `json:"page"`
		PerPage  int  `json:"per_page"`
		Offset   int  `json:"offset,omitempty"`
		Limit    int  `json:"limit,omitempty"`
		NumPages int  `json:"total_pages"`
		Start    int  `json:"start,omitempty"`
		End      int  `json:"end,omitempty"`
		HasPrev  bool `json:"has_prev,omitempty"`
		HasNext  bool `json:"has_next,omitempty"`
	}
)

func NewPaginator(ctx *Context) (p *Paginator, err error) {
	page := ctx.GetPage()
	perPage := ctx.GetPerPage()
	offset := ctx.GetOffset()
	limit := ctx.GetLimit()

	if page < 1 {
		err = errors.New("invalid page number")
		return
	}

	p = &Paginator{
		Page:    page,
		PerPage: perPage,
		Offset:  offset,
		Limit:   limit,
	}

	p.Start = p.Offset
	p.End = p.Start + perPage

	return p, nil
}

func (p *Paginator) Json(totalItems int64) (page *Paginator) {
	if p.Page > 1 {
		p.HasPrev = true
	}

	if p.End < int(totalItems) {
		p.HasNext = true
	}

	p.NumPages = int(math.Ceil(float64(totalItems) / float64(p.PerPage)))

	if p.Page > p.NumPages {
		return
	}

	page = p
	return
}

const PerPage = 10

// GetPage get params page for pagination
func (c *Context) GetPage() int {
	p := c.QueryParam("page")

	if p == "" {
		return 1
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		page = 1
	}
	return page
}

func (c *Context) GetLimit() int {
	p := c.QueryParam("limit")

	if p == "" {
		return 10
	}
	limit, err := strconv.Atoi(p)
	if err != nil {
		limit = 1
	}
	return limit
}

func (c *Context) GetOffset() int {
	p := c.QueryParam("offset")

	if p == "" {
		return 0
	}
	offset, err := strconv.Atoi(p)
	if err != nil {
		offset = 1
	}
	return offset
}

// GetPerPage get params per_page for pagination
func (c *Context) GetPerPage() int {
	p := c.QueryParam("per_page")
	if p == "" {
		return PerPage
	}
	perPage, err := strconv.Atoi(p)
	if err != nil {
		perPage = PerPage
	}
	return perPage
}
