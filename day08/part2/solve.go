package main

import (
	"aoc2019/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func solve(s string, w int, h int) string {
	s = strings.TrimSpace(s)
	ppl := w * h
	layers := make([]map[utils.Point]int, 0, len(s)/ppl)
	i := 0
	for _, char := range s {
		if i%ppl == 0 {
			i = 0
			layers = append(layers, make(map[utils.Point]int, ppl))
		}
		layers[len(layers)-1][utils.Point{
			X: i % w,
			Y: i / w,
		}] = utils.HandledAtoi(string(char))
		i++
	}
	image := make(utils.HashGrid, ppl)
	for _, layer := range layers {
		for p, v := range layer {
			if _, ok := image[p]; ok {
				continue
			}
			switch v {
			case 0:
				image[p] = false
			case 1:
				image[p] = true
			case 2:
				continue
			default:
				log.Fatalf("Invalid pixel type: %d\n", v)
			}
		}
		if len(image) == ppl {
			continue
		}
	}
	bhg := utils.BoundedHashGrid{
		Grid: image,
		W:    w,
		H:    h,
	}

	return bhg.GetBoundedHash()
}

func main() {
	var inputPath string
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	} else {
		_, currentFilePath, _, _ := runtime.Caller(0)
		dir := filepath.Dir(currentFilePath)
		dir = filepath.Dir(dir)
		inputPath = filepath.Join(dir, "input.in")
	}
	contents, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Error reading file %s:\n%v\n", inputPath, err)
		return
	}
	fmt.Println(solve(string(contents), 25, 6))
}
