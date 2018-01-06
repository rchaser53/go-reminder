package main

import (
	"encoding/json"
	"os"

	"goReminder/Scrape"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
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

func Handle(evt json.RawMessage, ctx *runtime.Context) (string, error) {
	println("nanisore-?")
	return "gya-n", nil
}
