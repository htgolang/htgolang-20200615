package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("user.csv")

	if err == nil {
		defer file.Close()

		writer := csv.NewWriter(file)

		writer.Write([]string{"编号", "名字", "性别"})
		writer.Write([]string{"1", "烟灰", "男"})
		writer.Write([]string{"2", "祥哥", "男"})
		writer.Flush()
	}
}
