package main

import (
	"net/http"
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

type tItemsInfo struct {
	Name string
}

type tResponseItemInfo struct {
	Key_skills []tItemsInfo
}

const cUrl = "https://api.hh.ru/vacancies"

func getReqestData(urlAPI string, result interface{}) (error) {
	client := &http.Client{} 
    req, err := http.NewRequest("GET", urlAPI, nil, 
    ) 
	if err != nil {
		return err
	}
    // добавляем заголовки
    req.Header.Add("User-Agent", "api-test-agent") 
  
    resp, err := client.Do(req) 
    if err != nil { 
        return err
    } 
    defer resp.Body.Close() 
	
	json.NewDecoder(resp.Body).Decode(&result)
	return nil
}

func getListVacancies() ([]tItem, error) {
	var responce tResponse

	err := getReqestData(cUrl, &responce)
	if err != nil {
		return []tItem{}, err
	}
	return responce.Items, nil
}

func getSkillsVacancie(id string) (string, error) {
	var itemInfo tResponseItemInfo

	err := getReqestData(cUrl + "/" + id, &itemInfo)
	if err != nil {
		return "", err
	}

	var result string

	for _, val := range itemInfo.Key_skills {
		result += val.Name + ","
	}
	
	if result != "" {
		result = result[:len(result) - 1]
	}

	return result, nil
}

func main() {
	listVacancies, err := getListVacancies();
	var a int

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		for key, val := range listVacancies {
			fmt.Printf("%d %s\n",key + 1, val.Name)
		}

		fmt.Print("Введите номер вакансии для просмотра необходимых скилов, для выхода введите 0\n")
		fmt.Scan(&a)

		if (a == 0) {
			break
		} else if (a > len(listVacancies)) {
			fmt.Print("Введен не корректный номер вакансии\n")	
		} else {
			skillInfo, err := getSkillsVacancie(listVacancies[a - 1].Id)
			if err != nil {
				fmt.Printf("Ошибка при получении информаиции о скалах\n%s", err)
			}

			if skillInfo != "" {
				fmt.Printf("Перечень скилов:%s\n",skillInfo)
			} else {
				fmt.Printf("Для выбранной вакансии перечень скилов не задан\n")
			}
			
			break
		}

	}
	

	
}


