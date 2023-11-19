package main

import "github.com/ashiruhabeeb/simpleTodoApp/cmd/server"

func main() {
	// Initialize server app
	app := server.App{}

	// Run app server
	app.AppRun()
}
