package main

import (
	"net/http"
	"strconv"
	"text/template"
)

type Weight struct {
	Date  string
	Value float64
}

type Data struct {
	Weights []Weight
}

func main() {
	router := http.NewServeMux()

	tmpl, err := template.ParseFiles("templates/layout.html")
	if err != nil {
		panic("template error")
	}

	data := Data{
		Weights: []Weight{
			{Date: "07-05-2025", Value: 96.4},
			{Date: "08-05-2025", Value: 92.1},
		},
	}
	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})

	router.HandleFunc("POST /weights", func(w http.ResponseWriter, r *http.Request) {
		weight, err := strconv.ParseFloat(r.FormValue("weight"), 64)
		date := r.FormValue("date")

		if err != nil {
			return
		}

		newItem := Weight{
			Date:  date,
			Value: weight,
		}
		data.Weights = append(data.Weights, newItem)

		tmpl, err := template.ParseFiles("templates/layout.html")
		if err != nil {
			panic("template error")
		}
		tmpl.ExecuteTemplate(w, "weights", data)

	})

	http.ListenAndServe(":8080", router)
}
