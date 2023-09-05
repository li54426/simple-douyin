package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "sync/atomic"
    "simple-douyin/models"
    "fmt"
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

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, _ := GenerateToken(username)   //username + password

    

    // 原本 的 usersLoginInfo[]代替了 数据库层
    userModel, _ :=  models.GetUserDao().GetUserByName(username)

    if userModel == nil {
        userid, _ := models.GetUserDao().CreateUser(&models.User{
            Name: username,
            Password: password,
        })
        c.JSON(http.StatusOK, UserLoginResponse{
                Response: Response{StatusCode: 0},
                UserId:   userid,
                Token:    token,
        })
            
    }
    
    
    
	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	// 	})
	// } else {
	// 	atomic.AddInt64(&userIdSequence, 1)
	// 	newUser := User{
	// 		Id:   userIdSequence,
	// 		Name: username,
	// 	}
	// 	usersLoginInfo[token] = newUser
        
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 0},
	// 		UserId:   userIdSequence,
	// 		Token:    username + password,
	// 	})
	// }
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

    userid := models.GetUserDao().CheckUsernamePassword(username, password)

    if userid == -1 {
        fmt.Println("uid==-1")
	 	c.JSON(http.StatusOK, UserLoginResponse{
	 		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
        })

        return
    }else{
        token, err :=  GenerateToken(username)
    
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

    username,err := ParseToken(token)

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
