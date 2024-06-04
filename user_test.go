package main

import (
	"barbershop/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	users  map[uint]models.User
	nextID uint
}

func (db *MockDB) CreateUser(user *models.User) error {
	db.nextID++
	user.ID = db.nextID
	db.users[user.ID] = *user
	return nil
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		CreateUser(c, mockDB)
	})
	return r
}

func CreateUser(c *gin.Context, db *MockDB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	if err := db.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

var mockDB = &MockDB{
	users:  make(map[uint]models.User),
	nextID: 0,
}

func TestCreateUser(t *testing.T) {
	router := SetupRouter()

	t.Run("valid user creation", func(t *testing.T) {
		newUser := models.User{
			Username: "testuser",
			Password: "testpassword",
			IsAdmin:  false,
		}
		jsonValue, _ := json.Marshal(newUser)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var createdUser models.User
		err := json.Unmarshal(w.Body.Bytes(), &createdUser)
		assert.NoError(t, err)
		assert.Equal(t, newUser.Username, createdUser.Username)
		assert.Equal(t, newUser.Password, createdUser.Password)
		assert.Equal(t, newUser.IsAdmin, createdUser.IsAdmin)
		assert.NotEqual(t, 0, createdUser.ID)
	})

	t.Run("missing username or password", func(t *testing.T) {
		newUser := models.User{
			Username: "",
			Password: "testpassword",
		}
		jsonValue, _ := json.Marshal(newUser)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		invalidJSON := `{"username": "testuser", "password": ` // Invalid JSON

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(invalidJSON)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
