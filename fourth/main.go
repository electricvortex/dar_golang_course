package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Asd struct {
	 Name string `json:"name"`
}

func main() {
	k := Asd{Name: "Four"}
	err := Total(k)
	if err == nil {
		k := Asd{Name:"world"}
		fmt.Println(k)
	}
	fmt.Println(k)
}

func Total(a Asd) error {
	a.Name = "hello"
	return nil
}

func PrintFilesInDir() {
	var a string
	_, err := fmt.Scanln(&a)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(a)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func JSON() {
	type None struct {
		Message *string `json:"message"`
		//Fields map[string]interface{} `json:"fields"`
		Fields struct{
			Name string `json:"name"`
			Surname int `json:"surname,omitempty"`
		} `json:"fields"`
	}

	a := None{}
	f, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	file, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	// Unmarshaling given json
	err = json.Unmarshal(file, &a)
	if err != nil {
		panic(err)
	}

	// From JSON to struct
	fmt.Println(a)

	l := "Go course"

	b := &None{
		Message: &l,
		Fields: struct {
			Name    string `json:"name"`
			Surname int `json:"surname,omitempty"`
		}{
			Name: "World",
		},
	}
	n, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fo, err := os.Create("output.json")
	if err != nil {
		panic(err)
	}
	_, err = fo.Write(n)
	if err != nil {
		panic(err)
	}
	err = fo.Close()
	if err != nil {
		panic(err)
	}
}