package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bockbone/podtask/driver"
	"github.com/bockbone/podtask/models"
	"github.com/bockbone/podtask/utils"
	_ "github.com/lib/pq"
)

type response struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

//GET all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := getAllUsers()

	if err != nil {
		log.Fatalf("Failed to fetch all user. %v", err)
	}

	//Parse the data into JSON
	json.NewEncoder(w).Encode(users)
}

//POST - Create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Failed to decode request body. %v", err)
	}
	user.Password = utils.HashPassword(user.Password)

	insertID := createUser(user)

	res := response{
		ID:      insertID,
		Message: "User created successfully!",
	}

	// //Parse the data into JSON
	json.NewEncoder(w).Encode(res)

}

func createUser(user models.User) string {
	db, err := driver.ConnectDB()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `INSERT INTO user_account (id,first_name,last_name,email,password, created_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	var id string

	err = db.QueryRow(query, utils.GenerateId(), user.FirstName, user.LastName, user.Email, user.Password, time.Now()).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func getAllUsers() ([]models.User, error) {
	//Connect to database
	db, err := driver.ConnectDB()

	if err != nil {
		panic(err)
	}

	//Close the DB connection
	defer db.Close()

	//Slices of users
	var users []models.User

	//SQL statement
	query := `SELECT id, first_name, last_name, email, password, created_at FROM user_account`

	//Execute the query
	rows, err := db.Query(query)

	if err != nil {
		log.Fatalf("Failed to execute sql. %v", err)
	}

	//Close the query statement
	defer rows.Close()

	//Iterate through each rows
	for rows.Next() {
		//Use the User model
		var user models.User

		//Unmarshal row into user
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)

		if err != nil {
			log.Fatalf("Failed to scan the row. %v", err)
		}

		//Append into users slices
		users = append(users, user)
	}

	return users, err

}
