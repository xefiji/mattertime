package main

import (
"gopkg.in/mgo.v2/bson"
"time"
)

// type TimeSpent struct {
// 	ID int `json:"id"`
// 	Date string `json:"date,omitempty"`
// 	Spent float64 `json:"spent,omitempty"`
// 	Task string `json:"tash,omitempty"`
// 	User string `json:"user,omitempty"`
//     // Address   	*Address `json:"address,omitempty"`
// }

type TimeSpent struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Date string `bson:"date" json:"date,omitempty"`
	Spent float64 `bson:"spent" json:"spent,omitempty"`
	Task string `bson:"task" json:"task,omitempty"`
	User string `bson:"user" json:"user,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at,omitempty"`
}

type TimeSpents []TimeSpent