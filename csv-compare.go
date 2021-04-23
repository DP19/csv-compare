package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"sync"
)

type Entry struct {
	Name string `csv:"name"`
	Email string `csv:"email"`
}

func main() {
	//string to file path of first csv file
	csvFileOne := os.Args[1]
	//string to file path of second csv file
	csvFileTwo := os.Args[2]


	// open first csv
	csv1, err := os.Open(csvFileOne)
	// panic and exit if it can't
	if err != nil {
		panic(err)
	}

	// setup structure
	csv1Entries := []*Entry{}
	defer csv1.Close()
	// read csv file into a list of our Entry structure
	if err := gocsv.UnmarshalFile(csv1, &csv1Entries); err != nil {
		panic(err)
	}

	// Same thing for the second file
	csv2Entries := []*Entry{}
	csv2, err := os.Open(csvFileTwo)
	if err != nil {
		panic(err)
	}
	if err := gocsv.UnmarshalFile(csv2, &csv2Entries); err != nil {
		panic(err)
	}

	// setup multi threading so we can loop through this really fast
	var wg sync.WaitGroup
	wg.Add(len(csv2Entries))

	for i := 0; i < len(csv2Entries); i++ {
		go func(i int) {
			defer wg.Done()
			for _, entry := range csv2Entries {
				// for each name in the first csv file, compare it to each of the names in the second.
				// if there is a match we print to the output the full entry we found a match on
				if csv1Entries[i].Name == entry.Name {
					fmt.Printf("%s in both entries", entry.Name)
				}
			}
		}(i)
	}
	wg.Wait()
}
