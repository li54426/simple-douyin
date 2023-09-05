package service

import (
    "simple-douyin/models"
)

// 服务端限制
var limVideoNum = 30
func VideoGet(lastTime string) []models.Video{
    return models.GetVideoDao().QueryVideo(&lastTime, limVideoNum)
}