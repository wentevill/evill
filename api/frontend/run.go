package main

import "evill/einit"

func main() {
	if _, err := einit.Init(einit.Log, "./config/api.yml"); err != nil {
		panic(err)
	}

	newRouter()
}
