package main

import (

	//"database/sql"
	"log"
	"strconv"

	//"strconv"

	"net/http"
	// "os"
	// "strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// 		}


func getList(c *gin.Context) {
    tasks, err := GetTasks()
    checkErr(err)
    if tasks == nil{
        c.IndentedJSON(http.StatusBadRequest,gin.H{"error":"Empty List"})

    } else {
        c.IndentedJSON(http.StatusOK,gin.H{"list": tasks})

    }
}
func addtask(c *gin.Context) {

	var json Todo

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := Addtask(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func updatetask(c *gin.Context) {

	var json Todo

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := Updatetask(json, taskId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func deletetask(c *gin.Context) {

	taskId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := Deletetask(taskId)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}


func main(){
    // database, _ := sql.Open("sqlite3", "./todo.db")
    // statement, _:= database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY,Title TEXT,Done BOOL,Body TEXT)")
    // statement.Exec()
    err := ConnectDatabase()
     checkErr(err)


	router := gin.Default()
////      v1 := router.Group("/api/v1")
//    {
//      v1.GET("list", getList)
//       v1.POST("list", addtask)
//       v1.PATCH("list/:id", updatetask)
//       v1.DELETE("list/:id", deletetask)

//      }
      router.GET("list", getList)
      router.POST("list", addtask)
      router.PUT("list/:id", updatetask)
      //router.DELETE("list/:id", deletetask)
router.Run("localhost:4000")

}


func checkErr(err error){
    if err != nil{
        log.Fatal(err)
    }
}