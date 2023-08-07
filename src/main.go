package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const (
	targetDir = "target_files"
	outputFilePath = "output/result.txt"
)

func main() {
	// ディレクトリ内のファイルパスを取得
	filePaths, err := getFilePaths(targetDir)
	if err != nil {
		fmt.Println("getFilePathsメソッドでエラーが発生しました：", err)
		return
	}

	// 結果ファイルを作成
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("出力ファイルを作成できませんでした：", err)
		return
	}
	defer outputFile.Close()

	// 中身が正規表現にマッチするかをチェックし、結果をファイルへ出力
	for _, filePath := range filePaths {
		outputRegExpMatching(filePath, outputFile)
	}

	fmt.Print(`処理が終了しました。` + outputFilePath + `を参照してください。`)
}

func outputRegExpMatching(filePath string, outputFile *os.File) {
	// 1. ファイルを開く
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("ファイルを開く処理中にエラーが発生しました：", err)
		return
	}
	defer file.Close()

	// 2. ファイルの内容をすべて読み込む
	fileContents, err := readAllLines(file)
	if err != nil {
		fmt.Println("ファイルの読み込み中にエラーが発生しました：", err)
		return
	}

	// 3. ファイルの内容に、正規表現にマッチする文字列が存在するかをチェック
	matchedStrings := findMatchedString(fileContents)

	// 4. 結果をファイルに出力
	writeResult(filePath, matchedStrings, outputFile)
}

func findMatchedString(fileContents []string) []string {
	pattern := `[0-9]+[A-Za-z]+`
	regExp := regexp.MustCompile(pattern)
	var matchedStrings []string

	for _, content := range fileContents {
		if regExp.MatchString(content) {
			matchedStrings = append(matchedStrings, content)
		}
	}

	return matchedStrings
}

// 正規表現にマッチしたかどうかの結果を出力
func writeResult(filePath string, matchedStrings []string, outputFile *os.File) {
	if len(matchedStrings) > 0 {
		outputFile.WriteString(fmt.Sprintf("対象文字列が存在します: %s\n", filePath))
		for i, str := range matchedStrings {
			outputFile.WriteString(fmt.Sprintf("  L%d: %s\n", i+1, str))
		}
		outputFile.WriteString("\n")
	} else {
		outputFile.WriteString(fmt.Sprintf("対象文字列は存在しません: %s\n", filePath))
	}
}

// ファイルの内容をすべて読み込んで行ごとの文字列のスライスを返す
func readAllLines(file *os.File) ([]string, error) {
	var lines []string
	 // 35MB 本番のファイルサイズと同じくらい
	buf := make([]byte, 35 * 1024 * 1024)

	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

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
