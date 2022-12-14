package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

// register visits to the site in a csv file and console, then redirect to URL
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// get IP address
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		// get current time
		t := time.Now()
		// get URL
		url := r.URL.String()
		// write to console
		fmt.Printf("%s, %s, %s", ip, t, url) // write to csv file on new line (append) or create file if it doesn't exist
		f, err := os.OpenFile("visits.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(fmt.Sprintf("%s, %s, %s", ip, t, url)); err != nil {
			fmt.Println(err)
		}
		// redirect to URL (change to your URL)
		http.Redirect(w, r, "https://www.google.com", 301)
	})
	// start server and print to console when it's running and listening on port 8080
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
