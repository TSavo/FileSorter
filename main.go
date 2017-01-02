package main

import (
 "io/ioutil"
 "fmt"
 "strings"
 "os"
)

func main() {
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
		for _, targ := range targets {
			if(strings.Contains(strings.ToLower(strings.Replace(f.Name(),"."," ", -1)), strings.ToLower(targ))){
				fmt.Println(f.Name() + " => " + targ)
				os.Rename("G:\\TV Shows\\" + f.Name(), "G:\\Videos\\" + targ + "\\" + f.Name())
				continue outer;
			}
		}
			fmt.Println("No home: " + f.Name());
	}
}