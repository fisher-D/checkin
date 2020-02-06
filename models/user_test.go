package models

import "fmt"
func InitUser(){
	res :=Userinit()
	fmt.Println(res)

}
func Login(){
	u := "1594250583"
	p := "daiyu2621300"
	res := LoginFunc(u,p)
	fmt.Println(res)
}
func Regist(){
	u := "18118160571"
	p := "yanyimingN.1"
	ph := "18118160571"
	e := "893908059@qq.com"
	err :=CommonRegister(u,p,ph,e)
	if err!=nil{
		//fmt.Println("Something Wrong")
		fmt.Println(err)
	}else{
		fmt.Println("注册成功")
	}
}
func Update(){
	c :=make(map[string]string)
	username := "1594250583"
	c["name"]="IamJack"
	c["phone"]="18020285982"
	err :=UpdateUser(username,c)
	if err !=nil{
		fmt.Println(err)
	}else {
		fmt.Println("更新成功")
	}
}
//TODO
//需要在注册和修改信息时
//添加更多的正则判断信息