package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("F:\\Siam的资料\\Code\\Go\\dtm_1.12.2_windows_amd64.tar\\dtm_1.12.2_windows_amd64\\dtm.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//rd := bufio.NewReader(f)
	//for {
	//	line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
	//
	//	if err != nil || io.EOF == err {
	//		break
	//	}
	//	fmt.Println(line)
	//}
	var lineText string
	scanner := bufio.NewScanner(f)
	max := 3
	nowLine := 0

	for scanner.Scan() {
		nowLine++
		if nowLine > max {
			break
		}
		lineText = scanner.Text()

		fmt.Println(lineText)
	}

	fmt.Println("扫描结束")
}
