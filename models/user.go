package models

import (
	"sync"
    "fmt"
    // "simple-douyin/controller"
    "time"
    // "github.com/jinzhu/gorm"
)

type User struct {
	UserId            int64 `gorm:"primaryKey, autoIncrement"`
	Name              string  
	FollowCount       int64  `gorm:"default:(0)"`
	FollowerCount     int64   `gorm:"default:(0)"`
	Password          string    //`gorm:"default:(-)"`
    BeFollow          bool     `gorm:"default:false"` 
    CreatedAt        time.Time  `gorm:"autoCreateTime"`
    UpdatedAt         time.Time
	// DeletedAt        time.Time
}
	// Avatar          string `gorm:"default:(-)"`
	// BackgroundImage string `gorm:"default:(-)"`
	// Signature string
	// TotalFavorited  int64  `gorm:"default:(-)"`
	// WorkCount       int64  `gorm:"default:(-)"`
	// FavoriteCount   int64  `gorm:"default:(-)"`
	// CreateAt        time.Time
	// DeleteAt        time.Time







type UserDao struct{}

var userDao *UserDao
var userOnce sync.Once

func GetUserDao() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		},
	)
	return userDao
}

/*
*
根据用户名和密码，创建一个新的User，返回UserId
*/
func (*UserDao) CreateUser(user *User) (int64, error) {
	/*user := User{Name: username, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := SqlSession.Create(&user)

	if result.Error != nil {
		return -1, result.Error
	}

	return user.UserId, nil
}





func (*UserDao) CheckUsernamePassword(username string, password string)(int64){
    user := User {Name: username, Password: password}
    result := SqlSession.Where("name =? and password=?", username , password).First(&user)

    err:= result.Error

    if err != nil {
        return -1
    }
    fmt.Println("username=", username, "id=", user.UserId)
    return user.UserId


    
}



/*
* 根据用户名，查找用户实体
*/
func (*UserDao) GetUserByName(username string) (User, error) {
	user := User{Name: username}

	result := SqlSession.Where("name = ?", username).First(&user)
	err := result.Error
	if err != nil {
		return User{}, err
	}
	return user, err
}

/*
* 根据用户id，查找用户实体
*/
func (d *UserDao) GetUserById(id int64) (User, error) {
	user := User{UserId: id}

	result := SqlSession.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		return User{}, err
	}
	return user, err
}
