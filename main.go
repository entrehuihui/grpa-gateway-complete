package main

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strings"
)

var grpcPath = "./service/myrpc/proto"

func main() {

	// 读取grpc文件夹文件
	grpcFileInfo, err := os.ReadDir(grpcPath)
	if err != nil {
		log.Fatalln(err)
	}

	grpcFileList := make([]fs.DirEntry, 0)
	for _, v := range grpcFileInfo {
		if strings.Contains(v.Name(), "_grpc.pb.go") {
			grpcFileList = append(grpcFileList, v)
		}
	}

	// 读取函数方法
	for _, v := range grpcFileList {
		log.Println("准备格式化: ====>>", v.Name())
		name := v.Name()
		name = name[:len(name)-11]
		readFuncMethod(name)

	}
	// 处理service文件
	readServiceFile(grpcFileList)
	// 处理server文件
	readServerFile(grpcFileList)
}

// 读取函数方法
func readFuncMethod(fileName string) {
	file, err := os.Open(grpcPath + "/" + fileName + "_grpc.pb.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	start := 0
	tysInfo := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text() // or
		if line == "// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream." {
			start = 1
			continue
		}
		if start == 1 {
			tysInfo = append(tysInfo, line)
		}
		if line == "}" {
			break
		}
	}
	for k, v := range tysInfo {
		value := v
		value = strings.Replace(value, "*", "*proto.", -1)
		value = strings.Replace(value, ", opts ...grpc.CallOption", "", -1)
		value = strings.Replace(value, "	", "", -1)
		tysInfo[k] = value
	}

	if tysInfo[0][:4] != "type" {
		log.Fatal("func检测不到type")
	}
	if tysInfo[len(tysInfo)-1][:1] != "}" {
		log.Fatal("func检测不到}")
	}

	tysInfo = tysInfo[1 : len(tysInfo)-1]
	funcInfoList := make([]FuncInfo, 0)
	// 首字母大写
	fileNameFunc := strings.ToUpper(fileName[:1]) + fileName[1:]
	fileNameLen := len(fileNameFunc)
	// func (s Service) DeviceReg
	for i := 0; i < len(tysInfo); i++ {
		if tysInfo[i][:fileNameLen] == fileNameFunc {
			info := FuncInfo{
				Data: tysInfo[i],
				Name: strings.Split(tysInfo[i], "(")[0],
			}
			if i != 0 && tysInfo[i-1][:2] == "//" {
				info.Details = tysInfo[i-1]
			}
			funcInfoList = append(funcInfoList, info)
		}
	}
	// 处理operate文件夹
	operateFile := "./service/operate" + "/" + fileName + ".go"
	operateBody := readService(operateFile)
	checkOperateFunc(operateFile, string(operateBody), funcInfoList)

	// 处理service文件夹
	serceFile := "./service/myrpc/service" + "/" + fileName + ".go"
	serviceBody := readService(serceFile)
	checkServiceFunc(serceFile, string(serviceBody), funcInfoList)
}

func readServerFile(grpcFileList []fs.DirEntry) {
	path := "./service/myrpc/server.go"
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// 注入GRPC服务
	dataList := strings.Split(string(body), "// 注入GRPC服务")
	dataList[1] = "// 注入GRPC服务\n"
	for _, v := range grpcFileList {
		name := v.Name()
		name = name[:len(name)-11]
		fileNameFunc := strings.ToUpper(name[:1]) + name[1:]
		dataList[1] += "\tproto.Register" + fileNameFunc + "Server(gs, server)\n"
	}
	dataList[1] += "\t// 注入GRPC服务"

	// 注入GW服务
	dataList = strings.Split(dataList[0]+dataList[1]+dataList[2], "// 注入GW服务")
	dataList[1] = "// 注入GW服务\n"
	for _, v := range grpcFileList {
		name := v.Name()
		name = name[:len(name)-11]
		fileNameFunc := strings.ToUpper(name[:1]) + name[1:]
		dataList[1] += "\terr = proto.Register" + fileNameFunc + `HandlerFromEndpoint(ctx, mux, cfg.grpcPort, opts)
	if err != nil {
		log.Fatal("启动GW错误:", err)
	}` + "\n"
	}
	dataList[1] += "\t// 注入GW服务"

	os.WriteFile(path, []byte(dataList[0]+dataList[1]+dataList[2]), 0666)
}

