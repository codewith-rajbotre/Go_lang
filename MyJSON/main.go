package main

import (
	"encoding/json"
	"fmt"
)

type collectDataOfCourse struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` //If you will provide a - in the field it will remove the field fron output
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello, Let us learn about JSON")
	//EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	courses := []collectDataOfCourse{
		{"ReactJSBootcamp", 299, "Udemy", "Paass@12", []string{"WebDev", "JS"}},
		{"Mobile Development", 200, "Udemy", "abc@19", []string{"Java", "MAD"}},
		{"Java Development", 1000, "W3Schools", "web@123", nil},
	}
	//Package this data as a JSON data
	fileJSON, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileJSON))

}
func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	  {
                "coursename": "ReactJSBootcamp",
                "Price": 299,
                "website": "Udemy",
                "tags": ["WebDev","JS"]
        }
	`)
	var courses collectDataOfCourse
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is Valid")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v\n", courses)
	} else {
		fmt.Println("JSON is not Valid")
	}
	//If we want to add a key value pair in your JSON data use the following code for it
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and the value is %v and Type is %T \n ", k, v, v)
	}
}
