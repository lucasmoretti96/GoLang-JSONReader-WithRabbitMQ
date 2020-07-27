package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"easynvest.com/treasurydirect/accountidsqueue/models"
	"easynvest.com/treasurydirect/accountidsqueue/queue"
)

func main() {
	fmt.Println("Getting file")

	jsonFile, err := os.Open("../users.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users models.Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("Username: " + users.Users[i].Username + " User Id: " + users.Users[i].AccountId)
	}

	sendAccountIds(users)
}

func sendAccountIds(users models.Users) {

	var accounts = models.AccountIds{}

	transaction := accounts

	for i := 0; i < len(users.Users); i++ {
		transaction.Id = users.Users[i].AccountId
		_json, _ := json.Marshal(transaction)
		queue.SendQueue(string(_json))
	}
}
