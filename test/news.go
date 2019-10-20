package test

import (
    "time"
)

type News struct {

    // 新闻id
    Id string `json:"id"`

    // 来源
    Source string `json:"source"`

    //  作者
    Author string `json:"author"`

    //  责任主编
    Editor string `json:"editor"`

    // 标题
    Title string `json:"title"`

    // 新闻内容
    Content string `json:"content"`

    // 状态:0-正常，1-禁用
    Status int `json:"status"`

    // 发布时间
    PushTime time.Time `json:"push_time"`

}

func (e News) TableName() string {
    return "News"
}
