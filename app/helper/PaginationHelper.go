package helper

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

type PaginationQuery struct {
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Order string `json:"order,omitempty"`
	Sort  string `json:"sort,omitempty"`
}

type PaginationLink struct {
	Current  string `json:"current_link,omitempty"`
	First    string `json:"first_url,omitempty"`
	Next     string `json:"next_url,omitempty"`
	Previous string `json:"previous_url,omitempty"`
	Last     string `json:"last_url,omitempty"`
}

func NewPagination() *Pagination {
	return &Pagination{
		TotalRows:  0,
		TotalPages: 0,
	}
}

func NewPaginationLink() *PaginationLink {
	return &PaginationLink{
		Current:  "",
		First:    "",
		Next:     "",
		Previous: "",
		Last:     "",
	}
}

func GeneratePaginationLink(appPATH string, pagination *Pagination, query *PaginationQuery, link *PaginationLink) {
	current, first, next, previous, last := "", "", "", "", ""

	if query.Page <= pagination.TotalPages {
		current = fmt.Sprintf("%s?page=%d&limit=%d&order=%s&sort=%s", appPATH, query.Page, query.Limit, query.Order, query.Sort)
	}

	if query.Page != 1 {
		first = fmt.Sprintf("%s?page=%d&limit=%d&order=%s&sort=%s", appPATH, 1, query.Limit, query.Order, query.Sort)
	}

	if query.Page+1 != pagination.TotalPages && query.Page < pagination.TotalPages {
		next = fmt.Sprintf("%s?page=%d&limit=%d&order=%s&sort=%s", appPATH, query.Page+1, query.Limit, query.Order, query.Sort)
	}

	if pagination.TotalPages > 2 && query.Page > 2 && query.Page <= pagination.TotalPages {
		previous = fmt.Sprintf("%s?page=%d&limit=%d&order=%s&sort=%s", appPATH, query.Page-1, query.Limit, query.Order, query.Sort)
	}

	if query.Page < pagination.TotalPages && pagination.TotalPages != 1 {
		last = fmt.Sprintf("%s?page=%d&limit=%d&order=%s&sort=%s", appPATH, pagination.TotalPages, query.Limit, query.Order, query.Sort)
	}

	link.Current = current
	link.First = first
	link.Next = next
	link.Previous = previous
	link.Last = last
}

func NewPaginationQuery(c *gin.Context) *PaginationQuery {
	// set default value
	page := 1
	limit := 10
	order := "id"
	sort := "desc"

	querys := c.Request.URL.Query()

	for key, value := range querys {
		qValue := value[len(value)-1] // query param value e.g => ?page=1 (1 is value)

	switchLabel: // labeling switch for break
		switch key {
		case "page":
			page, _ = strconv.Atoi(qValue)
			break switchLabel
		case "limit":
			limit, _ = strconv.Atoi(qValue)
			break switchLabel
		case "order":
			order = qValue
			break switchLabel
		case "sort":
			sort = qValue
			break switchLabel
		}
	}

	return &PaginationQuery{
		Page:  page,
		Limit: limit,
		Order: order,
		Sort:  sort,
	}
}

func (p *PaginationQuery) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}

	return p.Page
}

func (p *PaginationQuery) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

func (p *PaginationQuery) GetOrder() string {
	if p.Order == "" {
		p.Order = "id"
	}

	return p.Order
}

func (p *PaginationQuery) GetSort() string {
	if p.Sort == "" {
		p.Sort = "desc"
	}

	return p.Sort
}

func (p *PaginationQuery) GetOffset() int {
	offset := (p.GetPage() - 1) * p.GetLimit()

	return offset
}
