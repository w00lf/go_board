package main

import "time"

// Post - post for board
type Post struct {
    ID           int
    Title        string  `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
    Body         string  `sql:"type:varchar(100);"`

    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    *time.Time
}

