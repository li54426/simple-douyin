package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "sync/atomic"
	// "fmt"
	"simple-douyin/models"
	"simple-douyin/service"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}


/**
* handler 
* 注册功能
* 
*/
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
    // 原本 的 usersLoginInfo[]代替了 数据库层
    userId := service.CanRegister(username, password)

    if userId == -1 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
            
    } else if userId == -2 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User creat fail"},
		})
        
    }else {
        token, _ := service.GenerateToken(username) 
        c.JSON(http.StatusOK, UserLoginResponse{
                Response: Response{StatusCode: 0},
                UserId:   userId,
                Token:    token,
        })
    }
    
    
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

    // userid := models.GetUserDao().CheckUsernamePassword(username, password)
    userid := service.CanLogin(username, password)
    
    if userid == -1 {
	 	c.JSON(http.StatusOK, UserLoginResponse{
	 		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
        })

    }else if userid == -2{
        c.JSON(http.StatusOK, UserLoginResponse{
	 		Response: Response{StatusCode: 1, StatusMsg: "wrong password"},
        })
    }else {
        token, err :=  service.GenerateToken(username)
        if err == nil {
            c.JSON(http.StatusOK, UserLoginResponse{
                   Response: Response{StatusCode: 0},
                   UserId: userid,
                   Token: token,
                  })
        }else {
            c.JSON(http.StatusOK, UserLoginResponse{
                Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
            })
        }  
    }




    
	//token := username + password

	// if user, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 0},
	// 		UserId:   user.Id,
	// 		Token:    token,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	// 	})
	// }
}







func UserInfo(c *gin.Context) {
	token := c.Query("token")

    username,err := service.ParseToken(token)

    if err != nil {
        c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
    }else{
        userModel, err := models.GetUserDao().GetUserByName(username)

        if err == nil{
            c.JSON(http.StatusOK, UserResponse{
                Response: Response{StatusCode: 0},
                User: User{userModel.UserId, userModel.Name, 0, 0, false},
            })            
        }

        
    }

	// if user, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 0},
	// 		User:     user,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	// 	})
	// }
}
