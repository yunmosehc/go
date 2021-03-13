package controllers

import (
	"encoding/json"
	"fabcar/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"strconv"
	"time"

	//"github.com/astaxie/beego/orm"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"io/ioutil"
	"os"
	"path/filepath"
	//"strconv"
	//"time"
	"errors"
)

// ArticleController 自定义控制器
type ArticleController struct {
	beego.Controller
}

type Car struct {
	Title   string `json:"title"`
	IpfsAddress string `json:"ipfsaddress"`
	OwnerAccountId string `json:"owneraccountid"`
	LastOwnerAccountId string `json:"lastowneraccountid"`
	AcquireDate string `json:"acquiredate"`
	OwnerName string `json:"ownername"`
	OwnerCardNumber string `json:"ownercardnumber"`
}

type QueryResult struct {
	Key    string `json:"Key"`
	Record *Car
}

//将拿到contract的部分单独封装，返回的contract是一个指针
var contract *gateway.Contract

func getContract(){
	//**************fabric部分******************
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract = network.GetContract("fabcar")
}

func getCookieHandler(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("accountid")
	if err != nil{
		beego.Info(w, "Cannot get cookie")
	}
	beego.Info(c1.Value)
}

// ShowIndex 展示首页并实现分页功能
func (a *ArticleController) ShowIndex() {
	//// 1.查询所有文章数据
	//o := orm.NewOrm()
	//var articles []models.Article
	//querySeter := o.QueryTable("Article").RelatedSel()
	////获取用户的账户编号
	//accountId := a.GetString("accountid")
	//accountid, err2 := strconv.Atoi(accountId)
	//if err2 != nil {
	//}
	//beego.Info("当前登录账户编号是：")
	//beego.Info(accountid)
	//
	//querySeter = querySeter.Filter("OwnerAccountId", accountid)
	///*_,err := querySeter.All(&articles)*/
	//count, _ := querySeter.Count()
	//// 2.设置每一页显示的数量，从而得到总的页数
	//var pageSize = 5
	//pageCount := math.Ceil(float64(count) / float64(pageSize)) // 向上取整，显示的页面不会出现小数
	//// 3.首页和末页
	//pi := a.GetString("pi")
	//pageIndex, err := strconv.Atoi(pi)
	//if err != nil {
	//	pageIndex = 1 // 首页没有传pageIndex的值，防止默认pageIndex为0
	//}
	//// 3.1每一页显示的个数
	//stat := pageSize * (pageIndex - 1)
	//_, err = querySeter.Limit(pageSize, stat).RelatedSel().All(&articles)
	//if err != nil {
	//	beego.Info("获取文章数据失败")
	//	beego.Info(err)
	//	a.Redirect("/article/index", 302)
	//	return
	//}
	//// 4.上一页和下一页限制(视图函数)
	//var isFirstPage = false
	//var isLastPage = false
	//if pageIndex == 1 {
	//	isFirstPage = true
	//}
	//if pageIndex == int(pageCount) {
	//	isLastPage = true
	//}
	//
	//a.Data["username"] = a.GetSession("username")
	//a.Data["accountid"] = a.GetSession("accountid")
	//a.Data["count"] = count
	//a.Data["pageCount"] = pageCount
	//a.Data["pageIndex"] = pageIndex
	//a.Data["isFirstPage"] = isFirstPage
	//a.Data["isLastPage"] = isLastPage
	//a.Data["articles"] = articles
	//a.TplName = "index.html"

	//
	//result, err = contract.SubmitTransaction("createCar", "CAR2", "联盟链开发实战",
	//	"https://ipfs.io/ipfs/QmQU2gS4gZ7TpiTECjDUxdQFd9bBBEWxDxPPfhLfYHVuei", "000002", "000000", "2020.10.20 18:20:30", "李白", "110100200101101201")
	//if err != nil {
	//	fmt.Printf("Failed to submit transaction: %s\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println(string(result))
	//
	//result, err = contract.EvaluateTransaction("queryCar", "CAR2")
	//if err != nil {
	//	fmt.Printf("Failed to evaluate transaction: %s\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println(string(result))
	//
	//_, err = contract.SubmitTransaction("changeCarOwner", "CAR2", "000003", "杜甫", "371510199002202838")
	//if err != nil {
	//	fmt.Printf("Failed to submit transaction: %s\n", err)
	//	os.Exit(1)
	//}
	//
	//result, err = contract.EvaluateTransaction("queryCar", "CAR2")
	//if err != nil {
	//	fmt.Printf("Failed to evaluate transaction: %s\n", err)
	//	os.Exit(1)
	//}
	////carStr = string(result)
	//fmt.Println(string(result))

	//#################测试部分####################
	//createCar的result是空的，要得到数据必须再querycar
	//result, err := contract.SubmitTransaction("createCar", "第二篇文章",
	//	"https://ipfs.io/ipfs/QmQU2gS4gZ7TpiTECjDUxdQFd9bBBEWxDxPPfhLfYHVuei", "0002", "0000", "2020.10.20 18:20:30", "李白", "110100200101101201")
	//if err != nil {
	//	fmt.Printf("Failed to submit transaction: %s\n", err)
	//	os.Exit(1)
	//}

	//如果contract还未初始化，先初始化contract
	if contract == nil{
		getContract()
	}

	result, err := contract.EvaluateTransaction("queryAllCars")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}

	//得出result中的数据数量
	queryResults := new([200]QueryResult)
	json.Unmarshal(result, queryResults)
	var count int
	for i:=0; i<200; i++ {
		if(queryResults[i].Record != nil) {
			count++
		};
	}
	//遍历queryResults装载articles
	var articles []models.Article
	for i:=0; i<count; i++ {
		art := queryResults[i].Record
		var article models.Article
		article.ArtID = queryResults[i].Key
		article.Title = art.Title
		article.IpfsAddress = art.IpfsAddress
		article.OwnerAccountId = art.OwnerAccountId
		article.LastOwnerAccountId = art.LastOwnerAccountId
		article.AcquireDate = art.AcquireDate
		article.OwnerName = art.OwnerName
		article.OwnerCardNumber = art.OwnerCardNumber
		articles = append(articles, article)
	}

	//art := new(Car)
	//_ = json.Unmarshal(result, art)
	//
	//var articles []models.Article
	//var article models.Article
	//article.ArtID = "1"
	//article.Title = art.Title
	//article.IpfsAddress = art.IpfsAddress
	//article.OwnerAccountId = art.OwnerAccountId
	//article.LastOwnerAccountId = art.LastOwnerAccountId
	//article.AcquireDate = art.AcquireDate
	//article.OwnerName = art.OwnerName
	//article.OwnerCardNumber = art.OwnerCardNumber
	//articles = append(articles, article)

	a.Data["username"] = a.GetSession("username")
	a.Data["accountid"] = a.GetSession("accountid")

	var r *http.Request
	c1, err := r.Cookie("accountid")
	beego.Info("********************")
	beego.Info(c1.Domain)
	beego.Info("********************")

	a.Data["count"] = 1
	a.Data["pageCount"] = 1
	a.Data["pageIndex"] = 1
	a.Data["isFirstPage"] = true
	a.Data["isLastPage"] = false
	a.Data["articles"] = articles
	a.TplName = "index.html"
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}

