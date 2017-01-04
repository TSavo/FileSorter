package main

import (
 "io/ioutil"
 "fmt"
 "strings"
	"regexp"
 "os"
)


func main() {
	regularTest := regexp.MustCompile(`(.*)( S\d*E\d*)(.*)`)
	files, _ := ioutil.ReadDir("G:\\Videos")
	targets := make([]string, len(files))
	for index, f := range files {
    		if(f.IsDir()){
    			targets[index] = f.Name()
    			index++
        	}
    	}
	files, _ = ioutil.ReadDir("g:\\TV Shows")
	outer: for _, f := range files {
		if(f.IsDir() || f.Name() == "Seeding"){
			continue;
		}
		if(regularTest.MatchString(strings.Replace(f.Name(),"."," ", -1))){
			dirName := regularTest.FindAllStringSubmatch(strings.Replace(f.Name(),"."," ", -1), -1)[0][1]
			fmt.Println("G:\\Videos\\" + dirName)
			os.MkdirAll("G:\\Videos\\"+dirName, os.FileMode(0522))
		}
		for _, targ := range targets {
			if(strings.Contains(strings.ToLower(strings.Replace(f.Name(),"."," ", -1)), strings.ToLower(targ))){
				fmt.Println(f.Name() + " => " + targ)
				os.Remove("G:\\Videos\\" + targ + "\\" + f.Name())
				os.Rename("G:\\TV Shows\\" + f.Name(), "G:\\Videos\\" + targ + "\\" + f.Name())
				continue outer;
			}
		}
			fmt.Println("No home: " + f.Name());
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return false, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}