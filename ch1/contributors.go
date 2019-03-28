/*
@Time : 2019-03-28 10:45 
@Author : zhangjun
@File : contributors
@Description:
@Run:
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//GET /repos/:owner/:repo/contributors
var contributorsURL = "https://api.github.com/repos/childeYin/Cultivate/contributors"

var reposContributorsUrl = "https://github.com/childeYin/Cultivate/graphs/contributors"

type User struct {
	Id int
	Login string
	HtmlUrl string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
	Contributions int
}

/*
 * 1.抓取地址信息
 * 2.解析数据
 */

func main()  {
	result, err := getContributorsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("repos Contributors url is 【%s】\n", reposContributorsUrl)

	for _, item := range result {
		fmt.Printf("【%s】constributors number is 【%d】\n", item.Login, item.Contributions)
	}
}

func getContributorsList( ) (contributor []User, err error) {

	var contributors []User
	resp, err := http.Get(contributorsURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed %s", resp.Status)
	}

	//content, err := ioutil.ReadAll(resp.Body)
	//respBody := string(content)
	//fmt.Print(respBody)


	if err := json.NewDecoder(resp.Body).Decode(&contributors); err != nil {
		resp.Body.Close()
		return nil, err
	}
	//fmt.Print(contributors)

	resp.Body.Close()
	return contributors, err






}
