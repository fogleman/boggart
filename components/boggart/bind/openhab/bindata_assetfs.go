// Code generated by go-bindata.
// sources:
// templates/layouts/input.html
// templates/views/input.html
// DO NOT EDIT!

package openhab

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/elazarl/go-bindata-assetfs"
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

var _templatesLayoutsInputHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x6d\x6f\xdb\xb6\x13\x7f\x9f\x4f\x71\x65\xf1\x47\x9d\xff\x2c\x3f\x24\x71\x11\xb8\x76\x80\x2e\xcb\x90\x02\x1d\x5a\xa4\xc9\xd0\xbd\xa4\xc5\x93\xcd\x85\x22\x15\xea\xe4\x87\x19\xfe\xee\x03\x25\x39\x91\xad\x07\x3b\x1d\x06\x4c\x7c\x21\xd3\x3c\xde\xfd\xee\xee\xc7\x3b\x71\xbd\x06\xc2\x30\x52\x9c\x10\xd8\x44\x71\xfd\xc8\xa0\x03\x9b\xcd\xc9\xc9\x7a\x0d\x02\x03\xa9\x11\x98\x6f\x34\xa1\x26\xe6\xfe\x1f\xbd\xf9\xe5\xcb\xf5\xfd\x1f\x5f\x6f\x60\x46\xa1\xba\x3a\x19\xb9\x17\x28\xae\xa7\x63\x86\x9a\x5d\x9d\x00\x00\x8c\x66\xc8\x45\xf6\x33\x9d\x86\x48\x1c\x66\x44\x91\x87\x4f\x89\x9c\x8f\xd9\x75\xa6\xd1\xbb\x5f\x45\xc8\x20\xd7\x3f\x66\x84\x4b\xea\x3a\x85\x1f\xc0\x9f\x71\x1b\x23\x8d\x1f\xee\x7f\xf5\x2e\xd9\xbe\xae\xed\x2a\x4b\x28\xa8\x58\x2e\x9a\xfa\xee\x3d\x7c\xf4\xae\x4d\x18\x71\x92\x13\x55\xb4\xf6\xe9\x66\x8c\x62\x8a\xa5\xdd\x9a\x87\x38\x66\x73\x89\x8b\xc8\x58\x2a\x6c\x58\x48\x41\xb3\xb1\xc0\xb9\xf4\xd1\x4b\x27\x6d\x90\x5a\x92\xe4\xca\x8b\x7d\xae\x70\xdc\x67\x57\x27\x2f\xda\x62\x5a\x29\x04\x5a\x45\x98\xbb\xe6\xc7\x71\xc1\x9a\x7b\x26\x46\xac\x60\xbd\xf3\x97\x7b\x22\x2e\x84\xd4\xd3\x61\xef\x43\x69\x29\xe4\x76\x2a\x75\xd5\x8a\x42\x22\xb4\x5e\x1c\x71\xdf\xed\x85\x3d\x91\xcd\xce\x4c\x86\xd3\x0a\xbb\x81\x32\x9c\x86\xa0\x30\xa0\xc6\xcd\x3a\x4a\xa8\x0d\xce\x27\x6e\x91\xd7\x2b\xb2\x72\x3a\xa3\x32\x52\xb7\xd1\xe3\x4a\x4e\xf5\x10\x7c\xd4\x84\x76\xdf\xda\xf3\x74\xbd\x06\x19\x00\x3e\x41\x27\x8b\x26\xe3\x5a\x58\x23\x45\x4a\xc6\xe2\x9e\xff\x57\xa0\x50\x52\xa3\x37\x43\x07\x62\x08\x17\xbd\x68\xd9\xe4\x54\x4d\x2a\x02\xa3\xc9\x8b\xe5\x5f\x38\x84\xfe\x65\xb3\x86\xea\x98\x66\x09\xf3\x5c\x4c\x87\x70\xb6\xaf\xc1\x3d\x29\x95\xaa\x00\xba\xe7\x58\xf8\x8a\x4f\x50\xd5\x53\x29\xb7\x3f\x38\xe0\x41\x96\xd8\x49\x42\x64\x74\x85\xb2\x2d\x98\xf3\xb3\x68\x09\x6f\x64\xe8\x0e\x08\xd7\x15\x19\x0e\xf9\xd2\x3b\x5a\xb8\x4c\x87\x46\xf1\x9d\xac\x36\x2a\xdf\x75\xae\x81\xae\x3b\x2a\xfb\x9d\x01\x86\x8d\x4c\xa9\x0b\x4f\x33\xeb\x77\x8c\x9c\x5d\x1e\x8d\xfb\xf0\x69\x73\xd1\xce\x49\x34\xe8\xfd\xaf\x22\x1b\x52\xd7\xae\x37\xda\x6a\xa0\xc2\x84\xfb\x8f\x53\x6b\x12\x2d\x3c\xdf\x28\x63\x87\xf0\xf6\x6c\xe0\x46\xd9\xfc\xc4\x58\x81\x76\x08\xfd\x68\x09\xb1\x51\x52\xc0\xdb\xe0\xc2\x8d\xb2\xe8\x56\x55\xdd\x7a\xa6\xca\xb3\x5c\xc8\x24\x1e\xc2\x45\xd5\x89\xc9\x2b\x24\x9c\x47\x4b\x27\x90\xbe\xab\x0a\xa9\xd4\x3b\x2c\xaa\x2d\x41\xb0\x53\x86\x68\x86\x21\x02\x13\xdc\x3e\x96\x6a\x10\xd4\x97\x91\x9a\x90\xf5\xfb\xfd\x32\xb2\x62\x20\xb8\xef\x46\x49\x64\xb3\x8f\x0e\x55\x8c\xff\x1c\x4e\x10\x04\xcd\x70\x7a\xbd\xde\x11\x58\xb4\x28\x42\xd9\xa2\xdb\xad\xe3\x13\x1e\x4b\x3f\x91\xaf\xaf\xe3\xe7\xef\x7f\xa8\x8e\x1f\xeb\x6d\xc1\xd3\xf2\x62\xda\x0c\x02\x1e\x4a\xb5\x1a\x02\xbb\x33\x13\x43\x86\xb5\xd9\x2d\xaa\x39\x92\xf4\x39\x6b\xb3\x8f\x56\x72\xc5\xda\x31\xd7\xb1\x17\xa3\x95\x15\x26\x8a\x2d\xa5\x92\xc1\xa9\xc0\x22\xf7\x77\x50\x05\xa4\xe9\xb4\x6f\x3f\x1e\xe0\xb2\x4c\xfc\x57\x37\xac\x7e\x29\xda\x05\xf3\xe5\x63\x03\x7b\x4d\xe2\x47\x3a\x56\x7d\xa2\xa1\xd4\xcf\xaa\xe1\x6d\x73\xf8\xbe\xef\xc6\xc1\x7a\xf7\x2a\x0c\xff\x52\x4f\x39\x5c\xe4\x2b\x08\x8c\xbe\x1b\xf5\x05\xf7\x50\xd5\xab\x0c\xde\x2b\x9a\x49\x7f\x50\xf9\xc9\x32\x47\xeb\x0e\x83\xda\xf6\xf4\x50\x0a\xa1\xb0\x81\xa8\x3d\x28\x7d\x59\xc1\x7f\xe9\xb4\x1d\xf8\xb2\x86\x97\x43\x63\xf3\xee\xde\xd0\x98\x72\xe6\x56\x4a\xec\xb5\xb7\xf3\x7a\xf2\x95\xcb\xec\xa8\x9b\x56\xd6\xfc\x06\xd6\x7d\xb9\x82\x8d\x5c\x45\xbc\x2a\xee\x9c\x28\xe3\x3f\x02\x0b\x8c\x0d\xb3\xeb\x5e\x41\x5d\xe1\xfa\xe2\x5b\x19\x51\x7e\x7f\xe1\x51\xa4\xa4\xcf\x49\x1a\xdd\xfd\x93\xcf\x79\xb6\xb8\x77\x9b\x09\x12\xed\x3b\x09\x88\x51\x8b\xdf\xb9\x4a\xb0\x75\x5a\xc1\xe4\x39\xb7\xb0\x0c\xd5\x2d\x51\x04\x63\xd0\xb8\x80\xef\xbf\x7d\x76\xb3\x3b\x7c\x4a\x30\xa6\xd6\x69\x39\x32\xb9\x7c\xc7\x44\xa8\x5b\xec\xeb\x97\x6f\xf7\xac\x0d\x6c\xbd\x86\x4e\xbe\xa9\xf3\x70\xf7\xb9\xf3\x8d\xac\xd4\x53\xd8\x6c\x58\x1b\xc8\x26\xd8\xa0\x28\x46\xca\x77\xde\x22\x17\x68\x5b\xcf\x97\x52\xe7\xb0\x53\x5e\x74\x79\xe9\x2d\x16\x0b\xcf\x05\xcc\x4b\xac\x42\xed\x1b\x81\x82\x35\xaa\xd7\xa2\xc5\xe6\x2e\x06\x63\x06\x3f\x81\x30\x7e\x12\xa2\xa6\xce\x14\xe9\x46\xa1\xfb\xf9\xf3\xea\x93\x68\xbd\x4b\x45\xde\x9d\x76\xd2\xf7\x69\x4d\xb6\x47\xdd\x2c\xde\xdb\xe4\x66\x19\x1d\x75\xb3\x3b\xf8\x73\xf2\xfe\x0e\x00\x00\xff\xff\xb9\xfa\x6c\x98\xce\x0f\x00\x00"

func templatesLayoutsInputHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLayoutsInputHtml,
		"templates/layouts/input.html",
	)
}

func templatesLayoutsInputHtml() (*asset, error) {
	bytes, err := templatesLayoutsInputHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/layouts/input.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsInputHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x8f\x9b\x40\x0c\xbd\xe7\x57\x58\x3e\x44\xc9\xa1\x44\xbd\xf5\x00\xe9\x61\xab\x5e\x1a\x6d\xab\xfd\xe8\x7d\x02\x26\x3b\x2a\x78\x76\x07\xb3\x6d\x84\xe6\xbf\x57\x06\x41\x20\x61\xd3\x6a\xe7\x10\x65\x6c\xf3\xfc\x9e\xfd\xa0\x69\x20\xa3\xdc\x32\x01\xe6\xce\x97\x08\x21\x2c\x62\xfd\x07\x8e\xef\xeb\x7d\x69\x25\x41\x4f\x52\x7b\x86\xdc\x14\x15\xe1\x76\x01\x00\xd0\x34\x60\x73\x60\x82\xc8\x0a\x95\xd1\x8d\x11\x3a\x38\x7f\x04\x6c\x01\xb4\x22\xb6\xe5\x01\x2a\x9f\x26\xd8\x34\x10\xdd\xd1\x4b\x4d\x95\x44\x8f\x77\xbb\xe8\x87\x91\x27\x08\xe1\xb3\x49\xc5\x3a\x4e\x6c\xea\x78\xa9\x3f\x89\x16\x4e\xe1\x42\x58\x56\x62\x84\x4e\xa9\x7b\xbd\x6a\x5c\x39\x1a\x49\xaa\xd7\xc3\xd2\xf0\xf1\x6b\x77\x13\x5f\x13\xc2\x66\xe0\x48\x9c\x29\x9f\x39\xca\x3b\xb3\xa7\x62\xcc\xb7\x68\x03\xb9\xf3\x09\xbe\x9a\xa2\x26\xdc\x0e\x4d\xbb\xda\x10\xe2\x4d\x5b\x74\x05\xde\x70\x06\x2b\x7a\x81\x48\x8e\xcf\x04\x28\xf4\x47\x8c\x27\x83\x6b\x58\x0d\x9d\x1f\xda\xd4\x17\x23\xf4\x60\x4b\x9a\x49\xdd\xd6\xe5\x9e\x3c\xae\x7b\x6a\x27\x78\x45\xae\xe4\x58\x10\xa0\xe1\xcc\x3b\x9b\xe1\xb8\xa8\xd5\xb1\xaf\x45\x1c\x83\xe3\xb4\xb0\xe9\xaf\x04\x2b\xe2\xec\xa7\x0a\x5a\xad\x5b\x49\xf6\xe3\x27\x06\xfc\xfe\x0d\x21\x6a\x25\x75\xf5\xdb\x71\xa7\xb1\xae\xb8\xd7\x70\x46\x45\x39\x7b\xf7\xbb\x1a\x8d\xb0\x3f\x1a\xee\xd6\xde\x16\x84\x80\x6f\x81\x4f\xe1\xfe\xa1\xcc\xf1\xcd\x93\xe1\x03\x4d\x25\x5d\x83\xd6\x63\xb3\x7e\x9f\xc0\xa6\xa4\xf1\x72\xc7\x76\xe8\x7c\x85\xb7\x8f\xbb\x9d\x36\xbe\xf0\xdb\x80\x1e\x6f\xfa\x81\x9c\x6c\x50\x54\x74\xb2\x3d\x3f\xd7\x32\xb3\xb6\xcb\xfd\x9e\xeb\x53\xcb\x24\xc8\x5d\x72\x22\x4b\xf1\x2f\x61\x06\x07\xcd\x03\x65\x46\x48\x6c\x49\x1f\x0a\x97\x9a\xe2\x12\x70\xf6\x21\xd5\xf6\x9f\xdb\x9a\x1b\xdb\x04\xb1\x1d\x74\xe7\x83\xc9\x28\xdf\xbb\xb2\x37\x0c\x00\xe7\xaf\x63\xbc\xd1\x6f\xc3\x76\x31\x44\xfe\x06\x00\x00\xff\xff\xcb\xb5\x17\xa2\xe4\x04\x00\x00"

func templatesViewsInputHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsInputHtml,
		"templates/views/input.html",
	)
}

func templatesViewsInputHtml() (*asset, error) {
	bytes, err := templatesViewsInputHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/input.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/layouts/input.html": templatesLayoutsInputHtml,
	"templates/views/input.html":   templatesViewsInputHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"layouts": &bintree{nil, map[string]*bintree{
			"input.html": &bintree{templatesLayoutsInputHtml, map[string]*bintree{}},
		}},
		"views": &bintree{nil, map[string]*bintree{
			"input.html": &bintree{templatesViewsInputHtml, map[string]*bintree{}},
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
