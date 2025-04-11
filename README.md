
WIP:


```golang
package main

import (
	"fmt"
	"log"
	"golang-dirs/dirs"
)

func main() {
	d := dirs.NewDirs()

	home, err := d.HomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Home Directory:", home)

	cache, _ := d.CacheDir()
	fmt.Println("Cache Directory:", cache)

	config, _ := d.ConfigDir()
	fmt.Println("Config Directory:", config)

```