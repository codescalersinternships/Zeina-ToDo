package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type database struct {
	db *sql.DB
    
}

// func NewDatabase() *database{
// 	d := &database{}
// //	db,_:=ConnectDatabase(d)
// 	return d
// }
func  NewDatabase() *database{
	
	d := &database{}
    db,_:=ConnectDatabase(d)
	return db
}


// d := NewDatabase()
// d.GetTasks()
// func (d *Database) GetTodos() ([]Todo, err) {

// }

// var DB *sql.DB
func ConnectDatabase(database *database)  (*database,error) {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		    err= err
	}
	database.db=db
	return database,nil
}
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	}

func (d *database) GetTasks() ([]Todo, error) {

				rows, err := d.db.Query("SELECT id,title,done from tasks" )
			
				if err != nil {
					return nil, err
				}
			
				defer rows.Close()
			
				tasks := make([]Todo, 0)
			
				for rows.Next() {
					singletask := Todo{}
				 err = rows.Scan(&singletask.ID, &singletask.Title, &singletask.Done)
			
					if err != nil {
						return nil, err
					}
			
					tasks = append(tasks, singletask)
				}
			
				err = rows.Err()
			
				if err != nil {
					return nil, err
				}
			
				return tasks, err
			}

func (d *database)  Addtask(newtask Todo) (bool, error) {

				tx, err := d.db.Begin()
				if err != nil {
					return false, err
				}
			
				stmt, err := tx.Prepare("INSERT INTO tasks (title,done) VALUES (?, ?)")
			
				if err != nil {
					return false, err
				}
			
				defer stmt.Close()
			
				_, err = stmt.Exec(newtask.Title, newtask.Done)
			
				if err != nil {
					return false, err
				}
			
				tx.Commit()
			
				return true, nil
			}
			
			
			
func(d *database)  Deletetask(taskid int) (bool, error) {

				tx, err := d.db.Begin()
			
				if err != nil {
					return false, err
				}
			
				stmt, err := d.db.Prepare("DELETE from tasks where id = ?")
			
				if err != nil {
					return false, err
				}
			
				defer stmt.Close()
			
				_, err = stmt.Exec(taskid)
			
				if err != nil {
					return false, err
				}
			
				tx.Commit()
			
				return true, nil
			}


func (d *database) Updatetask(task Todo, taskId int) (bool, error) {

				tx, err := d.db.Begin()
				if err != nil {
					return false, err
				}
			
				stmt, err := tx.Prepare("UPDATE tasks SET  done = ? where id = ?")
			
				if err != nil {
					return false, err
				}
			
				defer stmt.Close()
			
				_, err = stmt.Exec(task.Done,taskId)
			
				if err != nil {
					return false, err
				}
			
				tx.Commit()
			
				return true, nil
			}