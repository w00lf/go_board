package main

import "time"

// Post - post for board
type Post struct {
    ID           int
    Title        string  `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
    Body         string  `sql:"type:varchar(100);"`
    ThreadID     int

    CreatedAt    time.Time
    UpdatedAt    time.Time
}

type Thread struct {
    ID           int
    Title        string  `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
    Posts        []Post

    CreatedAt    time.Time
    UpdatedAt    time.Time
}
