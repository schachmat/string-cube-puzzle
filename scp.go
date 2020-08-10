package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coords struct {
	X, Y, Z int
}

var (
	null = coords{0, 0, 0}
	xPos = coords{1, 0, 0}
	xNeg = coords{-1, 0, 0}
	yPos = coords{0, 1, 0}
	yNeg = coords{0, -1, 0}
	zPos = coords{0, 0, 1}
	zNeg = coords{0, 0, -1}
)

func Add(a, b coords) coords {
	return coords{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (c coords) String() string {
	if c.X == 1 && c.Y == 0 && c.Z == 0 {
		return "+X"
	} else if c.X == -1 && c.Y == 0 && c.Z == 0 {
		return "-X"
	} else if c.X == 0 && c.Y == 1 && c.Z == 0 {
		return "+Y"
	} else if c.X == 0 && c.Y == -1 && c.Z == 0 {
		return "-Y"
	} else if c.X == 0 && c.Y == 0 && c.Z == 1 {
		return "+Z"
	} else if c.X == 0 && c.Y == 0 && c.Z == -1 {
		return "-Z"
	}
	return fmt.Sprintf("{%d %d %d}", c.X, c.Y, c.Z)
}

type space struct {
	cell []bool
	w, h, d int
}

func NewSpace(w, h, d int) *space {
	return &space {
		cell: make([]bool, w*h*d),
		w: w,
		h: h,
		d: d,
	}
}

func (s* space) Set(c coords, val bool) error {
	x, y, z := c.X, c.Y, c.Z
	if x < 0 || y < 0 || z < 0 || x >= s.w || y >= s.h || z >= s.d {
		return fmt.Errorf("coordinates %d %d %d out of range", x, y, z)
	}
	if s.IsOccupied(c) {
//		log.Printf("coordinates %d %d %d already occupied", x, y, z)
	}
	s.cell[x + s.w*y + s.w*s.h*z] = val
	return nil
}

func (s* space) Occupy(c coords) error {
	return s.Set(c, true)
}

func (s* space) Free(c coords) error {
	return s.Set(c, false)
}

func (s* space) IsOccupied(c coords) bool {
	x, y, z := c.X, c.Y, c.Z
	if x < 0 || y < 0 || z < 0 || x >= s.w || y >= s.h || z >= s.d {
//		log.Printf("coordinates %d %d %d out of range", x, y, z)
		return true
	}
	return s.cell[x + s.w*y + s.w*s.h*z]
}

func (s* space) IsFree(c coords) bool {
	return !s.IsOccupied(c)
}

func cubicRoot(c int) int {
	for n := 1; n<10; n++ {
		if n*n*n == c {
			return n
		}
	}
	return 0
}

type dimension int

const (
	noDim dimension = iota
	dimX
	dimY
	dimZ
)

func (s* space) SegmentFree(start coords, dir coords, length int) {
	cur := Add(start, dir)
	for i := 1; i < length; i++ {
		if err := s.Free(cur); err != nil {
			log.Printf("Error freeing segment node %s: %v", cur, err)
		}
		cur = Add(cur, dir)
	}
}

func (s* space) SegmentOccupyIfFree(start coords, dir coords, length int) (bool, coords) {
	if length <= 1 {
		return true, start
	}
	c := Add(start, dir)
	if s.IsOccupied(c) {
		return false, null
	}

	allFree, newStart := s.SegmentOccupyIfFree(c, dir, length-1)
	if allFree {
		if err := s.Occupy(c); err != nil {
			log.Printf("Error while occupying free segment node %s: %v", c, err)
			return false, null
		}
	}
	return allFree, newStart
}

func (s* space) Recurse(start coords, lastDim dimension, snake []int) (string, error) {
	if len(snake) == 0 {
		return "done", nil
	}
	if lastDim != dimX {
		if free, next := s.SegmentOccupyIfFree(start, xPos, snake[0]); free {
			res, err := s.Recurse(next, dimX, snake[1:])
			if err == nil {
				return fmt.Sprintf("%d*%s -> %s", snake[0], xPos, res), nil
			}
			s.SegmentFree(start, xPos, snake[0])
		}
		if free, next := s.SegmentOccupyIfFree(start, xNeg, snake[0]); free {
			res, err := s.Recurse(next, dimX, snake[1:])
			if err == nil {
				return fmt.Sprintf("%d*%s -> %s", snake[0], xNeg, res), nil
			}
			s.SegmentFree(start, xNeg, snake[0])
		}
	}
	if lastDim != dimY {
		if free, next := s.SegmentOccupyIfFree(start, yPos, snake[0]); free {
			res, err := s.Recurse(next, dimY, snake[1:])
			if err == nil {
				return fmt.Sprintf("%d*%s -> %s", snake[0], yPos, res), nil
			}
			s.SegmentFree(start, yPos, snake[0])
		}
		if free, next := s.SegmentOccupyIfFree(start, yNeg, snake[0]); free {
			res, err := s.Recurse(next, dimY, snake[1:])
			if err == nil {
				return fmt.Sprintf("%d*%s -> %s", snake[0], yNeg, res), nil
			}
			s.SegmentFree(start, yNeg, snake[0])
		}
	}
	if lastDim != dimZ {
		if free, next := s.SegmentOccupyIfFree(start, zPos, snake[0]); free {
			res, err := s.Recurse(next, dimZ, snake[1:])
			if err == nil {
				return fmt.Sprintf("%d*%s -> %s", snake[0], zPos, res), nil
			}
			s.SegmentFree(start, zPos, snake[0])
		}
		if free, next := s.SegmentOccupyIfFree(start, zNeg, snake[0]); free {
			res, err := s.Recurse(next, dimZ, snake[1:])
			if err == nil {
				return fmt.Sprintf("%d*%s -> %s", snake[0], zNeg, res), nil
			}
			s.SegmentFree(start, zNeg, snake[0])
		}
	}
	return "", errors.New("no solution found :(")
}

func main() {
	snake := []int{}
	sum := 1
	for _, a := range os.Args[1:] {
		slen, err := strconv.Atoi(a)
		if err != nil {
			log.Fatalf("failed to parse cmdline arg %q to integer segment length", a)
		}
		snake = append(snake, slen)
		sum += slen - 1
	}
	fmt.Println("snake segment lengths:", snake)

	size := cubicRoot(sum)
	if size == 0 {
		log.Fatalf("segment length sum %d is not a cubic number", sum)
	}

	s := NewSpace(size, size, size)
	start := coords{0, 1, 0}

	s.Occupy(start)
	res, err := s.Recurse(start, dimZ, snake)
	if err != nil {
		log.Fatalf("Recursion failed: %v", err)
	}

	fmt.Println(s)
	for z := 0; z < size; z++ {
		for y := size-1; y >= 0; y-- {
			for x := 0; x < size; x++ {
				if s.IsOccupied(coords{x, y, 0}) {
					fmt.Printf("X ")
				} else {
					fmt.Printf("o ")
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	fmt.Println(res)
}
