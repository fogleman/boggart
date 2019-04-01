// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package mercury

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x6d\x93\xdb\xb6\x11\xfe\xee\x5f\xb1\xc3\x39\xd7\xa7\x24\xd4\x7b\xdc\xf6\x22\x29\x1f\xea\x76\x9a\x99\xa4\x93\xb1\xaf\xed\xb4\x49\x7a\x03\x89\xe0\x09\x31\x09\xd0\x00\xa8\xf3\x55\xa3\xff\x9e\x01\xc1\x17\x50\x47\x49\x24\x45\xda\x56\x4c\xcd\xf8\x2c\x12\xc0\x2e\xf0\xec\x83\x5d\x0a\x5c\x60\xbb\x05\x07\xbb\x84\x62\xb0\x56\x8c\x4a\x4c\xa5\x05\xbb\xdd\xb3\x99\x58\x71\x12\x48\x90\x8f\x01\x9e\x5b\x28\x08\x3c\xb2\x42\x92\x30\x3a\xf8\x15\x6d\x90\x2e\xb4\x16\xcf\x00\x00\xdc\x90\xae\x54\x09\x6c\x90\x17\xe2\xbf\x31\xee\x23\x79\x4d\x7b\xb0\x8d\x4a\xd5\x87\x63\x19\x72\x0a\x14\x3f\xc0\x77\x54\x7a\xfd\x7f\x84\xfe\x12\xf3\xb8\x66\xaf\xef\x26\x4d\xbe\x89\x5a\xec\x9e\xcd\x06\x5a\xc1\xe2\xd9\xcc\x21\x1b\x58\x79\x48\x88\xb9\xc5\xd9\x43\xac\xd1\xbc\xbb\x62\x9e\xed\x3b\xf6\x68\x0c\xea\x9b\xf0\x93\x6f\xef\x85\x3d\x1a\xc7\xf5\xf7\xdb\xbc\xbf\x0b\x10\xc5\x9e\x51\xfa\xb4\x46\x02\x46\xbe\x4e\x5a\x8f\x33\x0f\xcf\x2d\x89\x96\x45\x92\xd2\x9a\xa1\x97\x08\xa4\x68\x03\x14\x6d\x6c\x89\x96\x02\x96\x88\xdf\xa9\x2f\x56\x26\xc6\x23\xa2\x48\x57\x2a\xc9\x23\x71\xdd\x80\x63\x81\xa9\x8c\x6c\x61\x6d\xb7\x40\x5c\xc0\xef\xa0\x8f\xb4\x09\x2c\x65\xbc\x44\xa7\xba\xb7\xc1\xaa\x12\xa6\x0e\xec\x76\x8b\x19\x82\x35\xc7\xee\xfc\x40\xbb\xcc\xb2\x37\x1b\x46\x9c\xeb\x61\xef\x1b\xd5\xd6\x13\x18\x76\xbb\xed\x16\xfa\xaf\xf1\xbb\x10\x0b\xd9\xff\xe7\xeb\xef\xfb\x3f\x22\xb9\xd6\xb7\xb5\x70\x6b\xa1\x84\x8e\xfe\x44\xc1\xfa\x4b\xc8\x39\xa6\x12\x7c\x8c\x44\xc8\xb1\x8f\xa9\x14\x16\xf4\x61\xb7\x9b\x0d\xd0\x62\x36\xf0\x48\x03\x03\xf5\x19\x95\x6b\xef\xb1\xee\x78\x8d\xe6\x75\x86\xfd\xad\x96\x33\x8f\xc5\xfc\x01\x79\xde\x7c\x54\x84\xc5\x0f\xba\x02\x08\x35\x10\x21\xc9\xaa\x79\x24\x1c\x22\x02\x0f\xd5\x46\xc2\x68\x7e\x0e\x12\xb1\x98\x22\x0c\x5e\x25\x1a\x1a\x1e\xf8\x9a\x79\xc4\x41\x8f\xa2\xee\xc8\xcd\xf6\xe7\x0c\x3d\x91\x53\x34\xf6\xbf\xa7\x3a\x4e\x0f\x7e\x36\x08\xbd\x03\x25\x86\x6b\x92\x68\x69\x1f\x76\x4e\xb9\x16\x7b\x4e\xca\x94\xa0\xee\x80\x8b\x1c\x0c\x1a\x31\x20\xf4\x88\x34\xf5\x39\x3a\x87\x8e\xb6\x8c\xfa\x13\xa4\xda\xf1\x7b\x69\xfb\xa1\xc4\x0e\xb8\x8c\x4a\x7b\x34\x01\xdf\x5e\xda\x93\xe1\x09\xfd\x46\x3f\x38\x7a\x80\xeb\xbc\xaf\x71\x90\xc4\x20\x89\x8f\x81\x51\x70\xf0\x86\xac\x30\x3c\x8f\x50\xa7\xc4\x4b\xff\x5d\x07\x9c\x50\x09\x56\x2e\xba\xfd\x7c\x20\xbc\xfd\x6c\x2d\x1c\xb6\x0a\x95\xfb\xea\x3f\x70\x22\xf1\xb5\x52\x72\xcb\xde\x48\x4e\xe8\xfd\xf5\x0b\x0b\xae\xfb\xea\x4e\x5f\x07\x32\xb0\xc6\xc3\xe1\x4b\x7b\x38\xb2\x87\xe3\xdb\xd1\xd7\x37\xc3\xe9\xcd\xf0\xeb\xff\x0e\xff\x78\x33\x1c\x5a\x3d\xb0\x5e\xf4\x7a\x69\x58\xb3\x7a\xbd\x52\x98\x0d\x82\xc5\xb3\xd3\xb5\x54\xf4\xc0\x86\x6d\x3d\x0c\xd1\x5f\x5b\x48\x4e\x02\xec\x28\x6c\x90\xbe\xef\x48\x9b\x63\x11\x30\x2a\x94\xc9\x29\x7b\xe0\x28\xb0\x40\xc8\x47\xc5\x93\x07\xe2\xc8\xf5\xcd\x68\x38\x7c\x5e\xd2\x12\x33\xb9\xc6\xc8\x29\x5b\x97\x97\xab\x18\x0b\xde\xf3\xa0\xc9\xf4\x91\xeb\x9a\x52\x6e\x47\x30\x80\x7f\x37\x20\x67\xdc\x90\x9c\x49\x43\x72\xa6\xb5\xe4\xcc\x06\x65\xed\xa1\x64\x56\xb0\xf2\x92\x39\x8f\x15\xe6\x31\xbd\xc7\x70\x45\xbe\x82\x2b\x15\x1f\x05\xdc\xcc\xa1\xaf\xbf\x95\x98\x1e\x5a\x21\xd7\x7e\x09\x51\x07\xae\xf1\xbb\x58\x50\x3f\x22\x0d\x5c\xe9\xf9\x19\x5d\xf4\xcc\xe2\xff\x60\xc4\x93\x52\xf5\xbd\x67\x44\x10\x11\xae\x56\x58\x08\x23\x84\x54\xb0\x8c\x93\x5a\xc6\xec\x48\x5f\xbb\x0c\xb8\x52\x6a\xb6\xdb\x5c\x2f\x22\xab\x95\x84\x37\xd1\x51\xba\x72\xd4\xa0\xd4\x93\xfc\x9e\xa7\x33\x1f\xe5\xb3\x0e\xdf\x8e\x60\xb7\x33\xfc\x58\xa5\x7e\x68\x3b\xdd\xcb\x4c\xd8\x2d\x57\x00\x0f\xcb\xda\x3a\x3f\xa8\x00\xd1\xc4\x62\x1e\x5a\x62\x0f\xa2\xbf\xb6\xa3\x28\xc5\xad\x45\xa3\xa3\x7e\x85\x3d\x89\xf2\x43\x87\x41\xb3\xc0\x6a\x2c\x72\x2a\x66\x03\x35\xc8\xca\x20\x47\xcf\x2c\xc4\x05\xaf\x65\xa4\x93\x79\xd2\x41\x0d\x99\x63\x69\x07\x6a\x42\x5d\xd6\xe1\x9c\x77\x1e\x7a\x00\x0d\x3b\x0f\xec\xa2\xd0\x93\xed\x43\x5d\x17\x87\x08\xcf\xf2\xa1\xe2\x72\x02\xcb\xb8\xc9\xc0\x32\xbe\x94\xc0\x32\x6e\x7f\x16\x8e\x5b\x0d\x2c\xad\x20\xdd\x4a\x60\xb9\x4c\xa8\x8d\xc0\xd2\x0a\xd4\xcd\x07\x96\xcb\xc4\xd9\x74\x1e\x17\x13\x58\x8a\xa0\xee\x02\xcb\x1e\x48\x93\x26\x03\xcb\xe4\x52\x02\xcb\xa4\xfd\x59\x38\x69\x35\xb0\xb4\x82\x74\x2b\x81\xe5\x32\xa1\x36\x02\x4b\x2b\x50\x37\x1f\x58\x2e\x13\x67\xd3\x79\x5c\x4c\x60\x29\x82\xba\x0b\x2c\x7b\x20\x4d\x9b\x0c\x2c\xd3\x4b\x09\x2c\xd3\xf6\x67\xe1\xb4\xd5\xc0\xd2\x0a\xd2\xad\x04\x96\xcb\x84\xda\x08\x2c\xad\x40\xdd\x7c\x60\xb9\x4c\x9c\x4d\xe7\x71\x31\x81\xa5\x08\xea\x4f\x2d\xb0\x94\x7f\x33\x56\xad\x17\xb3\x41\xc9\x77\x63\xb3\x41\xf4\xae\xf6\xe4\xcb\x78\x63\xc6\x15\xe5\x72\x9c\xd6\xe3\x32\xee\x03\x71\xe6\x59\xab\x98\x07\xaa\xc0\x5e\x33\x4e\xfe\xcf\xa8\x44\x1e\x44\xd7\x9a\x16\x1e\x76\xa5\x15\xbd\x50\xb6\x25\xbb\xbf\xf7\xf0\xdc\xda\x20\x8f\x38\x48\x32\x6e\x81\x8f\xe5\x9a\x39\x73\x2b\x60\x47\x93\x99\x72\xbd\x30\x32\x1b\x22\x3d\xf7\x9c\x85\x41\xc9\xc6\x91\x00\xcd\xda\x34\x1b\x8c\x4a\xce\x3c\x3b\xbe\xa9\x73\xc3\xa6\x49\x6a\xd8\xd4\xc8\x0c\x53\xa3\x9a\x5b\x3e\x73\xf0\x9d\x1c\x55\xd0\x07\x71\xfc\x8e\xde\xbe\xbe\x59\xb3\x07\xb8\x1d\xe9\xb7\xaf\xf9\xc9\xc4\xf1\xbb\x90\x70\xec\x58\x8b\x2f\xaa\x52\x7c\x36\x88\xba\x5f\xa1\xc1\xd3\x7c\xb8\x97\xc9\x98\x5f\x16\x66\xc3\x95\x92\x4a\x68\x10\x26\x13\x7f\xb5\xc6\xab\xb7\x4b\xf6\x3e\x25\xc9\xaf\xc2\x16\x0f\x44\xae\xd6\x16\x50\xe4\xe3\x0c\xca\x88\x52\xe9\x45\x9a\x35\x12\xa7\x5e\xf4\xe3\x92\x1e\x58\x92\x87\x58\xa7\xee\x28\xd9\xd8\x49\xe7\x13\x0c\xaa\x60\xe5\x90\x4d\xd9\x99\x5d\xbe\xea\xa7\xc0\xca\xaa\xf6\xda\x63\xe5\xb8\x63\x25\xce\xa0\x34\x58\x39\x3e\xc8\xca\x71\xc7\xca\x93\xac\x9c\x9c\xc7\xca\x49\xc7\x4a\x9c\x41\x69\xb0\x72\x72\x90\x95\x93\x8e\x95\x27\x59\x39\x3d\x8f\x95\xd3\x8e\x95\x38\x83\xd2\x60\xe5\xf4\x20\x2b\xa7\x1d\x2b\x4f\xb1\x12\xf9\x2c\x3c\x9a\xa4\x5b\xf4\xc9\x33\x33\x16\xd1\xb1\x33\x0f\x69\xc6\xd0\xe4\x46\x31\x4b\x75\x69\xc7\xd4\x53\x4c\x0d\xd8\x03\xe6\x67\x11\x55\x4b\xe8\x78\x9a\x03\x34\xa3\x69\x7c\x5d\xcc\xd2\xa8\xb0\x23\xe9\xc9\x20\x4f\x7c\x7c\x16\x47\x23\x01\x1d\x45\x4d\x38\x8d\x50\x1f\x5d\x1e\x08\xf6\xc4\xc7\x1d\x3f\x4f\xf1\xd3\x41\xf2\x3c\x7e\x46\x02\x3a\x7e\x9a\x70\x66\xfc\xd4\x97\x4f\xf8\xa9\x6e\x77\xd4\x3c\x42\x4d\x35\x77\xcf\x58\xe1\xfc\x8e\x4a\xcc\x37\x7a\x11\x18\x28\x93\xb0\x8a\x37\x5c\x49\xc4\x89\xeb\x7e\x06\x74\xa5\xd1\x66\xe5\x84\x9a\x09\x9c\x11\x31\xd3\x8b\xe8\x7d\x47\xb4\xc5\xb0\x1f\xdf\x83\xdd\xce\x82\x00\x49\x89\x39\x9d\x5b\xff\xfb\xc9\xfe\xf2\x97\x6f\x7f\x1a\xda\x7f\xfe\xe5\x8b\x2b\xab\xa3\xa4\x7f\xce\xf2\x66\x8e\x92\x1d\x1d\xb3\xe5\xcd\xf4\xe2\x09\x1d\xc7\x1d\x1d\x4f\xd1\xb1\xf6\xba\x66\x8e\x8e\xd1\xa3\xfc\x57\x7a\x23\x2a\xa2\xce\xe7\x12\xd2\x8b\x48\x39\x31\x49\x39\x29\x20\xe5\xa4\x23\xe5\x29\x52\xd6\x5e\xd6\xcc\x91\x12\xb9\x12\x73\x08\x38\x16\x02\x96\x68\xf5\x16\x96\xa1\x94\x8c\x7e\xa6\xbc\x9c\x9a\xbc\x9c\x16\xf0\x72\xfa\xd1\x78\x59\x99\x98\x1e\xbd\x13\xcc\x23\x8e\xb5\xf8\xa0\xcc\x2e\x67\xb8\x84\xe5\xcc\x75\x05\x96\x76\x55\x07\x3b\xd3\x2c\x8d\x2d\xc9\xb1\xc0\x32\xfd\xc5\xb0\x94\x14\x96\x92\xea\x6c\xa0\xc4\x82\xba\x4a\xb6\x35\xfa\xb5\x6e\xa2\x37\x46\x6b\x61\xe7\xf4\x40\x84\x4b\x9f\x3c\xed\x42\x92\xfb\x95\xf4\x42\xa0\x0d\x36\x3a\xf1\x46\x5d\xd6\xec\x43\xf3\x8e\x6d\x36\x50\x06\xaf\x95\x5a\x62\x1e\x96\x71\x5a\xcf\xef\xfc\x4c\x82\x57\xd9\x69\x26\xb5\x77\xee\xd7\x3a\xd7\xe0\x13\xd8\xb7\x0f\x4f\xf6\xee\x3b\xe8\x31\xda\xb9\x9f\x50\xa4\x52\x62\x56\x15\x1b\x00\xa4\x7b\xec\x95\xd2\xec\xe0\x0d\xab\xf2\x36\x7a\x43\x94\xde\xae\xaf\xe4\x3d\xd9\xac\x5f\x31\x1f\xb9\xb4\x71\xe0\x13\x49\x1f\xfb\x6c\xa7\x72\x32\x1e\xdf\xb1\x55\x94\x1a\x9b\x27\x36\x65\xa7\x56\xd5\x9c\xe3\x47\x84\xff\x4b\x05\x89\xb3\x5d\xc7\x2b\xac\x93\x27\x49\xf2\x10\x77\x71\x0e\xa4\xee\xb4\x4f\xdc\xaf\xd4\x3f\xaf\xb2\x15\xfd\xea\xb3\x5f\x2f\x25\x46\x0b\x88\x4a\x4c\x1f\x73\xce\x78\xad\x1c\x5a\xe9\xe4\xce\x36\x4a\xf2\xee\xd5\x13\x65\x5e\x7a\xff\xaf\xb1\x8e\x5a\xbd\x2d\x3b\x63\x0b\xfa\x57\x4d\x57\xda\xb0\x4e\xea\x6f\xfe\x8c\xa4\x1c\x06\xd1\x13\x52\x89\xd3\x92\x60\xb7\x7b\x51\x77\x6b\x45\xf2\xd1\xd6\xa5\xd8\x50\xef\xd4\x4f\x93\xce\x10\x91\x9c\xd1\xfb\x42\x73\xd7\x16\x9a\xee\x06\x69\xb4\xab\x90\xfd\x1c\xfc\xd1\x0b\x39\xf2\xc0\xfa\xf2\xb9\x03\x02\xaf\x18\x75\x2c\xf3\x42\x58\x4f\x34\x5f\xe5\x4e\xcd\xda\x2f\x3d\xa3\x57\x67\xd0\xf8\xc8\xc0\xcc\x71\x7d\xb4\x61\x55\x4b\x48\xdf\xff\xcc\x06\x9a\x5a\xb5\xd9\x5e\x53\x7f\x3d\x4f\x54\x5d\x57\xce\x7f\xe7\x0e\x6d\x4b\xfc\x38\x30\x0a\x3e\x96\x49\x16\x41\x9b\x8f\x5e\xe7\x45\x9f\x1f\xd0\x5b\x6c\x2c\xec\xd5\x8e\x3c\x3e\x7a\xab\x5f\x6b\xb5\x13\x7a\xf6\xc4\x7f\xbc\xd8\x93\xef\xcc\xa1\x20\x50\xef\x11\xbe\x01\x32\x9e\x65\xce\x0f\xc9\xbb\xef\x91\x90\x7a\x81\x19\x58\xf2\xaa\xa3\x36\xf9\x3c\x24\xa4\x4e\x3c\xb9\x63\xae\x7b\xd7\xee\x53\xd0\x51\x65\x1f\x8f\x98\x0d\x3d\xdf\x1c\x1a\x5d\xcd\xc7\x9d\x0f\x39\x09\x2e\x83\xec\xb4\x41\xae\xd3\x0f\x48\xf5\x7d\x5d\xbf\x2b\xa6\xd3\x8e\xe8\xad\x2c\xcb\x1c\x1f\xdf\x91\x75\xde\x03\x45\x05\xb7\xf7\x6e\x19\x97\xf1\xd7\xf8\xbf\xb4\x3b\xbf\x05\x00\x00\xff\xff\x9f\x48\x9e\x01\x24\x5e\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x6f\x6c\x5b\x57\x15\xff\x15\xe2\x98\x38\x69\x1c\x36\xfe\x8e\xb2\xde\xc1\x5a\x3a\x5a\xb7\xb6\x93\xb2\xe1\x36\x2d\xa5\x7f\xd6\x76\xcb\x56\x5a\x33\xf8\x84\x78\xb5\xaf\x63\x53\xfb\x3d\xf3\xde\x73\xbb\x48\x43\x4a\x5c\xda\x6e\x6a\x58\x10\x30\x75\x12\x83\x09\x18\xe2\x0b\x12\x6e\x52\xaf\x6e\x12\x3b\x5f\x10\x0c\x69\x1f\xce\x45\xa0\x31\x24\x10\x1a\x08\xc1\x3e\x20\x04\x12\xda\x24\x24\xd0\x79\xf7\x36\x76\x12\x07\x85\x0f\xad\xd4\x9c\x73\xcf\x39\xf7\x9c\xdf\x39\xe7\x77\x9f\xfc\xfb\xbb\x7a\x9e\x03\x80\x07\x01\x6c\x01\xf0\xab\x4d\xc0\x09\x00\xd3\xef\x44\xf0\xaf\xd6\x03\xbc\x07\xc0\xf5\x1e\xe0\x7d\x00\x7e\xd6\x03\x84\x00\xfc\xae\x07\xe8\x05\xf0\xc7\x1e\x60\x07\x80\x37\x7b\x80\x0f\x01\xb8\x3b\xa4\xcf\x7b\x42\xfa\xfc\xf9\x10\x70\x2f\x00\xdf\xc8\x6f\x86\x74\xbe\xd9\x10\xb0\x19\xc0\xcf\x43\x00\x97\xfa\x75\x08\x78\x17\x80\xd7\x43\x40\x3f\x80\x37\x42\x40\x18\xc0\xdf\x8c\xfd\x5f\x21\x60\x1b\x80\xff\x18\xb9\xad\x17\xb8\x0f\xc0\x23\xbd\xda\x5f\xea\x05\xee\x07\x70\xbe\x57\xd7\xfd\x76\x2f\xb0\x15\xc0\x4d\x13\xf7\x5a\xaf\xce\xf7\x76\x2f\xd0\x03\x60\x53\x58\xcb\xde\x30\x30\x08\xa0\x3f\xac\xf1\xdc\x13\x06\xfa\x00\xec\x08\xeb\x3e\x93\x61\x8d\xef\xa1\xb0\xc6\xb5\xdf\xd8\x8f\x87\x81\xbb\x00\x3c\x16\xd6\xf5\xad\xb0\xce\xff\x25\xe3\x77\x4d\xfe\x27\x4d\xbe\xaf\x18\xff\x15\x23\x67\x8c\xbc\x66\xe4\x8b\x26\xff\x8f\x4c\xfc\x5c\x18\x88\x00\x68\x9a\xf3\x2b\x26\xee\x97\x46\xfe\xd6\xc8\x37\x8c\x7c\xd3\xd4\xfd\x47\x18\xc8\x6f\x02\xde\x0a\x03\xc3\x8c\xab\x0f\xf8\x04\x80\xe7\xfa\x80\x01\x00\xaf\x1a\xf9\x7a\x1f\xf0\x45\x00\x7f\xee\x03\x0e\x02\x78\x38\x02\xe4\x18\x5f\x04\x38\xca\x73\x8f\x00\xfb\x00\x8c\xf4\x03\x23\x00\xca\xfd\xc0\x47\x01\xfc\xb0\x5f\xf3\xe0\x17\xfd\xba\xef\x3f\xf4\xeb\xf9\xfd\xb5\x1f\x18\x02\xf0\x56\xbf\x9e\x4b\xdf\x80\xb6\x6f\x19\xd0\x75\x76\x0e\xe8\xfc\xd3\x03\xc0\x17\x00\xfc\x73\x00\xf8\x30\x80\xd3\x9b\x81\x93\x9c\x7f\x33\x70\x0a\xc0\xc2\x66\x20\x0d\x60\x60\x50\x9f\xad\x41\x8d\x77\x76\x50\xd7\x5b\x32\xf2\x95\x41\xe0\x63\x00\xd4\xa0\xe6\xc3\xdb\x83\x7a\xcf\x5b\xa3\xda\x7f\x28\xaa\xf9\x79\x32\xaa\xf9\xf6\x99\xa8\x9e\x67\x2e\xaa\xe3\xbf\x1c\xd5\x79\xaf\x45\x35\xce\x97\xa2\x1a\xff\x5c\x54\xe3\xff\xa9\x39\xbf\x16\x05\xde\x0f\xe0\x2f\x46\xb2\x91\xe5\x07\x8c\x8c\x0f\x69\x9e\x1d\x1f\xd2\xfd\xe4\x87\x74\xfe\xcb\xc6\xfe\xe3\x21\xbd\xd7\x57\x8d\xfc\x8d\x91\x7f\x32\xf2\xef\x43\xba\xce\xbf\x87\x80\x4d\xd0\x18\xf7\x42\x63\x63\xde\xf1\x0e\x77\x41\xef\x96\xfb\xe2\x5d\x24\x01\x7c\x04\xc0\x7b\xb9\x3e\x80\x18\xcf\x17\x9a\xf7\x51\x00\xbb\x19\x1f\x34\x8e\xce\x7f\xef\x58\x75\xe6\xbe\xb7\x1b\xfd\xae\x0e\x3b\xcf\x6b\x0f\x80\x04\x00\x01\xcd\xe3\x0f\x42\xe3\x7d\x00\x7a\xce\x43\x26\x76\x4b\xc7\xbd\xb0\x91\xcc\x0f\x9e\xfd\xbb\xcd\xf9\x1e\x23\x99\x87\xf7\x76\xc4\xf3\xfe\x78\x86\x77\x43\x7f\x33\xee\x37\x76\x9e\xe3\xc7\xa1\xdf\xe3\x56\xf3\x6d\x62\xee\xf1\x3b\xc0\xb6\xac\xf0\x64\xc6\xb1\xb3\x6d\xcd\xc3\xce\xb6\x75\x67\x87\xf9\x50\xd9\x2d\x14\x71\xa8\x32\x5e\xf1\x7c\x1c\xce\x5b\xf6\xb8\x14\xd9\x82\x57\x2e\x5a\x13\xa2\xe4\x64\xa5\xc8\x59\x85\xa2\xcc\x8a\x0b\x05\x3f\x2f\xa4\xeb\x3a\xae\xd8\xe6\x75\x0d\xf4\x2a\x99\x8c\xf4\xd6\xf8\xfc\x42\x69\x63\x49\x82\xc0\xe5\x24\x15\xd7\x95\xb6\x2f\xb2\x96\x2f\x85\x65\x67\xb5\xd7\xb1\x45\x49\xfa\xd2\x5d\xe9\xbe\xed\xca\xca\xf3\x85\x8c\x0c\x32\x1b\x77\x49\x5a\x5e\xc5\x95\x25\x69\xfb\x1e\x8e\x74\xa6\xc2\x11\x6b\x02\x47\x64\x46\x96\xce\x4a\x17\x47\xa4\x97\x71\x0b\x65\xbf\xe0\xd8\x38\xa2\xe1\xe0\x98\x3c\xeb\x56\x2c\x77\x02\x0f\x4b\x7f\x03\x13\xe9\x8c\x5a\xbf\x65\x8e\xca\x3b\xc5\x42\xd6\x9a\xf0\xba\x47\x1c\x37\x5e\x9c\xb0\x7d\xe9\x9e\xb7\x8a\x22\xe7\xb8\xc2\xca\xf9\xd2\x15\x65\x57\x7a\x9e\x38\x6b\x65\xce\x89\xb3\x15\xdf\x77\xec\x95\x41\x19\xd3\xb5\x6f\xb9\x85\x5c\x6e\xa5\xcf\x76\xfc\xff\xe9\x2f\x3b\x17\xa4\xbb\x4b\x23\xe7\x19\xf1\x60\x71\xd2\xb2\x83\x09\x9c\xac\x14\xf9\x8f\x2d\xf1\xa8\xe5\xf9\x3a\x56\x38\xb9\xdc\x8a\xa3\x8d\x31\xeb\x9c\xd4\x17\xc7\x2c\x37\x93\xc7\x98\x35\x81\xb1\xf6\x06\x30\xe6\xd8\x7e\x5e\xff\x2d\x4e\x08\xcf\xb7\xfc\x82\xe7\x17\x32\x1e\x1e\x73\xce\xeb\x3d\x3c\x9e\xf1\x1d\x96\xa7\xa5\x27\x7d\x9c\xb1\xce\x4b\x9c\x91\x65\x5f\x3b\xcf\xe4\x9d\x0b\x22\x9d\x30\x32\x69\xe4\xb0\x91\x23\x5a\x5a\x25\xa7\x62\xfb\x5a\x0f\xa0\x04\x5a\x00\x51\xab\xc1\xf2\xd3\x09\xb1\x47\x7c\x2e\x8f\x74\xd2\xc8\x61\x23\x47\xb4\x7c\xc2\x2a\x56\x24\x4e\xcb\xb2\xe3\xfa\xb1\x31\x6f\xbc\x90\x8d\x7d\xba\x32\xee\xc5\xd2\x4e\x8a\x59\xf6\xa9\x73\x85\xbc\x55\x72\x76\xbb\x95\xc8\xa9\xc7\xd3\xb1\xc3\xae\xb4\x98\x3a\x31\x66\x58\x4a\x24\xe3\x89\x4f\xc6\xe2\xc3\xb1\xe4\x83\x22\x39\x9c\xda\xbb\x77\x67\x7c\x38\x1e\x8f\xf0\xa8\x62\x69\xd7\xb2\xbd\xa2\xe5\x3b\x6e\x4a\x3c\x12\xe4\x10\x63\x15\xd7\x2a\x39\x59\x47\xec\x5f\x91\xf8\x40\xe4\x51\xcb\x1e\xaf\x58\xe3\x32\x96\x96\x56\x29\x25\x96\xcf\x29\x71\xba\xe2\x79\x05\xcb\x8e\x8c\x9d\x18\x3b\x1a\x7b\x42\xba\x5e\xc1\xb1\x53\x22\xb1\x3b\x1e\x39\xec\xd8\xbe\xb4\xfd\x58\x7a\xa2\x2c\x53\xc2\x97\x4f\xfa\x7b\xca\x45\xab\x60\xef\x13\x99\xbc\xe5\x7a\xd2\x1f\xfd\x6c\xfa\x58\xec\xa1\x76\x1c\xe3\xc9\x49\x37\x76\xd4\xce\x38\xd9\x82\x3d\x9e\x12\x91\x53\xc5\x8a\x6b\x15\x63\xc7\x1c\xb7\xe4\xa5\x84\x5d\x0e\x8e\xde\xe8\xf0\x3e\xa1\xd5\x51\x7b\x5b\x22\x3e\x3a\x9a\x10\xdb\xb7\x0b\x56\xe3\xf7\x8d\x26\x12\xe2\xa0\x88\x8b\x54\x70\x3e\x30\x9a\xbc\xed\xda\x3f\x3a\xc2\xea\x8e\x20\x6c\x7f\x22\x2e\x9e\x7a\x4a\x5f\x39\x30\x9a\x8c\x3f\x20\x0e\x8a\x84\x48\x89\xe4\x3e\xfe\x4a\xa9\x29\xaa\xd3\xbc\xba\x48\x4d\xba\x41\xb5\xd5\x16\x75\x75\xb5\x25\xf8\x86\xad\xba\xb4\xda\xa4\xae\xae\x31\x81\xbe\x4e\x4b\x6a\x92\xea\xb4\xa0\xa6\xf9\x30\x4b\x73\xea\xa2\x9a\x52\x55\xd0\xf3\x74\x93\x16\xa9\x4e\xcd\xe0\x7f\x83\xea\x22\x08\x7c\x99\x1a\xb4\x48\x2d\x9a\x15\xd4\x08\x0a\x35\x68\x9e\x6a\xea\x32\x35\xa8\x21\xe8\x26\xd5\x68\x96\xea\x6a\x52\x3d\x4d\x0d\x5a\xa0\x96\x9a\x52\xd3\x42\x4d\x09\x6a\x05\x96\xeb\x34\x4f\x2d\xba\xc5\x2f\x9b\x7e\xa0\xa6\x68\x89\xea\xea\x69\x6a\x52\x8b\xb3\x75\xd4\x53\x57\x3b\xaa\xa9\xab\x5d\x6a\x75\x03\xc8\x51\xaa\xca\xe5\x69\x96\x6a\x5c\xfe\xce\xc3\x5c\x5d\x73\x1d\xb0\xdf\xd7\x43\x57\xcf\x68\xa0\x37\xa8\xa6\xaa\x54\x13\x0c\x66\x36\xe8\x74\x51\xcd\x08\x6a\x52\x4d\xa8\x29\x75\x85\xea\xaa\xaa\xae\x04\x19\xea\x9d\x77\x6b\x1c\xb4\x7c\x37\x88\x0e\xb6\xa5\x26\x19\x2e\x6b\xdc\x56\x00\x7b\x55\xc1\x25\x6a\x31\x1c\xee\x9c\xd1\xd2\x02\x83\xfa\x56\x37\x14\x6c\xe6\xd6\xa6\xb5\xc2\x97\xae\xab\x49\x3e\x7e\x97\x96\xa8\xa1\xa6\xa8\xa6\xc7\xcd\x1b\x58\xd1\x27\xdf\x7d\x89\xea\x41\xaa\x9a\x66\xd4\x8b\xd4\xa2\x05\x75\x91\x3b\xba\x33\x24\x5a\x5b\xe0\x8e\x90\xa0\x4b\x99\xa5\xa0\xcb\x9b\x74\x83\x0d\xea\x0a\x93\x41\x5d\x12\xc1\xb1\x4e\xb7\xfe\xcf\xf4\x6b\x73\x69\x96\x34\x35\xd1\x57\x34\xd4\xb5\x99\x25\xce\x4f\x0b\x7c\xab\x49\x35\x7a\x39\xd8\x6c\x83\xe9\x32\xcf\xbc\xa5\x25\x9a\xdf\x60\x2a\xf5\x35\xee\x40\x4d\xd2\x9c\x3e\x56\x97\x89\x54\xa7\x39\x6a\xb1\xa1\xa6\x26\xa9\xa1\xbe\x4a\xb5\x8d\x61\x6b\xf2\xda\xdb\x69\x1a\xea\x52\x67\x12\xbd\x1e\xa6\xfd\x45\x5a\xe4\xc7\xbe\xb1\xa4\x8b\xd4\x52\xcf\x70\x6b\x01\xeb\x1b\xbb\x96\x19\xac\x37\x14\x70\x3a\x78\x29\x9c\xf0\x27\xd4\xe4\x54\x9a\xc6\xcf\xab\x67\x35\x39\x59\x69\x1a\x96\xea\xe1\x99\xed\xd5\x79\x47\x55\x9a\xa7\x05\xf5\x6c\x7b\xe5\x5d\xe3\x66\xd7\x46\x2d\xbf\xaa\xa5\xe0\x59\xf2\x07\x63\x96\x5a\x74\xc3\x3c\xcf\x1a\xe8\x85\x00\x4a\x35\x50\xe8\x96\x66\xd7\x8a\xc7\xc9\xa0\x5e\xa0\xba\x9a\x52\x33\xea\x72\x5b\xbd\xc2\xbb\x55\x33\x22\xc8\xa4\x17\x1c\x34\xcf\x97\x41\xdf\xa1\x96\x9a\x69\xbf\xd5\x79\x55\x6d\x1f\xbf\xc7\x4a\x80\xbe\xa1\xaa\xda\xd0\x52\x97\x02\xd6\x35\xdb\x26\x7e\xf6\x1d\x97\x96\x61\xa9\xab\xc1\xf4\xaa\x6a\x9a\x7f\x70\x74\xb7\x27\xd7\xb1\x0f\xaf\x63\x1f\xe9\x6e\x5f\xa6\xc1\xc5\xee\x7e\xb3\xd3\xf5\xbc\x2b\x58\xb1\x5e\x0f\x1d\xdf\xba\xe0\x97\x0f\x7d\x83\xbf\xb6\xfa\xd7\xcf\x6d\x7d\xb8\x43\x1f\x69\xeb\x74\x2d\x58\x41\x7b\xdb\xff\x0d\x00\x00\xff\xff\x3d\xc5\x77\x82\x0d\x11\x00\x00"

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
