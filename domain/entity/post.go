package entity

import (
	"time"
)

type Post struct {
	Id        string
	CreatedAt time.Time
	Level     int32
	Content   string
}
