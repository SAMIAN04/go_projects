package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
  CourseId   string `json:"courseid"`
  CourseName string `json:"coursename"`
  Price      int    `json:"price"`
  Author     *Author `json:"author"`
}
type Author struct{
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}
//fake db
var courses = []Course{
	{CourseId: "2", CourseName: "ReactJS", Price: 299, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}},
	{CourseId: "3", CourseName: "MERN Stack", Price: 199, Author: &Author{Fullname: "Jane Smith", Website: "janesmith.com"}},
}
//middleware, helpers - files
func (c *Course) IsEmpty() bool {
  return c.CourseId == ""  && c.CourseName == ""
}
func main()  {
	
}

//controllers - files

//routes - files
func HomeServer(w http.ResponseWriter ,r *http.Request){
	w.Write([]byte(" <body style='background-color: black;'> <h1 style='color: blue;'>Welcome to the Home Page!</h1> </body>"))
}

//server - files	
func getAllCourses(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//encode the data and send as response
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	//loop through courses and find matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
}
func createOneCourse(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
		
	}
    rand.Seed(time.Now().UnixNano())
	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	//generate unique id , string
	course.CourseId =  string(rand.Intn(100))
	//append new course to courses
	courses = append(courses, course)
	fmt.Printf("New course added: %+v\n", course)
}