package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"GO_API/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setupTestRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/employees", CreateEmployee())
	router.GET("/employees/:id", GetAEmployee())
	router.PUT("/employees/:id", EditAEmployee())
	router.DELETE("/employees/:id", DeleteAEmployee())
	router.GET("/employees", GetAllEmployees())
	return router
}

func clearTestDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Clear the employees collection before each test
	_, err := employeeCollection.DeleteMany(ctx, primitive.M{})
	if err != nil {
		panic(fmt.Sprintf("Failed to clear database: %v", err))
	}
}

func createTestEmployee() models.Employee {
	return models.Employee{
		Id:         primitive.NewObjectID(),
		FirstName:   "John",
		LastName:    "Doe",
		Position:    "Developer",
		Email:       "john.doe@example.com",
		Phone:       "1234567890",
		Department:  "IT",
		DateOfHire:  "2023-01-01",
	}
}

func TestCreateEmployee(t *testing.T) {
	clearTestDatabase()
	router := setupTestRouter()

	employee := createTestEmployee()
	jsonData, _ := json.Marshal(employee)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["message"])
}

func TestGetEmployee(t *testing.T) {
	clearTestDatabase()
	router := setupTestRouter()

	// First create an employee
	employee := createTestEmployee()
	insertResult, err := employeeCollection.InsertOne(context.Background(), employee)
	assert.NoError(t, err)

	employeeID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employees/"+employeeID, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["message"])
}

func TestUpdateEmployee(t *testing.T) {
	clearTestDatabase()
	router := setupTestRouter()

	// First create an employee
	employee := createTestEmployee()
	insertResult, err := employeeCollection.InsertOne(context.Background(), employee)
	assert.NoError(t, err)

	employeeID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	// Update employee details
	updatedEmployee := employee
	updatedEmployee.FirstName = "Jane"
	updatedEmployee.LastName = "Smith"

	jsonData, _ := json.Marshal(updatedEmployee)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/employees/"+employeeID, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["message"])

	// Verify the update
	var updatedDoc models.Employee
	err = employeeCollection.FindOne(context.Background(), primitive.M{"_id": insertResult.InsertedID}).Decode(&updatedDoc)
	assert.NoError(t, err)
	assert.Equal(t, "Jane", updatedDoc.FirstName)
	assert.Equal(t, "Smith", updatedDoc.LastName)
}

func TestDeleteEmployee(t *testing.T) {
	clearTestDatabase()
	router := setupTestRouter()

	// First create an employee
	employee := createTestEmployee()
	insertResult, err := employeeCollection.InsertOne(context.Background(), employee)
	assert.NoError(t, err)

	employeeID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/employees/"+employeeID, nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["message"])
}

func TestGetAllEmployees(t *testing.T) {
	clearTestDatabase()
	router := setupTestRouter()

	// Insert multiple employees
	employees := []models.Employee{
		createTestEmployee(),
		{
			Id:         primitive.NewObjectID(),
			FirstName:   "Alice",
			LastName:    "Johnson",
			Position:    "Manager",
			Email:       "alice.johnson@example.com",
			Phone:       "0987654321",
			Department:  "HR",
			DateOfHire:  "2022-05-15",
		},
	}

	for _, emp := range employees {
		_, err := employeeCollection.InsertOne(context.Background(), emp)
		assert.NoError(t, err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employees", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response["message"])

	data := response["data"].(map[string]interface{})["data"].([]interface{})
	assert.Len(t, data, 2)
}