// ShowAdd 展示添加文章界面
func (a *ArticleController) ShowAdd() {
	// 展示下拉类型
	//o := orm.NewOrm()
	//var types []models.ArticleType
	//_, err := o.QueryTable("article_type").RelatedSel().All(&types)
	//if err != nil {
	//	beego.Info("获取文章类型错误")
	//	a.Redirect("/article/index", 302)
	//	return
	//}

	//a.Data["types"] = types
	a.Data["username"] = a.GetSession("username")
	a.Data["accountid"] = a.GetSession("accountid")
	a.TplName = "add.html"
	//a.Redirect("/addArticle",302)
}

// HandleAdd 处理添加文章业务
func (a *ArticleController) HandleAdd() {
	//var filePath string
	// 1.拿到前端数据
	//artName := a.GetString("artname")
	//artContent := a.GetString("artcontent")
	//typeName := a.GetString("select")

	//获取提交的数据
	title := a.GetString("title")
	ipfsaddress := a.GetString("ipfsaddress")
	//当前用户id
	accountid, err2 := a.GetInt("accountid")
	if err2 != nil {
	}
	//当前时间
	t := time.Now()
	time_now :=fmt.Sprintf("%4d.%02d.%02d %02d:%02d:%02d(%02d)\n",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second(),t.Nanosecond())
	ownername := a.GetString("ownername")
	ownercardnumber := a.GetString("ownercardnumber")

	//获取session中的accountid并转成整数
	//accountId := a.Ctx.GetCookie("accountid")
	//beego.Info("string is : " + accountId)
	//accountid, err2 := strconv.Atoi(accountId)
	//beego.Info("int is : ")
	//beego.Info(accountid)
	//
	//if err2 != nil {
	//	beego.Info("accountId字符串转整数失败")
	//}

	//// 2.校验数据（是否为空）
	//if title == "" || ipfsaddress == "" || ownername == "" || ownercardnumber == "" {
	//	beego.Info("数据不能为空")
	//	a.Redirect("/article/add", 302)
	//	return
	//}
	//// 3.校验文件
	////file, head, err := a.GetFile("artfile")
	////if file != nil {
	////	defer file.Close()
	////	if err != nil {
	////		beego.Info("上传文件失败")
	////		a.Redirect("/article/add", 302)
	////		return
	////	}
	////	// 3.1限制文件的格式.jpg/.png/.gif
	////	ext := path.Ext(head.Filename) //获取文件拓展名
	////	if ext != ".jpg" && ext != ".png" && ext != ".gif" {
	////		beego.Info("文件格式不正确！")
	////		a.Redirect("/article/add", 302)
	////		return
	////	}
	////	// 3.2限制文件大小
	////	if head.Size > 20<<20 {
	////		beego.Info("文件不能大于20M")
	////		a.Redirect("/article/add", 302)
	////		return
	////	}
	////	//3.3给文件重命名
	////	unix := time.Now().Format("20060102_150405") + ext
	////	_ = a.SaveToFile("artfile", "./static/img/"+unix) //注意文件路径./开头
	////	filePath = "/static/img/" + unix
	////}
	//// 4.将数据插入数据库
	//o := orm.NewOrm()
	////aType := models.ArticleType{TypeName: typeName} //初始化一个ArticleType对象
	////err = o.Read(&aType, "TypeName")
	////if err != nil {
	////	beego.Info("获取文章类型失败")
	////	a.Redirect("/article/index", 302)
	////	return
	////}
	////art := models.Article{ArtName: artName, ArtContent: artContent, ArtImg: filePath, ArtType: &aType}
	//
	//art := models.Article{
	//	Title:              title,
	//	IpfsAddress:        ipfsaddress,
	//	OwnerAccountId:     accountid,
	//	LastOwnerAccountId: 0,
	//	AcquireDate:        time.Time{},
	//	OwnerName:          ownername,
	//	OwnerCardNumber:    ownercardnumber,
	//}

	//如果contract还未初始化，先初始化contract
	if contract == nil{
		getContract()
	}

	_, err := contract.SubmitTransaction("createCar", title,
		ipfsaddress, strconv.Itoa(accountid), "0000",time_now, ownername, ownercardnumber)
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		os.Exit(1)
	}

	//var err error
	//_, err = o.Insert(&art)
	//if err != nil {
	//	beego.Info("添加文章至数据库失败")
	//	beego.Info(err)
	//	a.Redirect("/article/add", 302)
	//	return
	//}
	// 5.跳转页面（index.html）
	//a.Redirect("/article/index", 302)
	a.Redirect("/article/index?accountid="+a.GetString("accountid"), 302)
	//a.TplName = "index.html"
}

