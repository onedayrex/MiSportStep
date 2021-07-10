package main

import (
	"encoding/json"
	"os"
	"sport"
	"sync"
)

func main() {
	userInfoJson := os.Getenv("XIAOMI_USER_INFO")
	var userInfoList = make([]src.UserInfo, 0)
	json.Unmarshal([]byte(userInfoJson), &userInfoList)
	var wg = sync.WaitGroup{}
	if len(userInfoList) > 0 {
		for _, info := range userInfoList {
			s := src.Sport{
				UserName: info.UserName,
				Password: info.Password,
				StepRang: info.StepRang,
			}
			wg.Add(1)
			go s.AsyncSport(&wg)
		}
	}
	wg.Wait()
}
