package controllers

import (
    "context"
    "GO_API/configs"
    "GO_API/models"
    "GO_API/responses"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "go.mongodb.org/mongo-driver/bson"
)

var employeeCollection *mongo.Collection = configs.GetCollection(configs.DB, "employees")
var validate = validator.New()

func CreateEmployee() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var employee models.Employee
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&employee); err != nil {
            c.JSON(http.StatusBadRequest, responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&employee); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        newEmployee := models.Employee{
            Id:           primitive.NewObjectID(),
            FirstName:     employee.FirstName,
			LastName:     employee.LastName,
            Position:    employee.Position,
			Email:     employee.Email,
			Phone:     employee.Phone,
			Department:     employee.Department,
			DateOfHire:     employee.DateOfHire,
        }

        result, err := employeeCollection.InsertOne(ctx, newEmployee)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.EmployeeResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
    }
}


func GetAEmployee() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        employeeId := c.Param("id") // Make sure the URL param matches the actual URL parameter (case-sensitive)
        var employee models.Employee
        defer cancel()

        // Convert the string ID to an ObjectID for MongoDB query
        objId, err := primitive.ObjectIDFromHex(employeeId)
        if err != nil {
            c.JSON(http.StatusBadRequest, responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid ID format"}})
            return
        }

        // Query MongoDB for the employee using the ObjectId
        err = employeeCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&employee)
        if err != nil {
            if err == mongo.ErrNoDocuments {
                c.JSON(http.StatusNotFound, responses.EmployeeResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Employee not found"}})
            } else {
                c.JSON(http.StatusInternalServerError, responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
            return
        }

        // Return the employee data
        c.JSON(http.StatusOK, responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": employee}})
    }
}


func EditAEmployee() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        employeeId := c.Param("id") // Ensure the URL parameter is lowercase
        var employee models.Employee
        defer cancel()

        // Convert the employeeId to ObjectId
        objId, err := primitive.ObjectIDFromHex(employeeId)
        if err != nil {
            c.JSON(http.StatusBadRequest, responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid ID format"}})
            return
        }

        // Validate the request body
        if err := c.BindJSON(&employee); err != nil {
            c.JSON(http.StatusBadRequest, responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}) 
            return
        }

        // Validate required fields
        if validationErr := validate.Struct(&employee); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
            return
        }

        // Prepare the update fields
        update := bson.M{
            "FirstName": employee.FirstName,
            "LastName":  employee.LastName,
            "Position":  employee.Position,
            "Email":     employee.Email,
            "Phone":     employee.Phone,
            "DateOfHire": employee.DateOfHire,
            "Department": employee.Department,
        }

        // Perform the update
        result, err := employeeCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        // Get the updated employee details
        var updatedEmployee models.Employee
        if result.MatchedCount == 1 {
            err := employeeCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedEmployee)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
                return
            }
        }

        c.JSON(http.StatusOK, responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedEmployee}})
    }
}



func DeleteAEmployee() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        employeeId := c.Param("id") // Ensure the URL parameter is lowercase
        defer cancel()

        // Convert the employeeId to ObjectId
        objId, err := primitive.ObjectIDFromHex(employeeId)
        if err != nil {
            c.JSON(http.StatusBadRequest, responses.EmployeeResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid ID format"}})
            return
        }

        // Perform the delete operation
        result, err := employeeCollection.DeleteOne(ctx, bson.M{"_id": objId})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        // Handle the case when no documents are deleted
        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound, responses.EmployeeResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Employee with specified ID not found!"}})
            return
        }

        c.JSON(http.StatusOK, responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Employee successfully deleted!"}})
    }
}


func GetAllEmployees() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var employees []models.Employee
        defer cancel()

        results, err := employeeCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleEmployee models.Employee
            if err = results.Decode(&singleEmployee); err != nil {
                c.JSON(http.StatusInternalServerError, responses.EmployeeResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
            }
        
            employees = append(employees, singleEmployee)
        }

        c.JSON(http.StatusOK,
            responses.EmployeeResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": employees}},
        )
    }
}

