package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
)
//JianJian 猫的名字
type JianJian struct{
	Cat
}
//GetInformation 获取猫的信息
func (j *JianJian)GetInformation()*Cat{
	session,err := mgo.Dial(URL)
	defer session.Close()
	target := new(Cat)
	if err != nil{
		fmt.Println("unable to dial with mongodb")
		return nil
	}
	coll :=session.DB(MongoDBName).C(Jcurrent)
	coll.Find(nil).One(target)
	return target
	//return j.Cat
}
//CreateJianJian 创建指定的猫
func (j *JianJian)CreateJianJian()error{
	JJ := GenerateCat(Jname,Jdes,Jsep,Jgroup,Jcount)
	session ,err:=mgo.Dial(URL)
	defer session.Close()
	if err!=nil{
	return fmt.Errorf("unable to dial mongodb")
}
	collhistory := session.DB(MongoDBName).C(Jhistory)
	collhistory.Insert(JJ)
	collJ :=session.DB(MongoDBName).C(Jcurrent)
	collJ.RemoveAll(nil)
	collJ.Insert(JJ)
	return nil
}
//UpdateData including group;count;name;description;special
func (j *JianJian)UpdateData(data map[string]interface{})error{
	session,err:= mgo.Dial(URL)
	defer session.Close()
	if err !=nil{
		return fmt.Errorf("error in dial with mongodb")
	}
	//target := new(Cat)
	gohost:= new(Cat)
	coll := session.DB(MongoDBName).C(Jcurrent)
	coll.Find(nil).One(gohost)

	for m,n:=range data{
	switch m {
	case "group":
		gohost.Group = n.(int8)
	case "count":
		gohost.Count = n.(int16)
	case "name":
		gohost.Name =n.(string)
	case "description":
		gohost.Description = n.(string)
	case "special":
		gohost.Special = n.(string)
	default:
		return fmt.Errorf("error! unsupported parameters")
		}
	}
	coll.RemoveAll(nil)
	coll.Insert(gohost)
	collhistory := session.DB(MongoDBName).C(Jhistory)
	collhistory.Insert(gohost)
	return nil
}

