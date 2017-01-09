package main

import (
	"fmt"
	"sort"
)

func main() {
	// createMap()
	// mapFunc()
	// slicePointerMap()
	// keyInMap()
	// orderMap()
	// showWeeks()
	// sliceMap()
	// sortMapByKey()
	inverseMap()
}

func inverseMap() {
	barMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"five":  5,
		"four":  4,
	}

	fmt.Println(barMap)
	invMap := make(map[int]string, len(barMap))

	for k, v := range barMap {
		invMap[v] = k
	}

	fmt.Println(invMap)
}

func sortMapByKey() {
	// 将map的key转换成slice，利用slice进行排序，然后再根据slice排序的结果输出map
	barMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"five":  5,
		"four":  4,
	}

	fmt.Println("unsorted: ", barMap)

	keys := make([]string, len(barMap))

	i := 0
	for k, _ := range barMap {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, barMap[k])
	}
}

func sliceMap() {
	// 这个语句创建了一个含有5个元素的切片，每个元素都是一个空的map
	items := make([]map[int]int, 5)

	for i, v := range items {
		fmt.Println(i, v)
	}

	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][5] = i
	}

	fmt.Println(items)
}

func createMap() {
	// 1. literal方式创建
	// 2. make创建
	// 3. 赋值的方式创建
	var mapLit = map[string]int{"one": 1, "two": 2}
	var mapAssigned map[string]int

	var mapCreated = make(map[string]float32)

	mapAssigned = mapLit

	mapLit["key"] = 1
	mapAssigned["key1"] = 3
	mapCreated["float32"] = 3.4

	fmt.Println(mapLit["key"])
	fmt.Println(mapAssigned["key1"])
	fmt.Println(mapCreated["float32"])
}

func mapFunc() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}

	fmt.Println(mf)
}

func slicePointerMap() {
	mp := map[int]*[]int{
		1: &[]int{1, 2, 3},
		2: &[]int{10, 20, 30},
		3: &[]int{100, 200, 300},
	}

	fmt.Println(mp)
}

func keyInMap() {
	mi := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}

	delete(mi, 3)

	for i := 0; i <= len(mi); i++ {
		if value, ok := mi[i]; ok {
			fmt.Println(value)
		} else {
			fmt.Printf("%d Does not exist in mi\n", i)
		}
	}
}

func orderMap() {
	// map的顺序是随机的
	capitals := map[string]string{"Italy": "Rome", "Japan": "Tokyo", "France": "Paris"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

func showWeeks() {
	mw := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}

	for k, v := range mw {
		fmt.Println(k, v)
	}
}
