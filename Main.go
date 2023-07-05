package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// TODO add access token here
const accessToken string = ""

type DomainResponse struct {
	Registrar Registrar `json:"registrar"`
}

type Registrar struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	ReferralUrl string `json:"referral_url"`
}

func main() {
	file, err := os.Open("domains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		getDomain(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getDomain(domain string) error {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("GET", "https://whoisjsonapi.com/v1/"+domain, nil)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	var domainResponse DomainResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&domainResponse)
	if err != nil {
		log.Fatal(err)
	}

	// Printing only Registrar part
	fmt.Println("Registrar:", domainResponse.Registrar)
	return nil
}
