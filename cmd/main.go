package main

import "github.com/ashiruhabeeb/simpleTodoApp/cmd/server"

func main() {
	app := server.App{}

	app.AppRun()
}