package test

import (
    "time"
)

type BookChapter struct {

    // 书目录id
    Id string `json:"id"`

    // 书id
    BookId string `json:"book_id"`

    // 章节名称
    Chapter string `json:"chapter"`

    // 章节内容
    Content string `json:"content"`

    // 状态:0-正常，1-禁用
    Status int `json:"status"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e BookChapter) TableName() string {
    return "BookChapter"
}
