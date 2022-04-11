package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	_, gitrepo := os.LookupEnv("GITREPO")
	if !gitrepo {
		fmt.Println("Environment variable GITREPO is mandatory")
		os.Exit(1)
	}
	_, environment := os.LookupEnv("ENVIRONMENT")
	if !environment {
		fmt.Println("Environment variable ENVIRONMENT is mandatory")
		os.Exit(1)
	}

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.Command("/scripts/sync.sh")
		out, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Printf("sync.sh failed with %s\n", err)
		}
		fmt.Println(string(out))
	})

	fmt.Println("Listening webhook on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
