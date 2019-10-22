package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

/*
 * @author: taohu
 * @date: 19-7-10
 * @time: 下午1:11
 * @desc: please describe function
**/

func main() {
	f, err := excelize.OpenFile("music_label.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows := f.GetRows("A")
	labelCategories := make(map[int]string, 0)
	labels := make(map[int]string, 0)
	for j, row := range rows {
		if j == 0 {
			for i, colCell := range row {
				if i > 1 {
					labels[i] = colCell
				}
			}
		}

		if j == 1 {
			temp := "场景"
			for i, colCell := range row {
				if i > 1 {
					labelCategories[i] = temp
					if colCell != "" && temp != colCell {
						temp = colCell
					}

					if colCell == "" {
						labelCategories[i] = temp
					}

				}
			}
		}

		var musicId string
		var musicName string
		if j > 1 {
			for i, colCell := range row {

				if i == 0 {
					musicName = colCell
				}
				if i == 1 {
					musicId = colCell
				}

				if i > 1 && colCell != "" {
					fmt.Printf("musicId: %s, musicName: %s, label_category: %s, label: %s \n",
						musicId, musicName, labelCategories[i], labels[i])
				}
			}
		}

		fmt.Println()
	}
	fmt.Printf("labelCategories: %+v", labelCategories)
	fmt.Printf("labels: %+v", labels)
}
