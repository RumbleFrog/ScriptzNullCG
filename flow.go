package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// SectionIndex - Per Page of Threads Indexing
// ThreadIndex - Per Section Page of Threads Indexing
var (
	SectionIndex int
	ThreadIndex  int
)

// Tracker - Keep track of index processing
var Tracker map[*Thread]int

func process(s *Section, t *Thread) {
	// Process all replies, check if we have enough, if not, fetch more
	index, ok := Tracker[t]

	if ok == false {
		Tracker[t] = 0
	}

	for ; index < len(t.Replies); index++ {
		//Send off to payload generator
		if credit <= 0 {

			byteValues, _ := json.Marshal(&Sections)

			ioutil.WriteFile("forum.json", byteValues, 0777)

			log.Fatal("Process Complete")
		}
		credit -= 3
	}

	Tracker[t] = index

	if t.Page >= t.Pages {
		addToCache(t)
		if ThreadIndex < len(s.Threads) {
			ThreadIndex++
			fetchReply(s, s.Threads[ThreadIndex])
		} else {
			if s.Page >= s.Pages {
				ThreadIndex = 0
				SectionIndex++
				fetchThreads(Sections[SectionIndex])
			} else {
				s.Page++
				fetchThreads(s)
			}
		}
	} else {
		t.Page++
		fetchReply(s, t)
	}
}
