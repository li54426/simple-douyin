package service

import (
	"simple-douyin/models"
)

// 返回值类型
// userid   success
// -1       用户已经存在
// -2       创建失败
func CanRegister(username string, password string)int64{
    _, err := models.GetUserDao().GetUserByName(username)

    if err == nil  { // 用户已经存在
        return -1;
    }else{// 用户不存在
        return CreateUser(username, password)
        
    }
    
}



func CreateUser(username string, password string) int64{
    userid, err := models.GetUserDao().CreateUser(&models.User{
            Name: username,
            Password: password,
    })
    if err != nil {
        return -2;
    }else{
        return userid;
    }

    
}



// -1     dont exist
//  -2  wrong password
func CanLogin(username string, password string) int64 {
    userModel, err := models.GetUserDao().GetUserByName(username)
    // user := GetUserByName(username)

    if err != nil{
        return -1
    }else if userModel.Password != password {
        return -2
    }else{
        return userModel.UserId      
    }
}