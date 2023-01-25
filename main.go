package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `josn:"website"`
}

func (c *Course) Isempty() bool {
	return c.CourseName == ""
}

var courses []Course

func main() {
	fmt.Println("API-Golang")
	r := mux.NewRouter()
	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 299, Author: &Author{FullName: "xxx", Website: "xxx.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "javascript", CoursePrice: 199, Author: &Author{FullName: "yyy", Website: "yyy.com"}})

	//routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	//Listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers to handle requests

func ServeHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page")
	w.Write([]byte("<h1>welcome to golang coding world</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("no data found for this ID")

}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("please enter data")
	}
	var Course Course
	json.NewDecoder(r.Body).Decode(&Course)

	if Course.Isempty() {
		json.NewEncoder(w).Encode("no data inside json")
		return
	}
	//generate unique course id for course
	rand.Seed(time.Now().UnixNano())
	Course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, Course)
	json.NewEncoder(w).Encode(Course)
}

func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no id is found")
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one resource")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("successfully deleted")
			break
		}
	}
}
