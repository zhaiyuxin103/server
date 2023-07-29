package make

import (
	"fmt"
	"server/pkg/console"
	"server/pkg/file"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeController = &cobra.Command{
	Use:     "controller",
	Short:   "Create api controller，example: make controller v1/user",
	Run:     runMakeController,
	Args:    cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	Example: "make controller v1/user",
}

func runMakeController(cmd *cobra.Command, args []string) {

	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")
	if len(array) > 3 {
		console.Exit("The directory supports up to three levels only, And the last name is the file name")
	}

	// 目录路径
	var dirPath string
	if len(array) == 1 {
		dirPath = "app/http/controllers"
	} else {
		dirPath = fmt.Sprintf("app/http/controllers/%s", strings.ToLower(array[0]))
		if len(array) > 1 {
			for i := 1; i < len(array)-1; i++ {
				dirPath = fmt.Sprintf("%s/%s", dirPath, strings.ToLower(array[i]))
			}
		}
	}

	// 创建目录
	if err := file.CreateDir(dirPath); err != nil {
		console.ExitIf(err)
	}

	model := makeModelFromString(array[len(array)-1])

	// 文件名称
	fileName := fmt.Sprintf("%s_controller.go", model.PackageName)

	var prefix string
	var Import = ""
	var baseController string

	if len(array) > 2 {
		Import = fmt.Sprintf("\"server/app/http/controllers/%s\"", strings.ToLower(array[0]))
	}

	switch len(array) {
	case 1:
		prefix = "controllers"
		baseController = "BaseController"
		break
	case 2:
		prefix = strings.ToLower(array[0])
		baseController = "BaseController"
		break
	case 3:
		prefix = strings.ToLower(array[len(array)-2])
		baseController = array[0] + ".BaseController"
		break
	}

	createFileFromStub(fmt.Sprintf("%s/%s", dirPath, fileName), "controller", model, map[string]string{
		"{{prefix}}":         prefix,
		"{{Import}}":         Import,
		"{{baseController}}": baseController,
	})

	console.Success(fmt.Sprintf("Created Successfully!\nPath: app/http/controllers/%s/%s", strings.ToLower(array[0]), fileName))
}
