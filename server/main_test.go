package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zeinaelkassas/todolist/models"
)
func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router 
}
func TestGetList(t *testing.T) {

    r:= SetUpRouter()
    App := app{}

    r.GET("/v1/list",App.getList)
    req, _ := http.NewRequest("GET", "/v1/list", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusOK, w.Code)
   //assert.NotEmpty(t, DB)
}
func TestAddtask(t *testing.T) {
	// ConnectDatabase()
    r := SetUpRouter()
    App := app{}

    r.POST("/v1/list", App.addtask)
    task := models.Todo{
		Title: "task4",
		Done: true,


	}
   jsonValue, _ := json.Marshal(task)
    req, _ := http.NewRequest(http.MethodPost, "/v1/list", bytes.NewBuffer(jsonValue))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
}

