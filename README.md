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





### 结构设计

#### 项目目录结构

其中：

- **configs** 目录包含项目的配置文件
- **controllers** 目录包含**控制器**文件
- **middleware** 目录包含**中间件**文件
- **models** 目录包含模型（ORM）文件
- **routes** 目录包含**路由**文件
- **services** 目录包含服务文件
- **utils** 目录包含多个工具文件
- **main.go** 是项目的入口文件，`README.md` 是项目的介绍文件。

```
├── configs
│   ├── config.yaml
│   └── db.yaml
├── controllers
│   ├── auth_controller.go
│   └── user_controller.go
├── middleware
│   ├── auth_middleware.go
│   └── logger_middleware.go
├── models
│   ├── db.go
│   ├── user.go
│   └── ...
├── routes
│   └── routes.go
├── services
│   ├── auth_service.go
│   └── user_service.go
├── utils
│   ├── response.go
│   └── ...
├── main.go
└── README.md
```



#### MVC架构

在传统应用程序中，我们通常采用经典的 MVC（Model-View-Controller）架构进行开发，它将整体的系统分成了 Model（模型），View（视图）和 Controller（控制器）三个层次，也就是将用户视图和业务处理隔离开，并且通过控制器连接起来，很好地实现了表现和逻辑的解耦，是一种标准的软件分层架构。

在遵循此分层架构的开发过程中，我们通常会建立三个 Maven Module, **从高到低**依次是：Controller、Service 和 Dao，它们分别对应表现层、逻辑层和数据访问层

在业务逻辑较为简单的应用中，MVC 三层架构是一种简洁高效的开发模式。然而，随着业务逻辑的复杂性增加和代码量的增加，MVC 架构可能会显得捉襟见肘。其主要的不足可以总结如下：

- **Service 层职责过重**：在 MVC 架构中，Service 层常常被赋予处理复杂业务逻辑的任务。随着业务逻辑的增长，Service 层可能变得臃肿和复杂。业务逻辑有可能分散在各个 Service 类中，使得业务逻辑的组织和维护成为一项挑战。
- **过于关注数据库而忽视领域建模**：虽然 MVC 的设计初衷是对数据、用户界面和控制逻辑进行分离，但它在面对复杂业务场景时并未给予领域建模足够的重视。这可能导致代码难以理解和扩展，因为代码更像是围绕数据库而不是业务需求进行设计。
- **边界划分不明确**：在 MVC 架构中，顶层设计上的边界划分并没有明确的规则，往往依赖于技术负责人的经验。在大规模的团队协作中，这可能导致职责不清晰、分工不明确等问题。
- **单元测试困难**：在 MVC 架构中，Service 层通常以事务脚本的方式进行开发，并且往往耦合了各种中间件操作，如数据库、缓存、消息队列等。这种耦合使得单元测试变得困难，因为要在没有这些中间件的情况下运行测试可能需要大量的模拟或存根代码。





#### DDD架构（Domain-Driven Design，领域驱动设计）

在 DDD 中，通常将应用程序分为四个层次，分别为**用户接口层（Interface Layer）**，**应用层（Application Layer）**， **领域层（Domain Layer）**，**基础设施层（Infrastructure Layer）**，每个层次承担着各自的职责和作用。分层模型如下图所示：

1. **接口层（Interface Layer）**：负责处理与外部系统的**交互**，包括 UI、Web API、RPC 接口等。它会接收用户或外部系统的请求，然后调用应用层的服务来处理这些请求，最后将处理结果返回给用户或外部系统。
2. **应用层（Application Layer）**：承担协调领域层和基础设施层的职责，实现具体的业务逻辑。它调用领域层的领域服务和基础设施层的基础服务，完成业务逻辑的实现。
3. **领域层（Domain Layer）**：该层包含了业务领域的所有元素，如实体、值对象、领域服务、聚合、工厂和领域事件等。这一层的主要职责是实现业务领域的核心逻辑。
4. **基础设施层（Infrastructure Layer）**：主要提供通用的技术能力，如**数据持久化**、缓存、消息传输等基础设施服务。它可被其他三层调用，提供各种必要的技术服务。









#### 架构改进 

```go
// 鉴权接口, controller
// 参数: token, 返回: username
func ParseToken(token string) (string, error)
```

改进

```go
// 鉴权接口, service层
// 参数: token, 返回: username
func ParseToken(token string) (string, error)
```



#### 所有对数据库的操作, 应该都由 `service`来完成

```go
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, _ := GenerateToken(username)   //username + password

    

    // 原本 的 usersLoginInfo[]代替了 数据库层
    userModel, err :=  models.GetUserDao().GetUserByName(username)

    if err == nil {
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
    
    
  
}

// 返回值类型
// userid   success
// -1       用户已经存在
// -2       创建失败
func CanRegister(username string, password string)int64{
    _, err := models.GetUserDao().GetUserByName(username)

    if err == nil  { // 用户已经存在
        return -1;
    }else{// 用户不存在
        return CreateUser(username, password)
        
    }
    
}
```
















### 问题

#### 1  用户登陆显示网络问题

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





#### 在设计dao层时, 有必要将返回值设置为指针么

如果 DAO 层返回的数据对象是一个复杂的对象，并且您希望在调用 DAO 方法时直接修改该对象的值，那么将返回值设置为指针可能是一个好的选择。这样，您可以直接操作该指针指向的内存地址，从而修改对象的值。

但是，如果您不需要直接修改对象的值，或者您不确定是否需要修改对象的值，那么将返回值设置为值类型可能更合适。这样，您可以直接使用返回值，而不需要担心指针的问题。

总之，将返回值设置为指针还是值类型取决于您的具体需求和设计要求。在设计 DAO 层时，您应该考虑到可能的场景，并选择最适合您的设计方案。



#### 返回值设计

- `nil `表示一个**指针**或**接口**类型的零值或空值。
- 当返回值为一个结构体时, 需要返回 `error`, 若返回值为指针, 那么, 可以不返回`error`

要解决这个错误，你需要确保在返回语句中使用的是一个有效的 User 值，而不是空值。这可以通过以下几种方式处理：

1. 检查你的代码，确保创建了一个有效的 User 对象，并将其赋值给返回值。
2. 如果有必要，可以使用指针类型来表示 User，然后在返回语句中使用该指针。这样，如果 User 对象尚未初始化，你可以将指针设置为 nil 值。
3. 如果你的函数定义中明确指定了返回类型为指针类型，那么你可以直接返回 nil，而无需返回一个有效的 User 对象。







#### dao接口

- 数据库返回的只有几个字段, 直接返回**所有字段**, 用什么自己查就好

```go
GetUserDao()

// 数据库接口---用户
func (*UserDao) CreateUser(user *User) (int64, error)
func (*UserDao) GetUserByName(username string) (model.User, error)
func (d *UserDao) GetUserById(id int64) (model.User, error) 

// 使用方式
_, err := models.GetUserDao().GetUserByName("qong");

// 数据库接口--- 视频
func (*VideoDao) CreateVideo(video Video) (model.Video, error) 
func (*VideoDao) GetVideoByUserId(userId int64) ([]model.Video, error) 
func (*VideoDao) GetVideo(date *string, limit int) []model.Video 

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



#### 接口

```go
// 鉴权接口
// 参数: token, 返回: username
func ParseToken(token string) (string, error)



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







