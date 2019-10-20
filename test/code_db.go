package test

import (
    "time"
)

type CodeDb struct {

    // ID
    Id int64 `json:"id"`

    // 数据库类型
    Driver string `json:"driver"`

    // 数据库主机
    Host string `json:"host"`

    // 数据库端口
    Port int `json:"port"`

    // 数据库名称
    DbName string `json:"db_name"`

    // 用户名
    Username string `json:"username"`

    // 密码
    Password string `json:"password"`

    // 其他参数
    Extra string `json:"extra"`

}

func (e CodeDb) TableName() string {
    return "CodeDb"
}
