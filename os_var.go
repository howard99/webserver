package main

import (
	"fmt"
	"os"
	"strconv"
)

func Getavg(s []int) int {
	sum :=0
	for i:=0; i<len(s); i++ {
		sum +=s[i]
	}
	avg :=sum /len(s)
	return avg
}
//外部取变量参数，[] string
//./os_var 89 90 78 56 ==> [./os_var 89 90 78 56]
//输入几个参数，求平均数

func main(){
var s []int
	//外部取变量参数，[] string 去掉地一个参数本身[/os_var]
	osArgs :=os.Args[1:]
	//fmt.Println(osArgs)

	//遍历数组 把string 转成int 切
for i:=0;i<len(osArgs);i++  {
	item, err :=strconv.Atoi(osArgs[i])
	if err!=nil{
		panic(err)
	}
	s =append(s,item)
}
fmt.Println(s)
fmt.Println("平均值是",Getavg(s))
}
