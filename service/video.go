package service

import (
    "simple-douyin/models"
)

// 服务端限制
var limVideoNum = 30
func GetVideo(lastTime string) []models.Video{
    return models.GetVideoDao().GetVideo(&lastTime, limVideoNum)
}