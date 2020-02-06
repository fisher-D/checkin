package models

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//User ...
type User struct {
	//Name can be changed
	Name string
	//UserName can be changed
	UserName string `json:"user_name" bson:"user_name"`
	//PassWord can be changed
	PassWord []byte `json:"pass_word" bson:"pass_word"`
	//Phone can be changed
	Phone string `json:"phone"`
	//Email can be changed
	Email      string `json:"email"`
	InviteCode string `json:"invite_code,omitempty" bson:"invite_code,omitempty"`
	//Head can be changed
	Head        string    `json:"head"`                       //The Path of Head picture
	Gender      string    `json:"gender"`                     //There will be some people are both gender
	AllDays     int16     `json:"all_days bson:"all_days"`    //Total Daka dates
	Continued   int16     `json:"continued" bson:"continued"` //Continued
	Collections []int8    //List of id from Collection
	Comment     []Comment //List of id from Comment
	status      bool      //禁言 true 为禁言，false为自由发言
}

//Admin ...
type Admin struct {
	AdminID int8 `json:"admin_id"`
	User
	Access bool `json:"access"`
}

//LoginFunc is for User to login
func LoginFunc(UserName, Password string) string {
	session, _ := mgo.Dial(URL)
	defer session.Close()

	coll := session.DB(MongoDBName).C("user")
	target := &User{}
	coll.Find(bson.M{"user_name": UserName}).One(target)
	err := bcrypt.CompareHashAndPassword([]byte(target.PassWord), []byte(Password))
	if err != nil {
		return "Wrong Password"
	}
	res, _ := json.Marshal(target)
	fmt.Println(string(res))
	return "登录成功"
}

//Userinit will initial an user in the db
func Userinit() error {
	session, _ := mgo.Dial(URL)
	defer session.Close()
	user := &User{}
	Password := "test_password"
	hash, _ := bcrypt.GenerateFromPassword([]byte(Password), 5)
	user.PassWord = hash
	user.InviteCode = "66666"
	user.Phone = "18036987412"
	user.Email = "1236987452@qq.com"
	user.Name = "Jack"
	user.AllDays = int16(265)
	user.Collections = []int8{2, 12, 15, 18, 29}
	pinglun := Comment{
		Date:    time.Now(),
		Content: "我就是我",
	}
	user.Comment = []Comment{pinglun}
	user.UserName = "1478569845"
	user.Continued = 25
	user.Gender = "Male"
	user.Head = "/opt/go/src/picture/123.jpg"
	coll := session.DB(MongoDBName).C("user")
	err := isExist(user.Name)
	if err != nil {
		return err
	}
	err1 := coll.Insert(user)
	if err1 != nil {
		fmt.Println(err1)
		return fmt.Errorf("omething wrong happened")
	}
	return fmt.Errorf("数据库初始化成功")
}
//Admininit will initial an Administrator
func Admininit() error {
	session, _ := mgo.Dial(URL)
	defer session.Close()
	user := &Admin{}
	Password := "daiyu2621300"
	hash, _ := bcrypt.GenerateFromPassword([]byte(Password), 5)
	user.PassWord = hash
	user.InviteCode = "66666"
	user.Phone = "18036987412"
	user.Email = "1478523698@qq.com"
	user.Name = "Jack"
	user.AllDays = int16(265)
	user.Collections = []int8{2, 12, 15, 18, 29}
	pinglun := Comment{
		Date:    time.Now(),
		Content: "我就是我",
	}
	user.Comment = []Comment{pinglun}
	user.UserName = "1236547896"
	user.Continued = 25
	user.Gender = "Male"
	user.Head = "/opt/go/src/picture/123.jpg"
	user.AdminID = 1
	user.Access = true
	coll := session.DB(MongoDBName).C("user")
	err := isExist(user.Name)
	if err != nil {
		return err
	}
	err1 := coll.Insert(user)
	if err1 != nil {
		fmt.Println(err1)
		return fmt.Errorf("omething wrong happened")
	}
	return fmt.Errorf("数据库初始化成功")
}

//CommonRegister register without inviteCode
func CommonRegister(UserName, Password, Phone, Email string) error {
	code := ""
	err := RegisterFunc(UserName, Password, Phone, Email, code)
	return err
}

