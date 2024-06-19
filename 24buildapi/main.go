package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model for course-file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrize int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var Courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId=="" && c.CourseName==""
	return c.CourseName == ""
}

func main() {
	fmt.Println("building api in golang")

	//creating a router
	r := mux.NewRouter()

	//seeding
	Courses = append(Courses, Course{CourseId: "2", CourseName: "golang", CoursePrize: 200, Author: &Author{Fullname: "Gaurang", Website: "advb@hauv.com"}})
	Courses = append(Courses, Course{CourseId: "4", CourseName: "JS", CoursePrize: 344, Author: &Author{Fullname: "Gaurang", Website: "advb@hauv.com"}})
	Courses = append(Courses, Course{CourseId: "5", CourseName: "MERN stack", CoursePrize: 300, Author: &Author{Fullname: "Gaurang", Website: "advb@hauv.com"}})
	Courses = append(Courses, Course{CourseId: "6", CourseName: "Cybersecurity", CoursePrize: 900, Author: &Author{Fullname: "Gaurang", Website: "advb@hauv.com"}})

	//routing
	r.HandleFunc("/", servHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers-file

// serve home route
func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello from gaurang brdv</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Courses) //will send the courses to whome made request
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("content-type", "application/json")

	//grab id from request
	params := mux.Vars(r) // vars provides all the variables also params is a folder of key value pair

	//loop through the courses,find matching id, and return the responce
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//if we dont find anything
	json.NewEncoder(w).Encode("no course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	//in the method we are geting json from frontend so we will use decoding method.
	fmt.Println("create one course")
	w.Header().Set("content-type", "application/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	//what if - data is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside the json")
		return
	}

	//generate the uniqueid
	course.CourseId = strconv.Itoa(rand.Intn(100))
	//push new course into the database
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("content-type", "application/json")

	//first - grab id from request
	params := mux.Vars(r)

	//loop through the value and match the id
	//deletion of that perticuler element in slices
	//addition of updated elementin slices with params id

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...) //removing
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: send responce when id is not found
	json.NewEncoder(w).Encode("id didnt found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			break
		}
	}
}
