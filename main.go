package main

import "sport"

func main() {
	s := src.Sport{
		UserName: "",
		Password: "",
		StepRang: "8888-9999",
	}
	s.Login()
	s.PushSetp()
}
