// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/views/widget.html
// locales/ru/LC_MESSAGES/widget.mo
package v1

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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x79\x6f\xe3\x36\x16\xff\x7f\x3e\xc5\x83\x90\xd9\x49\xda\xca\x77\xbb\xbb\xa9\x93\xa2\xd8\xb4\xd8\x01\xda\xc5\x60\x9a\x9d\x62\xd1\x76\x0d\xda\xa2\x63\x75\x24\x52\x43\xd1\x76\xb2\x41\xbe\xfb\x82\xa4\x0e\xca\x96\x6d\x9d\x8e\x95\x91\x81\xc9\x58\xd7\x7b\xe4\x8f\xbf\x77\x48\x7a\xa4\x1f\x1f\xc1\xc2\x73\x9b\x60\x30\x66\x94\x70\x4c\xb8\x01\x4f\x4f\xaf\xc6\xfe\x8c\xd9\x1e\x07\xfe\xe0\xe1\x2b\x03\x79\x9e\x63\xcf\x10\xb7\x29\xe9\xfe\x89\x56\x48\x1d\x34\xae\x5f\x01\x00\xcc\x97\x64\x26\x8e\xc0\x0a\x39\x4b\xfc\x23\x65\x2e\xe2\xe7\xe4\x02\x1e\xe5\x51\xf1\x61\x98\x2f\x19\x01\x82\xd7\xf0\x96\x70\xa7\xf3\xaf\xa5\x3b\xc5\x2c\x38\xf3\xa2\x33\x0f\x2f\xf9\x56\x5e\xf1\xf4\x6a\xdc\x55\x0a\xae\x5f\x8d\x2d\x7b\x05\x33\x07\xf9\xfe\x95\xc1\xe8\x3a\xd0\xa8\xef\x9d\x51\xc7\x74\x2d\xb3\x3f\x00\xf1\xcd\x77\xc3\x6f\xf7\xbe\xd9\x1f\x04\xe7\x6f\x5e\x73\x3f\xf1\x10\xc1\x8e\x76\x74\xfb\x8c\x10\x8c\xe4\x39\xd1\x79\x8c\x3a\xf8\xca\xe0\x68\x9a\x26\x29\x3a\x73\xe9\x84\x02\x09\x5a\x01\x41\x2b\x93\xa3\xa9\x0f\x53\xc4\x26\xe2\x8b\x11\x8b\x71\x6c\x3f\x4d\x57\x24\xc9\xb1\x83\x73\x3d\x86\x7d\x4c\xb8\x1c\x0b\xe3\xf1\x11\xec\x39\xe0\x4f\xd0\x41\x6a\x08\x0c\x31\x78\xa1\x4e\xb1\x6f\x85\xc5\x49\x98\x58\xf0\xf4\x74\x3d\x46\xb0\x60\x78\x7e\xb5\xe3\xba\x78\x64\x2f\x57\xd4\xb6\xce\x7b\x17\xdf\x8a\x6b\x1d\x1f\xc3\xd3\xd3\xe3\x23\x74\xde\xe3\x4f\x4b\xec\xf3\xce\xbf\xdf\xff\xd4\x79\x87\xf8\x42\xed\x56\xc2\x8d\x6b\x21\xb4\xff\x37\x02\xc6\x3f\x96\x8c\x61\xc2\xc1\xc5\xc8\x5f\x32\xec\x62\xc2\x7d\x03\x3a\xf0\xf4\x34\xee\xa2\xeb\x71\xd7\xb1\x2b\xe8\xa8\x4b\x09\x5f\x38\x0f\x45\xfb\xab\x5d\x5e\xa4\xdb\xdf\x29\x39\x57\x81\x98\xbf\x20\xc7\xb9\xea\xa7\x61\xf1\xb3\x3a\x01\x7c\xd1\x11\x9f\xdb\xb3\xea\x91\xb0\x6c\xdf\x73\x50\x61\x24\xb4\xcb\xcb\x20\x11\x88\x49\xc3\xe0\x26\xd4\x50\x71\xc7\xf1\x4a\x30\xcb\xa4\xc4\xa4\xf3\x79\xd1\xee\x6f\x09\x29\x03\x42\x42\x58\x1a\x14\x3f\xc8\x13\x80\x92\xae\xd4\x56\x13\x20\x1e\x26\xe6\xcc\xa1\x3e\x2e\x0b\x4a\x42\x50\x15\xc0\x44\x02\xf7\x81\xe3\x61\xd2\x0d\xb4\x56\x0c\xd0\x82\x3a\xb6\x85\x1e\xfc\xa2\xb8\xe8\xd7\x97\x81\x23\x94\x93\x86\xc2\x3f\x23\x1d\x87\x3b\x3f\xee\x2e\x9d\x1d\x47\xb4\x60\xc6\xd1\xd4\xdc\x1d\xce\x12\x57\x6c\x84\x35\x5d\x82\xd8\x03\x73\x64\x61\x50\x88\x81\x4d\xf6\x48\x13\x9f\xbd\x5e\x77\xef\x95\xb2\x3d\x5e\xa4\x1d\xdf\x73\xd3\x5d\x72\x6c\xc1\x9c\x12\x6e\xf6\x87\xe0\x9a\x53\x73\xd8\x3b\xa0\x5f\x6b\x07\x43\x6b\x38\x4f\x46\x27\x0b\x71\x0c\xdc\x76\x31\x50\x02\x16\x5e\xd9\x33\x0c\xaf\x25\xea\xc4\x76\xa2\x7f\xe7\x1e\xb3\x09\x07\x23\x91\x0f\xfd\xbe\x23\x21\xfa\xdd\xb8\xb6\xe8\x6c\x29\x02\x5e\x67\xcd\x6c\x8e\xcf\x85\x92\x5b\xfa\x0b\x67\x36\xb9\x3b\x7f\x63\xc0\x79\x47\xec\xe9\xa8\xd4\x07\x8c\x41\xaf\xf7\x8d\xd9\xeb\x9b\xbd\xc1\x6d\xff\xeb\xcb\xde\xe8\xb2\xf7\xb5\xd9\xfb\xeb\x65\xaf\x67\x5c\x80\xf1\xe6\xe2\x22\x4a\x84\x8c\x8b\x8b\x4c\x98\x75\xbd\xeb\x57\x87\xcf\x12\xf9\x06\xd6\xc6\xd6\xc1\x20\xff\x9a\x3e\x67\xb6\x87\x2d\x81\x0d\x52\xfb\x2d\x6e\x32\xec\x7b\x94\xf8\x62\xc8\x09\x5d\x33\xe4\x19\xe0\xf3\x07\xc1\x93\xb5\x6d\xf1\xc5\x65\xbf\xd7\x7b\x9d\x71\x24\xc6\x7c\x81\x91\x95\xf5\x5c\x96\xed\xc4\x40\xf0\x46\xcc\x0d\xcd\x87\x2f\xb2\x4b\x51\x84\xbd\xe3\x70\xd6\xe1\x88\xd9\xf3\xf9\x64\x46\x97\x84\x43\x4f\x48\xd2\x35\xdc\xf6\xa1\x0b\xbf\xea\x3a\x22\x4b\x2e\xad\xac\xbf\xa5\x6c\x50\x9f\xb2\xc1\x96\xb2\x61\x7d\xca\x86\x5b\xca\x46\xe5\x94\x8d\xbb\x59\x39\x22\x84\xe7\x60\xde\x94\x5a\x0f\x39\x7c\x0b\xb9\xc3\x70\x66\x7f\x05\x67\x22\xcb\xf3\xe1\xf2\x0a\x3a\xea\x5b\xd6\x7e\x70\xa6\x30\x43\xc4\x82\x73\xfc\x29\x10\xd4\x91\x44\x86\x33\xe5\x33\xe4\xc6\x85\x7e\xf8\x3f\x18\xb1\xf0\xa8\xf8\x7e\xa1\x45\x35\x7f\x39\x9b\x61\xdf\xd7\xc2\x5a\x0e\x5b\xb2\xa2\x21\xd2\x1b\xd2\x51\x6e\x0c\xce\xa4\x63\x34\x2c\xe4\x2f\xa6\x14\x31\x4b\x06\xd3\xc7\xc7\x44\xab\xe4\x70\x66\x84\x1b\x0e\x58\x5e\xae\x86\x67\x3e\x59\x5e\x90\xe9\x26\x77\xc3\xa5\xeb\x77\xb9\x71\xaf\x3f\x88\xbd\x7e\xe7\x56\x36\x5e\x98\xb0\xe6\xbc\x73\xb5\x29\x46\x42\x09\xbe\xed\xdf\x32\x31\x82\xb9\x90\x88\x3b\xe8\x21\x12\x52\xc2\x41\x53\xec\x80\xfc\x6b\x5a\x82\xb3\xcc\xb8\xae\x10\x81\xdb\xfe\x0d\x76\x38\x4a\x76\x1d\xba\x55\x82\x1c\x62\x91\x50\x31\xee\x8a\x4e\xe6\x06\x59\x26\x6a\xf6\x1c\x9c\x9a\x91\x0e\x0d\xb1\x85\x1a\x62\xcf\x55\x0f\xd4\x36\x99\xd3\x16\xe7\xa4\xf3\x50\x1d\xa8\xd8\x79\xe0\x39\x5a\x3a\xbc\x7e\xa8\x8b\xe2\x90\x2f\x5d\xc9\x1d\xa9\x2a\x4c\xf4\x9a\x17\xdb\x06\x55\xc6\xb6\x41\x53\x62\xdb\xa0\x7e\x47\x30\xa8\x35\xb6\xd5\x82\x74\x2d\xb1\xad\x99\x50\x6b\xb1\xad\x16\xa8\xab\x8f\x6d\xcd\xc4\x59\x77\x1e\x8d\x89\x6d\x69\x50\xbf\xec\xd8\x36\x68\x64\x6c\x1b\x56\x19\xdb\x86\x4d\x89\x6d\xc3\xfa\x1d\xc1\xb0\xd6\xd8\x56\x0b\xd2\xb5\xc4\xb6\x66\x42\xad\xc5\xb6\x5a\xa0\xae\x3e\xb6\x35\x13\x67\xdd\x79\x34\x26\xb6\xa5\x41\xfd\xb2\x63\xdb\xb0\x91\xb1\x6d\x54\x65\x6c\x1b\x35\x25\xb6\x8d\xea\x77\x04\xa3\x5a\x63\x5b\x2d\x48\xd7\x12\xdb\x9a\x09\xb5\x16\xdb\x6a\x81\xba\xfa\xd8\xd6\x4c\x9c\x75\xe7\xd1\x98\xd8\x96\x06\x75\xf3\x63\x5b\xf6\xb7\xa6\xb9\xe5\x66\x7b\x6f\x3a\xee\xca\xda\x82\x83\xc5\x23\x9a\x8d\xee\x2e\xd7\xda\x2b\xe4\x79\x2a\x1d\x8e\x53\xe5\x70\x83\x38\x2e\x52\xe4\x90\x10\xf2\x0b\x2f\x20\x25\x1b\x81\x32\xbf\x72\x2f\xfe\xba\x5d\x92\x41\xbe\x6e\x57\xb4\xc8\xf1\xbe\x3d\xdf\xeb\xf0\x42\x4e\x25\x59\xfb\x23\xdc\x8a\x6c\x65\xe7\xd6\x76\xb3\x54\x00\xc1\xd3\xd3\x9b\x84\xdf\xc9\xe3\x11\xc2\x77\xf8\xf3\x50\xa9\x1c\x68\x55\x98\xa6\x46\x9e\x12\x03\xce\x82\x02\xde\xa8\x66\x2d\x38\x24\x2c\xeb\x4c\x2f\xee\xcd\xae\xbc\x0e\xe7\x92\xc9\xb1\x54\xe1\x54\x12\xe5\x8e\xad\x63\x69\x1d\x4b\xeb\x58\xd2\x1a\x7d\xc0\xb1\x78\x78\xb7\x6b\x09\xac\xeb\x73\x72\x2e\x5a\x7d\xfd\xe1\xc6\xcc\x29\x73\xc1\xb6\xae\xe2\xab\x02\x17\x23\x0e\x98\x0b\xca\xec\xff\x51\xc2\x91\x03\x72\x5b\xa5\xbe\x0e\x9e\x73\x43\xfa\x1b\x93\xd3\xbb\x3b\xe1\x55\x56\xc8\xb1\x2d\xc4\x29\x33\xc0\xc5\x7c\x41\xad\x2b\xc3\xa3\x7b\x27\x98\x24\x5a\xa1\xd5\x0e\x4b\x3d\x77\x8c\x2e\xbd\x8c\x17\x4b\x01\x2a\x33\x8f\x66\xe8\x10\xce\xa8\x63\x06\x3b\xd5\x7c\x9d\x51\x38\x5d\x67\xa4\xcd\xd6\x11\xbd\xba\x32\x5c\x6a\xe1\x09\xef\xe7\xd0\x07\xc1\x33\x0a\xe5\x79\x16\x74\x0d\xb7\x7d\xe5\x7b\x92\x37\x0c\x0c\x7f\x5a\xda\x0c\x5b\xc6\xf5\x17\x79\xd3\xf8\x71\x57\x36\x3f\xc7\x05\xdb\x73\x94\xbe\x09\xfb\xfc\x4d\xea\x0c\xa5\x4c\x52\x6d\xe2\x2d\x43\x77\x31\x5b\xe0\xd9\xc7\x29\xbd\x8f\x48\xf2\xa7\x6f\xfa\x6b\x9b\xcf\x16\x06\x10\xe4\xe2\x18\x4a\x49\xa9\x68\x23\xaa\xcb\x0e\x8a\x9b\x3b\xe2\x48\xe7\xad\x1f\x14\xaf\x5d\x80\xc1\xd9\x32\x98\x3b\x20\x54\x60\x2b\x32\x27\xe8\xe6\x81\xcc\xb2\x57\x59\x0d\x3b\xfb\xa9\xa7\x40\xce\xbc\xc3\xb6\x41\xce\x41\x4b\x4e\x1c\x43\xa9\x91\x73\x70\x88\x9c\x83\x96\x9c\x07\xc9\x39\x2c\x47\xce\x61\x4b\x4e\x1c\x43\xa9\x91\x73\x78\x88\x9c\xc3\x96\x9c\x07\xc9\x39\x2a\x47\xce\x51\x4b\x4e\x1c\x43\xa9\x91\x73\x74\x88\x9c\xa3\x96\x9c\x87\xc8\x89\x5c\xba\xdc\x3b\x45\x2e\xed\x93\x24\x68\x20\xa2\x25\x69\x12\xd2\x98\xa8\xe1\x8e\x5d\x64\xfd\x5e\x1e\x6f\xb9\x7a\x88\xab\x1e\x5d\x63\x56\x8a\xaa\x4a\x42\xcb\xd4\x04\xa0\x31\x51\x83\xed\x5d\x3c\x7d\x27\x0e\xb7\x34\x3d\x18\xef\x6d\x17\x97\x62\xa9\x14\xd0\x92\x54\x87\x53\x8b\xfa\x72\x73\x67\xdc\xb7\x5d\xdc\x32\xf4\x10\x43\x2d\xc4\xcb\x31\xd4\x8a\x1e\x74\xb7\x0c\x0d\xe0\x8c\x19\xaa\x36\x77\x31\xf4\x06\xf1\x96\xa1\xfb\x18\x2a\x0c\xbc\xc4\xa3\xd0\xb7\x84\x63\xb6\x52\x4f\x8b\x81\x50\x0e\xb3\x60\xed\x03\x55\xcf\xf6\x19\xb0\x96\xc8\x95\xa6\x42\x86\x86\x70\x4a\x7e\x46\x1b\xb2\xf8\x43\xae\xf6\xd1\x11\xfb\xb4\x39\xbc\x06\x78\x88\x73\xcc\xc8\x95\xf1\xdf\xdf\xcc\x2f\xff\xf8\xee\xb7\x9e\xf9\xf7\x3f\xbe\x38\x33\x5a\x5e\xba\x65\x9e\x82\x26\x78\xd9\x72\x32\x7e\x0a\x1a\x6d\xa4\x73\x72\xd0\x72\xf2\x10\x27\x0b\x3f\xfc\x4c\x70\x52\x26\xff\x5f\xa9\xd5\x61\x10\xb1\x3e\x97\x18\x9f\xc6\xcc\xa1\xce\xcc\xe1\x2e\x66\x0e\x5b\x66\x1e\x62\x66\xe1\x27\x9f\x09\x66\xa2\x39\xc7\x0c\x3c\x86\x7d\x1f\xa6\x68\xf6\x11\xa6\x4b\xce\x29\xf9\x4c\xc9\x39\xd2\xc9\x39\xda\x45\xce\xd1\xb3\x91\x33\x37\x3b\x1d\x32\xf1\xa9\x63\x5b\xc6\xf5\x51\xe9\x9d\x6d\xf4\x42\xaa\xd3\xf9\xdc\xc7\xdc\xcc\xeb\x6a\xc7\x8a\xaa\xc1\x70\x32\xec\x63\x1e\xdd\x4c\x4c\x39\x81\x29\x27\xaa\x5c\x3a\x1c\x46\x75\x4a\x5c\x53\xf4\x5e\x5d\xa2\x6a\x8a\x94\xb0\x32\x2d\xf0\x97\x53\xd7\xde\x6e\x42\x58\x1c\x1f\xb6\xc2\x47\x2b\xac\x35\xe2\x17\xb1\x59\xb0\x0d\xd5\x7b\xb7\x71\x57\x0c\x78\xa1\xba\x14\x7d\x2d\xbb\x0c\xa5\x54\x2f\x7b\xc9\xb0\x9b\x78\x79\xca\xc2\x25\x6f\x85\x96\x1d\x3b\x81\x25\xac\x60\xab\xfc\xcd\x42\x0f\xb2\xf8\x2d\xa4\x48\xbe\x09\x4f\x39\xc6\x00\xe2\x8a\x32\xa1\x34\x2e\x5e\x33\x72\xaf\x20\xa5\x89\x52\x2b\x57\x09\x79\x07\xd7\xad\xca\x59\xfc\x96\x79\xb0\xe0\x44\xaa\xe7\x3f\x5b\xd3\x0e\xfb\xe3\x5a\xa6\x88\x5a\x03\x7d\x49\xde\x78\x59\xe2\x82\x36\xbf\x47\xb8\x9c\x73\x57\xda\x95\xdc\x60\x55\x9c\x69\x87\x99\x5d\xe3\x1c\x4a\x51\x37\x10\x97\x20\xcb\x1b\xaf\xf8\xf1\x7f\x7e\x6f\xa0\x9e\x3a\xca\xe5\xf1\x64\x32\x88\x19\xa3\xac\xd0\xa4\x23\x6e\x25\x96\x22\x0d\x27\x2a\x8a\x34\x33\x29\xbd\xf3\x43\xa0\xa3\x50\x6b\xb3\x5a\x6c\x4a\xfb\xf2\xe9\x8a\x2e\xac\xa4\xfa\x38\xc6\x40\x66\x4c\xf9\x2b\x90\x0b\xb5\x5d\x8d\x2e\xc1\x9a\x7a\xab\xf8\xbc\xb2\x18\x11\xce\x28\xb9\x4b\x1d\xee\xc2\x42\xa3\xe9\xb3\x95\x36\x15\xe2\x7b\xc4\x77\xce\x92\x21\x07\x8c\x2f\x5f\x5b\xe0\xe3\x19\x25\x96\xa1\x6f\xf8\xc6\x96\xe6\xb3\xc4\x22\xb7\x9b\x47\x4b\xb4\xaa\x04\x8d\xf7\x74\x4c\xef\xd7\xb3\x75\x2b\xdf\x0c\xbe\xcd\xcf\xb8\xab\xa8\x55\x98\xed\x05\xf5\x17\xf3\x44\xf9\x75\x25\xfc\x77\x62\x8d\xe5\xd0\x8f\x03\x25\xe0\x62\x1e\x16\x1d\xd4\x9c\x7a\x09\xd7\xef\x21\x86\xdc\x89\x83\x7c\x3e\x99\x2d\x84\x15\x4f\x44\x2a\x73\x8c\x6c\x56\xc1\xf0\x4e\xe8\x07\xa1\x1f\x94\xfe\x92\xa1\x2c\xbd\x3f\xf5\x04\xb6\x7d\xba\x9e\x2f\xcc\x55\x14\xb1\x76\x74\xae\x60\xfc\x7a\x16\x13\xdb\xe2\x96\xf6\x90\xfa\x74\xee\x6a\xa0\xb4\x0d\xfd\x8c\x3e\x16\xef\x1a\xc4\xb6\xe3\xa2\x8f\xea\x75\x74\x3d\xe6\xb2\x21\xfe\xf9\x2c\x24\xd9\x98\x5d\x8c\x2e\x76\x7f\x5d\x01\x6d\x4b\x0d\x67\x3e\xa6\x96\xe3\xdd\x8f\x36\x73\xd7\x88\x61\x58\x61\xe6\xeb\xf7\x61\xc5\xe8\x17\x48\xa9\x87\x7c\x09\xe1\xcf\x4b\xbd\xb0\x29\x92\x78\xf0\x5c\x2c\xab\x62\xf0\x8e\x49\xb6\x9f\x84\x1b\x97\x2f\x1d\x41\xfb\x5d\x98\x82\x54\x93\x61\x4d\x0a\x9b\xd0\xf9\x7c\x52\xef\xfd\xef\x5e\x65\x8d\xcf\x13\x76\xf5\xae\x01\x89\x42\x33\xc8\x5e\xd6\xad\xea\x03\x44\x8e\x48\xf5\x4d\x5d\x2f\x8a\xe9\xa4\x25\x7a\x24\xbf\x34\xd1\xe5\x8c\x77\x98\x21\xaf\x0a\xa6\x4b\x61\x93\x19\xf2\x8e\x41\xf5\x5d\xca\x5e\x06\xd7\x53\x7a\xd7\x92\xbd\x04\xd9\x3f\x50\x87\xa3\xd2\x4f\x38\x56\x4a\x4a\x4d\x69\xb2\x2e\xfc\x59\xd3\x64\x05\xd9\x6b\x0b\x3e\x6c\xfe\xde\x58\xd4\x48\x95\x40\x3f\x53\xfe\x1c\x3e\xbf\x5b\x15\x1f\xd4\x42\x0f\xec\x02\x7d\x13\x17\xdd\x1f\xef\x09\x43\xc0\x5c\x70\xd1\xbd\xed\x2e\x5d\x98\x21\xf2\x86\x03\xd3\x8b\x4d\xca\xb1\x59\x74\xa7\x56\x46\xc7\x0a\x4e\x9c\xd5\xb2\xa1\x21\xb3\xe3\x0a\x37\x14\xcf\x98\xac\xc4\xc1\xeb\xda\x32\xfe\xde\xde\x0b\xf5\xe9\x45\x57\xbd\xd5\x21\x9c\x48\x53\x78\x56\x7b\xac\xda\x1a\x55\x97\x6a\xb7\x49\x5d\x4d\x03\x2c\x33\x68\xee\x51\xed\x33\xd0\xd9\x5a\xe9\x31\x9f\xa9\x7f\xef\x7a\x98\x95\x4f\xd6\x50\x20\xa6\x1e\x3b\x4a\x4a\x3f\x05\xf3\xe9\x0c\xe6\xf0\xfd\x96\x01\x45\xed\x3c\x89\x8c\x0d\x95\x18\xda\x42\x29\x5b\xa8\xf0\xb8\x39\x5b\xc8\xe0\x1a\x92\x36\xbd\x43\xf5\x32\xfb\xb4\xd2\xb6\xfd\xec\xd6\x32\xb7\x5a\x43\x43\x42\x5f\x1b\x15\x0a\xe5\x6e\x3a\x86\xc7\x4e\xde\xd2\x0c\xb3\x72\xb3\xac\x33\x7d\xdb\xa5\xa7\x19\x26\x7a\xb4\x14\x2e\x45\x69\x6b\xad\xc7\xcc\xe1\xde\xc5\xcb\xb7\x14\xaf\x27\x12\x32\x6a\x2a\x1f\x8a\x45\x9f\x82\xe5\x58\xf0\xeb\x96\xdd\xa8\x26\x9e\x44\xd6\xe6\x15\x1d\xcc\x62\x65\x71\xf2\x35\xcf\x51\xf3\x35\xc9\xd6\x1a\x92\xb5\xa8\x2b\x35\xb2\xf8\xb4\xd2\xb4\x9d\x4c\x3e\xd2\xb3\xb5\x58\x57\xeb\xef\x0b\x65\x67\x11\x80\xc7\x4e\xcd\xb6\x6c\xb0\x5a\x0b\xac\x33\x29\x4b\x55\x72\xf2\xd6\x78\xb4\x64\x6c\x53\x63\x6b\x99\xc7\xcc\xc4\x6e\xd7\x14\x7c\x4c\x7c\xca\xfc\xa2\xb6\x04\x5a\x9d\x2a\xb5\xb0\x33\xe1\x6b\x3a\x09\x64\x16\xb7\x28\x38\x54\xb3\x9a\xae\xaa\x84\x5d\x41\xf9\x19\x20\xf1\x0f\x0c\xa4\xb4\x2f\x34\xa5\x08\xfb\x07\x1c\x60\x9e\xf2\x5b\x03\x84\xc6\x87\x72\xfe\xd0\xc0\x66\x7f\x4a\xa6\x79\x82\x21\xe1\x32\x37\x25\x98\x72\xcc\xd7\xf9\xef\xb1\xa3\x4d\x89\x2e\x4b\x67\x26\xa4\xd5\x4d\x64\x4d\xc9\x89\x51\x58\xb5\xac\xa1\xe4\x7d\x4b\x08\x66\xc0\x8a\xf2\x21\xd7\x7c\xd4\x8a\x67\x40\xef\xef\xf9\x9e\x25\x16\x76\x1c\x4a\xd9\xbd\xb1\x4b\xdb\x0c\xbe\x06\xff\xc5\xcd\xf9\x7f\x00\x00\x00\xff\xff\x37\x03\x76\x3c\x71\x97\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x5b\x6c\x1c\x67\xf5\xff\xf5\x62\xef\x7a\x9d\x38\xf1\x25\xfd\xb7\x7f\xa0\x7c\xa5\x4d\xef\x1b\x5f\xa1\x65\x13\x27\xad\x93\x14\x7a\x71\x55\xa5\x6e\x8a\x04\x2a\x9d\xee\x8e\xed\x25\xbb\x33\xab\x99\x59\xa7\x46\x05\xf9\xd2\x34\xa9\x1c\xe2\xaa\x82\x12\x44\x20\x4a\x41\xf0\x80\x04\x1b\x37\x9b\x6c\x7c\x59\x3f\xf1\x46\xd1\xf9\xde\x78\x40\x2d\xbc\x70\x15\x2f\x48\x08\x54\x21\x84\xce\x9c\xcf\xeb\xbd\x25\x75\x1f\xca\x3c\xcc\xf9\xce\xe5\x3b\xdf\xb9\xfc\xbe\x33\xf3\x5e\xd7\xcd\xdf\x01\x80\x43\x00\x3e\x09\x60\xdf\x8d\x80\x05\xe0\x2b\x37\x23\x7c\x7e\xd7\x02\xb4\x02\xf8\x7d\x0b\xc0\xa2\x3f\x19\xfa\xb7\x16\xa0\x07\xc0\xdf\x5b\x80\x5b\x00\x44\x5a\x81\x28\x80\xdb\x5b\x81\xff\x07\x70\x4f\xab\xf8\x1b\x69\x05\xee\x05\xf0\xd5\x56\xe0\x13\x00\x4e\x19\xfe\xc7\x86\x7f\xb7\x15\xd8\x09\xe0\xcf\xad\xc0\xed\x00\x3e\x30\xf4\xb6\x88\xf8\x7f\x38\x02\x6c\x07\x30\x1a\x01\x76\x00\xf8\x72\x44\xce\x3f\x66\xe4\xb9\x08\x70\x13\x80\x6f\x46\x80\x76\x00\x73\x11\x20\x02\xe0\x8c\xd1\xbf\x15\x01\x3a\xf9\xbc\x08\xb0\x1b\xc0\x65\x43\x7f\x1b\x01\xee\x00\x70\x63\x54\xe2\xbe\x2b\x0a\xdc\x05\x20\x1e\x95\xb8\x8e\x44\x81\x4f\x03\x70\xa2\x62\x77\x26\x0a\x74\x00\x28\x18\xba\x1e\x15\xff\xbf\x8e\x02\x6d\xec\x2f\x2a\xe7\xff\x21\x0a\xb4\x70\x5d\xa2\x40\x17\x80\x7f\x19\xbe\xbd\x0d\xf8\x3f\x00\x9d\x6d\xc0\x6d\x00\xee\x6b\x13\xf9\xc1\x36\xc9\xe7\xf1\x36\x89\xfb\x69\x43\x8f\x1a\xfa\x82\xa1\xe3\x6d\xe2\xdf\x69\x93\xf3\xbe\xde\x06\xc4\x00\xbc\x6a\xf8\x45\xe3\xef\x7b\xc6\xfe\xbc\xa1\x3f\x35\xf4\x17\x86\x5e\x31\x76\xeb\x86\xff\xa5\x89\xe7\xdd\x36\xc9\xfb\x8f\x6d\xc0\x8d\x00\x6e\x8a\x89\x5d\x6b\x4c\xea\xbb\x2d\x26\xf2\x9d\x31\x89\xb7\x3b\x06\x1c\xb8\x01\xb8\x35\x26\x7e\xbe\xd4\x2e\xf2\x17\xdb\x85\xb7\xdb\x81\x41\x00\xd9\x76\xe0\x73\x00\x2e\xb4\x0b\x8e\x7e\xd3\x0e\x0c\x00\x78\xbf\x1d\x78\x08\x40\xd7\x36\xe0\x45\x00\x4f\x6c\x03\x0e\x00\xf8\xd9\x36\x60\x1c\xc0\x5f\xb7\x01\x87\xb9\x0f\xdb\x01\x05\xe0\xec\x76\x60\x2f\x80\xab\xdb\x81\x21\x00\xff\xd8\x0e\xdc\x09\xe0\xfe\x0e\xe9\xd3\xd3\x1d\xc2\x3b\x1d\xd2\xcf\x37\x3b\x04\x97\xdf\x37\xfc\x52\x87\xe0\x6c\xa5\x43\xfa\xf2\xab\x0e\xc1\xd9\xfb\x1d\x40\x1f\x80\xff\x74\x48\x1c\xbd\x3b\xe4\xfc\x37\x76\x00\x2f\x00\xf8\xe7\x0e\xe0\x53\xdc\x8f\x9d\xc0\x13\x00\x82\x9d\xc0\x33\x5c\xbf\x9d\xc0\x18\xd7\xa3\x53\x78\xbb\x13\x78\x94\xf1\xd5\x09\xdc\xc3\x7d\xef\x14\x9c\xa9\x2e\xc1\xcf\x63\x5d\x82\xe7\xc9\x2e\xe9\xdb\x37\xba\x44\x7f\xaa\x4b\xe2\xfa\x79\x17\x90\x60\xbf\x5d\x52\x87\x0f\x8c\x3c\xd1\x2d\xf1\x3e\xd9\x2d\xf8\x79\xc1\xd0\x29\x43\xcf\x18\xfa\x93\x6e\xa9\x43\xb9\x5b\xe2\x7d\xaf\x5b\xfc\xff\xdb\xc8\x3f\xd3\x23\x7e\x46\x7a\x04\x47\xcf\x19\x6a\x19\x9a\x35\x74\xba\x47\xce\x3d\xd1\x23\xf1\xbe\xd5\x03\xec\xe7\xbe\xf4\x00\x07\x01\xfc\xa5\x47\x70\xb0\x67\x97\xd8\xf5\xef\x12\xbb\x91\x5d\xe2\xff\x79\x23\xff\xda\x2e\xe0\x06\xc6\x0d\x36\x1f\xae\x37\xdf\xb1\xdb\x0c\xdf\x52\xa5\xbb\xd5\xd0\x83\x55\x32\x8e\x9f\x6b\xda\x0b\xc9\xa1\x1f\x82\x1d\xae\x1f\x63\xe9\x61\x63\xd7\x61\x28\xf7\x91\x73\xe7\xfa\xdd\xcf\x31\x56\xf9\x62\xec\x30\x1e\x19\xe3\xf7\x19\xd9\x67\x01\x3c\x08\x99\x2b\xbb\x20\x79\xf0\xc3\x78\xbf\xb9\x6a\x6f\xa7\x99\x8d\x23\x10\x2c\xf1\xac\x7c\x00\x82\x73\x7e\x38\xef\x3b\xcd\x7a\x98\x6b\x0d\xe9\xf9\xdd\x90\x59\xf7\x90\xd1\x71\xdd\x6e\xa9\xf2\xcb\xf3\xe3\x11\xc8\xbc\xe4\x1a\x73\x1f\x19\xeb\xdd\x3c\x87\x21\xf7\x3a\x8e\xda\x87\xef\xd2\xe7\x21\x77\xa0\xbd\x4a\x1e\x62\x05\x82\xbd\x9a\x67\xf7\x9e\x81\x71\xf5\x28\x76\xa7\xd4\x51\x7e\x3d\xcf\x2f\xdf\x4e\xba\x4e\x6a\x73\xe5\xe3\x81\x4d\xe9\x03\x55\xe2\x47\xb3\x39\xdb\xb3\x26\xec\xca\x42\x65\xad\x97\xd3\xd9\x7c\x56\x25\x2d\x47\x79\xb6\x6f\x07\x4d\x55\xf7\x04\x46\x79\x70\xd2\x72\x26\x6c\x95\x4a\xfb\xb9\x8c\x35\xad\xb2\x6e\xca\x56\xe3\x56\x3a\x63\xa7\xd4\xf1\x74\x30\xa9\x6c\xcf\x73\x3d\xb5\xdb\x6f\x6a\xe8\xe7\x93\x49\xdb\x6f\xd0\x05\xe9\xec\xd6\x9c\x84\x86\x15\x27\x79\xcf\xb3\x9d\x40\x59\x1b\x29\x6d\x08\x52\x56\x60\x2b\xcb\x49\x89\xb9\xeb\xa8\xac\x1d\xd8\x5e\xad\x7a\x43\x95\xb2\xa7\xd2\x49\x3b\x3c\xca\xa8\xb3\xb6\xe5\xe7\x3d\x3b\x6b\x3b\xc1\xa6\x30\xe7\x1e\xaf\xf2\x30\xe5\x66\x02\x3e\xf0\x90\x15\xc8\xab\x72\x1a\x0e\x59\xd3\x38\x64\xfb\x49\x2f\x9d\x0b\xd2\xae\x83\x43\x12\x39\x0e\x4f\xb1\x3f\xe5\x3a\xbd\xee\xf8\x78\x85\xcb\xd9\x4e\x6f\x32\xe3\xfa\x36\xbe\x60\x07\x5b\xa8\x69\xb5\xd5\xb5\x8b\xc6\x56\x93\x6e\x26\x9d\xb2\xa6\xfd\xe6\x16\x5f\x34\x5a\x3c\xee\x04\xb6\x37\x65\x65\xd4\xb8\xeb\x29\x6b\x3c\xb0\x3d\x95\xf3\x6c\xdf\x57\x2f\x59\xc9\x63\xea\xa5\x7c\x10\xb8\x4e\xad\x51\xd2\xd4\x20\xb0\xbc\xf4\xf8\x78\xad\xce\x71\x83\xeb\xea\xc3\x2a\x3e\x28\x91\x73\xc5\xb8\x13\x78\xca\xf2\x03\x15\x16\x41\x25\xad\x9c\xb0\xa1\xa1\xe2\x4a\x55\xb3\x0e\x46\xad\x63\xb6\xec\x1a\xdd\xec\x12\x46\x5d\x27\x98\x94\x77\x66\x5a\xf9\x81\x15\xa4\xfd\x20\x9d\xf4\xf1\x4c\xd8\xb5\xf0\xdd\x04\xe6\x0d\xf2\x0a\xc6\x8f\x84\xef\x67\xad\x29\x1b\xcf\x4e\xba\xc7\xd5\x58\xbf\xa1\x03\x86\x0e\x1a\x3a\x24\xd4\xca\xba\x79\x27\x90\x75\x18\x5d\xb8\x12\xcc\x84\xcb\x10\x1a\xcf\x06\xac\x1a\xeb\x57\xbd\xea\xf9\x49\x8c\x0d\x18\x3a\x68\xe8\x90\xd0\xa3\x56\x26\x6f\xe3\xa8\x41\x98\xa1\x4d\xa2\x6f\xa2\xa9\xc4\x6f\x05\x10\x54\x71\x05\x5d\x07\x8c\x33\x1c\xb1\x73\xae\x17\xc4\x47\xfd\x89\x74\x2a\x3e\x92\x9f\xf0\xe3\x63\x6e\x82\xf1\xff\xc8\xb1\xf4\xa4\x95\x75\xf7\x78\xf9\x18\x57\x3b\x3e\xe6\x59\x8e\x9f\xb1\x02\xd7\x4b\xa8\x27\x43\x95\x1a\xcd\x7b\x56\xd6\x4d\xb9\x6a\x5f\x8d\xfd\xfe\xd8\x53\x96\x33\x91\xb7\x26\xec\xf8\x98\x6d\x65\x13\xaa\xc2\x27\xd4\x91\xbc\xef\xa7\x2d\x27\x36\xfa\xf8\xe8\xe1\xf8\x51\xdb\xf3\xd3\xae\x93\x50\xfd\x7b\xfa\x62\x07\x5d\x27\xb0\x9d\x20\x3e\x36\x9d\xb3\x13\x2a\xb0\x5f\x0e\x7a\x73\x19\x2b\xed\xec\x55\xc9\x49\xcb\xf3\xed\x60\xf8\xb9\xb1\xc7\xe2\x0f\x6f\xda\x71\x3c\xe3\xb6\x17\x3f\xec\x24\xdd\x54\xda\x99\x48\xa8\xd8\x33\x99\xbc\x67\x65\xe2\x8f\xb9\x5e\xd6\x4f\x28\x27\x17\xb2\xfe\xf0\xe0\x5e\x25\xcb\x61\x67\x77\x7f\xdf\xf0\x70\xbf\xba\xfb\x6e\xc5\xcb\xbe\x3b\x86\xfb\xfb\xd5\x01\xd5\xa7\x12\x21\xbf\x7f\x78\x60\x43\xb5\x6f\x78\x88\x97\xf7\x86\x66\xfb\xfa\xfb\xd4\x2b\xaf\xc8\x96\xfd\xc3\x03\x7d\xf7\xa9\x03\xaa\x5f\x25\xd4\xc0\x5e\x19\xbb\xf4\x06\xcf\xd8\x11\x7e\xd1\x9b\x7a\x8e\xa9\x9e\xa5\x22\x2d\xeb\x79\x5a\xa3\x4b\x54\xa8\x97\xe8\x85\x7a\x49\x38\x8e\xeb\x36\xd5\x8b\xf4\x42\x83\x08\xf4\x36\x95\x69\x19\x74\x8e\x0a\xb4\xac\x67\xa9\x44\xab\x7a\x9e\x56\x95\x9e\x63\x39\x15\xd8\xfa\xa2\x9e\xa1\x82\x9e\xd5\x0b\xb4\x44\x05\x2a\xd2\xaa\x5e\xa0\xab\xd7\xd9\x42\x6b\x54\xbc\xce\xbe\xb3\x74\x85\x56\xa9\xc8\x56\xb4\x46\x25\xb6\x9d\xa1\x22\x5d\x66\x47\x54\xa6\x25\x45\xa5\x30\x81\x12\x3b\xd3\xaf\x51\x89\x4a\x8a\xae\x50\x81\x96\xa8\xa8\x67\xf4\x29\x2a\xd1\x0a\x95\xf5\xac\x3e\xad\xf4\xac\xa2\x72\x28\xb9\x48\xcb\x54\xa6\xab\x3c\x86\xe8\x47\x7a\x96\xd6\xa9\xa8\x4f\xd1\x1a\x95\xd9\x5b\xd5\x79\x7a\xa1\xea\x34\xbd\xd0\xe4\xac\x66\x01\xb2\x95\x9e\xe3\xe3\xc3\x54\x56\xfe\x17\x61\xd6\x9f\x79\x8d\x60\xdf\x96\x66\xea\xd7\xa9\xa0\x17\x55\xd8\x8e\x15\x6e\x9b\x69\x46\xb5\x41\x98\xc9\x25\x2a\xe8\x39\x6e\x52\x49\xd1\x52\x58\x8a\x55\xbd\xc8\x2d\xe3\x56\xeb\x93\x54\xd4\x73\xfa\x64\x78\x44\xb1\xde\xf9\xe6\xde\xd0\x7a\x5e\xcf\xea\x39\x3d\xc3\xf9\xf0\x8a\xf3\x0e\xf3\xaa\x3b\x70\x5d\x02\xe1\xd2\x70\x3a\xb4\xd2\x18\x35\xf7\x5d\xbf\xce\x65\x08\x5d\x9e\xae\xd6\x17\xd9\xc7\x1a\x15\x68\x5d\xcf\xe8\x45\xba\xbc\xd1\x13\xd0\xb7\x25\x9a\xca\xa2\x36\x25\x16\x73\x21\x4f\x83\x7e\x48\xeb\x54\xd2\xb3\x54\xd8\xd8\x79\xb6\xb6\x8c\x6c\x7c\x81\xca\x74\x51\x2f\xe8\x39\x66\x15\x2d\xd1\x32\xad\xf4\xd2\x92\x5e\xe0\x45\x83\xba\xac\xe7\x68\x59\xcf\x88\x80\x8a\xbd\x61\xdf\xab\x04\xa0\xf3\x54\xa6\x15\x3d\xcf\x05\xfd\x78\x40\xde\x78\xc0\xc7\x02\xd2\x26\xc7\xac\xf3\xb5\xa6\x2b\x74\x89\x05\xfa\x24\x83\x55\x9f\x50\x21\x5b\xa4\xab\x1f\xd1\x7d\xa3\x2f\x01\xe9\x9a\x5c\xc4\x9a\x84\x9a\x26\xb3\xce\xfe\x69\x65\x03\x25\x97\x43\x2c\x84\x4d\x5a\x66\x40\xd1\x3a\x2d\x6f\xd1\x95\xfe\x16\x67\xa0\x67\xe8\x1d\x61\xe7\xaa\x30\xf8\x0e\x95\x59\x50\xd0\x33\x54\xd2\xaf\x32\xe8\xb6\x12\x5b\x38\x07\x37\xdd\x94\xf4\x89\x6a\x27\xd2\x1e\xbe\x75\xf3\xb4\xca\xc3\x68\x6b\x4e\x6b\xee\x0a\x95\x1e\xac\x60\x5e\x3a\x14\xde\x82\xf0\xa2\xb2\xc3\xf3\x1b\xd5\x31\xed\x29\xaa\x7a\xa8\x2a\xc3\x9c\xe2\x3a\xd5\x8f\x80\x42\x53\x0f\x21\xfa\x69\x45\x9f\xd9\x44\x45\x53\xbb\xa5\x46\xab\xca\x55\x5d\x0f\x07\x07\xcf\xbc\x25\x2a\xd3\x25\x33\x40\xe4\xbc\xda\x71\xc1\x37\xf8\x1c\x15\xf5\xac\x5e\xd4\xaf\x6d\x2e\x4f\x72\xbb\xc3\x89\x17\x96\x94\x7b\x1e\xd6\x43\xa2\x3e\x57\x3f\x50\xea\x3f\x56\xf5\x65\xfc\x08\xdf\xb9\x86\xad\x1f\xf2\xbd\xbb\x10\x2a\xb8\x3a\x25\x89\xe5\x02\x95\xf5\x89\x10\xf8\x6b\x1b\xa2\x4a\xd6\xb2\x99\x85\xfc\x7f\xd8\x5c\x3e\x70\x0d\xf9\xe0\x35\xe4\x43\xcd\xe5\x15\xe0\xcd\x37\xd7\x1b\x14\x5d\x4b\xdb\x30\xb3\x9b\x5a\x55\xcd\xe3\x0b\xa6\x60\x65\xbd\x68\xe0\x10\xfe\xb5\xf2\x1f\x8f\x3e\x29\x7f\xae\x1b\xeb\xc1\xaa\xf5\xd0\xe6\x9a\xbe\x1b\x36\xbd\x0a\x4f\x3f\x68\xf6\x61\x68\x68\x58\xc3\xe7\x43\xbe\x94\x5b\xed\x77\xb3\xed\x1f\xd2\xf3\x25\x34\x7e\x12\xcc\xc7\xa4\xf6\x42\x34\xb9\x23\x75\x1f\x17\xfc\x37\x00\x00\xff\xff\x40\xe9\x9e\xde\x7b\x15\x00\x00"

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
