package main

import (
	"flag"
	"fmt"
)

func main() {

	poemType := flag.Int("type", 1, "请选择类型")
	poemTang := flag.String("poemTang", "流波将月去，潮水带星来", "请输入一句唐诗")
	poemSong := flag.String("poemSong", "会挽雕工如满月，西北望，射天狼", "请输入一句宋词")

	flag.Parse()
	if *poemType == 1 {
		fmt.Println(*poemTang)
	}

	if *poemType == 2 {
		fmt.Println(*poemSong)
	}

}
