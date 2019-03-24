/*
@Time : 2019-03-24 15:19 
@Author : zhangjun
@File : fetch
@Description:fetch url content
@Run: go run fetch.go https://test.sorryu.cn
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func  main()  {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :%v\n", err)
			os.Exit(1)//进程收到信号挂起，终止执行
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch :reading %s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
