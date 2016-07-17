package cdnjslib

import "net/url"

// Library is the root element of the cdnjs json api.
type CdnJsLibrary struct {
	Keywords    interface{} `json:"keywords"`
	Homepage    string      `json:"homepage"`
	Author      interface{} `json:"author"`
	AutoUpdate  interface{} `json:"autoupdate"`
	Name        string      `json:"name"`
	Filename    string      `json:"filename"`
	Description string      `json:"description"`
	License     string      `json:"license"`
	Assets      []Asset     `json:"assets"`
	Version     string      `json:"version"`
	Repository  interface{} `json:"repository"`
	Namespace   string      `json:"namespace"`
}

// Assets are the version of the library and the files in that version.
type Asset struct {
	Version string   `json:"version"`
	Files   []string `json:"files"`
}

// These next three are not currently used in the  Library struct.
// They are instead replaced with interface{} since the json documents are not
// very consistent on what types are returned for what values.

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Url   string `json:"url"`
}

type Autoupdate struct {
	Type   string `json:"type"`
	Target string `json:"target"`
}

type Repository struct {
	Type string `json:"string"`
	Url  string `json:"url"`
}

// LibraryInfo is the information about a library being requested.
type LibraryInfo struct {
	Name         string
	Version      string
	AssetFilters []string
	CdnJsLibrary *CdnJsLibrary
	AssetUrls    []AssetUrl
}

type LibraryCollection struct {
	Libraries []LibraryInfo `json:"libraries"`
}

type AssetUrl struct {
	Path string
	Url  *url.URL
}
