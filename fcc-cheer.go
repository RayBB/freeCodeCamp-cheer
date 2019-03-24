package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	var oldPoints int
	userPtr := flag.String("user", "raybb", "a string")
	delayPtr := flag.Int("sec", 10, "an int")
	flag.Parse()

	username := *userPtr
	delay := *delayPtr
	fmt.Println("Starting for username: " + username + " delay: " + strconv.Itoa(delay))

	for true {
		userInfo, err := getPoints(username)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Duration(delay) * time.Second)
			continue
		}

		if oldPoints < userInfo.points {
			nameAndScore := userInfo.username + ": " + strconv.Itoa(userInfo.points)
			fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM") + " " + nameAndScore)
			err := beeep.Notify("FreeCodeCamp", nameAndScore, "assets/information.png")
			if err != nil {
				fmt.Println(err)
			}
			oldPoints = userInfo.points
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func getPoints(username string) (UserInfo, error) {
	var userInfo UserInfo

	var fccJSON map[string]interface{}
	err := getJSON("https://www.freecodecamp.org/api/users/get-public-profile?username="+username, &fccJSON)
	if err != nil {
		//fmt.Println(err)
		return userInfo, err
	}

	entities := fccJSON["entities"].(map[string]interface{})
	user := entities["user"].(map[string]interface{})
	vb := user[username].(map[string]interface{})
	points := int(vb["points"].(float64))

	userInfo.points = points
	userInfo.username = username
	return userInfo, nil
}

// UserInfo contains the info from JSON request that we want to use
type UserInfo struct {
	username string
	points   int
}

func getJSON(url string, result interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(result)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
