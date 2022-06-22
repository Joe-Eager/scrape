package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
	// for range time.Tick(time.Minute * 15) {
	// 	go func() {
	dt := time.Now()
	scraper := twitterscraper.New()
	output, _ := os.Create("output.json")
	output.WriteString(fmt.Sprintf("{\"lastUpdated\":\"%s\",\"CMPmtb\":{", dt.Format("Mon 02 Jan 06 03:04 pm")))
	re := regexp.MustCompile(`\r?\n|\"`)
	i := 0
	for tweet := range scraper.GetTweets(context.Background(), "CMPmtb", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		output.WriteString(fmt.Sprintf("\"%d\":{\"time\":%d000,\"tweet\":\"%s\"},", i, tweet.Timestamp, re.ReplaceAllString(tweet.Text, " ")))
		i++
	}
	output.WriteString(fmt.Sprintf("\"%d\":{\"time\":\"\",\"tweet\":\"\"}},", i))
	output.WriteString("\"SMPmountainbike\":{")
	i = 0
	for tweet := range scraper.GetTweets(context.Background(), "SMPmountainbike", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		output.WriteString(fmt.Sprintf("\"%d\":{\"time\":%d000,\"tweet\":\"%s\"},", i, tweet.Timestamp, re.ReplaceAllString(tweet.Text, " ")))
		i++
	}
	output.WriteString(fmt.Sprintf("\"%d\":{\"time\":\"\",\"tweet\":\"\"}},", i))
	output.WriteString("\"VulturesKnobMTB\":{")
	i = 0
	for tweet := range scraper.GetTweets(context.Background(), "VulturesKnobMTB", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		output.WriteString(fmt.Sprintf("\"%d\":{\"time\":%d000,\"tweet\":\"%s\"},", i, tweet.Timestamp, re.ReplaceAllString(tweet.Text, " ")))
		i++
	}
	output.WriteString(fmt.Sprintf("\"%d\":{\"time\":\"\",\"tweet\":\"\"}},", i))
	output.WriteString("\"CVNPmtb\":{")
	i = 0
	for tweet := range scraper.GetTweets(context.Background(), "CVNPmtb", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		output.WriteString(fmt.Sprintf("\"%d\":{\"time\":%d000,\"tweet\":\"%s\"},", i, tweet.Timestamp, re.ReplaceAllString(tweet.Text, " ")))
		i++
	}
	output.WriteString(fmt.Sprintf("\"%d\":{\"time\":\"\",\"tweet\":\"\"}},", i))
	output.WriteString("\"medina_trails\":{")
	i = 0
	for tweet := range scraper.GetTweets(context.Background(), "medina_trails", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		output.WriteString(fmt.Sprintf("\"%d\":{\"time\":%d000,\"tweet\":\"%s\"},", i, tweet.Timestamp, re.ReplaceAllString(tweet.Text, " ")))
		i++
	}
	output.WriteString(fmt.Sprintf("\"%d\":{\"time\":\"\",\"tweet\":\"\"}}", i))
	output.WriteString("}")
	// 	}()
	// }
}
