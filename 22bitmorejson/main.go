package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string //`json:"coursename"` this third perameter will print new name insteade of Name in json data
	Prise     int
	Plateform string
	Password  string   `json:"-"`              //this third peramether will not show password in output
	Tags      []string `json:"tags,omitempty"` // this perameter will use if any data have Tag value nil then it will not be printed in output.
}

func main() {
	fmt.Println("json in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJs Bootcamp", 200, "abcd@learn.dev", "123efb", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "abcd@learn.dev", "1f3efb", []string{"fullstack", "js"}},
		{"GOBootcamp", 400, "abcd@learn.dev", "123eeb", nil},
	}

	//package this data as json data

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
		{
                "Name": "ReactJs Bootcamp",
                "Prise": 200,
                "Plateform": "abcd@learn.dev",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	var course course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("json was valid")
		json.Unmarshal(jsonData, &course)
		fmt.Printf("%#v\n",course)
	} else {
		fmt.Println("json was no valid")
	}

	//some cases where you want to add some data to key value
	var myData map[string]interface{}
	json.Unmarshal(jsonData,&myData)
	fmt.Printf("%#v\n",myData)

	for k, v := range myData {
		fmt.Printf("key is %v and value is %v and type is %T\n",k,v,v)
	}
}
