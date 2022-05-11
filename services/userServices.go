package services

import (
	"errors"
	"fmt"
	"time"
	"web_project/dao"
	"web_project/models"
)

// JudgeUserIsExist 判断用户是否存在
func JudgeUserIsExist(userName string) (bool, error) {
	res, err := dao.GetUserInfo(&models.User{
		UserName: &userName,
	})
	if err != nil {
		return false, errors.New(fmt.Sprintf("查询用户信息错误：%+v", err))
	}
	if len(res) == 0 {
		return false, nil
	}
	return true, nil
}

// CreateOneUserAccount 创建一个用户
func CreateOneUserAccount(userName string, passWord string) error {
	if userName == "" || passWord == "" {
		return nil
	}
	now := time.Now()
	status := 1
	user := &models.User{
		UserName:   &userName,
		PassWord:   &passWord,
		CreateTime: &now,
		UpdateTime: &now,
		Status:     &status,
	}
	return dao.InsertBatchUserAccount([]*models.User{user})
}
