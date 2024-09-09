package cmd
//验证exe文件模块
import "os"

func PathExists(path string) (bool, error) { //判断文件是否存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func Reshell(path string)bool { //验证shell是否生成成功.
	b, _ := PathExists(path)
	if b {
		return true
	}
	return false
}