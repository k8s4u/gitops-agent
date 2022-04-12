package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type azStatusContext struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
}
type azStatusBody struct {
	State       string          `json:"state"`
	Description string          `json:"description"`
	Context     azStatusContext `json:"context"`
	TargetUrl   string          `json:"targetUrl"`
}

func main() {
	_, gitrepo := os.LookupEnv("GITREPO")
	if !gitrepo {
		fmt.Println("Environment variable GITREPO is mandatory")
		os.Exit(1)
	}
	envName, environment := os.LookupEnv("ENVIRONMENT")
	if !environment {
		fmt.Println("Environment variable ENVIRONMENT is mandatory")
		os.Exit(1)
	}

	// FixMe: We This should be provided as secret
	pat, _ := os.LookupEnv("PAT")

	azdev_repo_api_url, _ := os.LookupEnv("AZDEV_REPO_API_URL")
	// _, github_update_status := os.LookupEnv("GITHUB_UPDATE_STATUS")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		/*
			cmd := exec.Command("/scripts/sync.sh")
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("sync.sh failed with %s\n", err)
			}
			fmt.Println(string(out))
		*/
		if azdev_repo_api_url != "" {
			url := azdev_repo_api_url + "/commits/090f64adff8adcb05a12ee63dba69be5b7f11772/statuses?api-version=6.0"
			AzureDevOpsStatus(pat, url, envName, true)
		}
	})

	fmt.Println("Listening webhook on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func AzureDevOpsStatus(pat, apiUrl, envName string, ok bool) {
	state := "failed"
	if ok {
		state = "succeeded"
	}
	bodyContext := azStatusContext{
		Name:  fmt.Sprintf("sync/%s", envName),
		Genre: "gitops",
	}
	body := azStatusBody{
		State:       state,
		Description: fmt.Sprintf("GitOps Sync To %s", envName),
		Context:     bodyContext,
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(requestBody))

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json;excludeUrls=true;enumsAsNumbers=true;msDateFormat=true;noArrayWrap=true")
	req.SetBasicAuth("user:", pat)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}
