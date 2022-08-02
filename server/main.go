package main

import (

	//"database/sql"
	"log"
	"strconv"

	//"strconv"
	//"server/middleware"
	"net/http"
	// "os"
	// "strconv"
	//"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gin-contrib/cors"
	"github.com/zeinaelkassas/todolist/models"
	//  modell "todolist/models"
	//"models.database"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zeinaelkassas/todolist/middleware"
	//"github.com/steinfletcher/apitest/x/db"
)

// func  NewDatabase() *database{

// 	d := &database{}
//     db,_:=models.ConnectDatabase(d)
// 	return db
// }
type app struct{
	db models.Database
}


func (App *app) getList(c *gin.Context) {
	App.db=*models.NewDatabase()
    tasks, err := App.db.GetTasks()
    checkErr(err)
    if tasks == nil{
        c.IndentedJSON(http.StatusBadRequest,gin.H{"error":"Empty List"})

    } else {
        c.IndentedJSON(http.StatusOK,gin.H{"list": tasks})

    }
}
func (App *app) addtask(c *gin.Context) {
	App.db=*models.NewDatabase()

	var json models.Todo

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := App.db.Addtask(json)

	if success {
		c.JSON(http.StatusCreated, gin.H{"message": "added successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func (App *app) deletetask(c *gin.Context) {
	App.db=*models.NewDatabase()

	taskId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := App.db.Deletetask(taskId)

	if success {
		c.JSON(http.StatusNoContent, gin.H{"message": "Deleted successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func(App *app) updatetask(c *gin.Context) {
	App.db=*models.NewDatabase()

	var json models.Todo

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := App.db.Updatetask(json, taskId)

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
    //err := ConnectDatabase()
     //checkErr(err)
	 //r := gin.New()
	//  r.Use(middleware.UploadRetrievalLog())
 
	router := gin.Default()
   App:= app{}
    
	router.Use(middleware.Log())
	// println(middleware.UploadRetrievalLog())
    router.Use(cors.New(cors.Config{
	 	AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:  []string{"content-type"},
		ExposeHeaders: []string{"X-Total-Count"},
	  }))

      v1 := router.Group("/api/v1")
   {
	
     v1.GET("list", App.getList)
      v1.POST("list", App.addtask)
      v1.DELETE("list/:id", App.deletetask)
	  v1.PUT("list/:id", App.updatetask)
	  v1.PATCH("list/:id", App.updatetask)

     }
    //   router.GET("list", getList)
    //   router.POST("list", addtask)
    //   router.DELETE("list/:id", deletetask)
        router.Run("localhost:4000")

}


func checkErr(err error){
    if err != nil{
        log.Fatal(err)
    }
}