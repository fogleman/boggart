// Code generated by go-bindata.
// sources:
// locales/ru/LC_MESSAGES/config.mo
// DO NOT EDIT!

package internal

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
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

var _localesRuLc_messagesConfigMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x51\x6b\x14\x31\x14\x85\xcf\x4a\x45\x98\xc7\x3e\xfb\x70\x7d\xb0\x28\x9a\x9a\xcc\x58\x28\xd9\xcd\x56\x5b\x5b\x14\xbb\xb8\x96\xd1\xf7\xdb\x9d\x74\x36\xb8\x9b\x0c\x49\x46\x14\xfa\x37\xfc\x59\xfe\x05\x7f\x8b\xcc\x2c\x5a\x7a\x9e\xce\xe1\x9e\xfb\x71\xb9\x7f\xf6\xf7\x7e\x01\xc0\x03\x00\x8f\x01\xbc\x04\xf0\x10\xc0\x0c\x3b\x2d\x01\xec\x03\xf8\x0c\x60\x3d\x01\x1a\x00\x8f\x00\xfc\x9e\x00\x13\xdc\xd7\xc0\xc0\x2a\xf8\x1b\xd7\x8a\xcc\xd7\x7b\xa1\xb3\x7e\xcd\xd7\xb8\xb2\x5d\x88\x59\x2c\x52\xeb\x1a\x71\xda\xb7\x49\xd4\x41\x53\x63\xbf\xbf\xf9\xe6\xd6\xbc\x0d\x87\xb1\x2f\x96\x9f\x6a\x71\x16\x2d\x67\x17\xbc\x78\xc7\xd9\x6a\x2a\xa5\x3a\x16\xb2\x12\x65\x45\x65\xa5\x8f\x8e\x5e\xc8\x4a\xca\xe2\x92\x53\x16\x75\x64\x9f\x36\x9c\x43\xd4\xf4\x71\x64\xd0\xa2\x8f\xbc\x0d\x4d\xa0\xd9\x3d\xf0\xbc\xb8\x64\xdf\xf6\xdc\x5a\x51\x5b\xde\x6a\xfa\x9f\x35\x5d\xf5\x29\x39\xf6\xc5\xe2\xc3\xe2\x5c\x7c\xb5\x31\xb9\xe0\x35\xa9\x43\x59\x9c\x05\x9f\xad\xcf\xa2\xfe\xd9\x59\x4d\xd9\xfe\xc8\xaf\xba\x0d\x3b\x3f\xa5\xd5\x9a\x63\xb2\xd9\x7c\xa9\x2f\xc4\xf1\x5d\x6f\xb8\xe7\xc6\x46\x71\xee\x57\xa1\x71\xbe\xd5\x54\x2c\x37\x7d\xe4\x8d\xb8\x08\x71\x9b\x34\xf9\x6e\x8c\xc9\x54\x53\xda\x59\xe3\x9f\x2a\x69\x8c\xa2\x83\x03\x1a\xac\x7c\x62\x94\xa2\x13\x92\xa4\xc7\x3c\x37\xe5\xbf\xd1\xcc\xbc\x1e\xec\xb3\xb1\x36\x53\x92\x6e\x6f\x77\x2b\x73\x53\xca\xe7\x74\x42\x8a\x34\x95\x53\x0c\x0f\x7f\xff\xf6\x14\x7f\x03\x00\x00\xff\xff\x21\x82\xc2\xa6\xd5\x01\x00\x00"

func localesRuLc_messagesConfigMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesConfigMo,
		"locales/ru/LC_MESSAGES/config.mo",
	)
}

func localesRuLc_messagesConfigMo() (*asset, error) {
	bytes, err := localesRuLc_messagesConfigMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/config.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"locales/ru/LC_MESSAGES/config.mo": localesRuLc_messagesConfigMo,
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
				"config.mo": &bintree{localesRuLc_messagesConfigMo, map[string]*bintree{}},
			}},
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
