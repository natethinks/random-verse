package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var bom BookOfMormon
	err := json.Unmarshal([]byte(bomjson), &bom)
	if err != nil {
		fmt.Println(err)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	bookIndex := r1.Intn(len(bom.Books) - 1)
	chapterIndex := 0
	if len(bom.Books[bookIndex].Chapters) != 1 {
		chapterIndex = r1.Intn(len(bom.Books[bookIndex].Chapters) - 1)
	}

	verseIndex := r1.Intn(len(bom.Books[bookIndex].Chapters[chapterIndex].Verses) - 1)
	fmt.Println(bom.Books[bookIndex].Book, bom.Books[bookIndex].Chapters[chapterIndex].Chapter, ":", bom.Books[bookIndex].Chapters[chapterIndex].Verses[verseIndex].Verse)
	fmt.Println(bom.Books[bookIndex].Chapters[chapterIndex].Verses[verseIndex].Text)

}
