/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"errors"
	_ "fabcar/models"
	_ "fabcar/routers"
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

//********************CS系统部分*************************
//var carStr string = ""
//
//type HomeController struct {
//        beego.Controller
//}
//
//func (this *HomeController) Get() {
//        this.Data["carString"] = carStr
//		this.TplName = "allCar.html"
//}

// ShowPrePage 实现上一页
func ShowPrePage(pi int) (pre string) {
	pageIndex := pi - 1
	pre = strconv.Itoa(pageIndex)
	return
}

// ShowNextPage 实现下一页
func ShowNextPage(pi int) (next int) {
	next = pi + 1
	return
}

// AutoKey id自增
func AutoKey(key int) int {
	return key + 1
}
//********************CS系统部分*************************

func main() {


	//********************CS系统部分*************************
	//beego.Router("/",&HomeController{})
	//映射视图函数,必须放在run函数前
	_ = beego.AddFuncMap("PrePage", ShowPrePage)
	_ = beego.AddFuncMap("NextPage", ShowNextPage)
	_ = beego.AddFuncMap("autoKey", AutoKey)
	beego.Run()
	//********************CS系统部分*************************
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
