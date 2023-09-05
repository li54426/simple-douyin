package test


import(
    "testing"
    "fmt"
    "simple-douyin/models"
    
)


func TestDao(t *testing.T){
    models.InitProject()
    userDao := models.GetUserDao()
    
    uid, err := userDao.CreateUser(&models.User{Name: "test", 
                      Password:"test",                   })

    if err != nil {
        fmt.Println("err")
    }else{
        fmt.Println("insert success")
        fmt.Println("uid=", uid)
    }

    
}