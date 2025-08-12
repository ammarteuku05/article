// Package pagination provides support for pagination requests and responses.
package pagination

var (
	// DefaultPageSize specifies the default page size
	DefaultPageSize = 10
	// MaxPageSize specifies the maximum page size
	MaxPageSize = 1000
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies the query parameter name for page size
	PageSizeVar = "per_page"

	// default sort "ASC"
	SORT = "asc"

	// default limit 10
	LIMIT = 10

	// default max limit 1000
	MAX_LIMIT = 1000

	// default offset 1
	OFFSET = 1
)

// Pages represents a paginated list of data items.
type Pages struct {
	Pagination
	Items interface{} `json:"items"`
}

// Pagination contains pagination information
type Pagination struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"total_page"`
	TotalCount int `json:"total_data"`
}

type RequestPagination struct {
	Sort      *string `query:"sort"`
	Offset    int     `query:"page"`
	Limit     int     `query:"per_page"`
	StartDate *string `query:"start_date"`
	EndDate   *string `query:"end_date"`
	Search    *string `query:"search"`
}

type ResponsePagination struct {
	Offset    int         `json:"page"`
	Limit     int         `json:"per_page"`
	TotalPage int         `json:"total_page"`
	TotalData int         `json:"total_data"`
	Items     interface{} `json:"items"`
}

func New(page, perPage int) *Pages {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	if page < 1 {
		page = 1
	}

	return &Pages{
		Pagination: Pagination{
			Page:    page,
			PerPage: perPage,
		},
	}
}

// SetData sets list of the data and set count of page and data
func (p *Pages) SetData(data interface{}, total int) {
	pageCount := -1
	if total >= 0 {
		pageCount = (total + p.PerPage - 1) / p.PerPage
	}

	p.PageCount = pageCount
	p.TotalCount = total
	p.Items = data
}
