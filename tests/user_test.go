package controller

import (
	"CloudSystem/controller"
	"CloudSystem/database"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gavv/httpexpect"
	"github.com/gin-gonic/gin"
)

func initTestDB() {
	var err error
	database.DB, err = sql.Open("postgres", "postgres://postgres:Torm22torm*@localhost:5432/cloudsystem?sslmode=disable")
	if err != nil {
		log.Panicf("Unable to connect to database: %v\n", err)
	}

	// Configure the connection pool
	database.DB.SetMaxOpenConns(10)
	database.DB.SetMaxIdleConns(5)
	database.DB.SetConnMaxLifetime(30 * time.Minute)

	// Verify the connection
	if err := database.DB.Ping(); err != nil {
		log.Fatalf("Unable to verify connection: %v\n", err)
	}
}
func TestLogin(t *testing.T) {
	initTestDB()
	defer database.DB.Close()
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", controller.LoginUser)
	server := httptest.NewServer(router)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	// Test valid login
	e.POST("/login").WithJSON(map[string]interface{}{
		"email":    "moham.ed.22@gmail.com",
		"password": "Password1234*",
	}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().Value("token").String().NotEmpty()

	// Test missing fields
	missingFields := []map[string]interface{}{
		{"email": "moham.ed.22@gmail.com"},
		{"password": "Password1234*"},
		{},
	}
	for _, body := range missingFields {
		e.POST("/login").WithJSON(body).
			Expect().
			Status(http.StatusBadRequest)
	}
}

func TestRegisterUser(t *testing.T) {
	initTestDB()
	defer database.DB.Close()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/user", controller.RegisterUser)
	server := httptest.NewServer(router)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	// Test user creation
	e.POST("/user").WithJSON(map[string]interface{}{
		"email":     "okaysqqqqs.22@gmail.com",
		"password":  "Password1234*",
		"firstName": "Luca",
		"lastName":  "Modric",
		"street":    "zafraan",
		"city":      "alex",
		"country":   "Egyptss",
	}).
		Expect().
		Status(http.StatusCreated)

	// Test user creation with existing email
	e.POST("/user").WithJSON(map[string]interface{}{
		"email":     "moham.ed.22@gmail.com",
		"password":  "Password1234*",
		"firstName": "Luca",
		"lastName":  "Modric",
		"street":    "zafraan",
		"city":      "alex",
		"country":   "Egyptss",
	}).
		Expect().
		Status(http.StatusConflict)

	// Test missing fields
	missingFields := []map[string]interface{}{
		{
			"password":  "Password1234*",
			"firstName": "Luca",
			"lastName":  "Modric",
			"street":    "zafraan",
			"city":      "alex",
			"country":   "Egyptss",
		},
		{
			"email":     "mohamqs.ed.2s2@gmail.com",
			"firstName": "Luca",
			"lastName":  "Modric",
			"street":    "zafraan",
			"city":      "alex",
			"country":   "Egyptss",
		},
		{
			"email":    "mohaqqm.edq.22@gmail.com",
			"password": "Password1234*",
			"lastName": "Modric",
			"street":   "zafraan",
			"city":     "alex",
			"country":  "Egyptss",
		},
		{
			"email":     "mohamqd.ed.2a2@gmail.com",
			"password":  "Password1234*",
			"firstName": "Luca",
			"street":    "zafraan",
			"city":      "alex",
			"country":   "Egyptss",
		},
		{
			"email":     "mohaqsm.ed.22q@gmail.com",
			"password":  "Password1234*",
			"firstName": "Luca",
			"lastName":  "Modric",
			"city":      "alex",
			"country":   "Egyptss",
		},
		{},
	}
	for _, body := range missingFields {
		e.POST("/user").WithJSON(body).
			Expect().
			Status(http.StatusBadRequest)
	}
}
