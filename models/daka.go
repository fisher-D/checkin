package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type Card struct{
	Name string `json:"user_name" bson:"user_name"`
	Days []int
	Count int `json:"count"bson:"count"`
	Continuous int
}

func Daka(name string){
	session,_:=mgo.Dial(URL)
	defer session.Close()
	year,month,day,count:= Today()
	var target Card
	coll := session.DB(year).C(month)
	coll.Find(bson.M{"user_name":name}).One(&target)
	res,str,num:=DoubleCheckin(target,count)
	fmt.Println(num)
	if res !=false{
		target.Name = name
		target.Days = append(target.Days,day)
		target.Count = count
		target.Continuous = num
		coll.Upsert(bson.M{"user_name":target.Name},&target)
	}
	fmt.Println(str)
}

//Today year,month,day,continuous
func Today()(string,string,int,int){
	year,month,day := time.Now().Date()
	count:= time.Now().YearDay()
	return strconv.Itoa(year),month.String(),day,count
}
//DoubleCheckin 由于数据库名称和collection名称与
//年和月份挂钩，故而每个collection不会出现跨年和月的问题
func DoubleCheckin(target Card,count int)(bool,string,int){
	if target.Continuous == 0{
		//等于0说明没有记录过
		return true,"打卡成功",1
	}
	if count ==target.Count{
		//说明已经签到
		return false,"已经打过卡了~~~",target.Continuous
	}
	if count>target.Count+1{
		//说明中间断了
		return true,"打卡成功，要坚持打卡哦",1
	}
	if count <target.Count{
		return false,"未知错误",0
	}
	continuous := target.Continuous+1
	em := Emoji(continuous)
	return true,"打卡成功,"+em,continuous
}

func Emoji(day int)string{
	if day<=5{
		return "坚持就是胜利"
	}
	if day >5{
		return "666666666666666"
	}
	return "加油加油"
}