package main

import (
	"github.com/gocolly/colly"
	"github.com/mmcdole/gofeed"
)

func get_text(url string) string {
	c := colly.NewCollector()
	data := ""
	c.OnHTML("p", func(e *colly.HTMLElement) {
		data += e.Text + "\n"
	})
	c.Visit(url)
	return data
}
func get_articles(limit int) map[string]string {
	data := make(map[string]string)
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://vnexpress.net/rss/tin-noi-bat.rss")
	var artinum int
	if limit == 0 {
		artinum = len(feed.Items)
	} else {
		artinum = limit
	}
	for i := 0; i != artinum; i++ {
		title := feed.Items[i].Title
		link := feed.Items[i].Link
		//fmt.Println(title,link)
		data[link] = title
	}
	return data
}
