package model

import (
	"time"
)

// Project /api/projects?name=cloud
type Project struct {
	Name string `json:"name"`
	ID   int    `json:"project_id"`
}

// Repo /api/repositories?project_id=2
type Repo struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Tag /api/repositories/cloud/demojava/tags  test-bcd2e9d
type Tag struct {
	Size    int64     `json:"size"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

// Tags implement the sort interface
type Tags []Tag //结构体类型的切片

//以下三个函数用于自定义排序，自定义排序必须实现这3个方法。表示根据创建时间
func (t Tags) Len() int           { return len(t) }
func (t Tags) Less(i, j int) bool { return t[i].Created.After(t[j].Created) } //时间比较，i在j之后，则交互i，j。可理解为按照时间戳大小正序排序
func (t Tags) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
