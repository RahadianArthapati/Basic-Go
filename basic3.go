package main

import (
	"fmt"
)

func main() {
	calculateByte()
	interface1()
}
func calculateByte() {
	fmt.Println(ByteSize(2048))       // Outputs: 2.00KB
	fmt.Println(ByteSize(3292528.64)) // Outputs: 3.14MB”
}
func interface1() {
	apple := Apple{"Golden Delicious"}
	orange := Orange{"Yellow", "large"}
	PrintFruit(apple)
	PrintFruit(orange)
}

/*
const (
    KB = 1024
    MB = KB * 1024
    GB = MB * 1024
    TB = GB * 1024
    PB = TB * 1024
)*/
const (
	/* in the first case, 1 is multiplied by 2 ten times—creating 1,024 for KB,
	   and in the second case, 1 is multiplied by 2 twenty times—creating 1,048,576 for MB.*/
	KB ByteSize = 1 << (10 * (iota + 1))
	MB
	GB
	TB
	PB
)

type ByteSize float64

func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return "Very Big"
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%dB", b)
}

type Fruit interface {
	AString() string
}

type Apple struct {
	Variety string
}

func (a Apple) AString() string {
	return fmt.Sprintf("A %s apple", a.Variety)
}

type Orange struct {
	Variety string
	Size    string
}

func (o Orange) AString() string {
	return fmt.Sprintf("A %s %s orange", o.Variety, o.Size)
}

func PrintFruit(fruit Fruit) {
	fmt.Println("I have this fruit:", fruit.AString())
}
