
### dirs

A minimal library that offers platform-specific paths for configuration files, cache, and other data directories.

### Installation
```bash
go get github.com/ansrivas/dirs@latest
```


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

### License
MIT License

### Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.