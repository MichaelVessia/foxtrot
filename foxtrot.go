package main

import (
	"fmt"
	"net/http"
	"strings"
)

var natoMap = map[string]string{
	"a": "Alpha", "b": "Bravo", "c": "Charlie", "d": "Delta",
	"e": "Echo", "f": "Foxtrot", "g": "Golf", "h": "Hotel",
	"i": "India", "j": "Juliett", "k": "Kilo", "l": "Lima",
	"m": "Mike", "n": "November", "o": "Oscar", "p": "Papa",
	"q": "Quebec", "r": "Romeo", "s": "Sierra", "t": "Tango",
	"u": "Uniform", "v": "Victor", "w": "Whiskey", "x": "X-ray",
	"y": "Yankee", "z": "Zulu",
}

func main() {
	http.HandleFunc("/", fileServerHandler)
	http.HandleFunc("/nato/", natoHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func fileServerHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "index.html")
	} else {
		http.NotFound(w, r)
	}
}

func natoHandler(w http.ResponseWriter, r *http.Request) {
	word := strings.ToLower(r.URL.Query().Get("word"))
	var natoWords []string
	for _, char := range word {
		if char < 'a' || char > 'z' {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		natoWord, exists := natoMap[string(char)]
		if !exists {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		natoWords = append(natoWords, natoWord)
	}
	response := strings.Join(natoWords, "\n")
	fmt.Fprint(w, response)
}

