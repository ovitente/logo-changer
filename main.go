package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/reiver/go-telnet"
)

type LogoName string
type LogoURL string
type LogosList map[LogoName]LogoURL

type Name string
type URL string
type Config map[Name]URL

// AppConfig : Init structure for config
type AppConfig struct {
	UpdTimeout  int
	FtpConfig   FtpConfigFields
	ServerQuery ServerQueryFields
}

// FtpConfigFields : Struct for holding ftp credentials
type FtpConfigFields struct {
	Host     string
	Login    string
	Password string
}

// ServerQueryFields : Struct for holding ServerQuery credentials
type ServerQueryFields struct {
	Host     string
	Login    string
	Password string
	Port     int
	ServerID int
}

// PrintConfig : Self-explanatory
func PrintConfig(Credentials *AppConfig) {
	fmt.Printf("%+v\n", Credentials)
}

// ReadConfig : Self-explanatory
func ReadConfig() (*AppConfig, error) {
	var Credentials = new(AppConfig)
	var err error

	Credentials.UpdTimeout, err = strconv.Atoi(os.Getenv("LOGO_UPDATE_TIMEOUT"))
	if err != nil {
		log.Fatal("Cant assign LOGO_UPDATE_TIMEOUT")
	}

	Credentials.FtpConfig.Host = os.Getenv("FTP_HOST")
	Credentials.FtpConfig.Login = os.Getenv("FTP_LOGIN")
	Credentials.FtpConfig.Password = os.Getenv("FTP_PASSWORD")

	Credentials.ServerQuery.Host = os.Getenv("SQUERY_HOST")
	Credentials.ServerQuery.Login = os.Getenv("SQUERY_LOGIN")
	Credentials.ServerQuery.Password = os.Getenv("SQUERY_PASSWORD")
	Credentials.ServerQuery.Port, err = strconv.Atoi(os.Getenv("SQUERY_PORT"))
	if err != nil {
		log.Fatal("Cant Assign SQUERY_PORT")
	}
	Credentials.ServerQuery.ServerID, err = strconv.Atoi(os.Getenv("SQUERY_SERVER_ID"))
	if err != nil {
		log.Fatal("Cant Assign SQUERY_ID")
	}

	return Credentials, nil
}

func main() {

	Credentials, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found. Skipping.")
	}

	// for i := 0; i < 11; i++ {
	for {

		// fmt.Printf("%v| WELL WELL WELL \n", i)
		LogoRandomizer()
		time.Sleep(time.Duration(Credentials.UpdTimeout) * time.Second)
	}

	//// DEBUG
	// var configFilePath = "./config.yml"
	// Credentials, err := ReadConfig(configFilePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("\n------------DEBUG------------\n")
	// PrintConfig(Credentials)

	// fmt.Printf(" ServerQueryHost: %s:%v\n SID: %v\n ServerQueryLogin: %s\n ServerQueryPassworf: %s\n",
	// 	Credentials.ServerQuery.Host,
	// 	Credentials.ServerQuery.Port,
	// 	Credentials.ServerQuery.ServerID,
	// 	Credentials.ServerQuery.Login,
	// 	Credentials.ServerQuery.Password)

	// fmt.Println("-----------------------------")
	// fmt.Printf(" FTP Host: %s\n FTP Login: %s\n FTP Password: %s\n",
	// 	Credentials.FtpConfig.Host,
	// 	Credentials.FtpConfig.Login,
	// 	Credentials.FtpConfig.Password)
	// fmt.Printf("\n------------DEBUG------------\n")

}

func random(min, max int) int {
	// rand.Seed(time.Now().Unix())
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

// LogoRandomizer self explanatory
func LogoRandomizer() string {

	var finalChosenLogo string

	switch v := random(0, 8); v {

	case 0:
		finalChosenLogo = "videos"
		logoType, logoDirName, logoID, logoURL := ChooseVideo()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

	case 1:
		finalChosenLogo = "videos"
		logoType, logoDirName, logoID, logoURL := ChooseVideo()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

	case 2:
		finalChosenLogo = "info"
		logoType, logoDirName, logoID, logoURL := ChooseInfo()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

	case 3:
		finalChosenLogo = "books"
		logoType, logoDirName, logoID, logoURL := ChooseBook()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

	case 4:
		finalChosenLogo = "practics"
		logoType, logoDirName, logoID, logoURL := ChoosePractic()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

	case 5:
		finalChosenLogo = "payment"
		logoType, logoDirName, logoID, logoURL := ChoosePayment()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

	case 6:
		finalChosenLogo = "payment"
		logoType, logoDirName, logoID, logoURL := ChoosePayment()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

	case 7:
		finalChosenLogo = "payment"
		logoType, logoDirName, logoID, logoURL := ChoosePayment()
		ChangeLogo(logoType, logoDirName, logoID, logoURL)

		// case 6:
		// 	finalChosenLogo = "links"
		// 	logoType, logoDirName, logoID, logoURL := ChooseLink()
		// 	ChangeLogo(logoType, logoDirName, logoID, logoURL)

	}

	return finalChosenLogo
}

// func Priority() {
// 	package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type Item struct {
// 	Value  string
// 	Weight int
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	var items = []Item{
// 		{"foo", 1},
// 		{"bar", 1},
// 		{"three", 10},
// 	}
// 	fmt.Println(getRandom(items))
// }

// func getRandom(items []Item) Item {
// 	set := []int{}
// 	for k, v := range items {
// 		for i := 0; i < v.Weight; i++ {
// 			set = append(set, k)
// 		}
// 	}
// 	fmt.Println(set)
// 	rindex := rand.Intn(len(set))
// 	fmt.Println(rindex)
// 	index := set[rindex]
// 	return items[index]
// }

// }

//ChangeLogo Changes logo on the server
func ChangeLogo(logoType string, logoDirName string, logoID string, logoURL string) {

	Credentials, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("\n------------DEBUG------------\n")
	// PrintConfig(Credentials)

	log.Printf("Changing logo.")

	conn, _ := telnet.DialTo(fmt.Sprintf("%s:%v",
		Credentials.ServerQuery.Host,
		Credentials.ServerQuery.Port))

	conn.Write([]byte(fmt.Sprintf("login %s %s\r\n",
		Credentials.ServerQuery.Login,
		Credentials.ServerQuery.Password)))

	conn.Write([]byte(fmt.Sprintf("use sid=%v\r\n",
		Credentials.ServerQuery.ServerID)))

	conn.Write([]byte(fmt.Sprintf("serveredit sid=%v virtualserver_hostbanner_gfx_url=http://%s/%s/%v.png virtualserver_hostbanner_url=%s\r\n",
		Credentials.ServerQuery.ServerID,
		Credentials.FtpConfig.Host,
		logoDirName,
		logoID,
		logoURL)))

	conn.Write([]byte("quit\r\n"))

	// DEBUG
	fmt.Printf("\nChosen Logo\n LogoType:    %s\n LogoDirName: %s\n LogoID:      %s\n LogoURL:     %s\n\n",
		logoType,
		logoDirName,
		logoID,
		logoURL)

	log.Printf("Logo has changed.") // DEBUG
}
