package main

import (
	"os"

	"goReminder/Scrape"
)

var TargetWord string = "Registrations Pause"

type HogeError struct {
	message string
}

func (hoge HogeError) Error() string {
	return hoge.message
}

func main() {
	url := os.Getenv("SlackURL")

	goReminder.ScrapeBinanceRegistratoinForm(TargetWord)
	err := goReminder.PostSlack(url)

	err = HogeError{"giepi-"}

	if err != nil {
		panic(err)
	}
}
