// Code generated by go-bindata.
// sources:
// web/contacts.html
// web/css/mdc.css
// web/index.html
// web/js/app.js
// web/js/contacts.js
// web/js/servers.js
// web/js/settings.js
// web/momentum/test.tmpl
// web/momentum/your-404-page.tmpl
// web/momentum/your-500-page.tmpl
// web/servers.html
// web/settings.html
// web/your-404-page.tmpl
// web/your-500-page.tmpl
// tmpl/momentum/ang.tmpl
// tmpl/momentum/jquery.tmpl
// tmpl/momentum/server.tmpl
// DO NOT EDIT!

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// webContactsHtml reads file data from disk. It returns an error on failure.
func webContactsHtml() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/contacts.html"
	name := "web/contacts.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webCssMdcCss reads file data from disk. It returns an error on failure.
func webCssMdcCss() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/css/mdc.css"
	name := "web/css/mdc.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webIndexHtml reads file data from disk. It returns an error on failure.
func webIndexHtml() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/index.html"
	name := "web/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webJsAppJs reads file data from disk. It returns an error on failure.
func webJsAppJs() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/js/app.js"
	name := "web/js/app.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webJsContactsJs reads file data from disk. It returns an error on failure.
func webJsContactsJs() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/js/contacts.js"
	name := "web/js/contacts.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webJsServersJs reads file data from disk. It returns an error on failure.
func webJsServersJs() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/js/servers.js"
	name := "web/js/servers.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webJsSettingsJs reads file data from disk. It returns an error on failure.
func webJsSettingsJs() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/js/settings.js"
	name := "web/js/settings.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webMomentumTestTmpl reads file data from disk. It returns an error on failure.
func webMomentumTestTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/momentum/test.tmpl"
	name := "web/momentum/test.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webMomentumYour404PageTmpl reads file data from disk. It returns an error on failure.
func webMomentumYour404PageTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/momentum/your-404-page.tmpl"
	name := "web/momentum/your-404-page.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webMomentumYour500PageTmpl reads file data from disk. It returns an error on failure.
func webMomentumYour500PageTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/momentum/your-500-page.tmpl"
	name := "web/momentum/your-500-page.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webServersHtml reads file data from disk. It returns an error on failure.
func webServersHtml() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/servers.html"
	name := "web/servers.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webSettingsHtml reads file data from disk. It returns an error on failure.
func webSettingsHtml() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/settings.html"
	name := "web/settings.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webYour404PageTmpl reads file data from disk. It returns an error on failure.
func webYour404PageTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/your-404-page.tmpl"
	name := "web/your-404-page.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// webYour500PageTmpl reads file data from disk. It returns an error on failure.
func webYour500PageTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/web/your-500-page.tmpl"
	name := "web/your-500-page.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tmplMomentumAngTmpl reads file data from disk. It returns an error on failure.
func tmplMomentumAngTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/tmpl/momentum/ang.tmpl"
	name := "tmpl/momentum/ang.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tmplMomentumJqueryTmpl reads file data from disk. It returns an error on failure.
func tmplMomentumJqueryTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/tmpl/momentum/jquery.tmpl"
	name := "tmpl/momentum/jquery.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tmplMomentumServerTmpl reads file data from disk. It returns an error on failure.
func tmplMomentumServerTmpl() (*asset, error) {
	path := "/Users/Adrian/gosapphire/src/github.com/cheikhshift/megalith/tmpl/momentum/server.tmpl"
	name := "tmpl/momentum/server.tmpl"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"web/contacts.html":               webContactsHtml,
	"web/css/mdc.css":                 webCssMdcCss,
	"web/index.html":                  webIndexHtml,
	"web/js/app.js":                   webJsAppJs,
	"web/js/contacts.js":              webJsContactsJs,
	"web/js/servers.js":               webJsServersJs,
	"web/js/settings.js":              webJsSettingsJs,
	"web/momentum/test.tmpl":          webMomentumTestTmpl,
	"web/momentum/your-404-page.tmpl": webMomentumYour404PageTmpl,
	"web/momentum/your-500-page.tmpl": webMomentumYour500PageTmpl,
	"web/servers.html":                webServersHtml,
	"web/settings.html":               webSettingsHtml,
	"web/your-404-page.tmpl":          webYour404PageTmpl,
	"web/your-500-page.tmpl":          webYour500PageTmpl,
	"tmpl/momentum/ang.tmpl":          tmplMomentumAngTmpl,
	"tmpl/momentum/jquery.tmpl":       tmplMomentumJqueryTmpl,
	"tmpl/momentum/server.tmpl":       tmplMomentumServerTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"tmpl": &bintree{nil, map[string]*bintree{
		"momentum": &bintree{nil, map[string]*bintree{
			"ang.tmpl":    &bintree{tmplMomentumAngTmpl, map[string]*bintree{}},
			"jquery.tmpl": &bintree{tmplMomentumJqueryTmpl, map[string]*bintree{}},
			"server.tmpl": &bintree{tmplMomentumServerTmpl, map[string]*bintree{}},
		}},
	}},
	"web": &bintree{nil, map[string]*bintree{
		"contacts.html": &bintree{webContactsHtml, map[string]*bintree{}},
		"css": &bintree{nil, map[string]*bintree{
			"mdc.css": &bintree{webCssMdcCss, map[string]*bintree{}},
		}},
		"index.html": &bintree{webIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"app.js":      &bintree{webJsAppJs, map[string]*bintree{}},
			"contacts.js": &bintree{webJsContactsJs, map[string]*bintree{}},
			"servers.js":  &bintree{webJsServersJs, map[string]*bintree{}},
			"settings.js": &bintree{webJsSettingsJs, map[string]*bintree{}},
		}},
		"momentum": &bintree{nil, map[string]*bintree{
			"test.tmpl":          &bintree{webMomentumTestTmpl, map[string]*bintree{}},
			"your-404-page.tmpl": &bintree{webMomentumYour404PageTmpl, map[string]*bintree{}},
			"your-500-page.tmpl": &bintree{webMomentumYour500PageTmpl, map[string]*bintree{}},
		}},
		"servers.html":       &bintree{webServersHtml, map[string]*bintree{}},
		"settings.html":      &bintree{webSettingsHtml, map[string]*bintree{}},
		"your-404-page.tmpl": &bintree{webYour404PageTmpl, map[string]*bintree{}},
		"your-500-page.tmpl": &bintree{webYour500PageTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
