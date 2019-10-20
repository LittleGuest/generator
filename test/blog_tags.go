package test

import (
    "time"
)

type BlogTags struct {

    // ID
    Id string `json:"id"`

    // 标签名称
    Name string `json:"name"`

    // 标签别名
    Alias string `json:"alias"`

    // 字体大小
    FontSize string `json:"font_size"`

    // 颜色
    Color string `json:"color"`

    // 父级ID
    Pid string `json:"pid"`

    // 是否删除:0-否，1-是
    Deleted int `json:"deleted"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e BlogTags) TableName() string {
    return "BlogTags"
}
