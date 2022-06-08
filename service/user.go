package service

import (
	"gin-project/global"
	"gin-project/model"
	"time"
)

func QueryUserByUsername(username string) (user model.User, notFound bool) {
	notFound = global.DB.Where("username = ?", username).First(&user).RecordNotFound()
	return user, notFound
}
func QueryUserByEmail(email string) (user model.User, notFound bool) {
	notFound = global.DB.Where("username = ?", email).First(&user).RecordNotFound()
	return user, notFound
}
func QueryUserByUserID(userID uint64) (user model.User, notFound bool) {
	notFound = global.DB.Where("user_id = ?", userID).First(&user).RecordNotFound()
	return user, notFound
}
func UpdateUser(user *model.User) error {
	err := global.DB.Save(user).Error
	return err
}
func CreateUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return
}
func GetBannedUsers() (users []model.User, err error) {
	err = global.DB.Where("bandate > ?", time.Now()).Find(&users).Error
	return users, err
}

var MAXLEVEL = 10
var EXPGAP = []int{0, 30, 60, 150, 450, 1050, 2100, 3600, 5500, 10500}

func UpdateUserExp(userID uint64, exp int) {
	user, _ := QueryUserByUserID(userID)
	user.Exp += exp
	if exp >= 0 {
		for i := user.UserLevel; i < MAXLEVEL; i++ {
			if user.Exp >= EXPGAP[i] {
				user.UserLevel = i + 1
				user.Exp -= EXPGAP[i]
			} else {
				break
			}
		}
	} else {
		for i := user.UserLevel; i > 0; i-- {
			if user.Exp < 0 {
				user.UserLevel = i - 1
				user.Exp += EXPGAP[i-1]
			} else {
				break
			}
		}
	}
	UpdateUser(&user)
}