//ShowContent 展示详情页面
//func (a *ArticleController) ShowContent() {
//	// 1.获取文章id
//	sid := a.GetString("id")
//	id, _ := strconv.Atoi(sid)
//	fmt.Println(id)
//	// 2.通过id查询数据库信息
//	o := orm.NewOrm()
//	article := models.Article{ArtID: id}
//	err := o.Read(&article)
//	if err != nil {
//		beego.Info("failed~")
//		a.Redirect("/article/index", 302)
//		return
//	}
//	// 3.将数据传给视图
//	a.Data["username"] = a.GetSession("username")
//	a.Data["article"] = article
//	// 4.跳转页面
//	a.TplName = "content.html"
//}

// ShowUpdate 展示产权转让页面
func (a *ArticleController) ShowUpdate() {

	//// 1.获取文章id
	//id, _ := a.GetInt("id")
	//// 2.根据id查询文章信息
	//o := orm.NewOrm()
	//// 2.1 获取文章类型
	////var types []models.ArticleType
	////_, err := o.QueryTable("ArticleType").RelatedSel().All(&types)
	////if err != nil {
	////	beego.Info("获取文章类型错误")
	////	a.Redirect("/article/index", 302)
	////	return
	////}
	//
	//article := models.Article{ArtID: id}
	//var err error
	//err = o.Read(&article)
	//if err != nil {
	//	beego.Info("查询数据信息失败")
	//	a.Redirect("/article/index", 302)
	//	return
	//}
	//// 3.将文章信息传给视图
	////a.Data["types"] = types
	a.Data["username"] = a.GetSession("username")
	a.Data["accountid"] = a.GetSession("accountid")
	a.TplName = "update.html"
}

