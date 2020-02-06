package models

import "fmt"
func InitUser(){
	res :=Userinit()
	fmt.Println(res)

}
func Login(){
	u := "test_account"
	p := "test_password"
	res := LoginFunc(u,p)
	fmt.Println(res)
}
func Regist(){
	u := "test_account1"
	p := "test_password1"
	ph := "18254789631"
	e := "5879641259@qq.com"
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
	username := "test_username"
	c["name"]="IamJack"
	c["phone"]="18023698741"
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