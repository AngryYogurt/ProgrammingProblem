package main

import (
	"fmt"
	"github.com/AngryYogurt/leetcode/tools"
	"path/filepath"
)

func main() {
	currPath, err := filepath.Abs("./Leetcode")
	if err != nil {
		fmt.Println(err)
	}
	tools.GenTemplate(currPath, 89, "gray-code")
}