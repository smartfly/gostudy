package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/*
 * @author taohu
 * @date 2019/9/12
 * @time 下午5:43
 * @desc 文件操作tools
**/

type Lines struct {
	values []string
}

// 读取大文件
func ReadFile(filePath string) ([]Lines, error) {
	batchSize := 2000
	var linesList []Lines
	var lines Lines

	fp, err := os.Open(filePath)
	defer fp.Close()
	if err != nil {
		fmt.Errorf("file: %s open fail: %s", filePath, err)
		return linesList, err
	}

	buf := bufio.NewReader(fp)
	iterator := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("meet file ending")
			} else {
				fmt.Println(err)
			}
			break
		}
		iterator++
		value := strings.TrimSpace(line)
		// 将文件切割成每2000行一份
		if iterator <= batchSize {
			lines.values = append(lines.values, value)
		} else {
			iterator = 1
			linesList = append(linesList, lines)
			lines.values = nil
			lines.values = append(lines.values, value)
		}
	}
	linesList = append(linesList, lines)
	fmt.Println("script end...", time.Now().Format("2006-01-02 15:04:05"))

	return linesList, nil
}
