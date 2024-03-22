package runner

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
	"todo_list/controllers"
)

type Cli struct {
	Factory
	DB gorm.DB
}

func (c *Cli) GetName() string {
	return "Command-line interface"
}

func (c *Cli) Run(args []string, db gorm.DB) {
	if len(args) < 3 {
		panic("Action method must be inputted.")
	}
	action, ok := controllers.GetActions(db)[args[2]]
	if !ok {
		panic("Action method does not exist.")
	}

	result, _ := action.Do(args)

	jsonData, _ := json.Marshal(result)

	lengthWritten, err := writeToFile(jsonData, args[2])
	if err != nil {
		return
	}

	if lengthWritten > 0 {
		fmt.Println(fmt.Sprintf("%s bytes written to file successfully.", strconv.Itoa(lengthWritten)))
	}
}

func writeToFile(jsonData []byte, method string) (int, error) {
	defaultFileName := method + time.Now().Format("2006-01-02_15-04") + ".json"
	file, err := os.Create(defaultFileName)
	if err != nil {
		panic(fmt.Sprintf("Could not write fetched data into file %s", defaultFileName))
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	return file.Write(jsonData)
}
