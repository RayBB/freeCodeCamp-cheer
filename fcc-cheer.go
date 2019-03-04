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
		userInfo := &UserInfo{}
		err := getPoints(username, userInfo)
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

func getPoints(username string, userInfo *UserInfo) error {

	//TODO: Add error handling for bad username

	var fccJSON map[string]interface{}
	err := getJSON("https://www.freecodecamp.org/api/users/get-public-profile?username="+username, &fccJSON)
	if err != nil {
		//fmt.Println(err)
		return err
	}

	//freeCodeCampJson := `{"entities":{"user":{"victorious-bean":{"about":"Potato :)","completedChallenges":[{"completedDate":1551487504630,"id":"587d78a4367417b2b2512ad3","files":[]},{"completedDate":1551487391148,"id":"587d78a4367417b2b2512ad2","files":[]},{"completedDate":1551487274107,"id":"587d78a3367417b2b2512ad1","files":[]},{"completedDate":1551487180251,"id":"587d78a3367417b2b2512ad0","files":[]},{"completedDate":1551487141957,"id":"587d78a3367417b2b2512acf","files":[]},{"completedDate":1551486839403,"id":"587d78a3367417b2b2512ace","files":[]},{"completedDate":1551486753921,"id":"587d781e367417b2b2512acc","files":[]},{"completedDate":1551485885856,"id":"587d781e367417b2b2512acb","files":[]},{"completedDate":1551485762498,"id":"587d781e367417b2b2512aca","files":[]},{"completedDate":1551485709326,"id":"587d781e367417b2b2512ac9","files":[]},{"completedDate":1551485542415,"id":"587d781d367417b2b2512ac8","files":[]},{"completedDate":1551474502614,"id":"587d781d367417b2b2512ac5","files":[]},{"completedDate":1551474482269,"id":"587d781c367417b2b2512ac4","files":[]},{"completedDate":1551474463908,"id":"587d781c367417b2b2512ac3","files":[]},{"completedDate":1551474420383,"id":"587d781c367417b2b2512ac2","files":[]},{"completedDate":1551474336158,"id":"587d781c367417b2b2512ac0","files":[]},{"completedDate":1551474251066,"id":"587d781c367417b2b2512abf","files":[]},{"completedDate":1551474148910,"id":"587d781b367417b2b2512abe","files":[]},{"completedDate":1551474060295,"id":"587d781b367417b2b2512abd","files":[]},{"completedDate":1551474025331,"id":"587d781b367417b2b2512abc","files":[]},{"completedDate":1551473937546,"id":"587d781b367417b2b2512abb","files":[]},{"completedDate":1551473899639,"id":"587d781b367417b2b2512aba","files":[]},{"completedDate":1551473847375,"id":"587d781a367417b2b2512ab9","files":[]},{"completedDate":1551473816324,"id":"587d781a367417b2b2512ab8","files":[]},{"completedDate":1551473778133,"id":"587d781a367417b2b2512ab7","files":[]},{"completedDate":1551473609625,"id":"587d7791367417b2b2512ab5","files":[]},{"completedDate":1551473589995,"id":"587d7791367417b2b2512ab4","files":[]},{"completedDate":1551472955393,"id":"587d7791367417b2b2512ab3","files":[]},{"completedDate":1551471907288,"id":"5a9d72ad424fe3d0e10cad16","files":[]},{"completedDate":1551471570028,"id":"5a9d72a1424fe3d0e10cad15","files":[]},{"completedDate":1551471533683,"id":"5a9d7295424fe3d0e10cad14","files":[]},{"completedDate":1551471394981,"id":"5b7d72c338cd7e35b63f3e14","files":[]},{"completedDate":1551471311675,"id":"5a9d7286424fe3d0e10cad13","files":[]},{"completedDate":1551471212176,"id":"5a9d727a424fe3d0e10cad12","files":[]},{"completedDate":1551471002139,"id":"5a9d726c424fe3d0e10cad11","files":[]},{"completedDate":1551463956983,"id":"5a9d725e424fe3d0e10cad10","files":[]},{"completedDate":1551463916124,"id":"bad82fee1348bd9aedf08721","files":[]},{"completedDate":1551463885215,"id":"bad87fee1348bd9aede08718","files":[]},{"completedDate":1551463841995,"id":"bad87fee1348bd9aedf08719","files":[]},{"completedDate":1551463806229,"id":"bad87fee1348bd9aedf08721","files":[]},{"completedDate":1551463733650,"id":"bad87fee1348bd9aedf08726","files":[]},{"completedDate":1551463702836,"id":"bad87fee1348bd9aedf07756","files":[]},{"completedDate":1551463665151,"id":"bad87fee1348bd9aedf06756","files":[]},{"completedDate":1551463630333,"id":"bad87fee1348bd8aedf06756","files":[]},{"completedDate":1551463576673,"id":"bad87fee1348bd9aedf04756","files":[]},{"completedDate":1551463534923,"id":"bad87fee1348bd9aedf08756","files":[]},{"completedDate":1551463463552,"id":"bad87fee1348bd9aedf08746","files":[]},{"completedDate":1551463405546,"id":"bad87fee1348bd9aedf08736","files":[]},{"completedDate":1551463384400,"id":"bad82fee1322bd9aedf08721","files":[]},{"completedDate":1551463300598,"id":"58c383d33e2e3259241f3076","files":[]},{"completedDate":1551462783037,"id":"bad87fee1348bd9afdf08726","files":[]},{"completedDate":1551462756702,"id":"bad87fee1348bd9aedf08826","files":[]},{"completedDate":1551462705084,"id":"bad87fee1248bd9aedf08824","files":[]},{"completedDate":1551462680495,"id":"bad87fee1348bd9aedf08824","files":[]},{"completedDate":1551462644781,"id":"bad87fee1348bd9aedf08823","files":[]},{"completedDate":1551462625455,"id":"bad87fee1348bd9aedf08822","files":[]},{"completedDate":1551462601269,"id":"bad88fee1348bd9aedf08825","files":[]},{"completedDate":1551462542272,"id":"bad87dee1348bd9aede07836","files":[]},{"completedDate":1551462478170,"id":"bad87eee1348bd9aede07836","files":[]},{"completedDate":1551462047480,"id":"bad87fed1348bd9aede07836","files":[]},{"completedDate":1551461989328,"id":"bad87fee1348bd9aedf08815","files":[]},{"completedDate":1551461975252,"id":"bad87fee1348bd9aedf08814","files":[]},{"completedDate":1551461948003,"id":"bad87fee1348bd9bedf08813","files":[]},{"completedDate":1551461873517,"id":"bad87fee1348bd9acdf08812","files":[]},{"completedDate":1551461703593,"id":"bad87fee1348bd9aedf08808","files":[]},{"completedDate":1551455872984,"id":"bad87fee1348bd9aedf08807","files":[]},{"completedDate":1551455730993,"id":"bad87fee1348bd9aede08807","files":[]},{"completedDate":1551455699941,"id":"bad87fee1348bd9aedf08806","files":[]},{"completedDate":1551455669178,"id":"bad87fee1348bd9aefe08806","files":[]},{"completedDate":1551455630535,"id":"bad87fee1348bd9aecf08806","files":[]},{"completedDate":1551454996031,"id":"bad87fee1348bd9aedf08805","files":[]},{"completedDate":1551454909702,"id":"bad87fee1348bd9aedf08803","files":[]},{"completedDate":1551454501817,"id":"587d78aa367417b2b2512aec","files":[]},{"completedDate":1551454274579,"id":"587d78aa367417b2b2512aed","files":[]},{"completedDate":1551454173382,"id":"bad87fee1348bd9aede08835","files":[]},{"completedDate":1551454109735,"id":"bad87fee1348bd9aedd08835","files":[]},{"completedDate":1551453910674,"id":"bad87fee1348bd9aedf08835","files":[]},{"completedDate":1551453546653,"id":"bad87fee1348bd9aedf08834","files":[]},{"completedDate":1551453191054,"id":"bad87fee1348bd9aedc08830","files":[]},{"completedDate":1551453139503,"id":"bad87fee1348bd9aedd08830","files":[]},{"completedDate":1551453099476,"id":"bad87fee1348bd9aede08830","files":[]},{"completedDate":1551453010833,"id":"bad87fee1348bd9aedf08830","files":[]},{"completedDate":1551452739134,"id":"bad87fee1348bd9aedf08829","files":[]},{"completedDate":1551451817545,"id":"bad87fee1348bd9aedf08828","files":[]},{"completedDate":1551451751753,"id":"bad87fee1348bd9aedf08827","files":[]},{"completedDate":1551451673397,"id":"bad87fee1348bd9aedf08820","files":[]},{"completedDate":1551451616781,"id":"bad87fee1348bd9aedf08817","files":[]},{"completedDate":1551451583402,"id":"bad87fee1348bd9aede08817","files":[]},{"completedDate":1551451449054,"id":"bad88fee1348bd9aedf08816","files":[]},{"completedDate":1551451197014,"id":"bad87fee1348bd9aedf08816","files":[]},{"completedDate":1551450675221,"id":"bad87fee1348bd9aedf08812","files":[]},{"completedDate":1551450570306,"id":"bad87fee1348bd9aecf08801","files":[]},{"completedDate":1551450488641,"id":"bad87fed1348bd9aedf08833","files":[]},{"completedDate":1551450473197,"id":"bad87fee1348bd9aedf08804","files":[]},{"completedDate":1551450443713,"id":"bad87fee1348bd9aedf08802","files":[]},{"completedDate":1551450422112,"id":"bad87fee1348bd9aedf08833","files":[]},{"completedDate":1551450385519,"id":"bad87fee1348bd9aedf08801","files":[]},{"completedDate":1551450363588,"id":"bad87fee1348bd9aedf0887a","files":[]},{"completedDate":1551450313357,"id":"bd7123c8c441eddfaeb5bdef","files":[]}],"githubProfile":"https://github.com/Eggyplant","isApisMicroservicesCert":false,"isBackEndCert":false,"isCheater":false,"isDonating":null,"is2018DataVisCert":false,"isDataVisCert":false,"isFrontEndCert":false,"isFullStackCert":false,"isFrontEndLibsCert":false,"isHonest":true,"isInfosecQaCert":false,"isJsAlgoDataStructCert":false,"isRespWebDesignCert":false,"linkedin":"","location":"Saint Petersburg","name":"","portfolio":[],"profileUI":{"isLocked":false,"showAbout":true,"showCerts":true,"showHeatMap":true,"showLocation":true,"showName":false,"showPoints":true,"showPortfolio":true,"showTimeLine":true},"twitter":"","username":"victorious-bean","website":"","yearsTopContributor":[],"isGithub":true,"isLinkedIn":false,"isTwitter":false,"isWebsite":false,"points":100,"calendar":{"1551329887":1,"1551450313":1,"1551450363":1,"1551450385":1,"1551450422":1,"1551450443":1,"1551450473":1,"1551450488":1,"1551450570":1,"1551450675":1,"1551451197":1,"1551451449":1,"1551451583":1,"1551451616":1,"1551451673":1,"1551451751":1,"1551451817":1,"1551452739":1,"1551453010":1,"1551453099":1,"1551453139":1,"1551453191":1,"1551453546":1,"1551453910":1,"1551454109":1,"1551454173":1,"1551454274":1,"1551454501":1,"1551454909":1,"1551454996":1,"1551455630":1,"1551455669":1,"1551455699":1,"1551455730":1,"1551455872":1,"1551461703":1,"1551461873":1,"1551461948":1,"1551461975":1,"1551461989":1,"1551462047":1,"1551462478":1,"1551462542":1,"1551462601":1,"1551462625":1,"1551462644":1,"1551462680":1,"1551462705":1,"1551462756":1,"1551462783":1,"1551463300":1,"1551463384":1,"1551463405":1,"1551463463":1,"1551463534":1,"1551463576":1,"1551463630":1,"1551463665":1,"1551463702":1,"1551463733":1,"1551463806":1,"1551463841":1,"1551463885":1,"1551463916":1,"1551463956":1,"1551471002":1,"1551471212":1,"1551471311":1,"1551471394":1,"1551471533":1,"1551471570":1,"1551471907":1,"1551472955":1,"1551473589":1,"1551473609":1,"1551473778":1,"1551473816":1,"1551473847":1,"1551473899":1,"1551473937":1,"1551474025":1,"1551474060":1,"1551474148":1,"1551474251":1,"1551474336":1,"1551474420":1,"1551474463":1,"1551474482":1,"1551474502":1,"1551485542":1,"1551485709":1,"1551485762":1,"1551485885":1,"1551486753":1,"1551486839":1,"1551487141":1,"1551487180":1,"1551487274":1,"1551487391":1,"1551487504":1},"streak":{"longest":1,"current":1},"picture":"https://i.imgur.com/cSiiIM6.jpg"}}},"result":"victorious-bean"}`
	//var result map[string]interface{}
	//json.Unmarshal([]byte(freeCodeCampJson), &result)

	entities := fccJSON["entities"].(map[string]interface{})
	user := entities["user"].(map[string]interface{})
	vb := user[username].(map[string]interface{})
	points := int(vb["points"].(float64))

	userInfo.points = points
	userInfo.username = username
	return nil
}

// UserInfo contains all the info from JSON request that we want to use
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
