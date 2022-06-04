package main

import (
	"net/http"
	// "log"
	"encoding/json"
	 "fmt"
)

type tItem struct{
	Id string
	Name string
}

type tResponse struct{
	Items []tItem 
}

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
        fmt.Println(err) 
        return
    } 
    defer resp.Body.Close() 

	var result tResponse
	
	json.NewDecoder(resp.Body).Decode(&result)

	for key, val := range result.Items {
		fmt.Printf("%d %s\n",key + 1, val.Name)
	}
	
}


