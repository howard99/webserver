package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

type (
	IpAddr [4]byte
)

func main() {
	fmt.Println("Hello\n")

	//const demo
	const PI float64 = 3.14159
	fmt.Println(PI, "\n")

	//var var string
	var var1 string = "test string"
	fmt.Println(var1, "\n")

	///////////////////////////////////////////////
	//                                           //
	//         array demo                        //
	//                                           //
	///////////////////////////////////////////////
	//array demo
	var array1 [5]int
	array1 = [5]int{1, 2, 3, 4, 5}
	////下面是简写，上面是标准写法
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2, "\n", array1, "\n")

	/////////////////////////////////////////////
	//                                        //
	//              slice demo                //
	//                                        //
	///////////////////////////////////////////

	var slice1 []int = make([]int, 5, 10)
	slice1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//////下面是简写，上面是标准写法
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice1, "\n", slice2, "\n")

	////###slice 使用时需要make
	var slice4 []int = make([]int, 5)
	fmt.Print("slice4 is", slice4, "\n")
	//add some int data
	var slice3 []int
	slice3 = append(slice2, 0, 0, 0, 1, 2, 3, 4)

	//slice2[0] is 1, slice2[5] is 6   ;;;slice2 := []int{1, 2, 3, 4, 5, 6}
	slice3 = append(slice3, slice2[5])
	fmt.Println(slice3, "\n\n")

	// demo 取切片第一个元素，陆陆续续取完
	fmt.Println("slice1 切片 总\n", slice1)
	for len(slice1) > 0 {
		r := slice1[0]
		slice1 = slice1[1:]
		fmt.Println("每次取走切片第一个数是r\n", r)
		time.Sleep(3 * time.Second)
		fmt.Println("剩下的切片数是 slice1\n", slice1)
	}

	//////////////////////////////////////////////////
	//                                              //
	//              map demo                       //
	//                                             //
	////////////////////////////////////////////////
	// need make
	var map_1 map[string]int = make(map[string]int)
	map_1["age"] = 22
	map_1["sum"] = 80
	fmt.Println("map1 is", map_1)

	// map2 demo
	var map_2 map[string]int = map[string]int{"age": 30, "sum": 80}
	map_2["total"] = 70
	fmt.Println("map_2 is", map_2)

	////// 下面是简写，上面是标准写法
	map_3 := map[string]int{"age": 20, "kkk": 30, "sum": 60}
	fmt.Println("map_3 is", map_3)

	hosts := map[string]IpAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name1, ip1 := range hosts {
		fmt.Printf("%v: %v\n", name1, ip1)
	}
	//////////////////////////////////////////////////////////
	//                                                     //
	//                 for loop demo                      //
	//                                                    //
	///////////////////////////////////////////////////////

	// way 1    for i
	varray_8 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i := 0; i < 10; i++ {
		fmt.Println("i val is", i, "\n")
		fmt.Printf("the array is %d  okok", varray_8[i], "\n")
		fmt.Println("the array_8 is", varray_8[i], "\n")
	}

	// way 2   for range
	for key, val := range varray_8 {
		fmt.Printf("key is %d,val is %d \n", key, val)
	}

	// way 3 for true  loop forevery
	// 时间格式必须是这个数 time.Now().Format("2006-01-02 15:04:05")

	for true {
		fmt.Println("time is", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(3 * time.Second)
		break
	}

	//////////////////////////////////////////////////////////
	//                                                     //
	//                  if     demo                        //
	//                                                     //
	////////////////////////////////////////////////////////

	var x int = 30
	if x < 20 {
		fmt.Println("x less 20", x)
	} else {
		fmt.Println("x great than 20")
	}

	if x > 5 {
		fmt.Println("x great 5", x)
	}

	// func if  demo
	fmt.Println(sqrt(9), sqrt(-9))

	//////////////////////////////////////////////////////////
	//                                                     //
	//              switch  case   demo                    //
	//                                                     //
	////////////////////////////////////////////////////////
	fmt.Print("GO runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Print("windows OS")
	default:
		//freebsd,openbsd, plan9..
		fmt.Printf("%s,\n", os)
	}

	//switch  way 2
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch way 3
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	//////////////////////////////////////////////////////////
	//                                                     //
	//            defer   demo                             //
	// defer 栈，后进先出；  chan 管道先进先出                 //
	////////////////////////////////////////////////////////
	defer fmt.Println("world!")
	fmt.Println("Hello")

	//////////////////////////////////////////////////////////
	//                                                     //
	//              point 指针  demo                       //
	//                                                     //
	////////////////////////////////////////////////////////
	i, j := 42, 2701
	p := &i         //point to i
	fmt.Println(*p) //read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i
	p = &j          // point to j
	*p = *p / 37    // divide j through the pointer
	fmt.Println(j)  // see the new value of j

	//////////////////////////////////////////////////////////
	//                                                     //
	//            struct  demo  結構                        //
	//                                                    //
	////////////////////////////////////////////////////////
	fmt.Println(struct1{50, 30})
	a := Person{"Howard", 32}
	b := Person{"Lucky", 31}
	fmt.Println(a, b)

	//////////////////////////////////////////////////////////
	//                                                     //
	//           interface 接口 demo                       //
	//                                                    //
	////////////////////////////////////////////////////////
	do(21)
	do("Hello")
	do(true)
}

//////////////////////////////////////////////////////////
//                                                     //
//            struct  demo  結構                        //
//                                                    //
////////////////////////////////////////////////////////

//struct
type struct1 struct {
	x int
	y int
}

type Person struct {
	Name string
	Age  int
}

//interface
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don'tnow about type %T!\n", v)
	}
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//func  if  Methods 方法
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
