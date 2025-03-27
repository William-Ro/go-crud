package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// Employee Constructs your Employee model under entities.
type Employee struct {
	ID         bson.ObjectID `json:"id"  bson:"_id,omitempty"`
	First_Name string        `json:"first_name" bson:"first_name"`
	Last_Name  string        `json:"last_name" bson:"last_name"`
	Email      string        `json:"email" bson:"email"`
	CreatedAt  time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt" bson:"updatedAt"`
}

// DeleteRequest struct is used to parse Delete Requests for Employees
type DeleteRequest struct {
	ID string `json:"id"`
}
