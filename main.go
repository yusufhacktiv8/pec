package main

import (
	"fmt"
	"os"

	"pecrsuh.id/pec/application"
)

func main() {
	a := application.App{}
	fmt.Printf("Running PEC Server...\n")
	a.InitDatabase(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	a.Router.Run(":8080")
}
