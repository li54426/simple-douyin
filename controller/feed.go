package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
    "strconv"
    "simple-douyin/service"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
// 返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多 30 个。
// message douyin_feed_request {
//   optional int64 latest_time = 1; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
//   optional string token = 2； // 可选参数，登录用户设置
// }
// douyin/feed/?latest_time=1693532816&token=eyJhbGc
// 转换以后是
// 2023-09-01 09:46:56
func Feed(c *gin.Context) {
    // 时间精度不用太小
    lstTime := time.Now().Format("2020-01-01 01:01:01")
    lstTimeStr := c.Query("latest_time")

    if lstTimeStr != "" {
        lstTimeNum , _ := strconv.ParseInt(lstTimeStr, 10, 64)

        lstTime = time.Unix(lstTimeNum, 0).Format("2023-09-01 09:46:56")
    }

    videoList:= service.VideoGet(lstTime)
    


    
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
