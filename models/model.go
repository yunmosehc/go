package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*数据库表结构的设计应该放在model文件夹下面*/

// UserInfo 用户信息结构体
type UserInfo struct {
	//Id       int
	//Name     string `orm:"unique"`
	//Password string
	//Articles []*Article `orm:"reverse(many)"`

	AccountId int    `orm:"pk;auto"`
	Username  string `orm:"unique"`
	Password  string
}

// Article 文章信息表
type Article struct {
	//ArtId       int       `orm:"pk;auto"`
	//ArtTitle    string    `orm:"size(48);unique"`
	//ArtCreateAt time.Time `orm:"auto_now;type(date)"`
	//ArtCount    int       `orm:"default(0)"`
	//ArtContent  string
	//ArtImg      string       `orm:"null"`
	//ArtType     *ArticleType `orm:"rel(fk)"`  //fk代表主键 和类型是一对多的关系
	//Users       []*UserInfo  `orm:"rel(m2m)"` //和用户是多对多的关系

	//ArtID              int       `orm:"pk;auto"`
	//Title              string    `orm:"size(500)"`
	//IpfsAddress        string    `orm:"size(500)"`
	//OwnerAccountId     int       `orm:"default(0)"`
	//LastOwnerAccountId int       `orm:"default(0)"`
	//AcquireDate        time.Time `orm:"auto_now;type(date)"`
	//OwnerName          string    `orm:"size(100)"`
	//OwnerCardNumber    string    `orm:"size(100)"`

	ArtID              string
	Title              string
	IpfsAddress        string
	OwnerAccountId     string
	LastOwnerAccountId string
	AcquireDate        string
	OwnerName          string
	OwnerCardNumber    string
}

// ArticleType 文章类型表
//type ArticleType struct {
//	Id       int
//	TypeName string     `orm:"size(500);unique"`
//	Articles []*Article `orm:"reverse(many)"`
//}

// init 连接数据库
func init() {
	// 1.设置数据库基本信息，参数二：驱动的名称"mysql""sqlite""postgres"
	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/AMS?charset=utf8", 30)
	beego.Info("Connect sql success")
	// 2.映射model 数据
	//orm.RegisterModel(new(UserInfo), new(Article), new(ArticleType))
	orm.RegisterModel(new(UserInfo), new(Article))
	// 3.生成表,参数一:数据库别名；参数二：是否强制更新表结构（若表结构该表需要切换成true）；参数三：创建过程在终端是否可见
	_ = orm.RunSyncdb("default", false, true)
	beego.Info("Create table successful")
}
