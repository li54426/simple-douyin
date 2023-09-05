# simple-demo

## 抖音项目服务端简单示例

具体功能内容参考飞书说明文档

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### 功能说明

接口功能不完善，仅作为示例

* 用户登录数据保存在内存中，单次运行过程中有效
* 视频上传后会保存到本地 public 目录中，访问时用 127.0.0.1:8080/static/video_name 即可

### 测试

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试






### 问题

#### 1  用户登陆现实网络问题

原来是没有使用 return 返回了两次 JSON 数据


#### 2 浏览器和app显示不一致
```
// 浏览器显示
{
    "status_code": 0,
    "video_list": [
        {
            "id": 1,
            "author": {
                "id": 1,
                "name": "TestUser"
            },
            "play_url": "https://www.w3schools.com/html/movie.mp4",
            "cover_url": "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
        }
    ]
}
```

app 却不显示

### Dao层

```go

GetUserDao()

```



#### dao层的函数返回的只有 自己定义的类型, 我想让他返回控制器层的类型怎么办


如果你希望 DAO 层的函数返回控制器层的类型，你可以通过定义接口来实现。

首先，在控制器层定义一个接口，该接口包含你希望 DAO 层返回的方法。例如：

```go
type UserController interface {
    CreateUser(user User) error
    GetUserByID(id int) (User, error)
    // 其他方法...
}
```
然后，在 DAO 层中，你可以实现这个接口并返回控制器层的类型。例如：

```go
type UserDAO struct {
    // 数据库连接等相关字段...
}

func (dao *UserDAO) CreateUser(user User) error {
    // 在数据库中创建用户的逻辑...
}

func (dao *UserDAO) GetUserByID(id int) (User, error) {
    // 根据ID从数据库中获取用户的逻辑...
}

```

最后，在控制器层中，你可以使用 DAO 层返回的实现了接口的对象。例如：

```go

func main() {
    dao := &UserDAO{}
    controller := dao // 将DAO层的对象赋值给控制器层的变量
    // 使用控制器层的方法进行业务逻辑处理...
}
```
通过这种方式，你可以在 DAO 层中返回实现了控制器层接口的对象，从而实现 DAO 层向控制器层的类型转换。


在控制器层中实现类型转换。控制器层负责处理请求和响应，它可以调用 DAO 层来获取数据，并将获取到的数据转换为控制器层的类型。
```go
type UserController struct {
    userDao *UserDAO
}

func (controller *UserController) GetUserByID(id int) (UserControllerType, error) {
    user, err := controller.userDao.GetUserByID(id)
    if err != nil {
        return UserControllerType{}, err
    }

    // 进行类型转换
    userControllerType := UserControllerType{
        ID:   user.ID,
        Name: user.Name,
        // 其他字段...
    }

    return userControllerType, nil
}
```

#### 简介





环境

- 需要提前安装`ffmpeg`





发布视频的流程

- 鉴权 ----- 使用别人的接口
- 存储视频   ---- (  demo中 有 )
- 抽取封面并存储  
- 视频 和 封面 上传到 `cdn` (  没有, 因此没有做 )
- 将收到的视频信息( 视频名称, 用户名, 播放地址, 封面的地址 )存储到数据库中   ---- 使用数据库的接口
- 返回消息 ( 发布成功)



因此最先做的工作就是

- 确定资源**存储的位置**, 通过查找**路由**得知, 有这样一个映射`r.Static("/static", "./public")`, 我们访问静态资源的时候, 网址为`xxx/static/fileName.txt`, 通过` c.Request.Host`字段, 来获得域名 : `2f781ee3592dd7a9ff0bbd0007fe40ce-app.1024paas.com`, 加头(协议)加尾(路径 + 文件名), 进行字符串拼接
- 使用`ffmpeg`来抽取某一帧来做封面, 需要设置 其 **环境**





```go
// controller/ publish.go
// 没有其他合作者的提供的接口
func Publish(c *gin.Context) {
    // token, 鉴权
 	
    // 存储
    err := c.SaveUploadedFile(data, saveFile);
    
    // 抽取并存储图片
    Vedio2Jpeg(saveFile, 6)
    
    // 拼接网址
    vedio_url := "https://"+ domain + "/static/"+ finalName
    
    // 放入数据库
    db.creat().....
    
}



```



抽取图片 demo

```go
//  将视频抽取一帧, 转化为流
package examples

import (
	"bytes"
	"fmt"
	"io"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
```



#### 队友的接口

```go
// 鉴权接口
// 参数: token, 返回: username
func ParseToken(token string) (string, error)


// 数据库接口---用户
func (*UserDao) FindUserByName(username string) (*User, error)
func (d *UserDao) FindUserById(id int64) (*User, error) 
_, err := models.NewUserDaoInstance().FindUserByName("qong");

// 数据库接口--- 视频
func (*VideoDao) CreateVideo(video *Video) (*Video, error) 

```





#### 使用了接口之后的伪代码

```go
func Publish(c *gin.Context) {
    // token, 鉴权
 	token := c.PostForm("token")
    user_name, err := ParseToken(token);
    user, err := models.NewUserDaoInstance().FindUserByName(user_name)
    var user_id = user.UserId
    
    // 存储
    err := c.SaveUploadedFile(data, saveFile);
    
    // 抽取并存储图片
    Vedio2Jpeg(saveFile, 6)
    
    // 拼接网址
    vedio_url := "https://"+ domain + "/static/"+ finalName
    
    // 放入数据库
    video1 := models.Video{UserId : user_id, PlayUrl : vedio_url, CoverUrl : img_url , }
    _, err = models.NewVideoDaoInstance().CreateVideo(&video1)
    
}
```







