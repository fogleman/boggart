// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package herospeed

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x5f\x6f\xe3\x36\x0c\x7f\xef\xa7\x20\xb4\x6e\x49\xb6\x3a\x4e\x8a\x3d\x0c\xad\xdd\x01\xdb\xbd\x0c\xd8\xc3\xa1\xd8\xf6\x1a\x28\x16\x1d\xeb\xa6\x48\x3e\x49\x76\x1a\x04\xf9\xee\x07\x59\xb6\xe3\xa4\xf9\x63\x17\x77\x0f\x57\x99\x22\x7f\x24\x7f\x14\x29\x65\xb7\x03\x86\x29\x97\x08\x24\x51\xd2\xa2\xb4\x04\xf6\xfb\xbb\x88\xf1\x12\x12\x41\x8d\x89\x89\x56\x1b\xf2\x72\x07\x00\xd0\x95\x26\x4a\x04\x6b\x16\xcc\x1f\xc1\xad\xcc\xba\x59\xbd\x99\x60\xfe\x58\xeb\x9f\xda\xbc\x2d\x72\x2a\x51\x74\x76\xdf\x6b\x34\x51\x1c\xeb\xb4\x7a\x5a\x09\x8c\x89\xa5\xcb\x73\x48\xad\x66\x21\x1a\x40\x49\x4b\x90\xb4\x0c\x2c\x5d\x1a\x58\x52\xbd\x70\x0b\x72\x80\x11\xdc\x9c\xf3\xd5\x22\x09\x5e\xeb\xe6\x1a\x0d\x4a\x4b\x2d\x57\x92\xec\x76\xc0\x53\xc0\xaf\x30\xa5\x89\x13\x00\x71\xac\x35\x3e\x9d\xac\x44\xa7\x84\x92\xc1\x7e\xff\x12\x51\xc8\x34\xa6\xf1\x05\xbb\x2f\xb4\xa4\x26\xd1\x3c\xb7\x4f\xa5\xe2\x6c\x3c\x9b\x3c\x3b\x5b\x61\x10\xf6\xfb\xdd\x0e\xa6\xaf\xf8\xb5\x40\x63\xa7\xff\xbe\xfe\x3d\xfd\x4c\x6d\xe6\xc5\x1e\x9c\xbc\x38\xd0\xf9\x6f\x12\xc8\x67\x8d\x25\xc7\x0d\x81\x29\xec\xf7\x51\x48\x5f\xa2\x50\xf0\xef\x90\x5b\xa2\x64\xca\x57\x85\xf6\xfb\x1f\x4c\xf4\x1d\xc8\x47\xb2\xfe\xdd\xa3\xc5\x47\x60\xe7\xa8\xf8\xf3\xd8\xdb\x6d\x42\xa2\xb0\x10\x17\x76\x3a\xc7\xd3\xd2\x65\x70\xf9\x80\x5e\x3b\xa8\x5d\x04\x27\x81\x94\x32\x04\x4f\x20\x70\x79\x05\xcd\xfd\xeb\x41\xe7\x55\xfb\xd3\x3c\xde\x16\x96\x5b\x81\x37\xbc\xb6\x96\xd9\xe3\x81\xd9\x3f\xa8\xe1\x09\x70\x99\x2a\xbd\x3e\x62\x37\x7b\xec\x89\xd6\x1d\x22\x02\xa9\x4e\xf9\x1b\x79\x89\x42\xc6\xcb\xdb\x00\xb5\xda\x6d\x3d\xd7\xda\xd8\x21\x5d\x20\x54\xff\x07\xc6\x6a\x9e\x23\x03\x46\x2d\xf5\x72\x66\x03\x8d\x26\x57\xd2\xb8\x5a\x48\xb5\xd1\x34\x27\x60\xec\xd6\x15\x70\xc3\x99\xcd\x9e\xe6\xb3\xd9\x8f\x7d\xc9\xb2\x19\x52\xd6\x57\x57\xf7\x53\xac\x81\x4f\x66\xef\xaf\xdd\xde\xa7\x9a\xae\xd1\xa2\x6e\xca\x61\xb3\x41\xd0\x07\xa4\xff\xa8\x28\x70\x28\x4a\x14\xf6\x4d\xc5\x61\x0e\x20\x68\xa9\xd8\xb6\x7f\x22\xbb\x1d\x68\x2a\x57\x08\xf7\x2a\x77\x67\xf3\x01\xee\x4b\x97\x0e\x3c\xc5\x30\x3d\x6a\x99\x3e\x1d\x73\x88\x62\x40\x99\xbc\x01\x73\x7c\xd6\x41\x78\x22\x7b\x66\x7c\x0a\xe1\xc3\x1f\x8a\xd0\xbf\x1c\xe0\x59\xf3\x03\xb4\x6f\x01\xfb\x15\x25\x0a\xab\xfe\xba\x39\xd9\xea\x91\x3f\x68\x80\xb9\xe9\x13\xac\xb4\x2a\x72\xc8\x0b\x21\x02\xcd\x57\xd9\xb5\x91\x7c\x09\x87\xcb\xbc\xb0\x1e\xa8\xa7\x75\x85\x20\xe8\x12\xdb\x27\xc6\xd2\xca\xc1\x08\x15\xca\xb2\xb0\x56\x49\xb0\xdb\x1c\x63\xe2\x3f\x48\x07\x14\x1c\x30\x4f\x54\xbd\x90\xa9\xaa\x16\x66\x4d\x40\xc9\x44\xf0\xe4\xff\xea\xe6\x76\x57\xfe\x2b\xa6\x1a\x4d\x36\x9e\x0c\x0c\xa1\x0a\x83\xb7\xa4\x52\x03\x29\x0d\xcc\x56\x26\x04\xaa\x0b\xc2\x5f\xe4\xd5\x58\xa8\x5d\xf8\xc1\xe0\xe6\xf5\x95\x97\xc5\x59\x37\xa1\xcf\x70\xa0\x55\xe7\x41\x71\xed\x41\x50\xf3\xf0\x13\x53\x1b\x29\x14\x65\xf1\xbc\x1f\x93\xdf\x81\xae\xc6\xe5\x7b\xca\x3e\xb5\x3b\x1f\xe6\x8c\x0e\xe9\xf9\xea\x54\xf6\x9d\xc2\x43\xae\xdb\x9b\x6a\x7c\xbd\x02\xa3\x93\xbe\x65\x22\x50\x5d\xab\x31\xa9\xee\x55\xc8\xd0\x75\x6f\xf3\xc5\x59\x7b\xac\x09\x84\xb7\xa7\xc7\xf5\xc9\x75\x25\x83\x0b\x5b\x67\xc4\x27\xa2\xce\x67\xbd\xac\xff\x1c\xc2\xb9\xeb\xfc\xae\xfa\x62\xda\xf7\x59\xe4\x1f\xbc\x75\xcb\xd3\x3c\x17\x3c\xa9\xae\xa3\xf0\xf0\x1a\xee\x9c\xc9\x0d\x97\x4c\x6d\xa6\xc7\x4d\x0e\x31\xa4\x85\xac\xf8\x1c\x4f\x60\x77\x14\xe8\xfd\x78\xf4\x43\xad\x3d\x9a\x4c\xa9\xb5\x7a\x3c\x32\x3a\x19\x3d\xc0\xa8\x67\x03\x2d\x6c\x3c\xfa\x65\x2c\x71\x03\x9f\xa8\xc5\xf1\x64\x32\x5d\xa1\xfd\x87\xaf\xdd\xf2\xb9\xf5\xb5\x7f\x3e\x3c\xc1\x8e\x7d\x6a\xa4\x6c\x3b\x6e\x02\x84\x77\x11\xd6\x29\x19\xb4\x7f\x49\x8b\xba\xa4\x62\x7c\x36\xcb\x07\x57\xdb\x46\xb6\xd0\x5e\xb8\xe0\xb5\x8d\xfb\x0d\xf2\x33\xcc\x67\xb3\x59\x37\xa8\x7a\x1d\x85\x9e\xc8\x4e\x3d\xbe\x05\x00\x00\xff\xff\x5f\x2b\x27\x8b\xe6\x0e\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xcd\x8a\x13\x41\x14\x85\x4f\xe2\x88\x1a\x50\xc4\x95\x0b\x17\xd7\x85\x83\xb3\xe8\xb1\x3b\xba\x90\x4e\x3a\x11\xe7\x07\xc4\x09\x0c\x21\xba\x2f\x26\x95\x4e\x63\xa7\x2a\x54\x75\xcf\x28\xcc\x22\x8e\xe2\x0f\x08\xae\x44\x70\xe1\xcf\xce\x65\xd4\x4d\x10\x27\xcf\x70\xeb\x05\x7c\x06\x1f\x41\xd2\x69\x15\x6b\x73\xce\xa9\xf3\x55\xdd\xfb\xf3\xc2\xca\x1b\x00\x38\x09\xe0\x12\x80\x4d\x00\xa7\x00\xa4\x58\x9e\x17\x00\xce\x02\x78\x09\xe0\x34\x80\xb7\x65\xff\xb1\xd4\xcf\x00\xda\x15\x60\x0a\xe0\x22\x80\x5f\x15\xe0\x1c\x80\x33\x55\xe0\xfc\xe2\xae\xd4\xb5\x2a\x50\x29\xff\x3c\x51\x6a\xb5\x9c\xbb\xb2\x08\x1b\x5a\x0d\x92\x38\x37\x22\x4b\xb4\xc2\xa6\x3e\x50\xa9\x16\x7d\xec\x1a\xb9\x9f\xc8\x03\x74\xe5\xc0\x48\x3b\x44\x57\x8e\xb5\xc9\xbc\x8e\x8d\x93\xbe\x77\x3b\x8f\xad\xd7\xd3\x21\xf5\xe5\xfe\xad\x07\xc9\x50\x8c\xf4\xba\xc9\x6b\x3b\xc2\x66\x5e\xcf\x08\x65\x53\x91\x69\x13\xd2\xdd\xa2\xa2\x4e\x6e\xc4\x48\xf7\x35\x35\xff\xe3\x5b\xb5\x1d\xa1\xe2\x5c\xc4\xd2\xeb\x49\x31\x0a\xe9\x6f\x0e\xa9\x9b\x5b\x9b\x08\x55\xeb\xdc\xe9\x6c\x79\xf7\xa5\xb1\x89\x56\x21\x05\xeb\x7e\x6d\x43\xab\x4c\xaa\xcc\xeb\x3d\x1a\xcb\x90\x32\xf9\x30\xbb\x36\x4e\x45\xa2\x1a\xb4\x37\x14\xc6\xca\x2c\xba\xd7\xdb\xf6\x6e\xfe\xe3\x16\xfb\x0c\xa4\xf1\xb6\xd4\x9e\xee\x27\x2a\x0e\xa9\xb6\x9b\xe6\x46\xa4\xde\xb6\x36\x23\x1b\x92\x1a\x17\xd1\x46\xd7\x1b\xb4\xb4\x91\xba\x12\xf8\x51\x14\xd0\xea\x2a\x2d\xac\x7f\x39\x0a\x02\x6a\x93\x4f\x61\x91\x5b\x51\xfd\x4f\xd5\x8c\x6e\x2c\xec\xd5\x02\x6b\x06\x3e\x1d\x1e\x2e\x9f\xb4\xa2\xba\xbf\x46\x6d\x0a\x28\xa4\x7a\x03\xfc\x8e\xe7\x7c\xec\x9e\xf2\x8c\xbf\xb9\x27\x6e\xc2\x53\xf7\x8c\x67\xee\x35\xf8\x13\x7f\xe7\xa9\x7b\xce\x53\x77\xe4\x5e\x81\x3f\xb8\x09\xcf\xdd\x63\xfe\xc1\x73\x77\xe4\x26\xe0\xf7\xfc\x85\x8f\x79\xce\x5f\x79\x56\x10\xbf\x03\x00\x00\xff\xff\x24\x96\x1f\x2a\x3a\x02\x00\x00"

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
