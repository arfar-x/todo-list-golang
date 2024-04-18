package cmd

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"todo_list/internal/runner"
)

func Run(runnerType string, args []string) {
	factory := runner.Factory{Name: runnerType}
	server := factory.Serve(runnerType)
	if server == nil {
		panic("Server type could not be found.")
	}
	server.Run(args, connectDB())
}

func connectDB() gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//defer func() {
	//	dbInstance, _ := DB.DB()
	//	_ = dbInstance.Close()
	//}()
	if err != nil {
		panic("Could not connect to database.")
	}

	return *DB
}
