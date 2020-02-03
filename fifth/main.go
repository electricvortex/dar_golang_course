package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"net/http"
)

var (
	configFile = ""
	port = "8080"

	flags = []cli.Flag{
		&cli.StringFlag{
			Name: "config",
			Aliases: []string{"c"},
			Usage: "config file path",
			Destination: &configFile,
		},
	}
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}/createfile", CreateFile("name"))
	http.Handle("/", r)
	fmt.Println("Server started")
	http.ListenAndServe(":" + port, r)
}

type Message struct {
	Body string
}

func CreateFile(name string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nm, ok := vars[name]
		if !ok {
			resp, _ := json.Marshal(&Message{Body:"Error"})
			w.Write(resp)
			return
		}
		a := "Answer"
		w.Write([]byte(a))
		fmt.Println(nm)
		b := "result"
		w.Write([]byte(b))
		//out, err := exec.Command("ls", "-a").Output()
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(string(out))
		//jsonAnswer, err := json.Marshal(out)
		//if err != nil {
		//	panic(err)
		//}
		//w.Write(jsonAnswer)
	}
}