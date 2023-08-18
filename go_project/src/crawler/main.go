package main

import (
	"go_project/src/crawler/engine"
	"go_project/src/crawler/zhenai/parser"
)

func main() {
	engine.Run(
		engine.Request{
			Url:       "https://www.zhenai.com/zhenghun",
			ParseFunc: parser.ParseCityList,
		})
}
