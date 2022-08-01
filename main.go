package main

import (
	"context"
	"fmt"
	"gofiberwork/db"
	"html/template"
	"log"
	"net/http"
)

type itms struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phoneno string `json:"contact"`
	Status  string `json:"status"`
}

//inserting data to database function.
func Toinsertdata(todo itms) (itms, error) {
	res, err := db.DB.Collection("uservalues").InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
		return itms{}, err
	}
	fmt.Println(res.InsertedID)
	return todo, nil
}

//main.
func main() {
	tmpl := template.Must(template.ParseFiles("form.html")) //calling html template.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		formitems := itms{
			Name:    r.FormValue("name"),
			Email:   r.FormValue("email"),
			Phoneno: r.FormValue("number"),
			Status:  r.FormValue("status"),
		}
		_ = formitems //just to avoid error

		tmpl.Execute(w, struct{ success bool }{true})

		fmt.Println("\nUserName: " + formitems.Name)
		fmt.Println("Email: " + formitems.Email)
		fmt.Println("Mobile-NO: " + formitems.Phoneno)
		fmt.Println("Current Working Status: " + formitems.Status)
		Toinsertdata(formitems) //entering user's values to the database

	})
	http.ListenAndServe(":2411", nil)

}
