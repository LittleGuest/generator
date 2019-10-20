package test

import (
    "time"
)

type Role struct {

    // 
    Id int64 `json:"id"`

    // 
    Name string `json:"name"`

    // 
    CreatedAt time.Time `json:"created_at"`

    // 
    UpdatedAt time.Time `json:"updated_at"`

}

func (e Role) TableName() string {
    return "Role"
}
