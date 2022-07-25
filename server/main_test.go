package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	//"github.com/steinfletcher/apitest"
// 	"github.com/stretchr/testify/assert"
// )


// const task =`{" {
// 				"id": 5,
// 				"title": "taskkk",
// 				"done": true,
// 				"body": "taskk2"
// 			}
// 			}`


// 			// type Todo struct{
// 			// 	ID int `json:"id"`
// 			// 	Title string `json:"title"`
// 			// 	Done bool `json:"done"`
// 			// 	Body string `json:"body"`}
			

// // func TestGetMessage(t *testing.T) {
// // 	// handler := func(w http.ResponseWriter, r *http.Request) {
// // 	// 	msg := task
// // 	// 	_, _ = w.Write([]byte(msg))
// // 	// 	w.WriteHeader(http.StatusOK)
// // 	// }
// // 	r := gin.Default()
// //     r.Handle("/v1/bookmark", getList)
// // 	apitest.New(). // configuration
// // 			HandlerFunc(handler).
// // 			Get("/list"). // request
// // 			Expect(t).       // expectations
// // 			Status(http.StatusOK).
// // 			End()
// // }
// func SetUpRouter() *gin.Engine{
//     router := gin.Default()
//     return router
// }
// func TestGetList(t *testing.T) {
//     r:= SetUpRouter()
//     r.GET("/list", getList)
//     req, _ := http.NewRequest("GET", "/list", nil)
//     w := httptest.NewRecorder()
//     r.ServeHTTP(w, req)

//     var todo [] Todo
//     json.Unmarshal(w.Body.Bytes(), &todo)

//     assert.Equal(t, http.StatusOK, w.Code)
//     assert.NotEmpty(t, todo)
// }
// func TestAddtask(t *testing.T) {
//     r := SetUpRouter()
//     r.POST("/list", addtask)
//     //taskId := id.New().String()
//     task := Todo{
// 		Title: "task4",
// 		Done: true,
// 		Body: "taskk2",
// 	}
	
//     jsonValue, _ := json.Marshal(task)
//     req, _ := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonValue))

//     w := httptest.NewRecorder()
//     r.ServeHTTP(w, req)
//     assert.Equal(t, http.StatusCreated, w.Code)
// }