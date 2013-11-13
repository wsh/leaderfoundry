package main

import (
	"fmt"
	"net/http"
	"os"
)

var apiConfig *TwitterAPIConfig

func main() {
	apiConfig = readconfig()
	fmt.Println(apiConfig.Key)
	http.HandleFunc("/", leaderserve)
	fmt.Println("Up.")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func leaderserve(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Leadership coming soon.\n")
}
