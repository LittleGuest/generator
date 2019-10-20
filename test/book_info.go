package test

import (
    "time"
)

type BookInfo struct {

    // 书id
    Id string `json:"id"`

    // 书名
    Name string `json:"name"`

    // 作者
    Author string `json:"author"`

    // 描述
    Description string `json:"description"`

    // 书分类id
    CategoryId string `json:"category_id"`

    // 状态:0-正常，1-禁用
    Status int `json:"status"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e BookInfo) TableName() string {
    return "BookInfo"
}
