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
		processFile(filePath)
	}
}

func processFile(filePath string) {
	// 1. ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Openメソッドでエラーが発生しました：", err)
		return
	}
	defer file.Close()

	// 2. 正規表現パターンを定義
	pattern := `[0-9]+[A-Za-z]+`
	regExp := regexp.MustCompile(pattern)

	// 3. ファイルの内容をすべて読み込む
	fileContents, err := readAllLines(file)
	if err != nil {
		fmt.Println("ファイルの読み込み中にエラーが発生しました：", err)
		return
	}

	// 4. ファイルの内容に、正規表現にマッチする文字列が存在するかをチェック
	matchedStrings := findMatchedStrings(regExp, fileContents)

	// 5. マッチする文字列が存在しているかどうかを出力
	printResult(filePath, matchedStrings)
}

func findMatchedStrings(regExp *regexp.Regexp, fileContents []string) []string {
	var matchedStrings []string
	for _, content := range fileContents {
		if regExp.MatchString(content) {
			matchedStrings = append(matchedStrings, content)
		}
	}
	return matchedStrings
}

// 正規表現にマッチしたかどうかの結果を出力
func printResult(filePath string, matchedStrings []string) {
	if len(matchedStrings) > 0 {
		fmt.Println("対象文字列が存在します:", filePath)
		for i, str := range matchedStrings {
			fmt.Printf("  L%d: %s\n", i+1, str)
		}
		fmt.Println()
	} else {
		fmt.Printf("対象文字列は存在しません: %s\n", filePath)
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
