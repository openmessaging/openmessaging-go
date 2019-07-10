package internal

import (
	"github.com/emirpasic/gods/maps/hashmap"
)

var properties = hashmap.New()

type DefaultKeyValue interface {
	KeyValue(key string, value bool)
}

//
//func Put(key string,value bool) message.KeyValue{
//	properties.Put(key,strconv.FormatBool(value))
//	return properties
//}
//
//func GetBool(key string)bool {
//	ok,err:=properties.ContainsKey(key)
//	if !ok{
//		   return ok
//	}
//	s,err:=strconv.ParseBool(key)
//	if err!=nil{
//		return false
//	}
//	return s
//}
//
//func GetBoolKV(key string,defaultValue bool)bool{
//	ok,err:=properties.ContainsKey(key)
//		if err!=nil{
//		return  false
//	}
//	return GetBool(key) || ok || defaultValue
//}

//func GetShort(key string)int16{
//	ok,_:=properties.ContainsKey(key)
//	if !ok{
//		return 0
//	}
//	fmt.Println(properties.Get(key))
//	_, ok = properties.Get(key)
//	fmt.Println(ok)
//	 return   strconv.FormatBool(ok)
//}
//
//func GetShortKV(key string,DefaultValue int16)int16{
//	ok,_ := properties.ContainsKey(key)
//	if !ok {
//		return 0
//	} else{
//	//	GetShort(key)
//	}
//	return  DefaultValue
//}
//
//func ContainsKey(key string)bool{
//	return ContainsKey(key)
//}
