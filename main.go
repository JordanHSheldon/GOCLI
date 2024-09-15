package main

import (
	"fmt"
	"math/big"
	"os"
)

var recursive bool
var human bool

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
)

func main() {
	// read command line arguments
	args := os.Args

	// validate we have enough arguments for execution.
	// if not give message on how to use application
	if len(args) < 2 {
		println("Usage: GOCLI [directories...] [--flag]")
		return
	}

	// skip first arg as we do not want to traverse our own program.
	// check for flags and set them, then extract directories
	directories := []string{}
	for i := 1; i < len(args); i++ {
		if args[i] == "--recursive" {
			recursive = true
			continue
		}
		if args[i] == "--human" {
			human = true
			continue
		}

		directories = append(directories, args[i])
	}

	// if no directories exit
	if len(directories) == 0 {
		println("please supply a directory.")
		return
	}

	total := big.NewInt(0)

	// count total of all directories
	// traverse and print out each directory's size if it exists
	for _, dir := range directories {
		dirSize := readDirRecursive(dir, big.NewInt(0))
		total.Add(total, dirSize)
		print("\nSize of " + dir + " : ")
		CustomPrint(dirSize)
		println("\n-----------------------------------")
	}

	print("\nSize of all directories: ")
	CustomPrint(total)
	println()
}

// Recursive function to read directory size
func readDirRecursive(path string, size *big.Int) *big.Int {
	dir, err := os.Open(path)
	if err != nil {
		println("file not found:")
		return big.NewInt(0)
	}
	defer dir.Close()

	names, _ := dir.Readdirnames(0)
	for _, name := range names {
		filePath := fmt.Sprintf("%v/%v", path, name)
		file, err := os.Open(filePath)
		if err != nil {
			println("could not open file:", filePath)
			continue
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			println("could not get file info:", filePath)
			continue
		}

		if fileInfo.IsDir() {
			dirSize := readDirRecursive(filePath, big.NewInt(0))
			size.Add(size, dirSize)
			if recursive {
				fmt.Print(filePath + ": ")
				CustomPrint(dirSize)
			}
		} else {
			fileSize := big.NewInt(fileInfo.Size())
			size.Add(size, fileSize)
			if recursive {
				fmt.Print(file.Name() + ": ")
				CustomPrint(fileSize)
			}
		}
	}

	return size
}

// CustomPrint function to print sizes in human-readable format or as raw bytes
func CustomPrint(size *big.Int) {
	if !human {
		fmt.Println(size.String())
		return
	}

	sizeFloat := new(big.Float).SetInt(size)
	tb := new(big.Float).SetFloat64(TB)
	gb := new(big.Float).SetFloat64(GB)
	mb := new(big.Float).SetFloat64(MB)
	kb := new(big.Float).SetFloat64(KB)

	switch {
	case sizeFloat.Cmp(tb) >= 0:
		result := new(big.Float).Quo(sizeFloat, tb)
		fmt.Printf("%.2f TB\n", result)
	case sizeFloat.Cmp(gb) >= 0:
		result := new(big.Float).Quo(sizeFloat, gb)
		fmt.Printf("%.2f GB\n", result)
	case sizeFloat.Cmp(mb) >= 0:
		result := new(big.Float).Quo(sizeFloat, mb)
		fmt.Printf("%.2f MB\n", result)
	case sizeFloat.Cmp(kb) >= 0:
		result := new(big.Float).Quo(sizeFloat, kb)
		fmt.Printf("%.2f KB\n", result)
	default:
		fmt.Printf("%s B\n", size.String())
	}
}
