package entities

import "time"

// User user entity
type User struct {
	ID        string     `json:"id" bson:"id"`
	FirstName string     `json:"first_name" bson:"first_name"`
	LastName  string     `json:"last_name" bson:"last_name"`
	Nickname  *string    `json:"nickname,omitempty" bson:"nickname"`
	Age       *int       `json:"age,omitempty" bson:"age"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Marshal, Unmarshal 時に null と ゼロ値を区別する方法
