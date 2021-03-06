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
	"strings"
)

//先读入文件
//按单词扫描
/**
 * map 协程下使用会报错，concurrent map iteration and map write，注意加锁
 * 单goroutines，死锁问题 all goroutines are asleep - deadlock!
 */
var sum = 0
var ch = make(chan string)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
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

	go func() {
		for word, num := range counts {
			sum += num
			setWordNumToCh(word, num )
		}
		close(ch)
	}()

	getWordNumFromCh();
	fmt.Printf("单词【%d】个\n", sum)
}

func countWordNum(word string, counts map[string]int) {
	word = strings.ToLower(word)
	counts[word]++
}

func setWordNumToCh(word string, num int){
	ch <- fmt.Sprintf("单词【%s】出现了【%d】次\n", word, num)
}

func getWordNumFromCh() {
	for data := range ch {
		fmt.Println(data)
	}
}
