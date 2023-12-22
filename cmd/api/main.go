package main

import (

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type product struct {
	Name string

	Price int64
}
func main() {
	route := chi.NewRouter()
	route.Use(myMiddleware)

	// route.Get("/{param}", func (res http.ResponseWriter, req *http.Request)  {
	// 	param := chi.URLParam(req, "param")
	// 	res.Write([]byte(param))
	// })
	route.Get("/", func (res http.ResponseWriter, req *http.Request)  {
		param := req.URL.Query().Get("param")
		res.Write([]byte(param))
	})

	route.Get("/json", func (res http.ResponseWriter, req *http.Request)  {
	
		obj := map[string]string{"message":"success"}

		render.JSON(res, req, obj)

	
	})
	
	route.Post("/teste", func (res http.ResponseWriter, req *http.Request)  {
		var product product

	
		render.DecodeJSON(req.Body, &product)

		render.JSON(res, req, product)

	
	})


	http.ListenAndServe(":3000", route)
}

func myMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		println("before")

		next.ServeHTTP(w,r)

		println("after")


	})
}