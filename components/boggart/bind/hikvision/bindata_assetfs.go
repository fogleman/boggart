// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package hikvision

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x59\x5f\x6f\xe3\x36\x12\x7f\xdf\x4f\x31\xd5\xe6\x4e\x76\x2f\x92\x92\xb4\x45\x0f\x59\x2b\x87\x5e\xf7\x7a\xb7\xc0\xde\x62\xb1\xdd\x3e\x15\x45\x40\x8b\x23\x8b\x1b\x9a\x54\x48\xca\x89\x61\xf8\xbb\x1f\x48\x4a\x8a\xfc\x5f\x71\xbc\x40\x71\x7e\x48\x28\x72\x66\x38\xf3\x9b\xe1\x70\x34\x5a\x2c\x80\x62\xce\x04\x42\x90\x49\x61\x50\x98\x00\x96\xcb\x57\x23\xca\x66\x90\x71\xa2\x75\x1a\x28\xf9\x10\xdc\xbc\x02\x00\xe8\xce\x66\x92\x47\x53\x1a\x5d\x5e\x81\x1d\xe9\x69\x33\x7a\xd4\xd1\xe5\x55\x4d\xbf\xce\xf3\x78\x5b\x12\x81\xbc\xb3\xba\x49\xd1\x68\xb1\x4a\xd3\xd2\x29\xc9\x31\x0d\x0c\x19\x6f\x93\xd4\x52\x56\xbc\x11\x28\xc8\x0c\x04\x99\x45\x86\x8c\x35\x8c\x89\xba\xb5\x83\xe0\x49\x0c\x67\x7a\xdb\x5e\xad\x24\xce\x6a\xda\x52\xa1\x46\x61\x88\x61\x52\x04\x8b\x05\xb0\x1c\xf0\x1e\x62\x92\xd9\x09\x08\x2c\x6a\xcd\x9e\x76\x6e\x86\x96\x08\x05\x85\xe5\xf2\x66\x44\xa0\x50\x98\xa7\x3b\xf8\xbe\x90\x19\xd1\x99\x62\xa5\xb9\x9e\x49\x46\x07\x17\xc3\x37\x96\x97\x6b\x84\xe5\x72\xb1\x80\xf8\x13\xde\x57\xa8\x4d\xfc\xdb\xa7\xf7\xf1\x47\x62\x0a\x3f\xed\x85\x07\x37\x56\xe8\xe5\xdf\x05\x04\x1f\x15\xce\x18\x3e\x04\x10\xc3\x72\x39\x4a\xc8\xcd\x28\xe1\xec\x04\xb6\xb1\x29\x99\xe0\xb1\x06\xb6\xcc\xc7\x58\xf9\x0f\x2f\x25\x75\x42\xb6\x99\xfc\xce\x4b\x3f\x6c\xf0\x28\xa9\xf8\x8e\x95\x4e\xf8\x19\x32\x8e\x76\x07\xe0\xbe\x40\xec\x4a\xb0\x33\x90\x13\x8a\xe0\x81\x02\x26\xf6\x48\xdb\x03\xd9\x4e\x9e\x9a\x4f\x11\x31\x41\x38\x63\xe7\x70\x96\x15\x44\x08\xe4\x70\x9d\x42\x5c\x8f\xf5\x21\x09\xeb\xd6\xd7\x47\xfa\x87\xe6\x44\xff\xb0\xf5\x40\xef\x95\xc6\xa6\x13\xd0\x2a\x73\x81\x70\xb6\xcf\xa3\xa5\x8f\xd5\xbf\xd6\xba\xa6\x97\x17\x96\xa3\x7e\x8a\xdf\xbd\xb5\x6e\x86\x07\x46\x4d\x91\x06\x97\x17\x17\x7f\x09\xa0\x40\x36\x29\x4c\xf3\xc4\xa8\x0b\x5b\x1f\xef\xc9\x61\xed\x46\x09\x65\xb3\x1e\x64\x9b\x70\xfc\xd8\xc0\xf1\xe3\xf3\xe1\xc8\xa5\x9a\xd6\x91\x62\x87\x01\x50\x62\x48\x64\xe4\x64\x62\xa7\x66\x84\x33\x4a\x8c\x54\xde\x1c\xe7\xf6\x48\xa3\x31\x4c\x4c\x74\x00\x35\x52\x07\xa1\xac\xc3\x65\x8a\xa6\x90\x16\x15\xb9\x37\xa3\xad\xff\x7c\xf8\x3d\x21\xaf\x7e\xae\xcc\x2f\x8c\x1b\x54\x7d\xc2\x67\x1b\x6e\x24\xcb\xa4\xa2\x36\x95\x74\xcc\xea\x4c\xae\xe6\x5e\x20\x8a\x91\x68\x5a\x71\xc3\x34\x72\xcc\xec\xbc\x5d\x56\x15\x3e\xc3\x8a\x75\x1d\xf6\xdd\x0e\x7b\x65\x90\x15\x09\x51\x81\x84\x32\x31\xe9\x28\xbd\xd5\xa8\x48\x3f\x30\x93\x15\x6b\xfe\xcd\x24\xe7\xa4\xd4\x58\x4f\x97\x44\xa1\x30\x69\xf0\x7a\x03\x11\x9f\x3b\xeb\xf9\x86\xab\x95\xe9\x00\xc2\xc7\x92\x08\x8a\xb4\x46\xc6\x4f\xda\x54\xa5\x24\xd7\x8d\x3e\xeb\xac\xcf\xb7\xdf\x61\x50\x7c\xbf\x0a\x82\x61\x86\x63\x27\xe9\xbe\x25\xf3\xe4\x83\x3d\x8c\xd0\xa8\x78\xe6\xf2\x6f\xf1\xfd\x11\x80\x27\xe4\x08\x26\xeb\xe9\x27\x37\x6c\x20\xb6\xa2\x7d\xb3\x0a\xed\x80\x75\x83\xb0\x4e\xdf\x0e\x4f\x4e\xc6\xc8\x39\xd2\xf1\x7c\xa7\x87\x8f\x84\x74\x3d\x34\xa3\xb1\xa4\xf3\x23\x85\x39\x81\x4e\x55\xc8\xa5\x4a\x03\xa6\xa2\xac\x32\x51\xee\x0e\x6d\x64\xe6\xe5\x73\x0f\xce\xfa\xef\x80\xa3\x61\xa4\x4b\x22\xda\xea\x10\xef\x2b\xa6\x90\x06\x37\xdf\x8e\x12\xbb\xf0\x02\xa3\x12\x67\xd5\x0b\x04\xf8\x0c\xe2\x43\x63\x13\x95\x46\x65\x9b\x8a\x9b\xb3\x03\x9e\xe5\x2a\x80\xc6\x90\xae\x49\x2f\x82\x71\x24\x4b\x77\x9f\xcf\x08\xaf\x30\x0d\x28\x99\x3f\x95\x47\xdb\xf2\x6d\xfc\x79\x5e\x22\x38\x3a\x8b\xb2\x57\xcc\xea\xd3\x8c\x3a\x05\x57\xd7\x45\xcd\xf1\xf3\xdb\x9d\x54\x67\x61\x9d\xdf\x4b\x6b\x4f\xd9\x5f\xef\x0f\x9e\xfe\xab\x69\x4e\x2a\x23\x7b\x29\xee\x08\xfb\xeb\xfd\x93\x23\xff\x6a\x6a\xeb\xac\x40\x5a\x71\xec\xa5\x7a\x4b\xdc\x5f\xfd\x5f\x5b\x96\xd3\x98\x30\x4a\xfc\x6e\x47\xa6\xc5\x7e\x45\xd9\x09\xd8\x8e\x61\xf9\x26\x8a\xfe\x2c\xf5\x47\x7b\x77\xd1\x83\x95\x88\x34\x05\xaa\xd3\x16\x22\xb5\xc8\xb5\x3a\x24\x27\x5c\x1f\x2c\x44\x3c\xeb\x69\xeb\x90\x9f\xbd\x70\x36\xe6\x08\xff\x56\xb2\x2a\xe1\x9d\xc1\xa9\x86\xd7\x57\x7f\x82\x12\xa4\xc6\x6a\x7f\x05\x72\x44\xf9\xf1\x22\x20\x4f\x5e\x7d\x94\x37\x23\x6d\x94\x14\x93\x15\x67\x58\x37\xc0\x95\x8b\xb1\x51\x52\xaf\xbf\x20\xb3\x94\xc7\x33\xff\x24\xd8\x14\x4a\x0b\xaa\xa9\x14\x64\x9c\x65\x05\x82\xc2\x52\x61\x81\x82\xa2\x62\xe6\x1c\xd0\xd2\x20\xab\xf4\x54\x52\x28\xd8\xa4\x00\xce\x72\xfb\xca\x9e\x55\x9a\x4c\x2b\x0d\x06\x95\x9a\x83\x62\x59\x41\x14\xd5\x52\x00\xa1\xa0\xef\x2b\x46\x63\xf8\x0e\x1e\x24\xcf\x61\x2a\xa5\x00\x99\xe7\x2c\x63\x04\x48\x65\xf0\x1c\x84\x14\x90\x55\xa5\x7d\xaf\x23\x06\xf4\x1d\x31\x38\x96\x44\x51\xa0\x92\x4b\x05\x63\x55\x89\xac\x88\xe1\x17\x29\x29\x18\x55\x65\x77\x70\x5f\x31\x21\x09\x08\xd4\x19\xab\x84\x01\x4e\xc6\x52\x55\xad\x6a\x31\xfc\xd3\xf1\xac\xec\x69\x70\x5a\x4a\x75\xfe\xff\x97\x74\xa3\xa8\x3f\xfd\x33\xc5\xb7\xb7\x61\xbf\x57\xf7\xc4\x56\x89\x27\xe9\x2b\x1c\xde\xf8\xa9\x0f\xb6\x57\x50\xf7\x1c\xbb\x1a\x76\xe2\xb2\x5f\x59\x71\x1e\x29\x57\x4f\x3d\xaf\xc1\xc1\x44\x59\x19\x2f\xa4\x6f\x43\xc3\xbf\x76\xd4\xfc\x63\x23\x9e\xc5\xed\x24\x8c\x2b\x63\x6c\x04\xcf\x4b\x4c\x03\xff\x10\x74\x04\x82\x15\xca\x32\x59\x0f\x44\x2e\xdd\x40\x4f\x03\x90\xc2\x1e\xe4\xbb\xb6\xef\xf3\x09\x73\x85\xba\x18\x0c\x9f\xdb\x27\x60\x2d\x88\x44\x43\x4e\x22\x3d\x17\x59\x00\xee\x76\xf1\x1d\x4c\x57\x2a\xd5\xe2\x7d\x73\x31\xb8\x19\x25\x7b\x5a\xa9\x1b\x5b\x24\xde\xb2\x67\x70\x74\x3a\xa8\x7d\xfa\x66\x54\x3e\x08\x2e\x09\x4d\x2f\xfb\xa1\xf7\x42\x88\x9a\xed\x36\x61\x7a\xdb\xae\x1c\x85\x53\xcf\xab\xb7\xf7\xab\x61\x8f\xf3\xd8\x87\xa4\xdb\xc7\xec\xe1\x8e\x93\xb4\x2a\x0f\xe7\x89\x3d\x9a\xef\x58\xda\x32\xbd\x36\xd5\x79\xac\x87\xf5\xbf\x27\x75\x5e\x75\x3e\x12\x7d\xd1\x6d\x63\xba\x4f\xe3\x7a\xb1\x00\x6d\x88\x61\xd9\x7f\x3e\xff\xf7\x3d\x0c\xfc\xf8\xb7\x4f\xef\x21\x48\x28\xd1\x85\xbb\x1a\x13\xa2\x35\x1a\x9d\xcc\x50\x50\xa9\x74\xd2\xb6\x45\x93\x2f\x9d\x87\x78\xca\x44\x6c\x77\x77\x75\xe7\xb0\xa3\x44\xa3\xa5\x33\xc1\x7f\x61\xa8\xd3\x0b\x29\x4b\xce\x32\xf7\x59\x23\x79\xfa\xfc\xd0\x39\x0b\x0f\x4c\x50\xf9\x10\xaf\x26\x14\x48\x21\xaf\x84\x33\x69\x30\x84\xc5\x0a\x78\x67\x83\xf0\x75\x4d\x1d\x0e\x63\x62\x8c\x1a\x84\x5a\x65\xe1\x39\x84\x3d\x0f\xee\xad\x49\xc3\xbf\x0d\x04\x3e\xc0\x5b\x62\x70\x30\x1c\xc6\x13\x34\x9f\xd9\xd4\x0e\xdf\xb4\x7b\x2d\xdf\xbc\x7a\xb5\x7d\x4f\x85\x84\xce\x07\x8d\x82\xb0\xa1\x61\x6d\x92\x46\xf3\x4e\x18\x54\x33\xc2\x07\x5b\xad\x3c\xb7\xd0\x35\x73\xb7\xca\x4f\xde\xb2\x9a\xc7\xbe\x51\x7e\x0b\x97\x17\x17\x17\x5d\xa5\x86\x2b\x5a\x51\x99\x55\x53\x14\xe6\xb0\x4e\xd6\x82\xd5\xee\x36\xb8\xbb\xe7\x77\x46\xff\x38\x87\xf5\x25\xff\x46\xb9\x39\x6f\xf0\xd1\x10\x85\xc4\x72\x85\x43\xf7\x95\x63\x82\x83\x9d\xae\xb2\xbf\x19\x51\x3b\x0f\x13\x42\x0a\x67\x03\x53\x30\x3d\xdc\x5d\x48\xd9\x32\x0e\x52\x58\x2c\x77\x93\xdc\x41\x0a\x18\x97\x4a\x96\x83\x90\xd1\x70\xb8\x41\xd8\x01\xad\xf9\xb1\x1c\x06\x0d\x8f\x0d\xd5\x70\x08\x69\x9a\x42\x98\x15\x98\xdd\x8d\xe5\x63\xb8\xcd\x98\x46\x9f\xdf\xef\xfe\xe8\x6c\xe9\x58\x90\x86\x43\xf8\x26\x85\x30\x7c\xb3\xc1\xb6\xf4\xb5\xc5\x61\x79\x36\x54\x36\xb5\x5f\x6e\x6a\x7f\x16\x93\x2f\xe4\x71\xb0\x5d\xa2\x35\xe7\x7a\x8b\xc7\xc3\x61\xad\xb0\xff\x4e\x11\xee\x00\xbd\x52\x7c\x2f\xb7\x3f\x4d\xbb\xb8\xad\x39\xd7\xee\xef\xf6\x75\x5d\x65\x19\x6a\x7d\x0d\x4f\xb1\xaa\x76\x41\x0d\xb5\x9f\x54\xac\x50\x57\xdc\x78\x0f\xe5\x84\x71\x87\xf6\x6e\x26\xfb\xb3\xe7\xfb\xe3\x07\x69\x58\x3e\xdf\x81\xd3\x0a\x66\xf6\x4e\xbd\x86\xf0\x5f\x4a\x49\x15\x1e\x2e\xeb\xed\x49\xb8\x06\x15\x4f\x51\x6b\x32\xc1\x1e\x0c\xce\x29\x21\xf6\x94\xaf\xcd\x9c\x33\x31\xb9\x86\x70\x2c\xa5\xd1\x46\x91\xf2\xbb\x70\x2f\xd7\x72\xb8\x19\x78\xed\x9a\x0f\x40\x8f\x65\xad\x32\x7c\x63\xc1\xac\x84\xbf\x56\xbe\x16\x9e\xbf\x7a\x77\x7f\x45\x44\x75\xef\x1d\x4e\x8c\xe9\xd6\x95\xcd\xd9\x75\x19\xcb\xb5\x74\x0e\xbe\x87\xe7\xee\xc5\xce\x95\xff\xbf\x00\x00\x00\xff\xff\xf4\xed\x46\x9c\x16\x22\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x4e\x4d\x4f\x13\x6d\x14\x3d\x6d\xe1\xe5\xb5\x46\x62\x8c\x31\xd1\x98\x78\x5d\x48\x50\x33\x30\xd3\x8a\x91\x81\x01\x95\x8f\x84\x48\x95\x60\x75\xff\xa4\x7d\x68\x27\xb6\x33\xcd\x7c\x80\x24\x2c\x0a\x1b\x36\x86\xb8\xd0\xc4\x8d\x12\xdd\xba\x29\x84\x2a\x41\xa8\x7f\xe1\x3e\x7f\x40\xff\x81\xd1\x8d\x0b\x57\x66\x3a\x83\xc6\x9b\xcc\x9c\x7b\xe6\x9e\x73\xe6\x7c\x39\xd3\xf3\x12\x00\xb2\x00\x2e\x02\xa8\x01\x38\x05\xe0\x10\xf1\x7c\x07\xd0\x03\xe0\x07\x80\x0c\x80\x9f\x00\x4e\x03\xf8\x05\xe0\x7f\x00\x27\x53\x40\x2f\x80\xb3\x09\x5e\x48\x01\x7d\x00\x28\xc1\xc1\x54\xac\xd3\x53\x40\x35\x05\x8c\x26\xfc\x5d\x3a\xc6\xf7\x69\xe0\x1a\x80\xbd\x34\xd0\x0f\xe0\x6b\x1a\x38\x07\xe0\x5b\x72\xef\xcd\xc4\xff\xeb\x4f\xf0\x52\x06\x38\x1f\xe5\x65\x80\x54\xd2\xbb\x2f\xe9\x9a\x4e\xba\x46\xbe\x13\x88\xfb\x20\xe9\x1d\xcd\x7f\xd1\xeb\x4e\x18\xb8\x98\x16\xab\xd1\x33\x7c\xdf\xae\x54\x03\xf2\x57\xec\xa0\x54\xc5\xb4\xbb\xe2\xd4\x5c\x51\xc6\x5c\x5d\x54\x24\xba\x37\x2c\x78\x72\xd9\x96\x2b\x58\x94\x4b\x9e\xf4\xab\x78\x58\xaa\xca\x72\x58\x93\x58\x94\x0d\xd7\x0b\xb4\x82\x5f\xb1\xcb\xda\xdd\xb0\xe2\x6b\x45\xd7\xa4\xb2\x5c\xbe\xfd\xc4\xae\x8a\xba\x3b\xe4\x85\xd9\x85\x07\x45\x6d\xca\x93\x22\xb0\x5d\x47\x9b\x16\x81\x34\x29\xa7\x1b\xa3\x9a\x9e\xd7\x8c\x9b\x94\xcb\x9b\x23\x23\xd7\xf5\xbc\xae\x67\xe7\x85\x1f\x68\x45\x4f\x38\x7e\x4d\x04\xae\x67\xd2\xbd\x6e\x06\x15\x42\x4f\xd4\xdd\xb2\x4b\xe3\xff\x04\x4f\x64\xe7\x85\x53\x09\x45\x45\x6a\x45\x29\xea\x26\xfd\xe1\x26\x2d\x86\xbe\x6f\x0b\x27\x5b\x98\x2b\xcc\x68\x8f\xa5\xe7\xdb\xae\x63\x92\x31\xa4\x67\xa7\x5c\x27\x90\x4e\xa0\x15\x57\x1b\xd2\xa4\x40\x3e\x0d\x86\x1b\x35\x61\x3b\x63\x54\xaa\x0a\xcf\x97\x81\xf5\xa8\x38\xab\xdd\xfa\xab\x8b\xfa\x2c\x49\x4f\x9b\x71\x4a\x6e\xd9\x76\x2a\x26\x65\x17\x6a\xa1\x27\x6a\xda\xac\xeb\xd5\x7d\x93\x9c\x46\x97\xfa\x56\x7e\x8c\xe2\xd5\x72\xae\x18\xba\x65\x19\x34\x30\x40\xd1\xaa\x5f\xb6\x0c\x83\x26\x49\x27\xb3\xcb\x27\xac\xdc\xf1\x69\xdc\xba\x11\xad\x83\x5d\xd9\xb8\xa1\xd3\xda\x5a\x6c\x99\xb0\x72\xfa\x55\x9a\x24\x83\x4c\xca\x8d\x81\x9f\xf3\xae\xda\xe0\x0e\xf8\x05\xb7\xf9\x48\x3d\x03\x6f\x73\x5b\x35\xb9\xcd\x07\xfc\x49\x6d\xa9\xcd\xe8\x33\xef\x73\x9b\x78\x2f\x56\x0c\xf3\x11\x77\xd4\x66\x24\x7d\xcb\x07\xdc\x52\x9b\xdc\x52\x1b\x11\x7d\xc5\x1f\xb9\xc3\x3b\xaa\xc9\x2d\xfe\x70\xec\x03\xbf\x3e\x96\x6f\xab\x26\x77\xd4\x3a\x1f\x72\x47\x6d\xa8\x26\xf8\x0d\xef\x44\x61\xbc\xcb\xfb\x71\xc2\x36\x77\x28\xb2\xab\x75\xfe\xcc\xfb\x6a\x9d\x5b\x51\x88\xda\xc2\xef\x00\x00\x00\xff\xff\x50\xcb\xd1\x3c\x4a\x03\x00\x00"

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
