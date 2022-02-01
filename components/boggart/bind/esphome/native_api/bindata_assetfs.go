// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/views/index.html
// templates/views/light.html
// locales/ru/LC_MESSAGES/light.mo
// locales/ru/LC_MESSAGES/widget.mo
package nativeapi

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

var _templatesViewsIndexHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdd\x6e\xdb\x3a\x12\xbe\xef\x53\x0c\x88\xb4\x48\x80\xca\x4a\x72\xb5\x70\xed\x2c\x8a\x0d\x16\x5b\x20\xdb\x53\xb4\xe9\x75\x31\x92\x46\x16\x5d\x8a\x54\x49\xda\x89\x6b\xf8\xdd\x0f\x48\xfd\x58\x8a\x65\x47\x4e\x93\xf6\x18\x68\x23\x91\xf3\x3f\x1f\x67\x86\x5a\xaf\x21\xa1\x94\x4b\x02\x16\x2b\x69\x49\x5a\x06\x9b\xcd\xab\x57\xeb\x35\xf0\x14\x46\x24\x2d\xb7\x9c\x8c\x5b\x9b\x24\x7c\x09\xb1\x40\x63\xa6\x4c\xab\x3b\x76\xf5\x0a\x00\xa0\xbd\x1a\x2b\x11\xe4\x49\x70\x71\x09\xee\xc9\xe4\xf5\xd3\xbd\x09\x2e\x2e\x2b\xfa\x87\x3c\xf7\xdf\x0a\x94\x24\x5a\xbb\xbb\x14\xb5\x65\x5d\x1a\x4f\x67\x31\x12\x54\x53\x96\x2f\xfe\xff\xc0\x58\xcd\x0b\x4a\x20\x41\x8b\xe5\x7a\x62\x03\x4d\xa6\x50\xd2\xf0\x25\x81\x54\x77\x1a\x0b\x06\xc6\xae\x04\x4d\xd9\x1d\x4f\x6c\x36\xbe\x38\x3f\x7f\xdd\xa3\xa5\xd4\x94\x11\x26\xfd\x7b\xe5\xbe\xde\xbf\x59\x09\xb8\x72\x51\xbd\xf8\x97\x04\xf6\x57\x34\xa7\xd8\xc2\x87\x6b\x06\x23\xd8\x6c\x26\xa1\xcd\x8e\xe0\xbe\x5d\x15\xf4\x24\xc6\x8f\x98\x3f\x8d\xf1\x8b\x45\xfb\x34\xce\xf7\xb1\xe5\x4a\x9a\x21\xbc\x93\x70\x5f\x0c\x1d\xdf\xde\xe8\x4f\x6c\xa4\x92\xd5\x7e\xb1\xeb\x35\x68\x94\x33\x82\x13\xfe\x16\x4e\x3c\xa2\x57\x30\x9e\x76\xc1\xbd\xdf\x97\x47\xd3\x9a\x38\x67\x2b\xb9\xa3\x32\xb1\x1f\xae\x4b\x67\x0f\xe0\xa5\x87\xd7\xa5\xf5\x29\x7c\x2e\xab\xc7\xf0\xf9\xc4\xd4\xcc\x3e\xb3\x70\x32\x98\xff\x20\x01\x3c\x38\xba\x91\x95\xc1\x4c\xab\x45\xc1\x40\x2b\x77\xcc\xca\x97\xc7\x85\x40\x99\x37\x9e\x02\xfd\xe8\xc6\x87\x09\x3e\xcb\xca\x22\x35\x44\xc8\x04\x21\xd3\x94\x4e\x99\x8b\xd7\xe8\x33\xfd\x58\x90\xb1\xa3\xaf\x9f\x6f\x46\x9f\xd0\x66\xb0\xd9\xfc\x1b\x3d\x40\xa7\xa5\x8e\x37\xca\xe7\x6f\xda\x9f\x51\xd6\xf2\x0b\x9c\x6f\x5c\xa6\xaa\x7c\x88\x55\xb9\x72\x6f\x06\x7a\xe7\x8d\xe3\xb5\xc0\x99\x58\x15\x99\x17\xd2\x3c\x05\xb4\xa2\x40\x15\x24\x19\x58\x6e\x5d\xf0\x9a\x33\x75\x4d\x16\xb9\x30\xcc\xa7\x8d\x5d\x4d\x42\x3e\x4c\xe7\x24\xc4\xc1\xa1\x27\x99\xf8\x46\x30\x90\x9c\xa7\xa0\x34\x9c\xee\xc9\xd6\x59\xcf\x8e\xb9\xe3\x36\xce\xd8\xd9\xd0\x4c\x42\x1f\x24\x4a\xf4\x32\x25\x07\x03\xa2\xfe\x0d\x06\x86\x71\x2a\x0e\xe3\xe2\x8d\xa7\x99\x9e\xef\xe0\x23\x71\x75\x47\xff\x02\x42\x1a\x6b\x0f\x22\xc5\x58\x55\xec\xa2\xe4\x8b\x0f\x30\xa8\x34\x3d\x1a\x28\x8d\xda\xa1\x80\x81\x0a\x34\xc2\xd0\x3f\x20\x11\x17\x7f\x2a\x11\x85\xc0\xd5\x6e\x22\x6e\x17\x5a\x82\x83\xe8\xef\xca\x42\x79\x74\x9f\x97\x7a\x12\x26\x7c\xf9\x48\x7b\x38\xd8\x40\xf6\x77\xf7\xc7\x0d\x99\x84\x7b\x7a\xfc\x24\xf4\xb3\xdd\x83\xe1\xb1\x6b\x6a\xeb\xb5\x7a\xac\xfe\xb4\xaa\xdc\x8b\xce\xb6\x4d\xd5\x1a\x29\x8b\xdf\xf4\x42\x4a\x2e\x67\x90\x62\x79\x58\x9c\x14\x81\x85\xa1\xa4\xb1\xe7\xe0\x30\xec\xd1\xd5\x37\x0a\x67\x97\xad\xf9\xf2\xf6\x7d\x3d\x71\x65\x97\x3d\xb4\x0b\x51\x0b\x94\xb8\x04\x89\xcb\x08\x75\xa0\x5d\xad\x06\x6f\xf2\x37\xab\x94\x88\xd4\xfd\xbe\x61\x58\xf0\xab\x09\xb6\x42\xe3\x1d\x08\x04\x97\xdf\xd9\xd5\xf6\x88\xa4\x08\x29\x06\x71\x46\x4b\xad\x64\xe0\x1a\xbf\xc3\xbe\x83\xf3\x24\x14\x3d\x87\x60\x12\x2e\x44\xcf\x6a\x3b\x0b\x82\x50\xa7\xfc\xde\x49\xda\xc1\x63\xdf\xd2\xb3\xdd\x23\x8e\xbf\x2f\x1c\x9e\x4a\x8f\xbb\x2f\xb8\x3e\xb7\x18\x34\x44\x43\x6b\xce\x4b\xbb\x88\xdb\x6c\x1a\x79\xd5\x52\x29\x70\x5b\xb8\xb7\xfb\x84\xc9\x6a\xbb\xeb\x51\xf9\x8b\xc7\xbb\xb6\x86\xb4\x56\xfa\x17\x47\xee\x56\x64\x6e\xd0\x58\xf0\x32\x8f\x8c\x4e\xc7\x98\x97\x2b\x5d\xc7\x7b\xf4\x49\xab\x99\x26\x73\x54\xb6\x8f\x9a\xca\x8b\x46\x41\x05\xe6\x1c\xf5\x8c\xcb\x20\x52\xd6\xaa\x7c\x7c\x3e\xb0\x2d\xf6\x89\x0c\x22\xd4\xd0\x7e\x69\xce\x4f\x2f\x1e\x3b\x94\x7e\x98\x76\xbd\x7e\x49\x5b\x44\x76\x28\xca\x2e\x5e\xc6\xdb\x4d\xe3\xe5\x9d\xa2\x26\x89\x50\x33\x40\xcd\x31\x58\xa2\x58\x90\x54\x77\xbe\x07\x7b\x95\x77\x9a\x5b\x4b\xd2\xcf\xf0\x5b\x92\x9c\xcb\x29\x3b\xef\xac\xe0\xfd\x96\xc9\x2a\x8b\xc2\xb3\x74\x0e\x7d\xbd\x5d\xeb\x85\xcd\x66\x5f\x11\xe8\x0d\x9b\x29\x50\xd6\x12\x05\x97\x14\x64\xe4\x0a\xef\xf8\xf2\xbc\xd8\x57\x6f\xf7\xfd\x1a\xcc\xbc\x4e\x20\x5a\x59\x32\xa0\x52\x78\x9d\x38\xe4\x48\x2e\x9a\x7f\x9d\x10\x74\x5c\x1b\x6e\x75\xe8\xcc\x1e\x7a\xcd\x78\x74\x54\x18\x48\xf6\x4c\x25\x27\xce\x28\xfe\x6e\x16\xf9\xf3\x9d\xd1\xff\x54\x12\x9f\x52\x73\x5a\xd6\xfc\xc9\x89\x69\x67\xe7\xf0\x98\xb2\x2b\x28\x55\x3a\x6f\x3a\xbd\xd2\x79\x90\x29\xcd\x7f\x2a\xe9\x90\xe5\xdf\x05\x46\x24\x02\x41\xa9\x85\x44\xab\xe2\xa7\x92\x54\x9f\x59\xb7\xcf\x20\x27\x9b\xa9\x64\xca\x0a\x65\x2c\x03\x9e\x4c\x99\xb2\xc8\xa0\x9a\xf7\x59\x3d\xf8\xfb\x45\xa9\x96\x28\x78\x82\x96\xae\x26\xa1\x63\xdf\x75\xad\x3f\x22\x4f\x98\x06\x5b\x1f\x45\x33\xc2\xa4\xb9\x5b\xae\xd7\xe0\x2e\x18\x3c\xfe\xdf\xed\xff\x6f\xe0\xb4\x7c\xfe\xfa\xf9\x06\x58\x98\xa0\xc9\x22\x85\x3a\x09\xd1\x18\xb2\x26\x5c\x92\x4c\x94\x36\x61\xf3\xfd\xd1\x8c\x24\xd9\x20\x32\x61\x6c\xca\xd5\xdb\x72\x35\x52\xca\x1a\xab\xb1\x18\xe5\x5c\x8e\x62\x57\x96\x7d\xcc\xcf\x9e\x51\xeb\xf6\xbb\x67\x6d\xc0\x76\xe5\x85\x0c\xa8\x12\x3e\xaf\xfc\xad\x5e\xfb\x54\xf4\x07\x7e\x6e\x9e\x31\xec\xe1\xdc\x84\xf3\x1f\x0b\xd2\xab\x51\x2b\xf2\xce\x96\xf9\x4b\x84\x3b\x32\x4e\xe1\xde\x1c\xbf\x88\xce\x6d\x42\x1f\xe8\x6e\x65\xfa\x37\x28\xaf\x7c\xdf\x0b\xaf\x67\x52\xbf\x45\xd7\xfc\x01\xb8\x76\x15\x4c\x4c\xac\x79\x61\xc1\xae\x0a\x9a\x32\x2c\x0a\xc1\x63\x74\x85\x25\x9c\xe3\x12\xcb\xcd\x56\xdf\x3d\x39\x4d\x54\xbc\xc8\x49\xda\xb3\x91\x1f\x83\x4f\xd3\x85\xf4\x85\x08\x4e\xcf\x60\xdd\xa9\x2d\xd7\xb5\x66\x55\xf8\xaf\xdc\xae\x72\xc2\xf4\x01\x91\xfb\x15\xa8\x51\x08\x12\x5f\x0b\xa1\x30\x31\x63\xb8\x78\xdb\x4b\x93\x7f\xc4\x9c\xc6\xc0\x52\xae\xf3\x3b\xd4\xc4\x76\xc9\x62\x4d\x68\xe9\x43\x8e\x33\xba\xcd\x16\x79\x24\x91\x0b\x33\x2e\x7d\xde\xa5\xc6\x38\xa6\xc2\x52\xf2\x5f\x2e\xc8\x8c\xa1\xe3\x7e\x8e\x71\xc4\x25\xea\xd5\xdb\x51\xc4\x25\xeb\xf0\x6e\xde\x35\xaf\x9b\xb3\x77\x55\xb9\x2c\x63\xd5\x2e\x95\x7f\x07\x00\x00\xff\xff\x9d\x35\xf5\xce\x3f\x1a\x00\x00"

func templatesViewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsIndexHtml,
		"templates/views/index.html",
	)
}

func templatesViewsIndexHtml() (*asset, error) {
	bytes, err := templatesViewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/index.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesViewsLightHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x99\x4d\x6f\xe3\x36\x13\xc7\xef\xcf\xa7\x18\x10\xc1\x83\x16\xa8\xec\xbc\x00\x8b\xa2\x8d\x0c\x34\x45\xda\x5b\x0f\xc9\xa2\x7b\x5c\xd0\xe2\xc8\x62\x4b\x91\x2a\x39\x4a\x9c\x1a\xf9\xee\x85\x48\x49\x91\x14\x2b\x51\xd2\x55\xd7\x49\x0e\x01\x29\x0d\x67\x86\xff\xf9\x91\xb4\xcd\xdd\x0e\x04\xa6\x52\x23\xb0\xc4\x68\x42\x4d\x0c\xee\xef\xff\x07\x00\xb0\xdb\x81\x4c\x61\x81\xd6\x1a\xdb\x3c\x3b\x17\xf2\x06\x12\xc5\x9d\x8b\x19\x57\x68\x09\xfc\xff\x48\x70\xbd\x41\xdb\x74\xa4\xcb\xa5\x73\x7c\xad\x90\xad\xfc\x30\x3f\x74\x5d\x12\x19\x0d\x74\x57\x60\xcc\x42\x87\x35\xbe\x12\x65\x1c\x32\x10\x9c\x78\x33\xbc\x0e\xc0\x80\x5b\xc9\xa3\x4c\x0a\x81\x3a\x66\x64\x4b\x64\xab\xff\x93\xcc\xd1\xfd\x78\xbe\x0c\x6e\x56\x4d\xbe\x83\x64\x97\x42\xde\xb4\xef\x50\x8b\xea\x45\x67\x6a\x5c\x0b\x58\xa0\x26\x49\x77\xb0\x70\xc4\x09\xf7\x4d\xd3\x9a\xdb\xee\x2c\x3a\x6f\x12\xa3\xa2\x5c\x44\x27\xa7\x50\xb5\x5c\xde\xb4\xb6\x2e\x3a\x39\xed\x8c\x19\x8e\xdb\x7e\x2e\xb8\x46\x35\xb0\x78\x6c\x45\x92\x7a\x0a\xf6\x2c\xb3\xd3\x95\x9f\xb0\x4f\x7f\xf1\x1b\xcf\xab\xe4\xcf\x97\xd9\xe9\x88\x7d\x37\x6f\x85\xdc\xa6\x72\xcb\x56\x1d\x85\x7a\xc6\x23\x8f\x7b\xd9\x35\xb8\x8c\xc4\x4b\x8d\xcd\x1b\xe3\xaa\x1d\x65\xc6\xca\xbf\x8d\x26\xae\xc0\xf7\x15\x5f\xa3\x8a\x14\xa6\xc4\xc0\x1a\x85\xc1\x8c\x41\x8e\x94\x19\x11\xb3\xc2\x38\x62\x20\x45\xa5\x73\x9e\x73\x2d\x18\x68\x73\xc3\x95\x14\x9c\x70\x7f\xd0\x61\x92\x92\x30\x0f\xc1\x36\xd6\x94\xc5\x48\xaa\xed\x48\x9f\x51\x65\x1f\x33\x4f\xc3\x03\x9d\x46\x93\x35\x2a\xa4\x0c\x75\xd9\xcf\x9a\xaa\x9f\x8d\x16\x7d\xf8\x57\x61\x77\xf2\xbd\x06\x76\x1d\xdc\x2f\x1a\xde\x46\x53\x5a\xfa\x90\xcf\xe4\xfd\x18\xc9\x0f\x4d\x6e\x1f\x26\xe7\x36\xf4\x34\xc1\x1c\x5a\xd1\xa6\xd9\x7a\x7b\xa9\x8b\x92\xea\x4d\x20\xc9\x30\xf9\x73\x6d\xb6\xad\xd0\x7f\xb8\xc8\xdd\x4a\x4a\x32\x06\x9a\xe7\xd8\xd4\x61\xb2\xf7\xfa\xaf\x82\xa6\xae\x60\xbd\x89\xf9\xde\xe2\xba\x5e\xe3\xe0\x03\xa3\x68\x77\x05\x58\x4e\x9c\xed\x94\x7a\xc0\xf8\x0a\x7a\x81\xc9\x13\xaf\x9b\x8d\x39\x2c\xfd\xeb\xb2\x28\x8c\x25\x77\x61\xe5\x26\x23\x8d\xce\x3d\x45\xd5\x97\x59\x1e\xeb\x36\xd6\x7c\x6b\xe4\xa2\x13\xe3\xe0\x16\x4a\x97\x62\x5b\x9d\x7d\x0c\x72\xa9\x63\x76\xcc\x20\xe7\xdb\x98\x9d\x30\x70\x84\x45\xcc\x16\xc7\x27\xac\xb7\x0d\xd6\x3a\x85\x9d\xad\xa3\xe3\x54\xc6\xc3\xba\xe8\x16\xe0\x86\xab\x12\x63\x56\x9d\x06\x01\xf3\x1e\x09\xec\x39\xb6\xff\x15\x87\xf5\xa1\x3a\xf2\x76\x0f\xa5\x57\x9b\xf5\xfc\x78\x5a\x14\xf3\x71\x79\x55\x39\x7f\xaf\x40\x7a\xe5\x02\x62\x55\x73\x2a\x94\x8f\x10\xbc\x42\x31\x2f\x7b\xf0\xc5\x60\xd9\x58\x44\x3d\x1f\x2e\xbf\x06\xf7\xef\x15\x98\x5a\xbd\x80\x4c\xe8\xbc\x1a\x1a\x2f\xd5\x5b\xc1\x66\xad\xca\x19\x3f\x20\x5e\x78\xef\xef\x15\x9a\xa0\x5d\x7d\x92\x55\xed\x57\x23\x53\xe9\x74\x78\x87\xdc\xa7\x4c\x12\xfe\xce\x43\x6e\x33\x73\x78\x5b\xc5\x9a\x0f\xc4\x4f\xc1\xfd\x7b\x25\xb1\x56\x2f\xa0\x18\x3a\xaf\x66\xd1\x4b\x75\x78\x30\xfe\x6c\x94\xb1\x1f\x31\x2f\xd0\x72\x2a\xed\x7f\x80\x64\x52\x45\x8c\xe8\x21\xe4\x7c\x78\xfa\xc9\x41\x2f\xd4\x7b\x45\xf5\xb1\xaa\x2f\xfb\xca\xb0\xa7\x2a\x8f\x10\xde\xc3\xca\x57\xa6\x59\x23\x7c\xa3\x50\xb7\x54\x5f\xa6\x29\x26\xe4\xbe\x85\xe3\xf9\x39\x46\x1f\x6b\x3e\x78\x2f\x6b\xff\x07\x47\xac\x43\x85\x09\x79\xea\x1a\x0d\x02\x43\x03\x45\xba\x90\x42\x18\x34\xc5\xff\x6e\x07\x7e\x1d\xc0\x51\xfa\x1d\x1c\x05\x9f\xf0\x43\x3c\xac\xf1\x73\xaa\xb4\xe9\x9a\x82\xa4\xd1\x1d\x9a\x1b\x9f\xf7\xf7\x2c\x50\x84\x7f\xb5\xcf\x8e\x6a\xd4\x2f\x1b\x93\x17\xfd\xb8\x14\x26\x89\x22\x66\x4d\x8b\xb5\x14\xaf\x9a\xc2\xb6\xa1\xfc\x2f\xb2\x21\xbb\x49\xaa\x3c\xb1\x1a\x7a\x13\x5e\x86\xe0\x87\xff\x21\x39\x55\xdc\x65\x91\x42\xbd\xa1\x6c\xbe\x75\xf4\x4b\x15\x05\x44\x69\x79\xa5\xf4\x01\xae\xa7\xee\x09\x40\xb8\xa5\x87\x03\x60\x74\xb7\xef\x2b\x17\x56\x5f\xff\x59\x8d\xfb\xf1\x9b\xf8\xb6\x44\x96\x6b\x27\xab\xf2\xcc\x4e\xc3\xc7\x36\xd4\xdb\x42\x62\x14\x85\x3d\xda\x05\x1e\xf6\xbc\x68\xa0\x70\x5f\xed\xd8\x1e\xea\xa4\xf4\x67\x67\x94\x14\xa3\x97\x4d\xfb\x06\x4d\x67\x6c\xac\x22\xb9\x88\x4c\x9a\x3a\xa4\xe8\x6c\x4a\x39\xea\xbb\x49\x7f\x71\x80\x5a\xb0\xba\x32\xae\x5c\xe7\xf2\xa1\x36\x6b\xd2\xb0\x26\x1d\xb9\x32\x49\xd0\x39\xd6\xee\xf8\xec\xa7\xa2\x50\x77\x01\xb3\xfe\xa5\xe4\x68\xc4\x57\xcb\x7f\xbe\xac\xd4\x99\x74\x69\x37\x78\xd4\xe9\xee\xbd\x1a\x7d\x68\xfd\x13\x00\x00\xff\xff\x7b\x4e\x20\xaa\x16\x1e\x00\x00"

func templatesViewsLightHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesViewsLightHtml,
		"templates/views/light.html",
	)
}

