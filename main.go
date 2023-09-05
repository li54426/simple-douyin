package main

import (
	"simple-douyin/models"
	"simple-douyin/service"
	"github.com/gin-gonic/gin"
    //"fmt"
)




// /douyin/user/register/?username=1234567&password=1234567



func main() {
	go service.RunMessageServer()

	// 初始化配置，如MySQL等
	err := models.InitProject()


    // ----------testing-----
    // userDao := models.GetUserDao()
    
    // uid, err := userDao.CreateUser(&models.User{Name: "testing", 
    //                   Password:"testing",                   })
    // uid2, err :=userDao.CreateUser(&models.User{Name: "testing", 
    //                   Password:"testing",                   })



    
    // if err != nil {
    //     fmt.Println("err")
    // }else{
    //     fmt.Println("insert success")
    //     fmt.Println("uid=", uid)
    // }

    // userModel1, err := userDao.GetUserById(uid)
    
    // userModel2, err := userDao.GetUserById(uid2)
    // fmt.Printf("%v", userModel1)
    // fmt.Printf("%v", userModel2)


    // return





    

    
	if err != nil {
		panic(err)
	}
	defer models.Close()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
