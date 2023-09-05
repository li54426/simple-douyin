package controller

import (
	//"fmt"
	"simple-douyin/models"
)



type UserController struct{
    // UserDao *models.UserDao
}

// type UserController interface{
//     CreateUser(user *models.User) (int64, error)
//     FindUserByName(username string)  (*models.User, error) 
//     FindUserById(id int64) (*models.User, error) 
// }



// type UserController interface{
//     GetUserById(id int64) (User, error)
// }



func GetUserById(id int64) (User, error){
    userModel, err := models.GetUserDao().GetUserById(id)
    if err != nil {
        return User{}, err
    }

    user := User{
        Id: userModel.UserId,
        Name: userModel.Name ,
        FollowCount : userModel.FollowCount,
        FollowerCount: userModel.FollowCount,
        IsFollow : userModel.BeFollow,
    }

    return user, nil
}





func GetVideoList(list []models.Video) []Video{
    
    
    var res = make([]Video, len(list))
    

    
    for i, v := range list {
        author, _ := GetUserById(v.UserId)

        
        
        res[i] = Video{
            Id: v.VideoId,
            Author:  author,
            PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
            IsFavorite: v.IsFavorite,
            
            
        }
    }
    return res
}




func GetUserByName(name string)(User, error){
    userModel, err := models.GetUserDao().GetUserByName(name)

    if err != nil {
        return User{}, err
    }

    user := User{
        Id: userModel.UserId, 
        Name: userModel.Name , 
        FollowCount : userModel.FollowCount, 
        FollowerCount: userModel.FollowCount,
        IsFollow : userModel.BeFollow,
    }
    return user, nil
    
}

func CreateUser(username string, password string)(int64, error){
    userModel := models.User{
        Name: username,
        Password: password,
    }
    userId , err:= models.GetUserDao().CreateUser(&userModel)

    if err != nil {
        return -1, err
    }

    return userId, nil
}

func CheckUsernamePassword(username string, password string)(int64){
    return models.GetUserDao().CheckUsernamePassword(username, password)

}
    
// type UserController interface{
//     CreateUser(user *models.User) (int64, error)
//     FindUserByName(username string)  (*models.User, error) 
//     FindUserById(id int64) (*models.User, error) 
// }



func GetVideoListById(userId int64, user *User)([]Video, error){
        videoListModel, err := models.GetVideoDao().GetVideoByUserId(userId)

    if err!= nil {
        return []Video{}, err
    }
    
    //user, err := GetUserById(userId)
    // fmt.Printf("list = %+v\n", videoListModel)
    videoList:= make( []Video, len(videoListModel), 3)

    if len( videoListModel) > 0{
           for i, video:= range(videoListModel){
                videoList[i] = Video{
                    Id:video.VideoId,
                    Author:*user,
                    PlayUrl:video.PlayUrl,
                        CoverUrl: video.CoverUrl,
                    FavoriteCount:video.FavoriteCount,
                    CommentCount:video.CommentCount,
                    IsFavorite: video.IsFavorite,
                }
               // fmt.Printf("con list = %+v\n\n", videoList[i])
        
                
            } 
    }
    


    return videoList, nil
    
    
}








// type VideoController interface{
//     CreateVideo(video *models.Video) (*models.Video, error)
//     FindVideoById(id int64) (*models.Video, error)
//     QueryVideoByUserId(userId int64) ([]*models.Video, error)
//     QueryVideo(date *string, limit int) []*models.Video
// }