// 产权转让业务处理
func (a *ArticleController) HandleUpdate() {
	// 1.获取页面数据
	artId := a.GetString("artid")
	newOwnerAccountId := a.GetString("newaccountid")
	lastOwnerAccountId := a.GetString("accountid")
	t := time.Now()
	time_now :=fmt.Sprintf("%4d.%02d.%02d %02d:%02d:%02d(%02d)\n",t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second(),t.Nanosecond())
	newOwnerName := a.GetString("newownername")
	newOwnerCardNumber := a.GetString("newownercardnumber")

	// 2.判断合法性
	//判断文章是否存在
	//car, err := s.QueryCar(ctx, carNumber)
	result, err := contract.EvaluateTransaction("queryCar", artId)
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		beego.Info("该文章编号不存在！")
		a.Redirect("/article/delete", 302)
		os.Exit(1)
	}
	//判断是否有操作权限
	car := new(Car)
	json.Unmarshal(result, car)
	accountId := a.GetString("accountid")
	if car.OwnerAccountId != accountId {
		beego.Info("您无法删除不属于您的文章！")
		a.Redirect("/article/update", 302)
		return
	}
	//判断新产权人账户是否存在
	o := orm.NewOrm()
	new_OwnerAccountId, err := strconv.Atoi(newOwnerAccountId)
	newowneraccountid := models.UserInfo{AccountId: new_OwnerAccountId}
	err = o.Read(&newowneraccountid)
	if err != nil {
		beego.Info("新产权人编号不存在！")
		a.Redirect("/article/update", 302)
		return
	}

	// 3.更新文章产权信息
	_, err = contract.SubmitTransaction("changeCarOwner", artId,
		newOwnerAccountId, lastOwnerAccountId, time_now, newOwnerName, newOwnerCardNumber)
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		beego.Info("changeCarOwner交易执行失败")
		a.Redirect("/article/update", 302)
		os.Exit(1)
	}
	a.Redirect("/article/index?accountid="+a.GetString("accountid"), 302)
}

// ShowEdit 展示产权信息编辑页面
func (a *ArticleController) ShowEdit() {
	a.Data["username"] = a.GetSession("username")
	a.Data["accountid"] = a.GetSession("accountid")
	a.TplName = "edit.html"
}

