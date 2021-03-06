package main

import (
	mtols "./mytools"
	"fmt"
	"math"
	"math/big"
	"time"
)

func mysort(arr [5]int) { //桶排序
	var bucket [5 + 1]int //6个桶
	for _, v := range arr {
		bucket[v]++
	}
	for i := 1; i < 6; i++ { //遍历桶
		for j := 0; j < bucket[i]; j++ {
			fmt.Print(i)
		}
	}
	fmt.Println()
}
func testarr() {
	var arr [5]int
	arr0 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(arr0[2:4])
	fmt.Println("请输入5个数")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &(arr[i]))
	}
	t1 := time.Now()
	mysort(arr)
	t2 := time.Now()
	fmt.Println(t2.Sub(t1).Seconds()*1000, "ms")
}
func pointarrtest() {
	//指向数组的指针
	var arr = []int{1, 2, 3}
	var arrp *[]int
	arrp = &arr
	fmt.Println((*arrp)[1])
	//保存指针的数组
	a0 := 1
	var parr = [3]*int{&a0}
	fmt.Println(*(parr[0]))
}

type Animal struct {
	Name string
	Type string
	Id   int
}

func structtest() {
	var a Animal = Animal{"Luck", "Dog", 10}
	a.Type = "Cat"
	fmt.Println(a)
	an := new(Animal)
	an.Type = "狗"
	fmt.Println(an)
}
func stringtest() {
	var str string = "0123456"
	fmt.Printf(string(str[2:]) + "\n")
	str = `D:\tt\1.txt`
	fmt.Println(str)
}
func packagetest() {
	mtols.SayHello("阿良")
}
func mysqrt(i float64) (r float64, ok bool) {
	if i < 0 {
		return
	}
	return math.Sqrt(i), true
}
func labeltest() {
LABEL1:
	for i := 0; i <= 5; i++ {
		if i == 4 {
			continue LABEL1
		}
		fmt.Println(i)
	}
}
func deferf() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func defertest() {
	fmt.Println(deferf())
}
func qiepiantest() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	qp := arr[5:8]
	qp[0] = 88
	fmt.Println(arr, "\r\n", qp, "len:", len(qp), "cap:", cap(qp))
}
func maptest() {
	mp := map[int]string{}
	mp[1] = "一"
	mp[2] = "二"
	mp[3] = "三"
	mp[4] = "四"
	fmt.Println(mp)
}
func bigtest() {
	a := big.NewRat(1, 2)
	b := big.NewRat(1, 3)
	a.Mul(a, b)
	fmt.Println(a)
}
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
func typetest() {
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}
func getsum(chs chan int) {
	s := 0
	for i := 0; i < 100; i++ {
		s += i
	}
	time.Sleep(time.Microsecond * 1000)
	chs <- s
}
func getxiang(ch chan float64) {
	for i := 1; i < 1000000; i += 2 {
		ch <- 4.0 / float64(i)
	}
	close(ch)
}
func getpi() {
	t1 := time.Now()
	defer func() {
		t2 := time.Now()
		t3 := t2.Sub(t1)
		fmt.Println(t3.Seconds() * 1000)
	}()
	s := float64(0.0)
	ch := make(chan float64)
	go getxiang(ch)
	i := 1
	for c := range ch {
		s += c * float64(i)
		i = -i
		//fmt.Println(s)
	}
	fmt.Println(s)
}
func channeltest() {
	chi := make(chan int) //不带缓冲

	//chi:=make(chan int,1)//带1个缓冲
	//chi<-1

	go getsum(chi)
	fmt.Print(<-chi)
}

func channeltemp() {
	ch := make(chan int, 1)
	ch <- 5
	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("无数据")
	}
	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("无数据")
	}
}
func arrtest() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr2 := append(arr[:3], arr[4:]...)
	fmt.Println(arr2)
}
func test() {
	//testarr()
	//pointarrtest()
	//stringtest()
	//packagetest()
	//labeltest()
	//qiepiantest()
	//maptest()
	//bigtest()
	//structtest()
	//typetest()
	//channeltest()
	//getpi()
	//channeltemp()
	arrtest()
}
func main() {
	test()
}
