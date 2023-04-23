package main

import (
	"fmt"
	"strconv"
)

// Write code here

//ExecutePipeline

func ExecutePipeline() {
	SingleHash(2)
	MultiHash(2)
	CombineResults(2)
}

//Hash calculating

func SingleHash(data int) string {
	stData := strconv.Itoa(data)
	crs32 := DataSignerCrc32(stData)
	crs32md5 := DataSignerMd5(stData)
	return crs32 + crs32md5
}

func MultiHash(data int) string {
	var result string
	for i := 0; i < 6; i++ {
		result += DataSignerCrc32(strconv.Itoa(i) + SingleHash(data))
	}
	return result
}

func CombineResults(data int) {
	fmt.Println(SingleHash(data) + "_" + MultiHash(data))
}
