package test

import (
    "time"
)

type Link struct {

    // 
    Id string `json:"id"`

    // 标题
    Title string `json:"title"`

    // 链接地址
    Url string `json:"url"`

    // 状态:0-正常，1-禁用
    Status int `json:"status"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e Link) TableName() string {
    return "Link"
}
