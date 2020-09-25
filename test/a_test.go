package test

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([][]int, 0)

	for i := 1; i <= 65; i++ {
		n := generateRandomNumber(r, 1, 13, 4)
		arr = append(arr, n)
	}

	delk := checkFront(arr)
	fmt.Println(delk)
	fmt.Println("t---------")

	newArr := reformat(arr, delk)

	fmt.Println(newArr)
	//fmt.Println( checkFront(newArr))

	j, _ := json.Marshal(newArr)
	if f, e := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777); e != nil {
		log.Println(e)
	} else {

		f.WriteString(string(j))
	}
}

func reformat(arr [][]int, delk []int) [][]int {

	if len(delk) == 0 {
		return arr
	}

	newArr := make([][]int, 0)

	for i2, v2 := range arr {
		if inArr(i2, delk) == false {
			newArr = append(newArr, v2)
		}

	}
	return newArr
}
func inArr(i int, arr []int) bool {
	for _, v2 := range arr {
		if v2 == i {
			return true
		}
	}
	return false
}
func checkFront(arr [][]int) (delKeys []int) {
	//allArr := make([int]string, 0)
	allArr := make(map[string]int, 0)

	for key, val := range arr {
		front, _ := json.Marshal(val[:3])
		allArr[string(front)] = key
		tail, _ := json.Marshal(val[1:])
		allArr[string(tail)] = key
	}

	for i, _ := range allArr {
		otherMap := otherArr(allArr, i)

		for k, v := range otherMap {
			if i == k {
				fmt.Println("delk", i, k, v)
				delKeys = append(delKeys, v)
			}
		}
	}
	delKeys = RemoveRepByLoop(delKeys)
	return
}
func otherArr(allArr map[string]int, val string) map[string]int {
	otherMap := make(map[string]int)
	for k, v := range allArr {
		if k != val {
			otherMap[k] = v
		}
	}
	return otherMap
}

func RemoveRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(r *rand.Rand, start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
