package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var path = "./"
var file = "data.json"
var pathFile = path + file

type User struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty"`
	Salary int    `json:"salary,omitempty"`
}

func main() {
	RunApplication()
}

func RunApplication() {
	output, err := readFile(pathFile)
	if err != nil {
		panic(err)
	}

	data := changeCurrencyToIDR(output)
	isDone := writeFile(data)

	if <-isDone {
		log.Println("Change currency is done...")
	}
}

// Read json file
func readFile(fileName string) (<-chan User, error) {
	output := make(chan User)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return output, err
	}

	users := []User{}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return output, err
	}

	go func() {
		for _, user := range users {
			output <- user
		}
		close(output)
	}()

	return output, nil
}

// Make worker for change currency USD to IDR
func changeCurrencyToIDR(dataCh <-chan User) <-chan User {
	output := make(chan User)
	go func() {
		for data := range dataCh {
			newData := data
			newData.Salary = newData.Salary * 15_000
			output <- newData
		}
		close(output)
	}()

	return output
}

// Write new file with content name as file name
func writeFile(dataCh <-chan User) <-chan bool {
	err := os.Mkdir("users", 7666)
	CheckError(err)

	wg := sync.WaitGroup{}
	isDone := make(chan bool)

	for data := range dataCh {
		wg.Add(1)
		go func(data User) {
			userByte, err := json.Marshal(data)
			CheckError(err)
			err = os.WriteFile(path+"users/"+data.Name+".json", userByte, 0666)
			CheckError(err)

			wg.Done()
		}(data)
	}

	go func() {
		wg.Wait()
		isDone <- true
		close(isDone)

	}()

	return isDone
}

// Check if error
func CheckError(err error) bool {
	if err != nil {
		log.Println(err.Error())
	}
	return (err != nil)
}
