package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CoursrId    string  `json: "courseid"`
	CourseName  string  `json :"coursename"`
	CoursePrice int     `json:"price"`
	Auther      *Auther `json:"author"`
}
type Auther struct {
	FullName string `json:"fullname"`
	Website  string `json :"website"`
}

// fake DB
var courses []Course

// middleware,Helper-file
func (c *Course) IsEmpty() bool {
	return c.CoursrId == "" && c.CourseName == ""
}
func main() {
	fmt.Println("Hello We will create API in this Folder: ")

}

//Controllers-file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welocom to API by RouteController</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}
