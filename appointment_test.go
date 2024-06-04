package main

// import (
// 	"barbershop/controllers"
// 	"barbershop/models"
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// type MockDB struct {
// 	appointments map[uint]models.Appointment
// }

// func (db *MockDB) FindAppointmentByID(id uint) (models.Appointment, bool) {
// 	appointment, exists := db.appointments[id]
// 	return appointment, exists
// }

// func (db *MockDB) SaveAppointment(appointment models.Appointment) error {
// 	db.appointments[appointment.ID] = appointment
// 	return nil
// }

// // Mocked database
// var mockDB = &MockDB{
// 	appointments: map[uint]models.Appointment{
// 		1: {ID: 1, Username: "testuser", Time: "2024-05-27 15:30:00", Master: "testmaster"},
// 	},
// }

// func setupRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.PUT("/appointments/:id", controllers.UpdateAppointment)
// 	return r
// }

// func TestUpdateAppointment(t *testing.T) {
// 	router := setupRouter()

// 	t.Run("valid update", func(t *testing.T) {
// 		appointmentID := uint(1)
// 		newAppointment := models.Appointment{
// 			ID:       appointmentID,
// 			Username: "updateduser",
// 			Time:     "2024-05-28 16:00:00",
// 			Master:   "updatedmaster",
// 		}
// 		jsonValue, _ := json.Marshal(newAppointment)

// 		req, _ := http.NewRequest("PUT", "/appointments/"+strconv.Itoa(int(appointmentID)), bytes.NewBuffer(jsonValue))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		router.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusOK, w.Code)

// 		var updatedAppointment models.Appointment
// 		err := json.Unmarshal(w.Body.Bytes(), &updatedAppointment)
// 		assert.NoError(t, err)
// 		assert.Equal(t, newAppointment.Username, updatedAppointment.Username)
// 		assert.Equal(t, newAppointment.Time, updatedAppointment.Time)
// 		assert.Equal(t, newAppointment.Master, updatedAppointment.Master)
// 	})

// 	// Test case: invalid ID
// 	t.Run("invalid ID", func(t *testing.T) {
// 		req, _ := http.NewRequest("PUT", "/appointments/invalid", nil)
// 		w := httptest.NewRecorder()
// 		router.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})

// 	t.Run("appointment not found", func(t *testing.T) {
// 		appointmentID := uint(1)
// 		newAppointment := models.Appointment{
// 			ID:       appointmentID,
// 			Username: "updateduser",
// 			Time:     "2024-05-28 16:00:00",
// 			Master:   "updatedmaster",
// 		}
// 		jsonValue, _ := json.Marshal(newAppointment)

// 		req, _ := http.NewRequest("PUT", "/appointments/"+strconv.Itoa(int(appointmentID)), bytes.NewBuffer(jsonValue))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		router.ServeHTTP(w, req)

// 		assert.Equal(t, http.StatusNotFound, w.Code)
// 	})
// }
