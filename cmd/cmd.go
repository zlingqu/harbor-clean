package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zlingqu/harbor-clean/clean"
)

var (
	url         string
	user        string
	password    string
	projectName string
	keepNum     int
)

func NewHarborCleanCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "harbor-clean",
		Short:   "harbor 仓库镜像清理",
		Long:    "harbor-clean 用于清理harbor的仓库中的tag，以释放存储资源",
		Example: "harbor-clean --url https://harbor.abc.com  --user admin --password FJLSDfdso3489X --projectName abc-web --keepNum 200",
		Run: func(cmd *cobra.Command, args []string) {
			clean.Clean(url, user, password, projectName, keepNum)
		},
	}
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "例如：https://harbor.abc.com")
	rootCmd.Flags().StringVarP(&user, "user", "U", "", "用户名，例如：admin")
	rootCmd.Flags().StringVarP(&password, "password", "p", "", "密码")
	rootCmd.Flags().StringVarP(&projectName, "projectName", "P", "", "项目名，all表示所有项目")
	rootCmd.Flags().IntVarP(&keepNum, "keepNum", "k", 0, "保留的tag数目，例如50")
	return rootCmd
}
