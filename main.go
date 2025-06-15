package main

import (
	"github.com/robfig/cron/v3"
	"go-notification/db"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	log.Println("Starting notification...")
	db.Init()

	c := cron.New()
	c.Start()

	log.Println("Starting cron...")
}
