// client.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Make GET request to server's /ping endpoint
	resp, err := http.Get("http://localhost:8081/ping")
	if err != nil {
		fmt.Println("Error calling server:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response from server:", string(body))

	// Call /greet with query parameter
	greetResp, err := http.Get("http://localhost:8081/greet?name=Ranjith")
	if err != nil {
		fmt.Println("Error calling greet:", err)
		return
	}
	defer greetResp.Body.Close()

	greetBody, _ := ioutil.ReadAll(greetResp.Body)
	fmt.Println("Greeting:", string(greetBody))
}
