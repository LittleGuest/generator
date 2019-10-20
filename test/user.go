package test

import (
    "time"
)

type User struct {

    // 
    Id int64 `json:"id"`

    // 
    RoleId int64 `json:"role_id"`

    // 
    Name string `json:"name"`

    // 
    CreatedAt time.Time `json:"created_at"`

    // 
    UpdatedAt time.Time `json:"updated_at"`

}

func (e User) TableName() string {
    return "User"
}
