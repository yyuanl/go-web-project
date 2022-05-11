package services

import (
	"errors"
	"fmt"
	"web_project/dao"
	"web_project/models"
)

// JudgeUserIsExist 判断用户是否存在
func JudgeUserIsExist(userName string, passWord string) (bool, error) {
	res, err := dao.GetUserInfo(&models.User{
		UserName: &userName,
		PassWord: &passWord,
	})
	if err != nil {
		return false, errors.New(fmt.Sprintf("查询用户信息错误：%+v", err))
	}
	if len(res) != 1 {
		return false, errors.New(fmt.Sprintf("没有找到唯一用户：%d", len(res)))
	}
	return true, nil
}
