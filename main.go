package main

import "CODE/goVue/src/system/app"

func main() {
	s := app.NewServer()
	s.Init()
	s.Start()
}
