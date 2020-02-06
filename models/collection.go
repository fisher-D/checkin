package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Collection ...
type Collection struct{
	IDnumber int8 `json:"id_number" bson:"id_number"`
	Name string 	`json:"name"`
	Picture string `json:"picture" bson:"picture"` //Save the path of the picture
	Price int8 `json:"price" bson:"price"` //Pay attention to the decimal
	Description string `json:"description" bson:"description"`
}

//Create 不需要指定物品ID
func (c *Collection)Create(info Collection)error{
	session,_ := mgo.Dial(URL)
	defer session.Close()
	coll :=session.DB(MongoDBName).C("collection")
	var res Collection
	coll.Find(nil).Sort("-id_number").One(&res)
	c.IDnumber = res.IDnumber + 1
	c.Description =info.Description
	c.Price = info.Price
	c.Picture = info.Picture
	c.Name = info.Name
	errI :=coll.Insert(c)
	if errI!=nil{
		return fmt.Errorf("无法创建新的物品；code:1001")
	}
	fmt.Println("成功创建")
	return nil
}
//Delete ...
func (c *Collection)Delete(number int8)error{
	session,_:=mgo.Dial(URL)
	defer session.Close()
	coll := session.DB(MongoDBName).C("collection")
	err :=coll.Remove(bson.M{"id_number":number})
	if err !=nil{
		fmt.Errorf("无法删除物品；code：1002")
	}
	fmt.Println("成功删除")
	return nil
}
//GetData ...
func (c *Collection)GetData()[]Collection{
	session,_ := mgo.Dial(URL)
	defer session.Close()
	coll := session.DB(MongoDBName).C("collection")
	var res []Collection
	err :=coll.Find(nil).All(&res)
	if err !=nil{
		fmt.Println("获取数据出错；code：1003")
		return nil
	}
	return res
}
//GetSpecific ...
func (c *Collection)GetSpecific(id int8)(Collection,error){
	session,_:=mgo.Dial(URL)
	defer session.Close()
	coll := session.DB(MongoDBName).C("collection")
	var Result Collection
	err :=coll.Find(bson.M{"id_number":id}).One(&Result)
	if err!=nil{
		fmt.Println("无法找到指定物品；code：1004")
		return Result,err
	}
	return Result,nil
}
//Update ...
func (c *Collection)Update(id int8,data map[string]interface{})error{
	session,_:=mgo.Dial(URL)
	defer session.Close()
	coll := session.DB(MongoDBName).C("collection")
	var target Collection
	coll.Find(bson.M{"id_number":id}).One(&target)
	//fmt.Println(target)
	for v,k:=range data{
		switch v {
		case "name":
			target.Name = k.(string)
		case "picture":
			target.Picture = k.(string)
		case "price":
			target.Price = k.(int8)
		case "description":
			target.Description = k.(string)
		default:
			return fmt.Errorf("unspported type")
		}
	}
	target.IDnumber = id
	err := coll.Update(bson.M{"id_number":id},&target)
	if err!=nil{
		return fmt.Errorf("无法更新；code:1005")
	}
	fmt.Println("更新成功")
	return nil
}
//GetCollectionNumber ...
func (c *Collection)GetCollectionNumber()(int,error){
	session,_:=mgo.Dial(URL)
	defer session.Close()
	num ,err:=session.DB(MongoDBName).C("collection").Count()
	if err!=nil{
		fmt.Println("数据出错")
		return 0,err
	}
	return num,nil
}