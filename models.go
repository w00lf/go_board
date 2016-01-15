package main

import (
    "time"
    "w00lf/go_board/Godeps/_workspace/src/github.com/jinzhu/gorm"
)

// Post - post for board
type Post struct {
    gorm.Model
    ID           int
    Title        string  `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
    Body         string  `sql:"type:varchar(100);"`
    Posts        []Post
    PostID       int
    BoardID      int

    CreatedAt    time.Time
    UpdatedAt    time.Time
}

// Board - board has_many posts
type Board struct {
    gorm.Model
    ID           int
    Title        string  `sql:"size:255"` // Default size for string is 255, you could reset it with this tag
    Body         string  `sql:"type:varchar(100);"`
    Posts        []Post

    CreatedAt    time.Time
    UpdatedAt    time.Time
}
