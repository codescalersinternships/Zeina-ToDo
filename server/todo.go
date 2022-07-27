package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
	Body string `json:"body"`}

func GetTasks() ([]Todo, error) {

				rows, err := DB.Query("SELECT id,title,done,body from tasks" )
			
				if err != nil {
					return nil, err
				}
			
				defer rows.Close()
			
				tasks := make([]Todo, 0)
			
				for rows.Next() {
					singletask := Todo{}
				 err = rows.Scan(&singletask.ID, &singletask.Title, &singletask.Done, &singletask.Body)
			
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

func Addtask(newtask Todo) (bool, error) {

				tx, err := DB.Begin()
				if err != nil {
					return false, err
				}
			
				stmt, err := tx.Prepare("INSERT INTO tasks (title,done, body) VALUES (?, ?, ?)")
			
				if err != nil {
					return false, err
				}
			
				defer stmt.Close()
			
				_, err = stmt.Exec(newtask.Title, newtask.Done, newtask.Body)
			
				if err != nil {
					return false, err
				}
			
				tx.Commit()
			
				return true, nil
			}
			
			
			
func Deletetask(taskid int) (bool, error) {

				tx, err := DB.Begin()
			
				if err != nil {
					return false, err
				}
			
				stmt, err := DB.Prepare("DELETE from tasks where id = ?")
			
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