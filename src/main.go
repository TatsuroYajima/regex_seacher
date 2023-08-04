package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	targetDir := "target_files"

	// ディレクトリ内のファイルパスを取得
	filePaths, err := getFilePaths(targetDir)
	if err != nil {
		fmt.Println("getFilePathsメソッドでエラーが発生しました：", err)
		return
	}

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Openメソッドでエラーが発生しました：", err)
			continue
		}
		defer file.Close()

		pattern := `[0-9]+[A-Za-z]+`
		regExp := regexp.MustCompile(pattern)

		// ファイルの内容をすべて読み込む
		fileContents, err := readAllLines(file)
		if err != nil {
			fmt.Println("ファイルの読み込み中にエラーが発生しました：", err)
			continue
		}

		var matchedStrings []error

		// 正規表現にマッチする行を探す
		matched := false
		for i, content := range fileContents {
			if regExp.MatchString(content) {
				matched = true
				matchedStrings = append(matchedStrings, fmt.Errorf("  L%d: %s", i+1, content))
			}
		}

		if matched {
			fmt.Println("対象文字列が存在します:", filePath)
			for _, error := range matchedStrings {
				fmt.Println(error)
			}
			fmt.Println()
		} else {
			fmt.Printf("対象文字列は存在しません: %s\n", filePath)
		}
	}
}

// ファイルの内容をすべて読み込んで行ごとの文字列のスライスを返す
func readAllLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getFilePaths(dirPath string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ファイルの場合のみパスを追加
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}
