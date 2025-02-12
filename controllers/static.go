package controllers

import (
	"html/template"
	"net/http"

	"calhoun-basic/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Abrakadabra?",
			Answer:   "Yes!",
		},
		{
			Question: "Which came first: the chicken or the egg?",
			Answer:   "T-Rex!",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:pivojoppin@gmail.com">pivojoppin@gmail.com`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
