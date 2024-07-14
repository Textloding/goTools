package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// 计算文件的SHA-256校验和
func fileChecksum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// 递归遍历目录并计算文件校验和
func walkDir(dirPath string, writer *csv.Writer) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			checksum, err := fileChecksum(path)
			if err != nil {
				return err
			}
			writer.Write([]string{path, checksum})
		}

		return nil
	})
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <directory> <output.csv>")
		return
	}

	dirPath := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"File Path", "SHA-256 Checksum"})

	if err := walkDir(dirPath, writer); err != nil {
		fmt.Println("Error walking the directory:", err)
	}
}
