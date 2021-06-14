package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type jedi struct {
	ID        int
	Name      string
	Email     string
	CreatedAt string
}

type jediCollection []jedi

var jedis = jediCollection{
	{
		ID:        1,
		Name:      "Master Yoda",
		Email:     "jedi.yoda.master@jediclub.com",
		CreatedAt: "10.06.2021",
	},
	{
		ID:        2,
		Name:      "Obi-Wan Kenobi",
		Email:     "obi-wan@jediclub.com",
		CreatedAt: "10.06.2021",
	},
}

type defaultValues struct {
	Name string
	Email string
}

var defaultForm defaultValues

var errors = make(map[string]string)

type response struct {
	Data jediCollection
	Errors map[string]string
	Default defaultValues
}

func createJedi(w http.ResponseWriter, r *http.Request) {
	defaultForm.Name = ""
	defaultForm.Email = ""

	var newJedi jedi

	r.ParseForm()

	newJedi.Name = r.FormValue("name")
	newJedi.Email = r.FormValue("email")

	isValid := validate(newJedi)

	if true == isValid {
		newJedi.ID = len(jedis) + 1
		today := time.Now()
		newJedi.CreatedAt = fmt.Sprintf("%02d.%02d.%d", today.Day(), today.Month(), today.Year())
		jedis = append(jedis, newJedi)
	} else {
		defaultForm.Name = newJedi.Name
		defaultForm.Email = newJedi.Email
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func getAllJedis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jedis)
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	var resp response
	resp.Data = jedis
	resp.Errors = errors
	resp.Default = defaultForm

	t.Execute(w, resp)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/jedis", createJedi).Methods(http.MethodPost)
	router.HandleFunc("/jedis", getAllJedis).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8888", router))
}
