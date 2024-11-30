package main

import (
	"aoc2019/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

type move struct {
	dir utils.Direction
	amt int
}

func solve(s string) int {
	s = strings.TrimSpace(s)
	wires := utils.Map(strings.Split(s, "\n"), func(line string) []move {
		return utils.Map(strings.Split(line, ","), func(mStr string) move {
			var dir utils.Direction
			switch mStr[0] {
			case 'R':
				dir = utils.RIGHT
			case 'U':
				dir = utils.UP
			case 'D':
				dir = utils.DOWN
			case 'L':
				dir = utils.LEFT
			default:
				log.Fatalf("Invlaid direction: %q\n", mStr[0])
			}
			return move{
				dir: dir,
				amt: utils.HandledAtoi(mStr[1:]),
			}
		})
	})

	wire0 := map[utils.Point]bool{}
	intersections := []utils.Point{}
	p := utils.ORIGIN()
	for _, m := range wires[0] {
		for range m.amt {
			p.MoveInDir(m.dir, 1)
			wire0[p] = true
		}
	}
	p = utils.ORIGIN()
	for _, m := range wires[1] {
		for range m.amt {
			p.MoveInDir(m.dir, 1)
			if wire0[p] {
				intersections = append(intersections, p)
			}
		}
	}

	return slices.Min(utils.Map(intersections, func(p utils.Point) int {
		return p.Manhattan()
	}))
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
	fmt.Println(solve(string(contents)))
}
