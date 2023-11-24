package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/nato/", natoHandler) // Define a variable route

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Go Web Server!")
}

var natoMap = map[string]string{
    "a": "Alpha", "b": "Bravo", "c": "Charlie", "d": "Delta",
    "e": "Echo", "f": "Foxtrot", "g": "Golf", "h": "Hotel",
    "i": "India", "j": "Juliett", "k": "Kilo", "l": "Lima",
    "m": "Mike", "n": "November", "o": "Oscar", "p": "Papa",
    "q": "Quebec", "r": "Romeo", "s": "Sierra", "t": "Tango",
    "u": "Uniform", "v": "Victor", "w": "Whiskey", "x": "X-ray",
    "y": "Yankee", "z": "Zulu",
}

func natoHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the letter from the URL path
	letterPath := strings.ToLower(r.URL.Path[len("/nato/"):])
	if len(letterPath) != 1 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	natoWord, exists := natoMap[letterPath]
	if !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, natoWord)
}
