package main

import (
	"flag"
	"fmt"
	"harbor-clean/harbor"
	"sort"
)

var (
	url         string
	user        string
	password    string
	projectName string
	keepNum     int
	help        bool
)

func init() {
	flag.BoolVar(&help, "h", false, "help message")
	flag.StringVar(&url, "url", "", "harbor地址，例如https://harbor.abc.com")
	flag.StringVar(&user, "user", "", "harbor账号")
	flag.StringVar(&password, "password", "", "harbor密码")
	flag.StringVar(&projectName, "projectName", "", "项目名。注意：all 表示全部")
	flag.IntVar(&keepNum, "keepNum", 0, "每个repo保留的tag个数")
}

func deleteTagByID(harborClient harbor.Client, projectID, keepNum int) (err error) {
	repoNames, err := harborClient.GetRepoNames(projectID)
	if err != nil {
		return
	}
	var size int64
	for _, repoName := range repoNames {
		tags, _ := harborClient.GetRepoTags(repoName)

		//tags内容类似于：[{2190824484 v1 2020-08-28 02:14:13.009841239 +0000 UTC} {53504535 v2 2020-08-14 00:36:48.610531148 +0000 UTC}]
		if len(tags) > keepNum { //tag大于keepNum才执行
			fmt.Printf("当前tag: %-4d，保留的tag: %-4d of %-40s ，开始执行删除\n", len(tags), keepNum, repoName)
			sort.Sort(tags)                ////自定义排序，根据tag的创建时间戳正序排列
			toDeleteTags := tags[keepNum:] //需要删除的tag切片
			for _, tag := range toDeleteTags {
				fmt.Printf("     删除image: %s:%s, 创建时间为: %s\n", repoName, tag.Name, tag.Created)
				err := harborClient.DeleteRepoTag(repoName, tag.Name)
				if err != nil {
					fmt.Printf("image: %s:%s DeleteRepoTag: %s\n", repoName, tag.Name, err)
					continue
				}
				size += tag.Size
			}
			fmt.Printf("repo: %s共清理: %.2f MB\n", repoName, float64(size)/1024/1024)
		} else {
			fmt.Printf("当前tag: %-4d，保留tag: %-4d of %-40s ,无需删除! \n", len(tags), keepNum, repoName)
		}

	}
	return
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if url == "" || user == "" || password == "" || projectName == "" {
		flag.Usage()
		return
	}

	harborClient := *harbor.NewClient(user, password, url)

	if projectName == "all" {
		allid, _ := harborClient.GetAllProjectID()

		for _, id := range allid {
			deleteTagByID(harborClient, id.ID, keepNum)
		}
	} else {
		projectID, _ := harborClient.GetProjectID(projectName)
		deleteTagByID(harborClient, projectID, keepNum)
	}
}
