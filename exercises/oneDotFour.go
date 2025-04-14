package main

import (
	"fmt"
	"os"
	"strings"
)

func Names() {
	path, _ := os.Getwd()
	var paths [3]string
	var files [3]string
	paths[0] = path + "/files/file1.txt"
	paths[1] = path + "/files/file2.txt"
	paths[2] = path + "/files/file2.txt"
	for i := range paths {
		p, _ := os.ReadFile(paths[i])
		files[i] = string(p)
	}
	dup01 := findDuplicatesFromFileToFile(files[0], files[1])
	dup12 := findDuplicatesFromFileToFile(files[1], files[2])
	dup02 := findDuplicatesFromFileToFile(files[0], files[2])
	fmt.Println(strings.Join(dup01, "\n"))
	fmt.Println(strings.Join(dup12, "\n"))
	fmt.Println(strings.Join(dup02, "\n"))
}

func findDuplicatesFromFileToFile(file1 string, file2 string) (linesDuplicates []string) {
	lines1 := strings.Split(file1, "\n")
	lines2 := strings.Split(file2, "\n")
	var result []string
	for i := range lines1 {
		for y := range lines2 {

			if lines1[i] == lines2[y] {
				result = append(result, lines1[i])
			}
		}
	}
	return result
}
