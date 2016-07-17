package cdnjslib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func WriteError(writer io.Writer, err error) {
	m := make(map[string]string)
	m["error"] = err.Error()
	writer.Write([]byte(ToJson(m)))
}

func ToJson(arg interface{}) string {
	b, e := json.MarshalIndent(arg, "", "  ")
	if e != nil {
		panic(e)
	}
	return string(b)
}

func ReadFile(filename string, v interface{}) error {
	cf, err := os.Open(filename)
	if err != nil {
		return err
	}
	return json.NewDecoder(cf).Decode(v)
}

func (library *LibraryInfo) LoadCdnjsData() error {

	jsonUrl := fmt.Sprintf("https://api.cdnjs.com/libraries/%s", library.Name)

	resp, err := http.Get(jsonUrl)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	library.CdnJsLibrary = &CdnJsLibrary{}

	err = json.NewDecoder(bufio.NewReader(resp.Body)).Decode(library.CdnJsLibrary)

	if err != nil {
		return err
	}
	return nil
}

func (library *LibraryInfo) getAsset() (Asset, error) {
	for _, asset := range library.CdnJsLibrary.Assets {
		if asset.Version == library.Version {
			return asset, nil
		}
	}
	return Asset{}, fmt.Errorf("unable to locate asset version %v in library %v", library.Version, library.Name)
}

func (library *LibraryInfo) shouldFilter(path string) bool {
	for _, filter := range library.AssetFilters {
		if strings.Contains(path, filter) {
			return true
		}
	}
	return false
}

func (library *LibraryInfo) LoadAssetUrls() error {
	libString := "https://cdnjs.cloudflare.com/ajax/libs/%v/%v/%v"

	asset, err := library.getAsset()
	if err != nil {
		return err
	}

	for _, f := range asset.Files {
		if !library.shouldFilter(f) {

			ls := fmt.Sprintf(libString, library.Name, library.Version, f)
			lu, err := url.Parse(ls)
			if err != nil {
				return err
			}

			library.AssetUrls = append(library.AssetUrls, AssetUrl{
				Path: f,
				Url:  lu,
			})
		}
	}
	return nil
}
