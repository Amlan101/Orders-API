package main

import (
	"context"
	"fmt"
	"github.com/Amlan101/Orders-API.git/application"
)

func main(){
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil{
		fmt.Println("Failed to launch the app", err)
	}
}
