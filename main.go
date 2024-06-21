package main

import (
	"CloudSystem/database"
	"CloudSystem/routes"
	"fmt"
)

// var result = make(map[string]Word)

// var totalFrequencyInAllFiles = 0

func main() {
	fmt.Println("Project start")
	// searchWords([]string{"Qmeu", "LD"})
	// fmt.Println(result)
	// Initialize the database connection pool
	database.Init()

	router := routes.RegisterRoutes()
	router.Run(":8080")
}

// type Word struct {
// 	totalFrequencyInAllFiles    int64
// 	DocumentFrequencyInAllFiles int64
// 	SearchCount                 int64
// 	name                        string
// }

// func searchWords(words []string) {

// 	files := [3]string{"./SearchFiles/file1.txt", "./SearchFiles/file2.txt", "./SearchFiles/file3.txt"}
// 	fmt.Println(words)
// 	// for _, file := range files {
// 	// 	CalculateTFAndDF(file, words)

// 	// }
// 	for _, word := range words {

// 		CalculateTFAndDF(files, word)
// 	}
// }

// func CalculateTFAndDF(files [3]string, word string) {

// 	var wordData Word
// 	wordData.name = word
// 	for _, file := range files {
// 		content, err := os.ReadFile(file)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fileTf, containsWord := searchWordInContent(word, string(content))
// 		wordData.totalFrequencyInAllFiles += int64(fileTf)
// 		if containsWord {
// 			wordData.DocumentFrequencyInAllFiles += 1
// 		}
// 		result[word] = wordData

// 	}

// 	// fmt.Println(string(content))
// }
// func searchWordInContent(word, content string) (int, bool) {
// 	tf := 0
// 	containsWord := false
// 	words := strings.Fields(content)

// 	for _, w := range words {
// 		if strings.EqualFold(w, word) {
// 			tf++
// 			containsWord = true
// 		}
// 	}

// 	return tf, containsWord
// }
