package main

import (
	"crypto/sha256"
	"bufio"
	"strings"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)


func main() {
	major:= 0
	minor:= 1
	build:= "07102017"

	var version = fmt.Sprintf("%d.%d.%s", major, minor, build)
	fmt.Printf("gomake version %s \n", version)

	//check for existence of Gomakefile.yml
	if _, err:= os.Stat("Gomakefile.yml");  ! os.IsNotExist(err) {
		fmt.Println("found makefile...                   OK")
	}else {
		fmt.Println("makefile NOT FOUND...            ERROR")
		return; 
	}

	loadMakefile()

	//if on windows, we want .exe extension
	if runtime.GOOS == "windows" {
		//BINOUTPUT = BINOUTPUT + ".exe"
	}

	//buildProgram()
}

func loadMakefile() {
	fmt.Println("reading Makefile..")

	inFile, _:= os.Open("Gomakefile.yml")
	defer inFile.Close()
	scanner:= bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
		
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#") != true {
			fmt.Println(scanner.Text())
		}
	
		
	}
	  
}

func buildProgram() {
	//go build [-o output] [-i] [build flags] [packages]
	cmd:= exec.Command("go", "build", "-o", "gomake.bin", "gomake.go")
	log.Printf("Running command and waiting for it to finish...")
	err:= cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

func clean() {

}

func test() {

}

//checksumFile will create sha256 checksum of whole file
func checksumFile() {
	f, err:= os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h:= sha256.New()
	if _, err:= io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}

//checksumString will create sha256 checksum of given string
func checksumString(mystr string) {
	sum:= sha256.Sum256([]byte(mystr))
	fmt.Printf("%x", sum)
}
