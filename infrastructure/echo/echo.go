package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/AJRDRGZ/hexagonal-architecture/domain/poetry"
)

// API allow print the poems from API rest
type API struct {
	echo         *echo.Echo
	poetryReader poetry.IRequestVerses
}

// New returns a new Echo
func New(poetry poetry.IRequestVerses) API {
	api := API{
		poetryReader: poetry,
		echo:         echo.New(),
	}

	api.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, api.poetryReader.GiveMeSomePoetry())
	})

	return api
}

// Start start of API server
func (a API) Start(port string) {
	a.echo.Logger.Fatal(a.echo.Start(port))
}
