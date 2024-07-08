package controller

import (
	"CloudSystem/database"
	"CloudSystem/models"
	"CloudSystem/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var result = make(map[string]Word)

func RegisterUser(context *gin.Context) {
	currentConnection, err := database.DB.Begin()
	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	defer func() {
		if err != nil {
			currentConnection.Rollback()
		}
	}()
	var user models.User
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// user.Id = context.GetInt64("userId")
	_, err = models.GetUserByEmail(user.Email)
	if err == nil {
		context.JSON(http.StatusConflict, gin.H{"message": "email already exists"})
		return
	}
	_, err = user.AddUser(currentConnection)
	fmt.Println("errorsQQ")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later. Adduser"})
		return
	}
	fmt.Println("in register PASSED")
	_, err = user.AddUserAddress(currentConnection)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later. address"})
		return
	}
	err = currentConnection.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created Successfully", "userId": user.Identifier})
}
func LoginUser(context *gin.Context) {
	// extract body
	body, err := utils.ExtractBodyFromRequest(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "SomeThing Went Wrong"})
		return
	}
	//check if email or password
	email, isEmailExists := body["email"]
	password, isPasswordExists := body["password"]
	if !isEmailExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing email in the request body"})
		return
	}
	if !isPasswordExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing Password in the request body"})
		return
	}
	user, err := models.GetUserByEmail(email.(string))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
		return
	}
	err = user.ValidatePassword(password.(string))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
		return
	}
	//generate token and get response
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

type WordsRequestBody struct {
	Words []string `json:"words"`
}

func SearchWords(context *gin.Context) {
	historyFilePath := `D:\LD Academy\Go-lang-Cloud-System\SearchFiles\history.json`
	files := [3]string{`D:\LD Academy\Go-lang-Cloud-System\SearchFiles\file1.txt`, `D:\LD Academy\Go-lang-Cloud-System\SearchFiles\file2.txt`, `D:\LD Academy\Go-lang-Cloud-System\SearchFiles\file3.txt`}

	var WordsRequestBody WordsRequestBody

	// Bind the JSON to the struct
	if err := context.ShouldBindJSON(&WordsRequestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wordChan := make(chan Word)

	// Read JSON from file into a map
	searchCounts, err := readJSONFromFile(historyFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, word := range WordsRequestBody.Words {
		searchCounts[strings.ToLower(word)]++
		if err := writeJSONToFile(historyFilePath, searchCounts); err != nil {
			fmt.Println("Error writing JSON file:", err)
			return
		}
		wg.Add(1)
		go func(word string) {
			defer wg.Done() // it decrement wait group by one
			CalculateTFAndDF(files, word, wordChan, searchCounts[strings.ToLower(word)], &mu)
		}(word)
	}

	go func() {
		wg.Wait()
		close(wordChan)
	}()
	for data := range wordChan {
		result[data.name] = data
	}

	context.JSON(http.StatusOK, result)
}

func CalculateTFAndDF(files [3]string, word string, wordChan chan Word, wordHistory int, mu *sync.Mutex) {
	var wordData Word
	wordData.name = word
	wordData.SearchCountHistory = int64(wordHistory)
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}
		fileTf, containsWord := searchWordInContent(word, string(content))
		mu.Lock()
		wordData.TotalFrequencyInAllFiles += int64(fileTf)
		if containsWord {
			wordData.DocumentFrequencyInAllFiles += 1
		}
		mu.Unlock()
	}
	wordChan <- wordData

}
func searchWordInContent(word, content string) (int, bool) {
	tf := 0
	containsWord := false
	words := strings.Fields(content)

	for _, w := range words {
		if strings.EqualFold(w, word) {
			tf++
			containsWord = true
		}
	}

	return tf, containsWord
}

type Word struct {
	TotalFrequencyInAllFiles    int64
	DocumentFrequencyInAllFiles int64
	SearchCountHistory          int64
	name                        string
}

// Function to write JSON to file
func writeJSONToFile(filePath string, data map[string]int) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, fileData, 0644)
}

func readJSONFromFile(filePath string) (map[string]int, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var searchCounts map[string]int
	if err := json.Unmarshal(file, &searchCounts); err != nil {
		return nil, err
	}

	return searchCounts, nil
}
