package runner

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"os"
	"reflect"
	"strconv"
	"time"
	"todo_list/controllers"
	"todo_list/internal/utils"
	"todo_list/models"
)

type Cli struct {
	Factory
	DB gorm.DB
}

func (c *Cli) GetName() string {
	return "Command-line interface"
}

func (c *Cli) Run(args []string, db gorm.DB) {
	if err := c.ValidateCliArgs(args); err != nil {
		panic("Could not validate CLI arguments.")
	}
	action, ok := controllers.GetActions(db)[args[0]]
	if !ok {
		panic("Action method does not exist.")
	}

	body := c.getArgument(args)

	result, _ := action.Do(body)

	jsonData, _ := json.Marshal(result)

	lengthWritten, err := writeToFile(jsonData, args[0])
	if err != nil {
		return
	}

	if lengthWritten > 0 {
		fmt.Println(fmt.Sprintf("%s bytes written to file successfully.", strconv.Itoa(lengthWritten)))
	}
}

func (_ *Cli) ValidateCliArgs(args []string) error {
	if len(args) < 1 {
		panic("Action method must be entered.")
	}
	return error(nil)
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

func (c *Cli) getArgument(args []string) controllers.ArgumentBody {

	body := make(controllers.ArgumentBody)

	typ := reflect.TypeOf(models.Task{})

	flags := utils.ParseFlags(args)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("flag")
		if tag != "" {
			if _, ok := flags[tag]; ok {
				body[tag] = flags[tag]
			}
		}
	}

	return body
}
