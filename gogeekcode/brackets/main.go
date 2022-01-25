package brackets

func main() {

	var name string
	var folder string
	var mod string

	// ...
	{
		prompt := &survey.Input{
			Message: "请输入目录名称：",
		}
		err := survey.AskOne(prompt, &name)
		if err != nil {
			return err
		}
		// ...
	}
	{
		prompt := &survey.Input{
			Message: "请输入模块名称(go.mod中的module, 默认为文件夹名称)：",
		}
		err := survey.AskOne(prompt, &mod)
		if err != nil {
			return err
		}
		// ...
	}
	{
		// 获取hade的版本
		client := github.NewClient(nil)
		prompt := &survey.Input{
			Message: "请输入版本名称(参考 https://github.com/gohade/hade/releases，默认为最新版本)：",
		}
		err := survey.AskOne(prompt, &version)
		if err != nil {
			return err
		}
		// ...
	}
}