func readServiceFile(grpcFileList []fs.DirEntry) {

	path := "./service/myrpc/service/service.go"
	body, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	dataList := strings.Split(string(body), "// ##继承")
	dataList[1] = "// ##继承\n"
	for _, v := range grpcFileList {
		name := v.Name()
		name = name[:len(name)-11]
		fileNameFunc := strings.ToUpper(name[:1]) + name[1:]
		dataList[1] += "\tproto.Unimplemented" + fileNameFunc + "Server\n"
	}
	dataList[1] += "\t// ##继承"
	log.Println("写入service服务")
	os.WriteFile(path, []byte(dataList[0]+dataList[1]+dataList[2]), 0666)
}

func checkOperateFunc(fileName string, body string, funcList []FuncInfo) {
	if body == "" {
		// 如果是空的  直接写入文件
		emptyBodyOperate(fileName, funcList)
		return
	}
	// 如果不是空
	funcInfo := append([]FuncInfo{}, funcList...)

	stringList := strings.Split(body, "func ")
	for _, v := range stringList {
		for k, value := range funcInfo {
			if value.Name+"(" == v[:len(value.Name)+1] {
				if len(funcInfo) == 1 {
					return
				}
				funcInfo = append(funcInfo[:k], funcInfo[k+1:]...)
				break
			}
		}
	}
	if len(funcInfo) == 0 {
		return
	}
	for _, v := range funcInfo {
		body += "\n" + createOpetateFunc(v)
	}
	// 写入文件
	os.WriteFile(fileName, []byte(body), 0666)
}

func emptyBodyOperate(fileName string, funcInfo []FuncInfo) {
	body := `package operate

import (
	"context"
)
	
`
	for _, v := range funcInfo {
		body += createOpetateFunc(v)
	}

	os.WriteFile(fileName, []byte(body), 0777)
}

func createOpetateFunc(v FuncInfo) string {
	body := v.Details + "\n" + "func " + v.Data + `{
	return nil,nil` + `
}` + "\n\n"

	return body
}

type FuncInfo struct {
	Details string
	Data    string
	Name    string
}

// 检测文件是否写入函数
func checkServiceFunc(fileName string, body string, funcList []FuncInfo) {
	if body == "" {
		// 如果是空的  直接写入文件
		emptyBody(fileName, funcList)
		return
	}
	// 如果不是空
	funcInfo := append([]FuncInfo{}, funcList...)

	stringList := strings.Split(body, "func (s Service) ")
	for _, v := range stringList {
		for k, value := range funcInfo {
			if value.Name+"(" == v[:len(value.Name)+1] {
				if len(funcInfo) == 1 {
					return
				}
				funcInfo = append(funcInfo[:k], funcInfo[k+1:]...)
				break
			}
		}
	}
	for _, v := range funcInfo {
		body += "\n" + createFunc(v)
	}
	// 写入文件
	os.WriteFile(fileName, []byte(body), 0666)

}

func emptyBody(fileName string, funcInfo []FuncInfo) {
	body := `package service

import (
	"context"
	"test/service/myrpc/proto"
	"test/service/operate"
)
	
`
	for _, v := range funcInfo {
		body += createFunc(v)
	}

	os.WriteFile(fileName, []byte(body), 0777)
}

func createFunc(v FuncInfo) string {
	body := v.Details + "\n" + "func (s Service) " + v.Data + `{
	return operate.` + v.Name + `(ctx, in)
}` + "\n\n"

	return body
}

func osStat(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Fatal(err)
	return false
}

func readService(fileName string) []byte {
	if !osStat(fileName) {
		// 如果不存在 创建文件
		return make([]byte, 0)
	}

	body, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