// 产权信息编辑处理
func (a *ArticleController) HandleEdit() {
	// 1.获取页面数据
	artId := a.GetString("artid")
	newOwnerName := a.GetString("newownername")
	newOwnerCardNumber := a.GetString("newownercardnumber")

	// 2.判断合法性
	//判断文章是否存在
	result, err := contract.EvaluateTransaction("queryCar", artId)
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		beego.Info("该文章编号不存在！")
		a.Redirect("/article/edit", 302)
		os.Exit(1)
	}
	//判断是否有操作权限
	car := new(Car)
	json.Unmarshal(result, car)
	accountId := a.GetString("accountid")
	if car.OwnerAccountId != accountId {
		beego.Info("您无法删除不属于您的文章！")
		a.Redirect("/article/edit", 302)
		return
	}

	// 3.更新文章产权信息
	_, err = contract.SubmitTransaction("changeCarOwner", artId, car.OwnerAccountId, car.LastOwnerAccountId, car.AcquireDate,
		newOwnerName, newOwnerCardNumber)
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		beego.Info("changeCarOwner交易执行失败")
		a.Redirect("/article/edit", 302)
		os.Exit(1)
	}
	a.Redirect("/article/index?accountid="+a.GetString("accountid"), 302)
}

//展示删除产权界面
func (a *ArticleController) ShowDelete() {
	a.Data["username"] = a.GetSession("username")
	a.Data["accountid"] = a.GetSession("accountid")
	a.TplName = "delete.html"
}

// Delete 删除业务处理
func (a *ArticleController) HandleDelete() {
	// 1.获取文章id
	artid := a.GetString("artid")
	// 2.查询出对应数据
	result, err := contract.EvaluateTransaction("queryCar", artid)
	// 3.合法性判断
	// 判断文章是否存在
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		beego.Info("文章编号不存在")
		a.Redirect("/article/delete", 302)
		os.Exit(1)
	}
	// 判断操作者是否是产权人
	car := new(Car)
	json.Unmarshal(result, car)
	accountId := a.GetString("accountid")
	if car.OwnerAccountId != accountId {
		beego.Info("您无法删除不属于您的文章！")
		a.Redirect("/article/delete", 302)
		return
	}

	//if car.OwnerName != a.GetString("ownername") || car.OwnerCardNumber != a.GetString("ownercardnumber") {
	//	beego.Info("产权人姓名或身份证号有误！")
	//	a.Redirect("/article/delete", 302)
	//	return
	//}

	// 4.删除文章
	_, err = contract.SubmitTransaction("deleteCarOwner", artid)
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		beego.Info("提交删除文章交易失败")
		a.Redirect("/article/delete", 302)
		os.Exit(1)
	}

	// 5.跳转列表页
	a.Redirect("/article/index?accountid="+a.GetString("accountid"), 302)
}

//ShowArtType 展示文章类型
//func (a *ArticleController) ShowArtType() {
//	a.TplName = "addType.html"
//	// 1.读取文章类型表
//	o := orm.NewOrm()
//	var types []models.ArticleType
//	_, err := o.QueryTable("article_type").All(&types)
//	if err != nil {
//		beego.Info("获取文章类型错误")
//	}
//	// 2.将数据传给前端视图
//	a.Data["username"] = a.GetSession("username")
//	a.Data["types"] = types
//}
//
//AddType 添加文章类型
//func (a *ArticleController) AddType() {
//	// 1.获取前端数据
//	typeName := a.GetString("typeName")
//	// 2.数据校验
//	if typeName == "" {
//		beego.Info("类型名为空")
//		a.Redirect("/article/addType", 302)
//		return
//	}
//	// 3.将数据插入类型表
//	o := orm.NewOrm()
//	artType := models.ArticleType{TypeName: typeName}
//	err := o.Read(&artType, "TypeName")
//	if err == nil {
//		a.Data["errMsg"] = "该类型已存在！"
//		a.Redirect("/article/addType", 302)
//		//a.TplName = "addType.html"
//		return
//	}
//	_, err = o.Insert(&artType)
//	if err != nil {
//		beego.Info("添加类型错误")
//		a.Redirect("/article/addType", 302)
//		return
//	}
//	// 4.跳转回类型页面
//	a.Redirect("/article/addType", 302)
//}
