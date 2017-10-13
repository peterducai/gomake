package main

import (
	"bufio"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var makefile = "Gomakefile.json"

type Makefile struct {
	target_binary string
	license       string
	author        string
	url           string
	repo          string
	version       struct {
		major int
		minor int
		build int
	}
	compiler struct {
		name string
		path string
	}
	dependencies struct {
		url string
	}
	build_configuration struct {
		name        string
		description string
		flags       []string
	}
}

// In Go:

// // web server

// type Foo struct {
//     Number int    `json:"number"`
//     Title  string `json:"title"`
// }

// foo_marshalled, err := json.Marshal(Foo{Number: 1, Title: "test"})
// fmt.Fprint(w, string(foo_marshalled)) // write response to ResponseWriter (w)
// In JavaScript:

// // web call & receive in "data", thru Ajax/ other

// var Foo = JSON.parse(data);
// console.log("number: " + Foo.number);
// console.log("title: " + Foo.title);

func main() {

	major := 0
	minor := 1
	build := "07102017"

	processArgs()

	var version = fmt.Sprintf("%d.%d.%s", major, minor, build)
	fmt.Printf("gomake version %s \n", version)

	checkMakefileExists()

	//if on windows, we want .exe extension
	if runtime.GOOS == "windows" {
		//BINOUTPUT = BINOUTPUT + ".exe"
	}

	//buildProgram()

}

func processArgs() {
	compilerFlag := flag.String("compiler", "default", "choose compiler version")
	configurationFlag := flag.String("configuration", "release_patch", "choose configuration")
	//numbPtr := flag.Int("configuration", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")
	flag.Parse()

	fmt.Println("compiler:", *compilerFlag)
	fmt.Println("configuration_flag:", *configurationFlag)
}

func checkMakefileExists() {
	//check for existence of Gomakefile.yml
	if _, err := os.Stat(makefile); !os.IsNotExist(err) {
		fmt.Println("found makefile...                   OK")
		loadMakefile()
	} else {
		fmt.Println("makefile NOT FOUND...            ERROR")
		initProj()
	}
}

func initProj() {
	fmt.Println("starting project GENERAtor...")

	//reading a string
	reader := bufio.NewReader(os.Stdin)
	var name string
	fmt.Println("Project name:")
	name, _ = reader.ReadString('\n')

	fmt.Println("Your name is ", name)
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
	cmd := exec.Command("go", "build", "-o", "gomake.bin", "gomake.go")
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

func clean() {

}

func test() {

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
