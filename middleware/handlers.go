package middleware

import (
	"database/sql"
	"fmt"
	// "encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"html/template"

	"todo_app/models"
	
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_"github.com/lib/pq"
)

var temp = template.Must(template.ParseFiles("./template/index.html"))

// connect to postgre DB
func connectDB() *sql.DB {
	// load env file
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	// open connection
	db, dbErr := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if dbErr != nil {
		panic(dbErr)
	}

	// check connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database successfully")
	return db
}

// function to insert todo
func insert(task string) int64 {
	db := connectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO todos (description) VALUES ($1) RETURNING todoid`

	var id int64

	err := db.QueryRow(sqlStatement, task).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to query, %", err)
	}

	fmt.Printf("Inserted a todo into the table, %v", id)

	return id
}

// function to get todo
// func get(id int64) (models.Todo, error) {
// 	db := connectDB()
// 	defer db.Close()

// 	sqlStatement := `SELECT * FROM todos WHERE todoid=$1`

// 	var todo models.Todo

// 	err := db.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.Desc)
// 	switch err {
//     case sql.ErrNoRows:
//         fmt.Printf("No rows were returned!")
//         return todo, nil
//     case nil:
//         return todo, nil
//     default:
//         log.Fatalf("Unable to query, %v", err)
//     }

//     return todo, err
// }

// function to get all todos
func getAll() ([]models.Todo, error) {
	db := connectDB()
	defer db.Close()

	var todos []models.Todo

	sqlStatement := `SELECT * FROM todos`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to query, %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.ID, &todo.Desc)
		if err != nil {
			log.Fatalf("Unable to query, %v", err)
		}

		todos = append(todos, todo)
	}
	return todos, err
}

// function to update todo
func update(id int64, task string) int64 {
	db := connectDB()
	defer db.Close()

	sqlStatement := `UPDATE todos SET description=$2 WHERE todoid=$1`

	res, err := db.Exec(sqlStatement, id, task)
	if err != nil {
		log.Fatalf("Unable to query, %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error getting affected rows, %v", err)
	}

	fmt.Printf("%v rows updated", rowsAffected)

	return rowsAffected
}

// function to delete todo
func delete(id int64) int64{
	db := connectDB()
	defer db.Close()

	sqlStatement := `DELETE FROM todos WHERE todoid=$1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to query, %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error getting affected rows, %v", err)
	}

	fmt.Printf("%v rows deleted", rowsAffected)

	return rowsAffected
}

// create todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("newtodo")

	id := insert(task)
	
	fmt.Printf("Successfully added a task, %v", id)

	http.Redirect(w, r, "/", 301)
}

// get todo
// func GetTodo(w http.ResponseWriter, r *http.Request) {
// 	// w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	// w.Header().Set("Access-Control-Allow-Origin", "*")
// 	// w.Header().Set("Access-Control-Allow-Methods", "GET")
// 	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Fatalf("Unable to convert string into int")
// 	}

// 	todo, err := get(int64(id))
// 	if err != nil {
// 		log.Fatalf("Unable to get user, %v", err)
// 	}

// 	json.NewEncoder(w).Encode(todo)
// }

// get all todos
func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := getAll()
	if err != nil {
		log.Fatalf("Unable to get all users, %v", err)
	}

	list := models.TodoList {
		Todos: todos,
	}
	_ = temp.Execute(w, list)
}

// update todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, intErr := strconv.Atoi(params["id"])
	if intErr != nil {
		log.Fatalf("Unable to convert string into int")
	}
	task := r.FormValue("edittask")
	fmt.Printf("see task, %v", task)

	updatedRows := update(int64(id), task)

	fmt.Printf("Successfully edited task, %v", updatedRows)

	http.Redirect(w, r, "/", 301)
}

// delete todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert string into int")
	}

	deletedRows := delete(int64(id))
	fmt.Printf("Amount of task/s deleted, %v", deletedRows)

	http.Redirect(w, r, "/", 301)
}