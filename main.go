package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/urfave/cli"
)

// Sections
// SectionCollector - Forum Categories Collector
// ThreadCollector - Threads Collector
// ReplyCollector - Reply Collector
var (
	ua               = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36"
	cookie           string
	credit           uint64
	client           = &http.Client{}
	Sections         []*Section
	SectionCollector = colly.NewCollector()
	ThreadCollector  = colly.NewCollector()
	ReplyCollector   = colly.NewCollector()
)

func main() {
	app := cli.NewApp()

	app.Name = "ScriptzNullCG"
	app.Usage = "ScriptzNull Credit Generator"

	app.Flags = []cli.Flag{
		cli.Uint64Flag{
			Name:        "c",
			Value:       100,
			Usage:       "Desired amount of credits",
			Destination: &credit,
		},
	}

	app.Action = func(c *cli.Context) error {
		if cookie = c.Args().Get(0); cookie == "" {
			log.Fatal("Cookie is not provided")
		}

		Tracker = make(map[*Thread]int)

		loadCache()
		fetchSections()

		Cache = append(Cache, &Thread{
			Name: "test",
		})

		saveCache()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func onRequest(r *colly.Request) {
	r.Headers.Set("User-Agent", ua)
	r.Headers.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	r.Headers.Set("Cookie", cookie)
}

func onError(r *colly.Response, e error) {
	log.Fatal("Try a fresh token perhaps?: ", e)
}

func onResponse(r *colly.Response) {
	log.Println(r.Body)
}
