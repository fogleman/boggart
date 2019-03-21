// Code generated by go-bindata.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
// DO NOT EDIT!

package homie

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xff\x6f\xdb\xb8\x15\xff\xbd\x7f\xc5\x1b\xeb\xce\xf2\x16\x49\x49\x7a\x18\x36\xc7\x76\x71\xb8\x0d\x38\x0c\x5d\x7b\xb8\x76\xb7\x1f\xba\x2e\xa0\xa5\x67\x8b\x0e\x4d\xea\x48\xca\x49\x66\xf8\x7f\x1f\x28\x4a\xb2\x6c\x4b\xb2\x93\xa5\xbd\x16\xa8\x80\x04\x14\xf9\xf8\xbe\x7e\xde\xe3\xa3\xbc\x5e\x43\x8c\x33\x26\x10\x48\x24\x85\x41\x61\x08\x6c\x36\xcf\x46\x31\x5b\x41\xc4\xa9\xd6\x63\x92\xd2\x39\xfa\x86\x19\x8e\x64\xf2\x0c\x00\xa0\xbe\x98\xcf\x5f\x73\x9c\x99\x62\x31\x27\x48\x5e\x4e\xd6\x6b\x60\x17\x7f\x16\x40\xfe\x8a\x2b\x16\x21\xbc\xd0\x04\x7a\x20\x18\xaf\xfe\x02\x41\x97\x08\x9b\xcd\x28\x4c\x5e\x16\x8c\xc3\x98\xad\x26\xcf\x5a\x84\x28\x36\x4f\x76\xa4\xd4\x28\x66\x52\x2d\xfd\xb9\x92\x59\x0a\x69\xc6\xb9\xbf\x4f\xbb\x4f\xcf\x44\x9a\x19\xb7\x61\x8f\x2a\xa7\xe4\x74\x8a\xfc\x70\x3e\x5f\x9b\x66\xc6\x48\x01\xe6\x3e\xc5\x31\xd1\xd9\x74\xc9\x0c\x29\xd9\x4e\x8d\x80\xa9\x11\x7e\x4c\xc5\x1c\x55\x3e\xd4\x4b\xd2\xc8\xa7\x7c\x62\x6a\xa8\x6f\xe4\x7c\xce\x71\x4c\x96\x32\xa6\x9c\x14\x73\x54\xcd\xd1\x8c\xc9\x73\x37\x79\x9c\x49\x4e\xe7\x02\x35\x26\x95\xfb\x7f\x90\x62\xc6\xd4\x12\x14\x6a\x34\x10\xe7\xc1\x20\x10\xc0\x66\x73\x32\xcb\xa9\x8c\xef\x6b\x1c\xff\xc5\x38\x87\x29\x16\x1c\x23\xcb\x7f\x5e\x30\x0e\xe0\x7b\x85\x70\x2f\x33\xd0\x99\xc2\x57\x0f\x94\x13\x51\xce\xa7\x34\xba\x19\x93\x9c\xb5\x37\xb8\x6a\xde\xba\xc5\xd6\xcf\x07\x46\x8d\x42\x17\x9f\x86\xa0\x86\x0d\x51\x2d\x10\xd7\xf0\x5a\x0c\x4b\x48\xd6\xc1\x13\x71\xa4\x6a\xc6\xee\xc8\xa4\x69\x55\xc9\xdb\x86\x3c\x89\x24\xf7\x97\xb1\x7f\x71\x09\x76\xa4\x97\xe5\xe8\x4e\xfb\x17\x97\x2d\xa0\xbe\xbb\x4e\xa9\x40\xde\x01\xe3\xbb\xeb\x7a\x5e\xee\x50\x25\x97\x07\x29\x48\x8d\x51\x6c\x9a\x19\xd4\xa5\xaf\x92\xcb\x86\x9d\x19\x2f\xd9\x0b\xba\x02\x41\x57\x53\xaa\x5c\x42\x41\xae\xcf\xb5\x91\x92\x4f\xe5\x5d\x83\xd4\x7c\x3f\x67\x93\x11\xad\xd9\xcd\x69\xaa\xd1\xe7\x4c\xdc\x90\xc9\x88\x55\x19\x4b\x61\x46\xfd\x28\xc1\x95\x92\xc2\xb7\x69\x38\x0a\xd9\x64\x14\xd2\xc9\x28\xe4\xac\x29\x7c\x59\x43\x46\x76\x86\xa5\x23\xd0\x87\x8e\x2c\xab\x5f\x83\x0c\x43\xa7\x1c\xab\x5a\x94\xbf\xe4\xff\x7d\x6d\x14\x4b\x31\xce\x31\xec\xe6\x63\xe3\x2b\xd4\xa9\x14\x9a\xad\x10\x84\xbc\x55\x34\x25\xa0\xcd\xbd\xcd\xc9\x5b\x16\x9b\x64\x78\x71\x7e\xfe\xa2\xcd\x75\x26\x41\x1a\xb7\xad\xa9\xe6\x85\x62\x63\xa9\xdf\x32\xf6\x2d\xae\x2e\xc9\x36\xfc\xdf\x97\x71\x2f\xc3\x6e\x92\x4e\x56\xdb\x9d\xbf\x50\x9e\x1d\xdd\x35\x0a\xdb\x54\xb3\x7b\x3a\x0c\xb2\x75\xa5\x5d\x91\xf5\x1a\x94\x2d\xa3\xd0\xb3\xe7\xc4\x19\xf4\x56\x56\x19\x18\x8e\x21\x70\x09\xaf\xaf\xb7\x80\xb6\x67\x56\xbb\x45\x1d\x9e\x73\x04\xb1\x35\xb9\x57\x9d\x47\xa6\x45\xe5\xfa\x86\x4e\x82\x42\x7f\x36\x03\xfc\xb5\xe0\x4b\xd8\x32\xe5\xb8\x44\x61\xa8\x61\x52\x04\xae\x6e\x92\x2e\xc5\x77\x44\xa6\x0a\x4b\x1c\xa5\x34\x8e\x99\x98\x0f\xcf\xc9\x64\x14\xc9\xb8\x82\xe6\x42\x4b\x41\x26\x23\x1d\x29\x96\x9a\xe2\x80\xa2\x69\xca\x59\x94\x8b\x0c\x17\x74\x45\xdd\x22\x99\xc4\x32\xca\xac\x32\xc1\xad\x62\x06\xbd\xbf\xbf\x7b\xfb\x26\xb0\x70\x16\x73\x36\xbb\x77\xaf\x29\x55\x1a\xbd\x3c\x0c\xb7\xa5\xf7\x37\x9b\xc1\x19\x88\x8c\xf3\x33\xb8\x1c\x0c\x46\xa1\x63\x37\x19\x85\x56\x8f\xc9\x28\x4c\x15\x9e\xe4\x19\xe4\x1a\x4f\x35\xdd\x86\xa6\x94\x7e\x12\x6f\x11\x1f\xa3\xec\x0e\x71\x3b\xa0\x8f\x4b\x18\x85\x2d\xb0\x1e\x85\x79\x85\x78\xe4\xf1\xe3\xc0\x54\x80\x66\xbf\x41\xfb\x02\x4f\x9c\x1f\x0a\x74\x7f\x3b\x66\x4e\x3e\x66\x8a\x10\xa3\x52\x52\xb5\xa2\xab\xc6\x8b\x72\x54\x06\xf2\xff\x65\xbf\x59\xbc\x30\xbd\x64\x5a\x5b\xb4\xb5\xf8\x0e\x0e\xfa\x58\xf7\x42\xb6\x36\x4a\x8d\x45\x1f\x5a\xb0\x2b\x04\x12\xa0\x8a\x51\x3f\x61\x71\x8c\x62\x4c\x8c\xca\x90\x4c\x7e\x6f\xd8\x12\xf5\x55\x7b\xef\x55\xd8\xe7\x8c\x0b\xfe\xd6\x69\xe2\xa1\x07\x61\x27\xeb\x0e\x83\x63\x9b\x7f\x50\xd2\x96\x46\x3b\x24\x40\x23\x5b\xef\xc6\xe4\xf9\xab\x62\x54\x16\xdb\x25\x9a\x44\xc6\x63\x92\x4a\x6d\x08\xb0\xd8\x02\xc7\xae\x68\xb2\xdb\x87\xaf\x28\x67\x31\x35\x52\xb5\x1e\xd4\x9f\xab\x25\x80\x23\x6d\x01\x9c\x76\xc0\xed\xb7\x07\x17\xb5\xf6\xe0\x6d\x6a\x7d\x74\x4a\x6f\x00\x8f\xea\x0f\xe0\x68\x49\xed\xec\x13\xe0\x41\xbd\x02\x3b\x83\x9e\xcc\x0d\xca\x3b\x85\x6d\xc1\x6c\x67\x7d\x62\x77\xe0\xb8\x06\x6f\x9e\xac\x49\xd8\x36\x08\x05\xeb\xf7\xf7\x29\x02\x99\x4a\xc9\x4f\x6a\x0b\x46\xf9\x25\xb6\x48\xe0\x28\xc1\xe8\xc6\x56\xcb\x6d\x33\xe0\xeb\x5b\x66\xa2\x84\x80\xed\x3f\xf2\xfb\xdb\x9e\x0d\x2e\x03\x9a\xe6\x2b\xd5\xbc\x54\x31\x61\x2a\x82\x3c\xde\x03\x70\x69\x0f\x9b\x0d\xe4\x62\x31\xae\xb2\x13\xc2\xa3\x36\xe7\x47\x3f\x9b\x81\x54\xe0\x1d\x18\xcf\x84\x21\x83\xe6\xf9\x3f\x7d\x47\x06\x0f\x76\x8b\xc8\x96\x53\x54\x64\xe7\x33\x81\x2d\xc1\xca\xfa\xb8\xdd\x2f\x79\xbb\xb1\xb3\xf4\x4b\xd1\x80\xb4\xfb\x2c\xa5\xc6\xa0\x12\x63\xf2\x9f\x0f\xfe\x1f\x3f\xbe\xfa\x70\xee\xff\xe5\xe3\x1f\x7a\xe4\xff\x74\x48\xd6\xe6\x91\xec\xeb\x72\xc9\xa3\x3c\x72\x60\xf4\x8c\x4b\x6a\xad\xfe\x4a\x8c\xde\xe2\xc0\xfb\x77\xe0\x06\x83\x57\x0f\x71\x00\x15\x71\x43\xe8\x5d\x9f\x6e\x51\x31\x47\xf0\x38\x8a\xfd\xfc\xbc\x38\x3f\x3f\x0d\x17\x06\xef\x0c\x55\x48\x9b\xfc\x02\x0a\x35\xfb\xaf\x3d\xc2\xae\x4b\x32\x02\x4a\xde\xea\x31\xf9\xee\x11\x35\x65\xd2\xe0\xc4\x51\x58\x72\x3e\xcd\x23\x0f\x8d\xba\xe5\xfe\x39\x62\x7e\x42\x40\x7f\xcb\x3b\x09\x74\xdd\x4b\xa0\xed\x6e\xe2\x16\xac\xd7\x1e\x77\x67\xa9\x75\x6c\xbb\xc3\xf2\x7b\xb3\x3d\xf2\xab\x4c\x5e\xaf\x41\xdb\x9b\x71\xf4\xe3\xfb\x7f\xbc\x06\xcf\x8d\xff\xf9\xf3\x6b\x20\x61\x4c\x75\x32\x95\x54\xc5\x21\xd5\x1a\x8d\x0e\x57\x28\x62\xa9\x74\x58\xf5\x58\x3a\x10\x68\xfc\xa9\x0e\x23\xed\x66\xdf\xbb\xd9\xa9\x94\x46\x1b\x45\xd3\x60\xc9\x44\x10\x69\x4d\x60\x46\xb9\xc6\xc1\x13\x4a\xdd\xf6\x76\xa5\x02\xdb\x99\x4f\xa3\x40\xc2\xe6\x09\xb7\xf7\xa4\x85\x93\x67\xe4\x52\x2a\x25\x6f\x9b\x64\xb8\xb8\xe4\xcd\x66\x2d\x27\xec\xb6\x5a\xbf\xf9\xbc\x68\x82\x21\x87\x08\xe4\x19\x74\xb6\x3f\x5b\x95\x8a\xf5\x0e\x1a\x5c\x03\x6b\x2b\xce\x8b\xab\x6a\x61\x53\xc0\x21\x97\xbb\x8f\x84\x32\xfc\x0b\xfd\x84\xc1\x0f\x17\x3a\x5c\xfc\x9a\xa1\xba\x0f\x6a\xf1\xb7\x0e\x59\x7c\x8a\xa0\x4f\xb5\x15\xd8\x8a\xb4\x4f\x22\x73\x0b\xab\x3d\xd9\x35\xbc\x7d\x06\xe1\x85\xed\xad\x20\x7f\x1a\xf1\x75\x8c\x2f\x6a\xaf\x41\x4a\xa3\x9b\x16\x39\xa7\x7d\xf8\xaa\x40\xda\xf3\xca\x6f\x60\x83\x40\x21\x8d\xef\xbd\x59\x26\xf2\xeb\x22\x78\x83\x3d\x94\x27\x7c\xa1\x03\x26\x98\xf9\xb1\x54\x84\x89\xf9\x5b\xf1\x5a\xd2\xd8\x1b\xd4\x70\x3f\xb8\xda\x5e\x4f\x2b\x6e\xc5\x8f\x19\x7b\x2c\x7b\x01\x5d\xd0\x3b\x6f\x7d\x50\x73\xad\xfa\x43\xe8\xff\xf4\xf6\xdd\xfb\xfe\xd9\xc1\x6a\xa6\xf8\x10\xfa\xe5\xbd\x36\x67\xdd\x40\xa5\xb3\x28\x42\xad\x87\x5b\x25\x3c\xb5\xaf\x40\xf9\xb0\x19\x78\xca\x82\x28\xe3\x06\xc6\xe3\x31\xf4\x67\x94\x71\x8c\xfb\x6d\x1b\xec\x23\xf0\x16\x7e\x7a\x23\x0d\x9b\xdd\x37\x98\xb0\x63\x0e\x33\xdc\xda\x93\x5f\xfa\x1b\x54\xdd\xa1\xc5\x3b\x33\x04\x15\x2c\x51\x6b\x3a\xc7\x23\xc4\xce\x4f\x78\x02\xdf\x84\xc5\x38\x74\x68\xe9\x26\xb4\x05\x8b\x89\xf9\x10\xfa\x15\xa8\x5f\xf6\x5b\x77\x6c\x6a\xa1\xdf\x99\xaf\xfa\x37\xaf\xb2\x05\x7e\x67\x3d\x9b\x09\x57\xfd\x9e\xda\xb9\xef\x5c\xbc\x3f\x81\x7b\xf5\x49\x9c\x3f\xb7\x83\x0f\x3f\x02\xd9\x47\xe1\x4c\xa1\x4e\x5c\x59\xf4\x1a\xf6\xee\xb6\x47\x75\xee\x35\x8e\x07\x9f\x5a\xcb\x85\x9e\xd7\xaf\xce\xc4\xfc\x8c\xfc\xc0\xe2\x8f\xf5\x73\xb2\x38\x21\xed\x74\x7f\x10\x44\x09\x15\x73\xac\x8a\xca\x41\x01\x58\x51\x75\xa0\x20\xc2\x18\x7a\x9e\x49\x98\x1e\x1c\xfa\xd1\x96\x64\x18\xc3\x7a\x73\xb8\x74\x03\x63\xc0\x20\x55\x32\xf5\xfa\x2c\xee\x0f\x76\x08\xae\x76\xbd\x65\x61\x59\xd2\xda\x20\xf7\x07\x2e\xeb\xcb\x2f\x08\x8d\xd0\xb4\xb2\x3f\xdc\x7c\xac\x89\x29\x6e\xfe\xfd\x81\x43\x76\x7f\xd7\xdb\x45\x0a\x74\x33\x5a\x51\xee\xed\xaa\xba\x17\xd8\x27\xa9\x91\x2e\x3c\x0d\x64\x56\x95\x61\xfe\xff\x5b\x01\xfd\xc2\xf2\xfb\x5b\x01\x6d\x7b\xbe\xd8\x02\x5a\x1b\x37\x7c\xab\xef\xfa\x71\xa3\x3b\x5a\xc7\xe0\xef\x62\xd2\xaf\x7e\x5d\x80\xcd\xa6\x89\xaa\x1b\xeb\x9d\xee\x3f\xee\xf2\x66\xe3\xc1\xdd\x81\xdc\x2f\xa4\xdb\x4b\xd0\xff\x02\x00\x00\xff\xff\x38\x74\xf6\xe6\x72\x25\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x4b\x6f\x1b\x55\x14\xc7\xff\x93\x69\x79\x18\x90\x00\x21\x95\x05\x8b\x53\xa1\x96\x97\xc6\xcc\xd8\x14\xc1\xc4\x13\x53\xd2\x16\x10\xb5\x5a\x45\xa6\x2c\xe1\xd6\xbe\x75\x46\x8c\x67\xac\x3b\x33\xa8\x95\xba\x70\x6c\x04\x91\x82\xc4\x4b\x08\x89\x05\x08\xf1\x05\x9c\x10\x83\x41\xd8\x6c\x59\x9e\xfb\x05\x90\xf8\x06\xc0\x96\x0d\x9a\x87\x13\x02\xa9\x37\xe7\xf9\xff\x9f\xdf\xb5\xe6\xb7\x87\x4f\x7c\x0e\x00\x15\x00\x8f\x01\x08\x00\x3c\x00\xe0\x67\x14\xbf\x3f\x00\xdc\x0b\xe0\x4f\x00\x77\x01\xf8\x1b\xc0\x23\x00\x56\x8c\xa2\x7f\xca\x00\x1e\x02\x70\xda\x28\xe6\x75\x03\xb8\x1f\x80\x6b\x00\x27\x01\xbc\x62\x00\x4f\x03\x68\x19\xc0\xa6\x01\xdc\x34\x80\x07\x01\xfc\xbe\x02\x3c\x0a\xe0\xaf\x15\xa0\x01\xe0\x3e\x13\x38\x05\xe0\x55\x13\x78\x02\xc0\xdb\x66\xc1\xb4\x6d\x02\x8f\x03\xf8\xcc\x2c\x74\x3f\x98\xc0\x5b\x00\x7e\x35\x01\xa3\x64\x3c\x91\xf1\x94\x7c\x27\xcb\xde\xdd\x65\xcc\x3c\xcc\xf2\x0d\xf7\x64\x8d\xf3\x49\xa2\xfc\xeb\x69\x22\xb1\x1e\x85\x37\xfc\x5e\x11\x54\x9f\x94\x8c\x65\x42\x5d\xf9\xae\xdf\x91\xb8\x90\x07\x3a\x13\x2f\x33\xb1\x94\xc5\xb8\x32\x48\xfc\x28\xc4\xc6\xbf\xf7\xaf\x89\x20\x95\x78\xd3\x0f\x02\xba\x2e\x4b\xab\x4e\xee\x5f\x6e\x54\xe9\xbc\x92\x74\x2b\x4a\x29\x4e\x95\x6c\x62\x43\x0e\x22\x95\x58\xad\xb8\xe7\x77\xad\x97\xd3\x5e\x6c\xb5\x23\x37\xdb\x7d\xe9\x1d\x7f\x53\xf4\xa3\xaa\x4a\x2b\x57\xaf\xb4\xad\x75\x25\x45\x76\xce\xba\x20\x12\xe9\x52\xcd\x76\x5e\xb4\xec\xba\xe5\x3c\x4f\xb5\xba\x7b\xee\xdc\x33\x76\xdd\xb6\x2b\x97\x45\x9c\x58\x6d\x25\xc2\x38\x10\x49\xa4\x5c\x7a\x3d\xf7\xa0\x56\xaa\x44\x3f\xea\x46\xd4\x38\x62\xbc\x56\xb9\x2c\xc2\x5e\x2a\x7a\xd2\x6a\x4b\xd1\x77\xe9\xa0\x76\x69\x23\x8d\x63\x5f\x84\x95\xd6\x6b\xad\x8b\xd6\x35\xa9\x62\x3f\x0a\x5d\x72\xaa\x76\x65\x3d\x0a\x13\x19\x26\x56\xfb\xd6\x40\xba\x94\xc8\x9b\xc9\xb3\x83\x40\xf8\xe1\x2a\x75\x36\x85\x8a\x65\xe2\xbd\xd1\xbe\x64\xbd\x70\xb8\x97\xf1\xdc\x90\xca\xba\x18\x76\xa2\xae\x1f\xf6\x5c\xaa\x5c\x0d\x52\x25\x02\xeb\x52\xa4\xfa\xb1\x4b\xe1\x20\x2f\x63\xaf\xbe\x4a\x45\xea\x85\x67\x1c\xdb\xf3\x1c\x3a\x7b\x96\xb2\xd4\x3e\xed\x39\x0e\x35\xc9\x26\x37\xaf\xd7\xbc\xda\x72\xd4\xf0\x9e\xcb\xd2\x27\xf3\xb5\x86\x63\xd3\xed\xdb\x85\x64\xcd\xab\xd9\x4f\x51\x93\x1c\x72\xa9\xb6\x0a\xfe\x58\x8f\xf4\x48\x0f\x79\xc6\xbb\x7a\xac\x47\xe0\x2f\x79\xc1\x73\xfd\x1e\xcf\xf8\x3b\x3d\xd6\x43\x9e\xe8\xf7\x79\xa6\x3f\x02\x7f\xcd\x0b\xde\xd7\x23\xde\xe3\xa9\x1e\xf2\xf7\xbc\xcf\x53\x9e\xf3\x8c\xa7\xa4\xb7\x78\x57\x0f\x79\xa1\xb7\x78\x42\x7a\xac\xb7\x72\xc7\x05\xff\x98\x65\xbc\xc7\x13\xf0\xb7\xff\x6b\x2e\xb2\xcf\xe7\xbf\xe7\xf5\xce\x1d\xf4\x5f\xf1\x2f\x4b\x8e\x6f\x0e\x8e\xcd\xf4\x48\x7f\x48\xb9\xee\x88\x33\xf8\x0b\x9e\xf3\x44\x7f\xb0\x24\x04\x7f\xa2\xc7\x19\xb0\x1e\x1d\xc2\x6e\xe7\xd3\x09\xf1\x4f\xc7\xbf\xf8\x58\x92\x2a\xf1\xa7\x39\x64\xf9\x37\x4c\x79\xae\x77\x9a\xf8\x27\x00\x00\xff\xff\x40\x07\xee\x74\x28\x04\x00\x00"

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
