package main

import (
	"blazeisclone/birthday-notifier/internal"
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	scheduler := cron.New()
	cronExpression := "* * * * *"

	toName := "Person Name"
	toEmail := "person@example.com"
	subject := "Reminder for Birthday"
	message := "Reminder for Birthday"

	scheduler.AddFunc(cronExpression, func() {
		internal.SendMail(toName, toEmail, subject, message)
	})

	scheduler.Start()

	log.Println("Scheduler started.")

	select {}
}
