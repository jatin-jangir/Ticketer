package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host        = "localhost"
	port        = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_NAME     = "postgres"
)

// DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

type Releases struct {
	ID                      string    `json:"movieid"`
	Features_in_Release     string    `json:"Features_in_Release"`
	PR_link                 string    `json:"PR_link"`
	Release_Date            time.Time `json:"Release_Date"`
	Release_Cluster         string    `json:"Release_Cluster"`
	Micro_Services          string    `json:"Micro_Services"`
	Major_Release           bool      `json:"Major_Release"`
	Migration_Required      bool      `json:"Migration_Required"`
	Documentation_Completed bool      `json:"Documentation_Completed"`
	Migration_hash          string    `json:"Migration_hash"`
	Author                  string    `json:"Author"`
	comment                 string    `json:"comment"`
}

type JsonResponse struct {
	Type    string     `json:"type"`
	Data    []Releases `json:"data"`
	Message string     `json:"message"`
}

// Main function
func main() {

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints

	// Get all releases
	router.HandleFunc("/releases/", GetReleases).Methods("GET")
	router.HandleFunc("/releases/sortbydate/", GetSortReleases).Methods("GET")
	// Create a releases
	router.HandleFunc("/releases/", CreateRelease).Methods("POST")

	// Delete a specific releases by the movieID
	router.HandleFunc("/releases/{releasesid}", DeleteMovie).Methods("DELETE")

	// Delete all releases
	router.HandleFunc("/releases/", DeleteMovies).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Get all movies

// response and request handlers
func GetReleases(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting movies...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM releases")
	// check errors
	checkErr(err)

	// var response []JsonResponse
	var releases []Releases

	// Foreach movie
	for rows.Next() {
		var ID string
		var Features_in_Release string
		var PR_link string
		var Release_Date time.Time
		var Release_Cluster string
		var Micro_Services string
		var Major_Release bool
		var Migration_Required bool
		var Documentation_Completed bool
		var Migration_hash string
		var Author string
		var comment string

		err = rows.Scan(&Features_in_Release, &PR_link, &Release_Date, &Release_Cluster, &Micro_Services,
			&Major_Release, &Migration_Required, &Documentation_Completed, &Migration_hash, &Author, &ID, &comment)
		fmt.Println("Release_Date")
		fmt.Println(Release_Date)
		fmt.Println("Release_Date type is -- ")
		fmt.Println(reflect.ValueOf(Release_Date).Kind())
		// check errors
		checkErr(err)
		releases = append(releases, Releases{ID: ID, Features_in_Release: Features_in_Release, PR_link: PR_link, Release_Date: Release_Date, Release_Cluster: Release_Cluster, Micro_Services: Micro_Services, Major_Release: Major_Release,
			Migration_Required: Migration_Required, Documentation_Completed: Documentation_Completed, Migration_hash: Migration_hash, Author: Author, comment: comment})
	}
	fmt.Println("data we send is ---- ")
	fmt.Println(releases)
	var response = JsonResponse{Type: "success", Data: releases}

	json.NewEncoder(w).Encode(response)
}
func GetSortReleases(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting movies...")

	// Get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM releases ORDER BY Release_Date DESC;")
	// check errors
	checkErr(err)

	// var response []JsonResponse
	var releases []Releases

	// Foreach movie
	for rows.Next() {
		var ID string
		var Features_in_Release string
		var PR_link string
		var Release_Date time.Time
		var Release_Cluster string
		var Micro_Services string
		var Major_Release bool
		var Migration_Required bool
		var Documentation_Completed bool
		var Migration_hash string
		var Author string
		var comment string

		err = rows.Scan(&Features_in_Release, &PR_link, &Release_Date, &Release_Cluster, &Micro_Services,
			&Major_Release, &Migration_Required, &Documentation_Completed, &Migration_hash, &Author, &ID, &comment)
		fmt.Println("Release_Date")
		fmt.Println(Release_Date)
		fmt.Println("Release_Date type is -- ")
		fmt.Println(reflect.ValueOf(Release_Date).Kind())
		// check errors
		checkErr(err)
		releases = append(releases, Releases{ID: ID, Features_in_Release: Features_in_Release, PR_link: PR_link, Release_Date: Release_Date, Release_Cluster: Release_Cluster, Micro_Services: Micro_Services, Major_Release: Major_Release,
			Migration_Required: Migration_Required, Documentation_Completed: Documentation_Completed, Migration_hash: Migration_hash, Author: Author, comment: comment})
	}
	fmt.Println("data we send is ---- ")
	fmt.Println(releases)
	var response = JsonResponse{Type: "success", Data: releases}

	json.NewEncoder(w).Encode(response)
}

// Create a movie

// CREATE TABLE releases (
//
//			Features_in_Release varchar(50) NOT NULL,
//			PR_link varchar(100) NOT NULL,
//			Release_Date DATE NOT NULL,
//			Release_Cluster  varchar(50) NOT NULL,
//			Micro_Services varchar[],
//			Major_Release BOOLEAN NOT NULL DEFAULT false,
//			Migration_Required BOOLEAN NOT NULL DEFAULT false,
//			Documentation_Completed BOOLEAN NOT NULL DEFAULT false,
//			Migration_hash varchar[],
//			Author varchar(50) NOT NULL,
//			id SERIAL,
//			comment varchar(200),
//	    PRIMARY KEY (id)
//
// );
// response and request handlers
func CreateRelease(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("ID")
	Features_in_Release := r.FormValue("Features_in_Release")
	PR_link := r.FormValue("PR_link")
	Release_Date := r.FormValue("Release_Date")
	Release_Cluster := r.FormValue("Release_Cluster")
	Micro_Services := r.FormValue("Micro_Services")
	Major_Release := r.FormValue("Major_Release")
	Migration_Required := r.FormValue("Migration_Required")
	Documentation_Completed := r.FormValue("Documentation_Completed")
	Migration_hash := r.FormValue("Migration_hash")
	Author := r.FormValue("Author")
	comment := r.FormValue("comment")
	var response = JsonResponse{}

	if ID == "" || Features_in_Release == "" || PR_link == "" || Author == "" || Release_Cluster == "" {
		response = JsonResponse{Type: "error", Message: "You are missing movieID or movieName parameter."}
	} else {
		db := setupDB()

		printMessage("Inserting movie into DB")

		fmt.Println("Inserting new movie with ID: " + ID + " and name: " + Features_in_Release)

		var lastInsertID int
		err := db.QueryRow("INSERT INTO movies(features_in_release , pr_link , release_date , release_cluster , micro_services , major_release , migration_required , documentation_completed , migration_hash , author , id , comment ) VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning id;", Features_in_Release, PR_link, Release_Date, Release_Cluster, Micro_Services, Major_Release, Migration_Required, Documentation_Completed, Migration_hash, Author, ID, comment).Scan(&lastInsertID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The release has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// response and request handlers
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	movieID := params["movieid"]

	var response = JsonResponse{}

	if movieID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing movieID parameter."}
	} else {
		db := setupDB()

		printMessage("Deleting movie from DB")

		_, err := db.Exec("DELETE FROM movies where movieID = $1", movieID)

		// check errors
		checkErr(err)

		response = JsonResponse{Type: "success", Message: "The movie has been deleted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all movies

// response and request handlers
func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Deleting all movies...")

	_, err := db.Exec("DELETE FROM movies")

	// check errors
	checkErr(err)

	printMessage("All movies have been deleted successfully!")

	var response = JsonResponse{Type: "success", Message: "All movies have been deleted successfully!"}

	json.NewEncoder(w).Encode(response)
}
