package main

import "fmt"

//map的长度是不固定的，也就是和slice一样，也是一种引用类型
//map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
func main() {
	simTest()
	advTest()
}

func simTest() {
	// 声明一个key是字符串，值为int的字典,
	var numbers = map[string]int{"five": 5, "four": 4, "three": 3, "two": 2}
	numbers["one"] = 1 //赋值
	println(numbers["three"])
	key := "nine"
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	v, ok := numbers[key]
	if ok {
		fmt.Println(key, v)
	} else {
		fmt.Println("no " + key)
	}
	for nk := range numbers {
		fmt.Println("Number", nk, "is", numbers[nk])
	}
}

func advTest() {
	/* 创建集合 */
	//make用于内建类型（map、slice 和channel）的内存分配,new用于各种类型的内存分配。
	countryCapitalMap := make(map[string]string)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* 使用 key 输出 map 值 */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 查看元素在集合中是否存在 */
	captial, ok := countryCapitalMap["United States"]
	/* 如果 ok 是 true, 则存在，否则不存在 */
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	/* 删除元素 */
	delete(countryCapitalMap, "Japan")
	fmt.Println("Entry for Japan is deleted")
	for k, v := range countryCapitalMap {
		fmt.Println("Capital of", k, "is", v)
	}
}
