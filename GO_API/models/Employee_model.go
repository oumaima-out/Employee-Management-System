package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
    Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"` 
    FirstName    string             `json:"first_name,omitempty" bson:"FirstName,omitempty" validate:"required"`
    LastName     string             `json:"last_name,omitempty" bson:"LastName,omitempty" validate:"required"`
    Position     string             `json:"position,omitempty" bson:"Position,omitempty" validate:"required"`
    Email        string             `json:"email,omitempty" bson:"Email,omitempty" validate:"required"`
    Phone        string             `json:"phone,omitempty" bson:"Phone,omitempty" validate:"required"`
    Department   string             `json:"department,omitempty" bson:"Department,omitempty" validate:"required"`
    DateOfHire   string             `json:"date_of_hire,omitempty" bson:"DateOfHire,omitempty" validate:"required"`
}
