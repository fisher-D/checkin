package models

import "time"

//Cat 猫的结构
type Cat struct{
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Group int8 `json:"group,omitempty" bson:"group,omitempty"`
	Count int16 `json:"count,omitempty" bson:"count,omitempty"`
	Special string `json:"special" bson:"special"`
}
//BuildUpdateInfo Example of Build Data
func BuildUpdateInfo(name,des,spe string)map[string]interface{}{
	updateData := make(map[string]interface{})
	updateData["group"]= int8(1)
	updateData["count"] = int16(0)
	updateData["name"] = name
	updateData["description"]= des
	updateData["special"] = spe
	return updateData
}
//GenerateCat init the Cat
func GenerateCat(name,des,special string,group int8,count int16)*Cat{
	cat := new(Cat)
	cat.Name = name
	cat.Description = des
	cat.Special =special
	cat.Group = group
	cat.Count = count
	return cat
}

//Bag ...
type Bag struct{
	Name string `json:"name"`
	Capacity string `json:"capacity"`
	Description string `json:"description"`
	Special string `json:"special" bson:"special,omitempty"`
}

//Comment ...
type Comment struct{
	Date time.Time `json:"date"`
	Content string `json:"content"`
}


