package models

import "time"

type User struct {
	Id         *int32     `json:"id"`
	UserName   *string    `json:"user_name"`
	PassWord   *string    `json:"pass_word"`
	Status     *int       `json:"status"`
	CreateTime *time.Time `json:"create_time"`
	UpdateTime *time.Time `json:"update_time"`
}
