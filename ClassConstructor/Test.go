package ClassConstructor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)
type TypeMap struct {
	intmap    map[string]int
	boolmap   map[string]bool
	stringmap map[string]string
	keys      []string
	kvmap     map[string]interface{}
}
func createStringFiled(name string) string {
	filed := name + " string\n"
	return filed
}
func createIntFiled(name string) string {
	filed := name + " int\n"
	return filed
}
func createBooleanFiled(name string) string {
	filed := name + " bool\n"
	return filed
}
func isIntvalue(value interface{}) bool {
	stringvalue := fmt.Sprintf("%v", value)
	for i := 0; i < len(stringvalue); i++ {
		charAt := stringvalue[i : i+1]
		if charAt <= "9" && charAt >= "0" {
			continue
		} else {
			return false
		}
	}
	return true
}
func readFile(file string, typemap TypeMap) TypeMap{
	userFile := file
	//kvmap:=map[string]interface{}{}
	//keys=make([]string,100,100)
	fin, err := os.Open(userFile) //打开文件,返回File的内存地址
	defer fin.Close()             //延迟关闭资源
	if err != nil {
		fmt.Println(userFile, err)
		return TypeMap{}
	}
	rd := bufio.NewReader(fin)
	var times int32 = 0
	for true {

		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		} else {
			nsplit := strings.Split(line, "\n")
			countSplit := strings.Split(nsplit[0], "=")
			typemap.keys[times] = countSplit[0]
			times++
			typemap.kvmap[countSplit[0]] = countSplit[1]
		}
	}
	//writeFileBytxt(typemap)
	return typemap
}
func writeFileBytxt(typeMap TypeMap,packagename string,structname string) TypeMap{
	gofileurl := "/home/mikasa/go/src/awesomeProject/src/MyClass/Tokyo.go"
	fin, err := os.Create(gofileurl)
	defer fin.Close()
	if err != nil {
		fmt.Println(gofileurl, err)
		return TypeMap{}
	}
	fin.WriteString("package "+packagename+"\n")
	fin.WriteString("type "+structname+" struct{\n")
	for i := 0; i < len(typeMap.keys); i++ {
		key := typeMap.keys[i]
		value := typeMap.kvmap[key]
		if value == "true" {
			typeMap.boolmap[key] = true
			fin.WriteString(createBooleanFiled(key))
		} else if value == "false" {
			typeMap.boolmap[key] = false
			fin.WriteString(createBooleanFiled(key))
		} else if isIntvalue(value) {
			stringvalue := fmt.Sprintf("%v", value)
			typeMap.intmap[key], _ = strconv.Atoi(stringvalue)
			fin.WriteString(createIntFiled(key))
		} else {
			stringvalue := fmt.Sprintf("%v", value)
			typeMap.stringmap[key] = stringvalue
			fin.WriteString(createStringFiled(key))
		}
	}
	fin.WriteString("}\n")
	//injectionValueAndDisplay(typeMap)
	return typeMap
}
func injectionValueAndDisplay(typemap TypeMap,mclass reflect.Type) {
	//mclass := reflect.TypeOf(MyClass.Myclass{})
	if mclass.Kind() == reflect.Ptr {
		mclass = mclass.Elem()
	}
	myclass := reflect.New(mclass)
	for i := 0; i < len(typemap.keys); i++ {
		if myclass.Elem().Field(i).Kind()==reflect.String{
			//fmt.Println(typemap.stringmap[typemap.keys[i]])
			myclass.Elem().Field(i).SetString(typemap.stringmap[typemap.keys[i]])
		} else
		if myclass.Elem().Field(i).Kind() == reflect.Int {
			//var myint int64=int64(typemap.intmap[typemap.keys[i]])
			//fmt.Println(typemap.intmap[typemap.keys[i]])
			myclass.Elem().Field(i).SetInt(int64(typemap.intmap[typemap.keys[i]]))
			//fmt.Println(int64(typemap.intmap[typemap.keys[i]]))
		} else
		if myclass.Elem().Field(i).Kind()==reflect.Bool {
			//fmt.Println(typemap.boolmap[typemap.keys[i]])
			myclass.Elem().Field(i).SetBool(typemap.boolmap[typemap.keys[i]])
		}
	}
	for j:=0;j<mclass.NumField();j++ {
		key:=mclass.Field(j).Name
		value:=myclass.Elem().Field(j)
		fmt.Println(key,"=",value)
	}
}
func main() {
	typemap := new(TypeMap)
	typemap.keys = make([]string, 100, 100)
	typemap.kvmap = make(map[string]interface{})
	typemap.stringmap = make(map[string]string)
	typemap.intmap = make(map[string]int)
	typemap.boolmap = make(map[string]bool)
	*typemap=readFile("/home/mikasa/go/src/awesomeProject/tokyo.txt",*typemap)//读配置文件
	packagename:="MyClass"
	structname:="Myclass"
	*typemap=writeFileBytxt(*typemap,packagename,structname)//生成结构体
	/*将go文件写入到goPath底下的src里*/
	//mclass:=reflect.TypeOf(MyClass.Myclass{})//
	//if mclass.Kind() == reflect.Ptr {
	//	mclass = mclass.Elem()
	//}
	//myclass:=reflect.New(mclass)//通过反射new一个对象去是使用这个类
	//injectionValueAndDisplay(*typemap,mclass)//注入值，打印里边的内容
}
