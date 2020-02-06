package models

import "fmt"

var (
	A Collection
)

func Create(){

	A.Description = "第一个物品"
	A.Price= 23
	A.Picture = "opt/gopath/src/123.jpg"
	err:=A.Create(A)
	if err!=nil{
		fmt.Println("出错了~~")
	}
}

func Updates(){
	m :=make(map[string]interface{})
	m["name"]="小甜甜"
	err:=A.Update(3,m)
	if err!=nil{
		fmt.Println("更新出错了~~~")
		fmt.Println(err)
	}
}

func Get(){
	res := A.GetData()
	fmt.Println(res)
}

func GetOne(){
	res,err :=A.GetSpecific(2)
	if err==nil{
		fmt.Println(res)
	}
}

func Delete(){
	err :=A.Delete(3)
	if err==nil{
		fmt.Println("删除成功")
	}
}

func GetCollectionNumber(){
	num,_ :=A.GetCollectionNumber()
	fmt.Println(num)
}