package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Sort       string      `json:"sort,omitempty"`
	TotalRows  int         `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

type PaginationRepository struct {
	Result interface{} // contains data and count data
	Error  error
}

type PaginationResponse struct {
	Links interface{} `json:"links"`
	Data  interface{} `json:"data"`
}

func GeneratePagination(c *gin.Context) *Pagination {
	// set default value
	page := 1
	limit := 10
	sort := "id desc"

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
		case "sort":
			sort = qValue
			break switchLabel
		}
	}

	return &Pagination{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}

	return p.Page
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}

	return p.Limit
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}

	return p.Sort
}

func (p *Pagination) GetOffset() int {
	offset := (p.GetPage() - 1) * p.GetLimit()

	return offset
}
