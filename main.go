package main

import (
	"encoding/json"
	"encoding/xml"
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
	// - parse RSS feed (don't want shorts included)

	testURLString := "https://www.youtube.com/@tanksfornothin/"
	testURL, err := url.Parse(testURLString)
	if err != nil {
		panic(err)
	}

	channelRSSURL := getRSSURLFromChannelURL(testURL)
	feed, _ := fetchRSSFeed(channelRSSURL)
	parsedFeed := parseRSSFeed(feed)
	out, err := json.MarshalIndent(parsedFeed, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
}

// getRSSURLFromChannelURL returns the URL of the RSS feed of a youtube channel when given that channel's homepage URL
func getRSSURLFromChannelURL(channelURL *url.URL) *url.URL {
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
	// limitedReader := io.LimitReader(response.Body, 800*1024) // only read 512KB (~0.5MB)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// regex pattern: channel_id{24 characters}
	r := regexp.MustCompile(`channel_id=\S{24}`)
	match := r.Find(body)
	channelId := string(match[len(match)-24:])

	RSSURL, err := url.Parse("https://www.youtube.com/feeds/videos.xml?channel_id=" + channelId)
	if err != nil {
		panic(err)
	}

	return RSSURL
}

// fetchRSSFeed returns the RSS feed contents of a given youtube channel RSS URL
func fetchRSSFeed(rssURL *url.URL) ([]byte, error) {
	// add User-Agent header to request because claude said youtube might be rate limiting me without one
	client := &http.Client{}
	request, err := http.NewRequest("GET", rssURL.String(), nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

func parseRSSFeed(rssBody []byte) feed {
	var resultFeed feed
	xml.Unmarshal(rssBody, &resultFeed)
	return resultFeed
}

// feed is an array of entries
// each entry has:
// 		- id -> string
//		- title -> string
//		- link -> URL
// 		- author -> string
//		- date -> date
// 		- media:thumbnail -> URL

type feed struct {
	// XMLName xml.Name `xml:"feed"`
	Entries []entry `xml:"entry"`
}

type entry struct {
	// XMLName     xml.Name   `xml:"entry"`
	ID          string     `xml:"id"`
	Title       string     `xml:"title"`
	Link        link       `xml:"link"` // TODO: format looks like: <link rel="alternate" href="https://www.youtube.com/watch?v=dUDpjxPm8os"/> and I just want to pull out the href attribute string
	Author      author     `xml:"author"`
	PublishDate string     `xml:"published"`
	MediaGroup  mediaGroup `xml:"http://search.yahoo.com/mrss/ group"`
}

type link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type mediaGroup struct {
	Thumbnail mediaThumbnail `xml:"http://search.yahoo.com/mrss/ thumbnail"`
}

type mediaThumbnail struct {
	URL string `xml:"url,attr"`
}

type author struct {
	// XMLName    xml.Name `xml:"author"`
	Name       string `xml:"name"`
	ChannelURI string `xml:"uri"`
}
