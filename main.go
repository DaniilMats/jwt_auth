package main

import (
	"jwt_auth/cmd"
	"jwt_auth/config"
)

func main() {
	c, err := config.GetConfig("./environment/.env")
	if err != nil {
		panic(err)
	}
	cmd.RunApplication(c)
}
