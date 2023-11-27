package main

import (
	"fmt"
	"os"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int { //return the index of entry
	entryCount++
	entry := fmt.Sprintf("%d %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {

	if index >= 0 && index < len(j.entries) {
		j.entries = append(j.entries[:index], j.entries[index+1:]...)
	}
}

// separation of concerns
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("Hello there")
	j.AddEntry("i m satyakr")
	j.AddEntry("Welcome to world of GO")
	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistence{lineSeparator: "\t"}
	p.SaveToFile(&j, "journal-two.txt")
}
