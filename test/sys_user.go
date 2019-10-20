package test

import (
    "time"
)

type SysUser struct {

    // 用户id
    Id string `json:"id"`

    // 账号
    Account string `json:"account"`

    // 密码
    Password string `json:"password"`

    // 手机号
    Mobile string `json:"mobile"`

    // 微信openid
    OpenId string `json:"open_id"`

    // 昵称
    NickName string `json:"nick_name"`

    // 头像
    AvatarUrl string `json:"avatar_url"`

    // 性别
    Gender int `json:"gender"`

    // 年龄
    Age int `json:"age"`

    //  简介
    Synopsis string `json:"synopsis"`

    // 详细介绍
    Introduce string `json:"introduce"`

    // 状态:0-正常，1-禁用
    Status int `json:"status"`

    // 创建时间
    CreatedAt time.Time `json:"created_at"`

    // 更新时间
    UpdatedAt time.Time `json:"updated_at"`

}

func (e SysUser) TableName() string {
    return "SysUser"
}
