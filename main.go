package main

import (
 "io/ioutil"
 "fmt"
 "strings"
 "os"
)

func main() {
	files, _ := ioutil.ReadDir("z:\\Videos")
	targets := make([]string, len(files))
	for index, f := range files {
    	if(f.IsDir()){
    		targets[index] = f.Name()
    		index++
        }
    }
	files, _ = ioutil.ReadDir("z:\\Videos\\Seeding")
	outer: for _, f := range files {
		if(f.IsDir() || f.Name() == "Seeding"){
			continue;
		}
		for _, targ := range targets {
			if(strings.Contains(strings.ToLower(strings.Replace(f.Name(),"."," ", -1)), strings.ToLower(targ))){
				fmt.Println(f.Name() + " => " + targ)
				os.Rename("z:\\Videos\\Seeding\\" + f.Name(), "z:\\Videos\\" + targ + "\\" + f.Name())
				continue outer;
			}
		}
			fmt.Println("No home: " + f.Name());
	}
}