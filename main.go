package main

import (
	"net/http"
	"log"
	"encoding/json"
)

const cUrl = "https://api.hh.ru/vacancies"

func main() {
	makeRequest()
}

func makeRequest() {

	client := &http.Client{} 
    req, err := http.NewRequest("GET", cUrl, nil, 
    ) 
    // добавляем заголовки
    req.Header.Add("User-Agent", "api-test-agent") 
  
    resp, err := client.Do(req) 
    if err != nil { 
        log.Println(err) 
        return
    } 
    defer resp.Body.Close() 
	var result map[string]interface{
	}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result["items"])
	
}