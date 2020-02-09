package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Test5")

}


func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8666", nil )
}


