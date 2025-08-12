package shared

import (
	"articles/logger"
	"database/sql"
	"net/http"
	"strconv"

	"go.uber.org/dig"

	"articles/shared/config"
	"articles/shared/pagination"
)

type (
	Deps struct {
		dig.In

		DB     *sql.DB
		Config *config.Configuration
		Logger logger.Logger

		CustomValidator          *CustomValidator
		NewPaginationFromRequest func(r *http.Request) *pagination.Pages
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewCustomValidator); err != nil {
		return err
	}

	if err := container.Provide(func() func(r *http.Request) *pagination.Pages {
		return func(r *http.Request) *pagination.Pages {
			// Example: get from query params
			page, _ := strconv.Atoi(r.URL.Query().Get("page"))
			if page <= 0 {
				page = 1
			}

			pageSize, _ := strconv.Atoi(r.URL.Query().Get("per_page"))
			if pageSize <= 0 {
				pageSize = pagination.DefaultPageSize
			}

			return pagination.New(page, pageSize)
		}
	}); err != nil {
		return err
	}

	return nil
}
