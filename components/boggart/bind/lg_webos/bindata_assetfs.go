// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package lg_webos

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\x4f\x6f\xdb\x3e\x0c\xbd\xf7\x53\x10\x3c\xfd\x7e\x03\x1c\x23\xc9\x50\xec\x60\xe7\xbc\x43\x7b\xd9\xba\x73\x20\x5b\x74\xa3\x42\x96\x32\x91\x49\xb3\x1a\xf9\xee\x83\xfc\x27\x75\x13\x2f\xcb\x61\xb9\x84\xb2\x9e\xc4\x47\xbe\x27\x36\x0d\x68\xaa\x8c\x23\xc0\xd2\x3b\x21\x27\x08\xc7\xe3\x5d\xa6\xcd\x1e\x4a\xab\x98\x73\x0c\xfe\x15\x57\x77\x00\x00\xe3\xaf\xa5\xb7\x49\xad\x93\xf9\x02\x62\xc4\xf5\x10\x1d\x38\x99\x2f\x7a\xfc\xf9\x99\xc3\x7a\xab\x1c\xd9\xd1\xee\x25\x42\x8c\x58\x3a\x43\xb4\xa8\xcd\x62\xd5\x34\x60\xe6\x5f\x1c\xe0\x93\x57\x2c\x08\x33\x38\x1e\xb3\x74\xb3\x98\x40\x8f\x99\x5a\x52\xa1\x32\x07\x5c\x65\xa9\x36\xfb\xb3\xe4\x13\x9f\x3e\xf0\x19\xba\x32\x91\xa3\xf2\xa1\x1e\x80\x31\x4e\x36\x3e\x98\x37\xef\x44\x59\x68\xd7\x56\x15\x64\x13\x4b\x95\x20\x04\x6f\xa9\x83\x21\xd4\x24\x1b\xaf\x73\xdc\xfa\x58\x85\xd1\x39\x32\x39\x8d\xe0\xfc\x5e\x59\xa3\x95\xd0\x65\xb6\x73\x66\x46\xa8\xee\xb2\x3c\x07\xbf\xdb\x4e\xf0\x3b\x9d\x6a\x69\x44\x6c\x8e\x35\x31\xab\x67\xc2\x77\x15\x9d\x04\x6f\x3b\xa6\xd0\x6b\xba\x1c\x24\x5d\x4e\x2a\x3a\xf5\x3b\x29\xf3\x38\x24\x88\xda\x40\xc6\x5b\xe5\x4e\x3e\xa2\x9f\x3b\x13\x48\xe3\xea\x53\x96\xc6\x8d\x2b\x94\xd3\x96\xd0\x15\xc0\xa5\x13\xef\x07\xd6\xf7\x37\xb3\xce\x84\x0e\xa2\x02\xa9\x0f\x2a\xf6\x3d\x81\x40\x6c\xde\x54\x61\x69\x3d\xc0\xa2\x8a\xaf\x9c\xe3\x67\x04\xa7\x6a\xca\x51\x3a\x1f\x46\x05\xfb\x70\x28\x72\x5c\x6e\xd3\xc0\xac\xdd\x6e\xed\x3a\x5c\x76\xad\xfa\x0b\x4f\xde\xb2\x35\x6a\x89\x75\x6b\xf6\xd6\xe8\x49\xcb\x4f\x1d\xb8\xcd\x48\x7f\x6a\x7a\xad\x13\x5f\x55\x4c\x92\x2c\xff\xd6\xf1\x62\x27\xe2\xdd\xc8\xf3\xf2\x6b\x4b\x39\xf2\xae\xa8\x8d\x9c\x6c\x59\x88\x83\x42\x5c\xc2\xbb\xb2\x24\x66\x7c\x7f\xf9\xdf\xdb\x43\xdd\xc3\xef\xee\xfa\x97\x7d\xcc\xd2\xd8\x88\xab\x13\x62\xb4\xec\xc3\xfe\xaf\x69\x80\x9c\x8e\xc3\xf3\x6e\x34\x54\x5f\xb8\x9d\xa7\xd0\xbd\x11\x16\x25\xa6\xfc\xfa\xf4\xf8\x00\xff\x75\xf1\x8f\x6f\x0f\x80\xa9\x56\xbc\x29\xbc\x0a\x3a\x55\xcc\x24\x9c\xee\xc9\x69\x1f\x38\xed\xe7\x81\x0f\xe9\xcb\x68\x31\xab\x8d\x9b\xc5\x9b\x2b\x65\x99\xfe\x8f\x09\x4e\xd9\x7f\x07\x00\x00\xff\xff\xe4\xce\x09\xe9\xd1\x05\x00\x00"

func templatesViewsWidgetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsWidgetHtml,
		"templates/views/widget.html",
	)
}

