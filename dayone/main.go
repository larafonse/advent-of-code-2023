package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	sum := 0
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		fmt.Println(line, getSum(line))
		sum += getSum(line)
	}

	fmt.Println(sum)
}

func getSum(line string) int {
	tempSum := 0
	left := 0
	right := len(line) - 1
	rightFound := false
	leftFound := false
	leftValue := 0
	rightValue := 0
	for left <= right {
		if right == left {
			if rightFound && !leftFound{
				leftValue = int(line[left]) - 48
				leftFound = true
				left ++
			}else {
				rightValue = int(line[right]) - 48
				rightFound = true
				right-- 
			}
		} else{

			if line[left] >= 48 && line[left] <= 57 && !leftFound {
				leftValue = int(line[left]) - 48
				leftFound = true
			} else if !leftFound {
				left++
			}
			if line[right] >= 48 && line[right] <= 57 && !rightFound {
				rightValue = int(line[right]) - 48
				rightFound = true
			} else if !rightFound {
				right--
			}
		}
		if leftFound && rightFound {
			break
		}
	}
	if leftFound && rightFound {
		tempSum += (leftValue * 10) + rightValue
	} else if leftFound && !rightFound {
		tempSum += (leftValue * 10) + leftValue
	} else if !leftFound && rightFound {
		tempSum += (rightValue * 10) + rightValue
	}
	return tempSum
}

func parseLine(line string) string{
	line = strings.ToLower(line)
	line = strings.Replace(line, "zero", "z0o",-1)
	line = strings.Replace(line, "one", "o1e",-1)
	line =strings.Replace(line, "two", "t2o",-1)
	line =strings.Replace(line, "three", "t3e",-1)
	line =strings.Replace(line, "four", "f4r",-1)
	line =strings.Replace(line, "five", "f5e",-1)
	line =strings.Replace(line, "six", "s6x",-1)
	line =strings.Replace(line, "seven", "s7n",-1)
	line =strings.Replace(line, "eight", "e8t",-1)
	line =strings.Replace(line, "nine", "n9e",-1)
	return line
}
