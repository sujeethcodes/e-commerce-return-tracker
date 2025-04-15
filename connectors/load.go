package connectors

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Creating env instance")
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("env connector error")
	}
}
