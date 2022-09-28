package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"regexp"
	"time"
)

type Config struct {
	host string
	port int
}

type Response struct {
	Response string `json:"response,omitempty"`
	Error    string `json:"error,omitempty"`
}

type User struct {
	name, gender, pronouns string
	age                    int
}

type Data struct {
	users []User
}

func (d *Data) findUserByName(name string) []User {
	list := make([]User, 0)
	for _, e := range d.users {
		time.Sleep(time.Millisecond * 100)
		if e.name == name {
			list = append(list, e)
		}
	}
	return list
}

func (d *Data) findUserByGender(gender string) []User {
	list := make([]User, 0)
	for _, e := range d.users {
		time.Sleep(time.Millisecond * 100)
		if e.gender == gender {
			list = append(list, e)
		}
	}
	return list
}

func (d *Data) findUserYoungerThan(age int) []User {
	list := make([]User, 0)
	for _, e := range d.users {
		time.Sleep(time.Millisecond * 100)
		if e.age < age {
			list = append(list, e)
		}
	}
	return list
}

func (d *Data) findUserOlderThan(age int) []User {
	list := make([]User, 0)
	for _, e := range d.users {
		time.Sleep(time.Millisecond * 100)
		if e.age < age {
			list = append(list, e)
		}
	}
	return list
}

func (d *Data) addUser(e User) {
	d.users = append(d.users, e)
}

func handleGetUserById() {}

func handleUser(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		idRegex, _ := regexp.Compile("/user/id/([0-9]+)")
		userId := idRegex.FindStringSubmatch(req.URL.Path)
		if userId != nil {

		}
		nameRegex, _ := regexp.Compile("/user/name/(.+)")
		userName := nameRegex.FindStringSubmatch(req.URL.Path)
		if userName != nil {

		}
	case "POST":
	case "DELETE":
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	p := Response{
		Response: "Hello, world!",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(p)
}

func config() Config {
	host := flag.String("host", "", "the host to listen on")
	port := flag.Int("port", 8080, "the port to listen on")

	config := Config{
		host: *host,
		port: *port,
	}
	return config
}

func main() {
	config := config()
	listenStr := fmt.Sprintf("%s:%d", config.host, config.port)
	http.HandleFunc("/user", hello)
	http.HandleFunc("/user/", hello)

	http.ListenAndServe(listenStr, nil)
}
