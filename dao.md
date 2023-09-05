


#### 原本的架构如下

```go

type UserDao struct{}

var userDao *UserDao
var userOnce sync.Once

func GetUserDao() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		},
	)
	return userDao
}

/*
*
根据用户名和密码，创建一个新的User，返回UserId
*/
func (*UserDao) CreateUser(user *User) (int64, error) {
	/*user := User{Name: username, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := SqlSession.Create(&user)

	if result.Error != nil {
		return -1, result.Error
	}

	return user.UserId, nil
}





func (*UserDao) CheckUsernamePassword(username string, password string)(int64){
    user := User {Name: username, Password: password}
    result := SqlSession.Where("name =? and password=?", username , password).First(&user)

    err:= result.Error

    if err != nil {
        return -1
    }
    fmt.Println("username=", username, "id=", user.UserId)
    return user.UserId


    
}



/*
*
根据用户名，查找用户实体
*/
func (*UserDao) GetUserByName(username string) (*User, error) {
	user := User{Name: username}

	result := SqlSession.Where("name = ?", username).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

/*
*
根据用户id，查找用户实体
*/
func (d *UserDao) GetUserById(id int64) (*User, error) {
	user := User{UserId: id}

	result := SqlSession.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

```


#### v2
要调用 DAO 层的函数，你需要按照以下步骤进行操作：

- 导入 DAO 包：在你要调用 DAO 函数的地方，首先导入 dao 包，通过 import "dao" 来实现。
- 创建 DAO 实例：使用 NewCountNumDAOImpl() 函数创建 DAO 层的实例。例如，daoInstance := impl.NewCountNumDAOImpl()。
- 调用 DAO 函数：通过创建的实例调用 DAO 函数。例如，result := daoInstance.AddNumInfo(ctx, info)，其中 AddNumInfo 是 DAO 接口定义的函数名。

```go
// 使用说明
func main() {
	daoInstance := impl.NewCountNumDAOImpl()
	service := NewCountNumService(daoInstance)

	// 调用服务层函数
	ctx := context.TODO()
	info := entity.NumInfo{}
	service.AddNumInfo(ctx, info)
	// 其他函数调用...
}
```

```go
package dao

import (
   "context"
   "count_num/pkg/entity"
)

type CountNumDAO interface {
   // 添加一个
   AddNumInfo(ctx context.Context, info entity.NumInfo) bool
   // 根据Key获取一个
   Ge tNumInfoByKey(ctx context.Context, url string) entity.NumInfo
   //查看全部
   FindAllNumInfo(ctx context.Context) []entity.NumInfo
   //  根据Key修改
   UpdateNumInfoByKey(ctx context.Context, info entity.NumInfo) bool
   // 删除一个
   DeleteNumInfoById(ctx context.Context, id int64) bool
   // 根据ID获取一个
   GetNumInfoById(ctx context.Context, id int64) entity.NumInfo
   // 根据ID修改
   UpdateNumInfoById(ctx context.Context, info entity.NumInfo) bool
}

package impl

import (
   "context"
   "count_num/pkg/config"
   "count_num/pkg/entity"
   "gorm.io/gorm"
)

type CountNumDAOImpl struct {
   db *gorm.DB
}

func NewCountNumDAOImpl() *CountNumDAOImpl {
   return &CountNumDAOImpl{db: config.DB}
}

func (impl CountNumDAOImpl) AddNumInfo(ctx context.Context, info entity.NumInfo) bool {
   var in entity.NumInfo
   impl.db.First(&in, "info_key", info.InfoKey)
   if in.InfoKey == info.InfoKey { //去重
      return false
   }
   impl.db.Save(&info) //要使用指针
   return true
}

func (impl CountNumDAOImpl) GetNumInfoByKey(ctx context.Context, key string) entity.NumInfo {
   var info entity.NumInfo
   impl.db.First(&info, "info_key", key)
   return info
}

func (impl CountNumDAOImpl) FindAllNumInfo(ctx context.Context) []entity.NumInfo {
   var infos []entity.NumInfo
   impl.db.Find(&infos)
   return infos
}

func (impl CountNumDAOImpl) UpdateNumInfoByKey(ctx context.Context, info entity.NumInfo) bool {
   impl.db.Model(&entity.NumInfo{}).Where("info_key = ?", info.InfoKey).Update("info_num", info.InfoNum)
   return true
}

func (impl CountNumDAOImpl) DeleteNumInfoById(ctx context.Context, id int64) bool {
   impl.db.Delete(&entity.NumInfo{}, id)
   return true
}

func (impl CountNumDAOImpl) GetNumInfoById(ctx context.Context, id int64) entity.NumInfo {
   var info entity.NumInfo
   impl.db.First(&info, "id", id)
   return info
}

func (impl CountNumDAOImpl) UpdateNumInfoById(ctx context.Context, info entity.NumInfo) bool {
   impl.db.Model(&entity.NumInfo{}).Where("id", info.Id).Updates(entity.NumInfo{Name: info.Name, InfoKey: info.InfoKey, InfoNum: info.InfoNum})
   return true
}

```
