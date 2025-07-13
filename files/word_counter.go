package files

import (
	"bufio"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/prasannakumar414/finder/models"
)

func WordCounter(filePath string, wordMapChan chan (map[string]int), wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error when opening file: %v", err)
	}
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	wordMap := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		wordMap[word]++
	}
	wordMapChan <- wordMap
}

func GetMostFrequentWordsFromMap(wordsCountMap map[string]int) []models.WordCount {
	var sortedWordCounts []models.WordCount
	for k, v := range wordsCountMap {
		sortedWordCounts = append(
			sortedWordCounts,
			models.WordCount{
				Word:  k,
				Count: v,
			})
	}

	// Sort the slice by Count in descending order
	sort.Slice(sortedWordCounts, func(i, j int) bool {
		return sortedWordCounts[i].Count > sortedWordCounts[j].Count
	})
	if len(sortedWordCounts) > 10 {
		return sortedWordCounts[:10]
	}
	return sortedWordCounts
}
