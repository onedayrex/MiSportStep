package main

import (
	"os"
	"sport"
)

func main() {
	userName := os.Getenv("XIAOMI_USERNAME")
	password := os.Getenv("XIAOMI_PASSWORD")
	stepRang := os.Getenv("XIAOMI_STEP_RANG")
	s := src.Sport{
		UserName: userName,
		Password: password,
		StepRang: stepRang,
	}
	s.Login()
	s.PushSetp()
}