//RegisterFunc register with inviteCode
func RegisterFunc(UserName, Password, Phone, Email, InviteCode string) error {
	session, _ := mgo.Dial(URL)
	defer session.Close()
	user := &User{}
	coll := session.DB(MongoDBName).C("user")
	eresult := VerifyEmailFormat(Email)
	if eresult != true {
		return fmt.Errorf("不是合法邮箱地址")
	}
	nume, _ := coll.Find(bson.M{"email": Email}).Count()
	if nume > 0 {
		return fmt.Errorf("该邮箱地址已使用，请更换")
	}
	presult := VerifyMobileFormat(Phone)
	if presult != true {
		return fmt.Errorf("不是合法手机号码")
	}
	nump, _ := coll.Find(bson.M{"phone": Phone}).Count()
	if nump > 0 {
		return fmt.Errorf("该手机号码已注册，请更换")
	}
	stru, _ := coll.Find(bson.M{"username": UserName}).Count()
	if stru > 0 {
		return fmt.Errorf("该用户名已使用，请更换")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 5)
	if err != nil {
		return fmt.Errorf("密码加密错误")
	}
	if InviteCode != "" {
		collcode := session.DB(MongoDBName).C("invite_code")
		result, _ := collcode.Find(bson.M{"invite_code": InviteCode}).Count()
		if result == 1 {
			user.InviteCode = InviteCode
			collcode.Remove(bson.M{"invite_code": InviteCode})
		} else {
			return fmt.Errorf("无效的邀请码")
		}
	}
	user.UserName = UserName
	user.Email = Email
	user.Phone = Phone
	user.PassWord = hash
	user.status = false
	coll.Insert(&user)
	return nil
}

//UserCheck Check whether the user is exist
func isExist(name string) error {
	seesion, _ := mgo.Dial(URL)
	coll := seesion.DB(MongoDBName).C("user")
	num, _ := coll.Find(bson.M{"name": name}).Count()
	if num > 0 {
		return fmt.Errorf("user exist")
	}
	return nil
}

//UpdateUser ...
func (u *User)UpdateUser(username string, data map[string]string) error {
	session, _ := mgo.Dial(URL)
	coll := session.DB(MongoDBName).C("user")
	target := &User{}
	finder := coll.Find(bson.M{"user_name": username}).One(target)
	if finder != nil {
		return fmt.Errorf("unable to find user")
	}
	for v, k := range data {
		switch v {
		case "name":
			target.Name = k
		case "password":
			hash, _ := bcrypt.GenerateFromPassword([]byte(k), 5)
			target.PassWord = hash
		case "phone":
			target.Phone = k
		case "email":
			target.Email = k
		case "head":
			target.Head = k
		}
	}
	err := coll.Update(bson.M{"user_name": username}, target)
	if err != nil {
		return fmt.Errorf("error in update")
	}
	return nil
}

func (a *Admin)DeleteUser(username string)error{
	session ,_:= mgo.Dial(URL)
	defer session.Close()
	collu := session.DB(MongoDBName).C("user")
	colla := session.DB(MongoDBName).C("delete_his")
	var target User
	err :=collu.Find(bson.M{"user_name":username}).One(&target)
	if err !=nil{
		colla.Insert(target)
		collu.Remove(bson.M{"user_name":username})
		return nil
	}
	return fmt.Errorf("指定用户不存在")
}
//SilentUser true 为禁言 false 为自由
func(a *Admin) SilentUser(username string,status bool)(string,bool){
	session ,_:= mgo.Dial(URL)
	defer session.Close()
	coll := session.DB(MongoDBName).C("user")
	coll.Update(bson.M{"user_name":username},bson.M{"$set":bson.M{"status":status}})
	res :=CheckSilent(coll,username)
	if res ==true{
		return "用户已被禁言",true
	}
		return "用户可以自由发言",false
}
//CheckSilent 
func CheckSilent(collection *mgo.Collection,username string)bool{
	var target User
	collection.Find(bson.M{"user_name":username}).One(&target)
	return target.status
}

func (u *User)MakeComment(data string)error{
	//TODO
	//以后再加
return nil
}

func (u *User)Checkin(){
	Daka(u.UserName)
}