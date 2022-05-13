package excel

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

// Convert excel file to csv

type Mark struct {
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

func Parsesheet(r io.Reader) []Mark {
	csvReader := csv.NewReader(r)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	marks := createmarks(data)
	return marks
}

func createmarks(data [][]string) []Mark {
	var marks []Mark
	for i, line := range data {
		if i > 0 { // omit header line
			var rec Mark
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 1 {
					s, _ := strconv.Atoi(field)
					rec.Score = int64(s)
				}
			}
			marks = append(marks, rec)
		}
	}
	return marks
}
