package web

import (
	"github.com/demisto/download/repo"
)

// AppContext holds the web context for the handlers
type AppContext struct {
	r *repo.Repo
}

// NewContext creates a new context
func NewContext(r *repo.Repo) *AppContext {
	ac := &AppContext{r: r}
	return ac
}

type session struct {
	User string `json:"user"`
	When int64  `json:"when"`
}
