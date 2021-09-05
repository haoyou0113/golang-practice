package main

import (
	"fmt"
	"net/http"

	"github.com/haoyou0113/myapp/pkg/handlers"
)

const port = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
