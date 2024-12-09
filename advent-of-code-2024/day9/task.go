package day9

import (
	"slices"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

func part1(input string) int {
	fileSystem, _ := createFileSystem(input)

	i, j := 0, len(fileSystem)-1
	for i < j {
		if fileSystem[i] != -1 {
			i++
		} else {
			if fileSystem[j] != -1 {
				fileSystem[i], fileSystem[j] = fileSystem[j], fileSystem[i]
			}
			j -= 1
		}
	}

	return calculateChecksum(fileSystem)
}

func part2(input string) int {
	fileSystem, largestFileId := createFileSystem(input)

	for fId := largestFileId; fId >= 0; fId-- {
		fStart, length := findStartAndLength(fileSystem, fId)
		freeIndex := findFreeSpot(fStart, fileSystem, length)
		if freeIndex >= 0 {
			for k := range length {
				swapI := freeIndex + k
				swapJ := fStart + k
				fileSystem[swapI], fileSystem[swapJ] = fileSystem[swapJ], fileSystem[swapI]
			}
		}
	}

	return calculateChecksum(fileSystem)
}

func findFreeSpot(fStart int, fileSystem []int, length int) int {
	free := 0
	for i := 0; i < fStart; i++ {
		if fileSystem[i] == -1 {
			free++
			if free >= length {
				return i - length + 1

			}
		} else {
			free = 0
		}
	}
	return -1
}

func findStartAndLength(fileSystem []int, fId int) (int, int) {
	fStart := slices.Index(fileSystem, fId)
	fEnd := fStart + 1
	for fEnd < len(fileSystem) {
		if fileSystem[fEnd] != fId {
			break
		}
		fEnd++
	}
	length := fEnd - fStart
	return fStart, length
}

func calculateChecksum(fileSystem []int) int {
	checksum := 0
	for i, v := range fileSystem {
		if v > 0 {
			checksum += i * v
		}
	}

	return checksum
}

func createFileSystem(input string) ([]int, int) {
	sizeOfFileSystem := 0
	for _, v := range input {
		sizeOfFileSystem += utils.GetAllNumbers(string(v))[0]
	}
	fileSystem := make([]int, sizeOfFileSystem)

	isFile := true
	fileId := 0
	fileSystemPointer := 0
	for _, v := range input {
		size := utils.GetAllNumbers(string(v))[0]
		for range size {
			if isFile {
				fileSystem[fileSystemPointer] = fileId
			} else {
				fileSystem[fileSystemPointer] = -1
			}

			fileSystemPointer++
		}
		if isFile {
			fileId += 1
		}
		isFile = !isFile
	}
	return fileSystem, fileId - 1
}
