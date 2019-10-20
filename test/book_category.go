package test

import (
    "time"
)

type BookCategory struct {

    // 书分类id
    Id string `json:"id"`

    // 分类名称
    Name string `json:"name"`

    //  别名
    Alias string `json:"alias"`

    // 状态:0-正常，1-禁用
    Status int `json:"status"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e BookCategory) TableName() string {
    return "BookCategory"
}
