package test

import (
    "time"
)

type BlogBase struct {

    // 博客id
    Id string `json:"id"`

    // 标签id
    TagId string `json:"tag_id"`

    // 标题
    Title string `json:"title"`

    // 副标题
    SubTitle string `json:"sub_title"`

    // 博客内容
    Content string `json:"content"`

    // 评论数
    Comments int `json:"comments"`

    // 点赞数
    Likes int `json:"likes"`

    // 是否删除:0-否，1-是
    Deleted int `json:"deleted"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e BlogBase) TableName() string {
    return "BlogBase"
}
