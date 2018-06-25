package main

import (
	"log"
	"math"
	"strconv"

	"github.com/gocolly/colly"
)

// Section ...
type Section struct {
	Name         string
	Href         string
	Page         uint64
	Pages        uint64
	Threads      []*Thread
	ThreadCount  uint64
	MessageCount uint64
}

var (
	tc  uint64
	mc  uint64
	err error
)

func fetchSections() {
	SectionCollector.OnRequest(onRequest)
	SectionCollector.OnError(onError)

	SectionCollector.OnHTML("li.node > div.nodeInfo > div.nodeText", func(e *colly.HTMLElement) {
		if tc, err = strconv.ParseUint(e.ChildText("div.nodeStats > dl:first-child > dd"), 10, 64); err != nil {
			return
		}

		if mc, err = strconv.ParseUint(e.ChildText("div.nodeStats > dl:last-child > dd"), 10, 64); err != nil {
			return
		}

		sections = append(sections, &Section{
			Name:         e.ChildText("h3.nodeTitle > a[href]"),
			Href:         e.ChildAttr("h3.nodeTitle > a[href]", "href"),
			Page:         1,
			Pages:        uint64(math.Ceil(float64(tc) / float64(20))),
			ThreadCount:  tc,
			MessageCount: mc,
		})
	})

	SectionCollector.OnScraped(func(r *colly.Response) {
		for _, s := range sections {
			log.Printf("Name: %s | Directive: %s | Threads: %d | Replies: %d", s.Name, s.Href, s.ThreadCount, s.MessageCount)
		}
	})

	SectionCollector.Visit(formatTarget(nil, nil))
}