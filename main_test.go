package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	r := gin.Default()
	setupGetAlbumsRouter(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)

	r.ServeHTTP(w, req)

	json := `[
		{
			"id": "1",
			"title": "Blue Train",
			"artist": "John Coltrane",
			"price": 56.99
		},
		{
			"id": "2",
			"title": "Jeru",
			"artist": "Gerry Mulligan",
			"price": 17.99
		},
		{
			"id": "3",
			"title": "Sarah Vaughan and Clifford Brown",
			"artist": "Sarah Vaughan",
			"price": 39.99
		}
	]`

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, json, w.Body.String())
}
