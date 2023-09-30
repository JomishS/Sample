package main

import (
	"example/Project3/internal/config"
	"example/Project3/cmd"

)
func main() {

	config.Loadconfig()
	cmd.Execute()

}
