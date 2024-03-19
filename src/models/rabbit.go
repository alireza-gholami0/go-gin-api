package models

import "time"

type LogLoginRabbit struct {
	Email string    `json:"email"`
	Time  time.Time `json:"time"`
}
