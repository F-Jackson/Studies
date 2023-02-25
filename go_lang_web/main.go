package main

import (
	"net/http"

	"go_lang_web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
