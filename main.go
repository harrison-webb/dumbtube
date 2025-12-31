package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func main() {
	// ok what do I need to do
	// - get RSS feed from main channel URL
	// - download RSS feed
	// -

	testURLString := "https://www.youtube.com/@jvscholz"
	testURL, err := url.Parse(testURLString)
	if err != nil {
		panic(err)
	}

	channelRSSURL := getRssFeedFromChannelUrl(testURL)
	fmt.Print(channelRSSURL)
}

func getRssFeedFromChannelUrl(channelURL *url.URL) *url.URL {
	// add User-Agent header to request because claude said youtube might be rate limiting me without one
	client := &http.Client{}
	request, err := http.NewRequest("GET", channelURL.String(), nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic(0)
	}

	// only read the first half of the body so that i dont have to parse the whole thing
	limitedReader := io.LimitReader(response.Body, 800*1024) // only read 512KB (~0.5MB)
	body, err := io.ReadAll(limitedReader)
	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile("channel_id=[a-zA-Z0-9]{24}")
	match := r.Find(body)
	channelId := string(match[len(match)-24:])

	RSSURL, err := url.Parse("https://www.youtube.com/feeds/videos.xml?channel_id=" + channelId)
	if err != nil {
		panic(err)
	}

	return RSSURL
}
