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
	postSlackBinanceReStart()
}

func postSlackBinanceReStart() {
	url := os.Getenv("SlackURL")
	err := goReminder.PostSlack(url)
	if err != nil {
		panic(err)
	}
}

func Handle(evt json.RawMessage, ctx *runtime.Context) (string, error) {
	if goReminder.IsBinanceReStart(TargetWord) {
		postSlackBinanceReStart()
	}

	return "succeed", nil
}
