package initialize

import "os"

func RuntimeDir() {
	// 判断 runtime 文件夹是否存在
	if _, err := os.Stat("runtime"); os.IsNotExist(err) {
		// 创建 runtime 文件夹
		os.Mkdir("runtime", 0755)
	}

	// 判断 runtime/logs 文件夹是否存在
	if _, err := os.Stat("runtime/logs"); os.IsNotExist(err) {
		// 创建 runtime/logs 文件夹
		os.Mkdir("runtime/logs", 0755)
	}

	// 判断 runtime/data 文件夹是否存在
	if _, err := os.Stat("runtime/data"); os.IsNotExist(err) {
		// 创建 runtime/data 文件夹
		os.Mkdir("runtime/data", 0755)
	}
}
