package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/aashi0000/search_engine_using_golang/utils"
)

func main() {
	var dump, query string
	//defining dump path and query
	flag.StringVar(&dump, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	flag.StringVar(&query, "q", "Small Wild Cat", "search query")
	flag.Parse()
	log.Println("full text search is in progress...")
	start := time.Now()
	//load documents
	docs, err := utils.LoadDocuments(dump)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded %d documents in %v", len(docs), time.Since(start))
	//indexing
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("indexed %d documents in %v", len(docs), time.Since(start))
	//searching for reqired docs
	start = time.Now()
	matched := idx.Search(query)
	log.Printf("search found %d documents in %v", len(matched), time.Since(start))
	//displaying query result
	for _, id := range matched {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
