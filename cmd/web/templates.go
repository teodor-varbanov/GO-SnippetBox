package main

import (
	"github.com/teodor-varbanov/snippetBox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
