package models

import "fmt"

func BuildAndUpdateTest(){
	var Jian JianJian
	updateData := BuildUpdateInfo("渐渐~~","可爱的小猫猫","今天可以抽取耳机哦")
	fmt.Println(updateData)
	err :=Jian.UpdateData(updateData)
	if err !=nil{
		res :=Jian.GetInformation()
		fmt.Println(res)
	}

}
func CreateAndGetTest(){
	var Jian JianJian
	err :=Jian.CreateJianJian()
	if err!=nil{
		fmt.Println("Go On")
	}
	res :=Jian.GetInformation()
	fmt.Println(res)
}