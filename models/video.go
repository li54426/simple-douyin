package models

import (
	"fmt"
	"sync"
	"time"
)

type Video struct {
	VideoId       int64 `gorm:"primaryKey"`
	UserId        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
    IsFavorite    bool
	// Title    string `gorm:"default:(-)"`
	CreateAt time.Time
	// UpdateAt time.Time
	// DeleteAt time.Time
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once //单例模式，只生成一个VideoDao实例，提高性能

func GetVideoDao() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}



/*
*
根据创建一个新的Video，返回Video实例
*/
func (*VideoDao) CreateVideo(video *Video) (*Video, error) {
	/*Video := Video{Name: Videoname, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := SqlSession.Create(&video)

	if result.Error != nil {
		return nil, result.Error
	}

	return video, nil
}





/*
*
根据videoid，查找video实体
*/
func (d *VideoDao) GetVideoById(id int64) (Video, error) {
	video := Video{VideoId: id}

	result := SqlSession.Where("user_id = ?", id).First(&video)
	err := result.Error
	if err != nil {
		return Video{}, err
	}
	return video, err
}

// 根据UserId，查出Video列表
func (*VideoDao) GetVideoByUserId(userId int64) ([]Video, error) {
	var videos []Video
	err := SqlSession.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		fmt.Println("查询Video列表失败")
		return nil, err
	}
    // fmt.Printf("dao list = %+v\n", videos)
	return videos, nil
}

// 根据时间和需要查询的条数，获取video列表
func (*VideoDao) GetVideo(date *string, limit int) []Video {
	fmt.Println("date=", *date)
	var VideoList = make([]Video, limit)
	SqlSession.Where("create_at < ?", *date).Order("create_at desc").Find(&VideoList)
	if len(VideoList) <= limit {
		fmt.Println(VideoList)
		return VideoList
	}
	// fmt.Println(VideoList)
	return VideoList
}