func templatesViewsWidgetHtml() (*asset, error) {
	bytes, err := templatesViewsWidgetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/widget.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x4e\xbd\x6b\x14\x41\x14\xff\xe5\x8c\x0a\x67\xa3\xc1\x52\xf0\x89\x18\x14\x99\xb8\x7b\x67\x24\x4c\xb2\x89\x98\x0f\x10\x73\x18\xe2\xa9\xf5\x70\x37\xb9\x2c\xde\xcd\x1c\x3b\xbb\xc1\x60\x8a\x18\x53\x28\x88\x9d\xad\x20\x36\xda\x48\x0c\x04\x12\x42\x36\x85\xff\xc0\x5b\xb0\xb6\xb3\xb0\xf1\x6f\x90\xdd\xcb\x87\x42\x5e\xf3\x7e\x5f\xf3\x9b\xf7\x73\xa0\xff\x3d\x00\x9c\x01\x70\x09\xc0\x2c\x80\x73\x00\x56\xd0\x9b\xcf\x00\x2e\x00\xf8\x02\xe0\x2c\x80\x5d\x00\xe7\x01\x7c\x07\x70\x11\xc0\x0f\x00\xfd\x00\x7e\x03\x58\xec\x03\xfe\x00\xb8\x0a\xe0\x69\x09\x18\x00\xf0\xa2\x04\x5c\x06\xf0\xae\x04\xdc\x01\xf0\xf5\x40\xff\x55\x02\xfa\x70\x3c\xa5\x83\x9e\xc3\x39\x7d\xb0\xf3\xbb\x4e\xe5\x60\x4a\x2f\x85\x0d\x4d\xa1\x23\xbb\xb0\xd0\x0e\x8d\x46\x4d\x3b\xa7\x5a\x47\x3b\xb7\x74\xa7\x1b\x2f\x1f\x09\x4e\x9b\x26\xb9\xa4\xd1\xd0\xce\xe1\x91\x36\x4d\xcc\xeb\xae\x8d\x62\x51\x73\xad\xb0\x29\xee\x25\x2d\x27\xea\x56\x52\x53\x2f\xdd\x7d\x16\x2e\xaa\x8e\x1d\x8a\x92\xf2\xdc\xc3\xba\x98\x8c\xb4\x8a\x43\x6b\xc4\x94\x8a\xb5\xa4\x8a\xe7\x8f\x08\xaf\x2a\x2a\x23\x54\xa9\xca\xe1\xe1\x9b\x5e\xd5\xf3\xca\xb3\xca\xc5\xa2\x1e\x29\xe3\xda\x2a\xb6\x91\xa4\x07\x45\x07\xd5\x92\x48\x75\x6c\xd3\xd2\xd8\x7f\xc5\xe3\xe5\x59\x65\x5a\x89\x6a\x69\x51\xd7\xaa\x23\xe9\x88\x4b\x9a\x4f\x9c\x0b\x95\x29\xd7\xee\xd7\xa6\xc5\x13\x1d\xb9\xd0\x1a\x49\xfe\x90\x57\x9e\xb4\x26\xd6\x26\x16\xf5\xe5\xae\x96\x14\xeb\xe7\xf1\xad\x6e\x5b\x85\x66\x94\x1a\x8b\x2a\x72\x3a\x0e\x1e\xd7\x67\xc4\xc8\x71\x2e\xbf\x67\x41\x47\x62\xda\x34\x6c\x33\x34\x2d\x49\xe5\xb9\x76\x12\xa9\xb6\x98\xb1\x51\xc7\x49\x32\xdd\x82\xba\xa0\x3a\x4a\x3d\x18\x98\x6b\xbe\x17\x04\x3e\x0d\x0e\x52\x0e\xbd\x2b\x81\xef\xd3\x04\x79\x24\x0b\x3e\x1e\x54\x0e\xad\xb1\xe0\x76\x0e\xaf\x17\xb1\x31\xdf\xa3\x95\x95\xde\x93\xf1\xa0\xe2\xdd\xa0\x09\xf2\x49\x52\x65\x14\xfc\x29\x7b\x99\xad\x65\xab\x9c\xf2\x4e\x8e\x78\x93\x53\xe2\x34\x5b\xcf\xd6\x79\x97\x37\x78\x87\xf7\xc0\x1f\x39\xe5\x94\xbf\x65\x6f\x78\x8b\xf7\x78\x9b\xb7\x4e\x90\x88\xf7\xb3\x57\x45\x43\x7a\xb2\x9f\xbb\xbc\xcf\x5b\xd9\x6b\xde\xeb\xfd\xb1\xc6\xfb\xd9\x2a\x6f\xf0\x26\xef\x16\xa9\x14\xfc\xe1\x1f\x71\x3b\x5b\xcb\xde\xe2\x6f\x00\x00\x00\xff\xff\x60\xec\xbe\xf3\xf8\x02\x00\x00"

func localesRuLc_messagesWidgetMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesWidgetMo,
		"locales/ru/LC_MESSAGES/widget.mo",
	)
}

func localesRuLc_messagesWidgetMo() (*asset, error) {
	bytes, err := localesRuLc_messagesWidgetMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/widget.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
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
	"templates/views/widget.html":      templatesViewsWidgetHtml,
	"locales/ru/LC_MESSAGES/widget.mo": localesRuLc_messagesWidgetMo,
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
	"locales": &bintree{nil, map[string]*bintree{
		"ru": &bintree{nil, map[string]*bintree{
			"LC_MESSAGES": &bintree{nil, map[string]*bintree{
				"widget.mo": &bintree{localesRuLc_messagesWidgetMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"widget.html": &bintree{templatesViewsWidgetHtml, map[string]*bintree{}},
		}},
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
