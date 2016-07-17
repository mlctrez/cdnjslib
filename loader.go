package cdnjslib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func (library *LibraryInfo) SaveLocal(dirname string) {

	checkError(library.LoadCdnjsData())
	checkError(library.LoadAssetUrls())

	wg := sync.WaitGroup{}

	msgChan := make(chan string)
	defer close(msgChan)

	go func(c chan string) {
		for m := range c {
			fmt.Println(m)
		}
	}(msgChan)

	for _, assetUrl := range library.AssetUrls {

		wg.Add(1)

		go func(assetUrl AssetUrl) {
			defer wg.Done()

			var err error

			localPath := filepath.Join(dirname, library.Name, library.Version, assetUrl.Path)

			err = os.MkdirAll(path.Dir(localPath), 0755)
			checkError(err)

			remotejs, err := http.Get(assetUrl.Url.String())
			checkError(err)
			defer remotejs.Body.Close()

			content, err := ioutil.ReadAll(remotejs.Body)
			checkError(err)

			err = ioutil.WriteFile(localPath, content, 0744)
			checkError(err)

			msgChan <- fmt.Sprint("retrieved ", assetUrl.Url.String())

		}(assetUrl)
	}
	wg.Wait()

}
