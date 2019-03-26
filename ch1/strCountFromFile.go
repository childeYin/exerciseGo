/*
@Time : 2019-03-26 10:23 
@Author : zhangjun
@File : strCountFromFile
@Description: count the words num from file
@Run: go run strCountFromFile.go  fileName
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

//先读入文件
//按单词扫描

var sum = 0
func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	ch := make(chan string)
	for _, file := range files {
		f, err := os.Open(file)
		if err == nil {
			fmt.Fprintf(os.Stderr, "strCountFromFile:%v\n", err)
		}
		input := bufio.NewScanner(f)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			word := input.Text()
			countWordNum(word, counts)
		}
	}

	for word, num := range counts {
		sum += num
		go setWordNumToCh(word, num , ch )
		getWordNumFromCh(ch)
	}
	fmt.Printf("单词【%d 】个\n", sum)

}

func countWordNum(word string, counts map[string]int) {
	counts[word]++
}

func setWordNumToCh(word string, num int, ch chan string){
	ch <- fmt.Sprintf("单词【%s 出现了【%d】次\n", word, num)
}

func getWordNumFromCh(ch chan string){
	fmt.Printf(<-ch)
}
