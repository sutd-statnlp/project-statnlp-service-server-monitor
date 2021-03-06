package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../resource"
	"github.com/stretchr/testify/assert"
)

func TestGetHostInfo(t *testing.T) {
	router := SetUpRouter()

	resource.InitHostRoutes(router)

	// Get host info
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/host/info", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}

func TestGetHostTemperture(t *testing.T) {
	router := SetUpRouter()

	resource.InitHostRoutes(router)

	// Get host temperature
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/host/temperature", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.True(t, true, w.Body.Len() > 0)

}
