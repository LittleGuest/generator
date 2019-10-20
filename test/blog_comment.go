package test

import (
    "time"
)

type BlogComment struct {

    // 博客评论id
    Id string `json:"id"`

    // 博客id
    BlogId string `json:"blog_id"`

    // 评论内容
    Content string `json:"content"`

    // 是否删除:0-否，1-是
    Deleted int `json:"deleted"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e BlogComment) TableName() string {
    return "BlogComment"
}
