package main

import ("fmt"
        "net/http")

func main() {
	
	http.HandleFunc("/", StartPage)
	http.HandleFunc("/ascii-art", SubmitTing)
	http.HandleFunc("/down",download)
	fmt.Println("http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}