func templatesViewsLightHtml() (*asset, error) {
	bytes, err := templatesViewsLightHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/views/light.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesLightMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x50\xcf\x8b\x1c\x45\x18\x7d\x9d\x19\x7f\x64\x8c\xbf\xe2\xc1\x4b\x0e\xe5\xc1\xa8\x48\xc7\xe9\x8d\x88\xf4\x6e\xef\xea\x86\x5d\x11\x77\x24\x8c\xa3\x1e\xa5\x32\x5d\x3b\xd3\xa6\xa7\x6a\xa8\xee\x96\x2c\x04\xdc\x9d\xac\x3f\x20\x81\x45\x44\x05\x8d\xc8\x22\x88\x17\x71\x76\x49\xeb\x66\xb3\x99\x3d\x79\xf3\xf0\xd5\x41\xf0\x20\x5e\x04\xff\x03\xf1\x24\x48\x75\xf7\xb8\x78\x4b\x5d\xea\x7b\xf5\xbd\xf7\xbe\xf7\xd5\xef\x27\xeb\x9f\x00\xc0\x43\x00\x4e\x01\xb8\x0e\xe0\x51\x5b\x3b\x28\xce\x25\x07\xb8\x0b\xc0\x9a\x03\xd4\x01\xbc\xeb\x00\x0d\x00\x9b\x0e\xf0\x30\x80\x2d\x07\xb8\x1b\xc0\xb6\x03\x3c\x00\xe0\x5b\x07\xb8\x07\xc0\x8f\x95\xee\xb0\xea\xff\x54\xe9\x7f\xae\xb0\x71\x80\x1a\x80\x5f\x2b\xfc\x9b\x03\x3c\x02\xe0\x8f\x4a\xf7\xb7\x03\x2c\x38\xc0\x3f\x0e\x70\x12\x40\x50\x2b\xfd\xcf\x57\xf7\x5b\x35\xe0\x09\x00\xb2\x06\x9c\x00\xf0\x69\x0d\xf0\x00\x7c\x5d\x2b\x77\xf9\xa5\xe2\xfd\x59\xf5\xff\xaa\x95\x73\x8e\xd5\x4b\xbf\xe3\xf5\xb2\x7f\xaa\x5e\xee\xf3\x74\x1d\x38\x0b\xe0\xb9\x0a\xbf\x5d\x07\xec\x17\xd8\x2c\x8d\xf2\x2b\x70\xac\xba\xad\xce\xee\x62\xf7\x7c\xb0\x7a\x3b\x81\xa3\x63\xf7\xba\x0f\xc0\xfd\xd5\xbf\xda\x73\x1c\xe5\xfc\x7b\xa7\xa4\x17\x87\xc3\x78\x0d\x8b\x71\x26\xb0\xa8\xa3\x5e\x3f\x95\x22\x49\x70\x4e\xc5\x4a\xb3\x54\x0c\x86\x42\xf3\x34\xd3\x02\x4b\xab\xab\xa2\x9b\x62\x39\xe6\x49\x9f\x85\x99\xe6\x69\xa4\x24\x96\xe3\xa8\x7b\x51\x68\xbc\xa4\x85\x90\x58\xe1\x83\x0b\x21\xc7\xab\x4a\x0a\xb4\xb9\x0c\xd5\x00\x6d\x11\xe2\xb5\x54\xab\x0b\x02\x1d\xcd\x65\x12\x59\xd9\x91\xfe\xcd\x7e\x94\x0a\xb4\xc5\x50\xe9\xd4\x6d\x25\xbd\x28\x74\x17\xb3\x5e\xe2\x76\x94\xcf\x42\xf1\xce\x0b\x17\xa3\x3e\x1f\xa8\x33\x3a\x6b\xac\xf0\x24\x75\x0b\x8b\x98\xa7\x4a\xfb\xec\x95\xa2\xc5\x5a\x99\xe6\x03\x15\x2a\x36\xf7\x3f\xfe\x7c\x63\x85\xcb\x5e\xc6\x7b\xc2\xed\x08\x3e\xf0\xd9\x7f\xd8\x67\xed\x2c\x49\x22\x2e\x1b\xad\x97\x5b\x4b\xee\x1b\x42\x27\x91\x92\x3e\xf3\xce\x34\x1b\xe7\x94\x4c\x85\x4c\xdd\xce\xda\x50\xf8\x2c\x15\x97\xd2\x67\x86\x31\x8f\xe4\x2c\xeb\xf6\xb9\x4e\x44\x1a\xbc\xde\x59\x76\x9f\x3f\xe2\xd9\x3c\xab\x42\xbb\x4b\xb2\xab\xc2\x48\xf6\x7c\xd6\x38\x1f\x67\x9a\xc7\xee\xb2\xd2\x83\xc4\x67\x72\x58\xc0\x24\x38\x3b\xcb\xca\x32\x90\x8f\x7b\xcd\x20\xf0\xd8\xe9\xd3\xcc\x96\xcd\xc7\x02\xcf\x63\x0b\xac\xc9\xfc\x02\xcf\x07\x33\xd3\xd6\x5c\xf0\xac\x2d\x9f\x2c\x68\x73\x5e\x93\x5d\xbe\x5c\x4a\xe6\x83\x99\xe6\x53\x6c\x81\x79\xcc\x67\x33\xb3\xa0\xaf\xcc\x3a\xed\xd1\x01\xe5\x74\x9b\xf6\xcc\xc8\x5c\x03\x7d\x4c\x13\xba\x65\xae\xd0\x0e\x4d\xe8\x26\xe8\x7b\xb3\x4e\xfb\x34\x31\x1b\x65\xf7\x1b\xda\xa5\xdc\x8c\x68\x42\xbb\x34\x36\x5b\xcc\x8c\x28\xa7\x03\x3a\xa4\xdc\xac\xd3\xd8\x8c\xcc\x15\x7b\x83\xbe\x33\x9b\x66\x93\x72\xda\x37\xa3\x72\xca\x84\x6e\x58\x63\xfa\xc1\xce\xa1\x9c\x6e\x99\x6b\x74\x7b\xea\xcb\x68\xd7\x6c\xd0\xa1\xb9\x6a\x3e\xa4\x7d\xda\x03\x5d\xb7\x7e\xe6\x7d\x1a\xdb\x5c\x94\x83\x3e\xb3\x12\x1b\xd3\x5c\xb5\xa9\xbe\x30\x5b\x74\x40\x3b\x74\xc3\xce\xfa\xd2\x26\x02\x6d\xdb\xd8\xe6\x03\x1a\xd3\xcd\x29\xed\xf3\x22\xd4\xc6\x14\x6e\x9b\x51\x91\x64\xe7\x8e\x22\x95\x3b\xe5\xe6\xbd\x82\x38\x06\x7d\x54\x70\xac\xd3\xbf\x01\x00\x00\xff\xff\x8e\x34\x3d\xdf\x75\x04\x00\x00"

func localesRuLc_messagesLightMoBytes() ([]byte, error) {
	return bindataRead(
		_localesRuLc_messagesLightMo,
		"locales/ru/LC_MESSAGES/light.mo",
	)
}

func localesRuLc_messagesLightMo() (*asset, error) {
	bytes, err := localesRuLc_messagesLightMoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/ru/LC_MESSAGES/light.mo", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x51\xbf\x6f\x1c\x45\x14\xfe\x76\x6f\xc3\x8f\xe3\x67\x4c\x41\x03\xd2\x44\xc2\x11\x48\x6c\xb8\x33\x29\x60\xed\xb5\x03\x71\x2c\x99\xd8\x21\x8a\x0f\x3a\x8a\xcd\xde\x78\xbd\xe4\x6e\xc6\xda\xd9\x05\x2c\xa5\xf0\x9d\x0b\xc2\x0f\xc5\x52\x04\x12\xd0\xa0\x04\x09\x5a\xe7\xcc\x29\x17\xe2\x9c\x41\xe2\x0f\x78\xd3\x21\x0a\x24\x4a\x0a\x1a\x68\x29\xd0\xec\x6c\x7e\xf9\x15\xbb\xef\x7d\xef\xfb\xde\x7c\xf3\xe6\x8f\x09\xef\x4b\x00\x78\x06\xc0\x73\x00\x7e\x04\xf0\x3c\x80\x25\x07\x65\xfc\xec\x00\x4f\x02\xf8\xc5\x01\x1e\x06\x40\x0e\xf0\x08\x80\xdf\xaa\xfa\x4f\x07\xa8\x01\xf8\xcb\x01\x3c\x00\x7f\x3b\xc0\xa3\x00\xfe\xa9\x78\xff\x39\xc0\x21\x00\x87\x5c\xe0\x21\x00\x8f\xb9\x76\xde\x84\x0b\xd4\x01\x1c\x71\xed\x9c\x97\x5d\xab\x3f\xee\xda\x79\xaf\xbb\x80\x0b\x20\x74\xad\xfe\x44\xc5\x5b\x70\x81\x39\x07\x38\xe3\x5a\xcf\x97\x6b\xc0\xd3\x00\xae\x55\xff\xdd\x1a\xf0\x38\x80\x5f\x6b\xf6\x7c\x5d\xb3\xe7\xfe\x5e\x03\x0e\x1b\xbf\x15\xef\xdf\x1a\x30\x01\xc0\xf3\x2c\xff\x59\x0f\x98\x04\xf0\x82\x67\xf1\xb7\x3c\xcb\x7b\xcf\xb3\x7a\x51\xe1\x45\x85\x5f\xaa\x74\x5f\x57\xf5\xf7\x1e\x60\x56\xf6\x84\xb9\xa3\x5d\x5d\x79\xcf\x3b\x61\xbc\x18\x7e\xbd\xaa\x27\xf0\x60\x1c\xbe\x2f\x37\x7b\x30\x33\x9f\x82\xbd\xb3\x5b\xbd\x8f\x09\xb3\x1b\xb3\x5f\xb3\x13\xe3\x0b\x93\x6d\x76\x7e\x23\xe7\x8a\xc9\x55\x36\xd9\xc6\x1b\x71\x9e\x4a\xa1\x70\x72\x8d\xc7\x17\x54\xd1\xc5\x3c\xcf\xa3\xb4\xa3\x70\x9a\x6f\xe0\x4c\xd4\xe5\x78\xfb\xfc\xfb\x3c\xce\xd9\xe2\x3c\xce\x66\x32\xc9\xb8\x52\x58\xc9\xa3\x9c\x97\xdf\x42\x61\xa5\x88\x63\xae\x14\xcb\x65\x92\x74\x38\x56\x3e\x4c\xf3\x78\x8d\xc9\xd5\x55\xb4\x8a\x4c\x30\x29\xd0\xda\x58\xe7\x30\x80\x14\xc8\x78\xd4\xde\x40\x56\x08\x91\x8a\x04\xe7\xf8\xba\xcc\x72\x7f\x59\x25\x69\xdb\x7f\xb3\x48\x94\xdf\x92\x01\x6b\xf3\x0f\x4e\x5c\x48\xd7\xa2\xae\x3c\x96\x15\xf5\xa5\x48\xe5\x7e\x2b\x8b\x84\xea\x44\xb9\xcc\x02\x76\xba\x6c\xb1\xe5\x22\x8b\xba\xb2\x2d\xd9\xcc\x03\xfc\xd9\xfa\x52\x24\x92\x22\x4a\xb8\xdf\xe2\x51\x37\x60\x77\xeb\x80\x9d\x2b\x94\x4a\x23\x51\x5f\x5e\x5c\x3e\xe5\xbf\xcb\x33\x95\x4a\x11\xb0\xe6\xb1\x46\xfd\xa4\x14\x39\x17\xb9\x6f\xac\x06\x2c\xe7\x1f\xe5\xaf\xac\x77\xa2\x54\x4c\xb3\x78\x2d\xca\x14\xcf\xc3\x77\x5a\x0b\xfe\x6b\xf7\x78\xc6\xcf\x2a\xcf\xfc\x53\x22\x96\xed\x54\x24\x01\xab\x9f\xed\x14\x59\xd4\xf1\x17\x64\xd6\x55\x01\x13\xeb\x65\xa9\xc2\x57\xa7\x99\x4d\x43\x31\xd9\x6c\x84\x61\x93\x1d\x3d\xca\x4c\xda\x38\x12\x36\x9b\x6c\x8e\x35\x58\x50\xd6\xb3\xe1\xd4\x9d\xd6\x4c\x78\xdc\xa4\x2f\x96\xb4\x99\x66\x83\x5d\xbc\x68\x25\xb3\xe1\x54\xe3\x25\x36\xc7\x9a\x2c\x60\x53\xd3\xe6\x31\xe9\x3a\xed\xd0\x4d\xdd\x67\x34\xa2\x1b\xe6\x45\xe9\x0b\x1a\xd2\x4d\xdd\xd3\x7d\x1a\xd0\x48\x6f\x83\x7e\xa0\x21\xfd\xa4\x7b\x7a\x8b\xf6\x68\x8f\x76\x4a\x86\xee\xd3\x0e\xdd\xa2\x11\xe8\x1b\xba\xa5\x2f\xeb\x8f\x41\x5f\xd1\x9e\xde\xc6\xe2\x3c\xa3\x31\x5d\xd7\x9f\x96\xa2\xbe\xa1\x7f\xab\x37\x69\x4c\xbb\x7a\x93\x86\xba\xa7\x7b\xa0\xab\x34\x2e\xe7\x8f\xf5\x36\xdd\xa6\x11\x0d\x41\x57\x0d\x55\xf7\xf5\x96\xe9\x7f\xa7\x7b\xb4\x4f\x43\x7d\x89\x6e\xd3\x98\x95\xe9\xa6\x19\x67\x4f\xa2\xa1\x81\x41\x57\xf4\x67\x77\xa1\x91\xee\xeb\xcf\x41\x57\x0e\x02\xd7\x68\x44\xfb\xa0\xc1\x7d\xd4\x4a\x3d\x38\x08\xec\xd2\xd8\x58\xa2\x81\x29\x6e\xd0\x0e\xed\xeb\x2d\xfd\x49\xd5\xfd\x3f\x00\x00\xff\xff\xc9\xba\x39\x6d\xb7\x04\x00\x00"

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
	"templates/views/index.html":       templatesViewsIndexHtml,
	"templates/views/light.html":       templatesViewsLightHtml,
	"locales/ru/LC_MESSAGES/light.mo":  localesRuLc_messagesLightMo,
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
				"light.mo":  &bintree{localesRuLc_messagesLightMo, map[string]*bintree{}},
				"widget.mo": &bintree{localesRuLc_messagesWidgetMo, map[string]*bintree{}},
			}},
		}},
	}},
	"templates": &bintree{nil, map[string]*bintree{
		"views": &bintree{nil, map[string]*bintree{
			"index.html": &bintree{templatesViewsIndexHtml, map[string]*bintree{}},
			"light.html": &bintree{templatesViewsLightHtml, map[string]*bintree{}},
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
