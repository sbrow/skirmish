// TDOD: Check if dataset file has changed before prompting in photoshop
package skirmish

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	_ "path/filepath"
	"strings"
)

func Leader(cardID string) (string, error) {
	// cards, err := os.Open(filepath.Join(os.Getenv("SK_SRC"), "data.txt"))
	cards, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(bufio.NewReader(cards))
	headMap := map[string]int{
		"id":     0,
		"Bast":   -1,
		"Igrath": -1,
	}
	headers, err := r.Read()
	if err != nil {
		panic(err)
	}
	for i, name := range headers {
		for key := range headMap {
			if strings.ToLower(key) == name {
				headMap[key] = i
				break
			}
		}
	}
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for key, value := range headMap {
			if key == "id" {
				if line[value] != cardID {
					break
				}
			} else {
				if line[value] == "true" {
					return strings.Title(key), nil
				}
			}
		}
	}
	return "", errors.New(fmt.Sprintf("\"%s\" was not found.", cardID))
}
