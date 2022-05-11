package dao

import "web_project/models"

// GetUserInfo 查询用户信息
func GetUserInfo(cond *models.User) ([]*models.User, error) {
	if cond == nil {
		return nil, nil
	}
	var res []*models.User
	query := DbInstance.UserAccountManagerDB.Table("user_info")
	if cond.UserName != nil {
		query = query.Where("user_name = ?", *cond.UserName)
	}
	if cond.PassWord != nil {
		query = query.Where("pass_word = ?", *cond.PassWord)
	}
	err := query.Find(&res).Error
	return res, err
}

// InsertBatchUserAccount 批量插入用户
func InsertBatchUserAccount(users []*models.User) error {
	if users == nil {
		return nil
	}
	db := DbInstance.UserAccountManagerDB.Table("user_info")
	err := db.Create(&users).Error
	return err
}
