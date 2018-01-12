package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var makefile = "makefile2go.json"

//About describes software, it's author and so on
type About struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Repository  string `json:"repository"`
	Author      string `json:"author"`
	Email       string `json:"email"`
	License     string `json:"license"`
	LicenseURL  string `json:"licenseurl"`
}

//Version describes version and versioning configuration
type Version struct {
	Major           string `json:"major"`
	Minor           string `json:"minor"`
	Build           string `json:"build"`
	Commit          string `json:"commit"`
	IncreaseVersion string `json:"increaseversion"` //increaseversion increment maj,min or build version
}

//Build describes build configuration
type Build struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Compilerpath string `json:"compilerpath"`
	Binary       string `json:"binary"`
	Flags        string `json:"flags"`
}

//Dependency hold list of dependencies to download thru go get
type Dependency struct {
	Name   string `json:"name"`
	DepURL string `json:"depurl"`
}

//Makefile struct that holds whole makefile data
type Makefile struct {
	About
	Version
	Builds       []Build
	Dependencies []Dependency
}

func getRunningDir() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
}

func printTimezones() {
	t := time.Now()

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("US/New York  : ", t.In(location))

	locberlin, _ := time.LoadLocation("Europe/Berlin")
	nowberlin := time.Now().In(locberlin)
	fmt.Println("Europe/Berlin: ", nowberlin)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	fmt.Println("Asia/Shanghai: ", now)
}

func main() {

	// mc := MicroLogger

	// mc.Warn("sldjflsjdf")
	// mc.Deprecated("ldskjf")

	current := time.Now()
	getRunningDir()

	var mk Makefile
	mk.About.Name = "mynewproject"
	mk.Version.Minor = "1"
	mk.Version.Major = "0"
	mk.Version.Build = current.Format("20060102150405")
	mk.Version.Commit = "0"

	fmt.Println("--------------------------------")
	fmt.Println("- GOMAKE                       -")
	fmt.Println("- builder and generator        -")
	var version = fmt.Sprintf("%s.%s.%s   -", mk.Version.Major, mk.Version.Minor, mk.Version.Build)
	fmt.Printf("- version %s \n", version)
	fmt.Println("--------------------------------")
	fmt.Println("args:")

	processArgs()
	fmt.Println("-----------------------")
	printTimezones()
	fmt.Println("-----------------------")
	checkMakefileExists()

	//fmt.Printf("%+v", mk)

	data, err := json.Marshal(mk)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	//buildProgram()

}

func processArgs() {
	compilerFlag := flag.String("compilerversion", "default", "choose compiler version")
	configurationFlag := flag.String("configuration", "release_patch", "choose configuration")
	//numbPtr := flag.Int("configuration", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()

	fmt.Println("compiler:", *compilerFlag)
	fmt.Println("configuration_flag:", *configurationFlag)
}

//checkMakefileExists check for existence of Gomakefile.yml
func checkMakefileExists() {
	if _, err := os.Stat(makefile); !os.IsNotExist(err) {
		fmt.Println("found makefile...                   OK")
		loadMakefile()
	} else {
		fmt.Println("makefile NOT FOUND...            ERROR")
		initProj()
	}
}

//TODO: design and deploy
func initProj() {
	fmt.Println("starting project GENERAtor...")

	//reading a string
	reader := bufio.NewReader(os.Stdin)
	var name string

	fmt.Println("--------------------------------")
	fmt.Println("Project name:")
	name, _ = reader.ReadString('\n')

	fmt.Print("creating project ", name)
	//create dir
	//cd dir
	//generate makefile2go.json
	fmt.Println("--------------------------------")

}

func loadMakefile() {
	fmt.Println("reading Makefile..")
	//fmt.Printf("working dir is %s", os.Getwd())
	inFile, _ := os.Open(makefile)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ":") == true {

			s := strings.Split(scanner.Text(), ":")
			if strings.Contains(s[1], "{") == true {
				//fmt.Println("GROUP -> ",s[0])
			} else {
				//param, value := s[0], s[1]
				//fmt.Println(param, "...",value)
			}

		}

	}

}

func buildProgram() {
	//go build [-o output] [-i] [build flags] [packages]
	cmd := exec.Command("go", "build", "gomake.go")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

// clean binaries
func clean() {

}

//test is to run all _test files
func test() {

}

// raiseBuildNum
func raiseBuildNum() {

}

//checksumFile will create sha256 checksum of whole file
func checksumFile() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}

//checksumString will create sha256 checksum of given string
func checksumString(mystr string) {
	sum := sha256.Sum256([]byte(mystr))
	fmt.Printf("%x", sum)
}
