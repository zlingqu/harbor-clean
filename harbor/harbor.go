package harbor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"

	"github.com/zlingqu/harbor-clean/model"
)

// Client
type Client struct {
	Client  *http.Client //http.Clinet类型，结构体嵌套
	BaseURL string
}

// NewClient
func NewClient(username, password, baseURL string) *Client {
	client := &http.Client{ //定义client变量，是一个结构体
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) { //匿名函数
				req.SetBasicAuth(username, password)
				return nil, nil
			},
		},
	}
	return &Client{
		Client:  client,
		BaseURL: baseURL,
	}
}

// getProjectID . 根据projectName 获取projectID

func (c *Client) GetProjectID(projectName string) (projectID int, err error) {
	resp, err := c.Client.Get(c.BaseURL + "/api/projects?name=" + projectName)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("response code is:%v", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var projects []model.Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return
	}
	for _, p := range projects { //返回的是模糊查询的结果，所以需要做个判断
		if p.Name == projectName {
			return p.ID, nil
		}
	}

	return 0, errors.New("not found")
}

func (c *Client) GetAllProjectID() (allid []model.Project, err error) {
	resp, err := c.Client.Get(c.BaseURL + "/api/projects")
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("response code is:%v", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var allProjects []model.Project

	err = json.Unmarshal(body, &allProjects) //解析json
	if err != nil {
		return
	}
	for _, a := range allProjects {
		allid = append(allid, a)
	}

	return allid, nil
}

// func (c *Client) GetRepo(projectId int) (repoNames []string, err error)

func (c *Client) GetRepoNames(projectID int) (repoNames []string, err error) {
	resp, err := c.Client.Get(c.BaseURL + "/api/repositories?project_id=" + strconv.Itoa(projectID))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var repos []model.Repo
	err = json.Unmarshal(body, &repos) //json转结构体
	if err != nil {
		return
	}
	for _, repo := range repos { //结构体转切片
		repoNames = append(repoNames, repo.Name)
	}
	return repoNames, nil
}

//获取tag列表
func (c *Client) GetRepoTags(repo string) (tags model.Tags, err error) {
	resp, err := c.Client.Get(c.BaseURL + "/api/repositories/" + repo + "/tags")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &tags) //json转结构体
	if err != nil {
		return
	}
	// 排序
	sort.Sort(tags)
	if !sort.IsSorted(tags) {
		return nil, errors.New("tags not sorted")
	}
	return
}

// DeleteRepoTag delete tags with repo name and tag.
func (c *Client) DeleteRepoTag(repo string, tag string) (err error) {
	request, err := http.NewRequest("DELETE", c.BaseURL+"/api/repositories/"+repo+"/tags/"+tag, nil)
	if err != nil {
		return
	}
	resp, err := c.Client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = fmt.Errorf("resp code=%v", resp.StatusCode)
		return
	}
	return
}
