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

var _templatesViewsWidgetHtml = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x7b\x73\xdb\x36\xf2\xff\xe7\x53\x6c\x59\x77\x24\xb5\x16\x29\x3b\x6d\x7e\xbf\x53\x24\x65\x72\xcd\x75\xda\x9b\x34\xc9\x34\x6e\x3b\x73\xb9\x9c\x07\x22\x21\x09\x0e\x08\x30\x00\x28\x59\xf1\xe8\xbb\xdf\x00\xe0\x4b\x7c\x48\x94\x2d\xa7\xc9\x4d\x38\x73\x39\x12\x58\xec\x02\xfb\xc6\xae\xdc\x9b\x1b\x08\xf0\x8c\x30\x0c\x8e\xcf\x99\xc2\x4c\x39\xb0\xd9\x3c\x18\x05\x64\x09\x3e\x45\x52\x8e\x9d\x08\xcd\x71\x5f\x11\x45\xb1\x33\x79\x00\x00\x50\x9c\x34\xe3\x97\x14\xcf\x54\x32\x69\x00\x16\x0f\xf3\x0f\xfd\xdc\xdc\x00\x39\xfb\x7f\x06\xce\x33\xbc\x24\x3e\x86\x6f\xa4\x03\x27\xc0\x08\xcd\xfe\xe7\x32\x14\x62\xd8\x6c\x46\x53\x01\xde\xe4\x41\x65\xf5\x0c\x5c\x8a\xa4\xba\x8c\xa3\x00\x29\x0d\xb8\x05\x61\x88\xca\x10\x51\x3a\x19\x49\x5f\x90\x48\x81\x5a\x47\x78\xec\xa0\x28\xa2\xc4\x47\x8a\x70\xe6\x5d\xa1\x25\xb2\x93\xce\x24\xe0\x7e\x1c\x62\xa6\xdc\x95\x20\x0a\x77\x35\xce\x0b\xfe\x5a\x09\xc2\xe6\xdd\xce\xcd\xcd\x16\x31\xf7\x27\x2e\x42\xa4\xc0\x39\x1f\x0c\x1e\xf5\x07\x67\xfd\xc1\xf9\xc5\xd9\x0f\xc3\xc1\xf7\xc3\xc1\x0f\xff\x1a\xfc\xdf\x70\x30\xd0\x2c\xeb\xf4\x7a\x23\xcf\xa2\x9f\x8c\x3c\xbb\x97\xf2\x29\x30\x0b\x8a\x3b\x1f\x79\x29\x9f\x46\x5e\x40\x96\xc9\xa9\xab\xdc\x15\x64\xbe\xd8\x62\x6f\x01\x62\xc6\x45\xd8\x9f\x0b\x1e\x47\x10\xc5\x94\xf6\xcb\xb0\x65\x78\xc2\xa2\x58\xd9\x05\x25\x28\x03\x49\xd1\x14\xd3\x14\x76\xaa\x58\x23\xa4\x81\x9e\xc6\x4a\x71\x96\xb0\xda\x7e\x38\x85\xc5\xa0\x11\x10\x9f\xdb\x97\x15\x12\x8c\xb0\xb9\x79\x97\xa1\x53\x8b\x31\x7d\x02\xa4\x50\x5f\xf1\xf9\x9c\xe2\xb1\x13\xf2\x00\x51\x27\x19\x43\x62\x8e\xd5\xd8\xf9\xda\x0e\xee\x47\x62\xe0\xac\xf6\x8e\x9d\x4c\x0d\x7f\xe4\x6c\x46\x44\x08\x02\x4b\x85\x84\x82\xc0\xa8\xa5\x03\x2e\x6c\x36\xad\x91\x4e\x79\xb0\x2e\xe0\xfc\x93\x50\x0a\x53\x5c\xc2\xe9\xc2\x53\x81\x61\xcd\x63\x90\xb1\xc0\x4f\x0e\x24\xe1\x23\x4a\xa7\xc8\x7f\x37\x76\x90\xaf\x75\xb8\xdb\x49\xb0\x77\x7a\x8f\x1b\x84\x62\x04\x43\x32\xf5\x40\x12\x66\xa8\x1f\xf1\x15\x16\x7d\x3e\x9b\x39\x50\x66\xc5\x6f\x16\x61\xb2\xb1\xc9\xc8\x23\x0d\xc2\xf6\xac\x80\x8f\xa1\x0a\x01\x62\x73\x2c\x3e\x35\x4d\xc0\xc7\xd5\x83\x15\x89\x30\xf8\x1a\xfd\xfc\xfe\x74\x01\x1f\xaa\x09\x58\x20\x89\x45\x55\x0d\xfe\x24\x11\xbe\x83\x0e\x8c\x3c\xe3\x39\x4a\x7e\xc7\x7a\xb5\x9a\xcf\xe4\x35\x75\x7b\x45\x07\xe5\x53\x8c\xc4\x8c\x5c\xeb\x6d\x54\x67\x05\x5f\xd5\x04\x21\x9f\xd3\x7e\x18\xf4\xcf\xce\x41\xbf\xc9\x30\x7d\xbb\x96\xfd\xb3\xf3\x06\xc7\x79\x7d\x19\x21\x86\xe9\x0e\x57\x79\x7d\x59\x0c\x7a\x5b\x50\x8b\xf3\x49\x39\xa4\x21\xa5\x04\x99\xc6\x0a\x4b\xcb\xc6\x91\xb7\x38\xaf\x59\x19\x67\xde\x95\xa1\x25\x30\xb4\x9c\x22\x61\x9d\x36\x98\xfd\x5c\x2a\xce\xe9\x94\x5f\x37\xb9\x5c\x4a\x26\x23\x54\x38\x37\x45\x91\xc4\x7d\x4a\xd8\x3b\x67\x52\x14\xb6\x96\xb5\xbf\xc0\x4b\xc1\x59\x5f\x3b\x70\x2d\xd3\x91\x87\x26\x23\x8f\xd6\x08\x77\xe4\xc5\xb4\x66\x74\x97\x58\x76\x08\xba\xca\xc8\x34\xb5\xa8\xa1\xa1\xd0\x94\xe2\x2c\xde\x99\x0f\xf3\x6f\x5f\x2a\x41\x22\x1c\x18\x13\xb0\xe3\x81\xea\x0b\x2c\x23\xce\x24\x59\x62\x60\x7c\x25\x50\xe4\x80\x54\x6b\xad\xc9\x2b\x12\xa8\xc5\xf0\x6c\x30\xf8\xa6\x89\x75\x6a\x81\x51\xd0\x34\x27\x76\xd8\x90\x5a\xa4\xfb\x0b\x83\xbe\xd6\xab\x73\x27\x17\xff\xd3\x54\xee\xa9\xd8\xd5\x62\x27\xaa\x7c\xe5\x1f\x88\xc6\x7b\x57\x8d\xbc\xa6\xad\xe9\x35\x3b\x0e\xa4\xbd\x52\xf3\x46\x6e\x6e\x40\x68\x27\x0c\x27\x3a\xef\x3a\x85\x93\xa5\xde\x0c\x0c\xc7\xe0\x5a\x5f\x25\x2f\x73\x85\xae\xcb\xb6\x72\x42\x3b\x38\x67\x01\x02\x7d\xe4\x93\x34\xbf\xf3\x54\xc3\x96\x8b\x0b\x76\x02\x40\x96\x12\xe2\xf7\x09\x5e\x87\x84\x11\xc5\x3a\xa1\x33\xa9\x9e\x6b\xdd\xae\xb3\x6b\xe3\x5b\x24\x23\x81\x53\x3d\x8a\x50\x10\x10\x36\x1f\x0e\x9c\xc9\xc8\xe7\x41\xa6\x9a\x57\x92\x33\xe7\x76\xd9\xe5\x3f\x5f\xbf\x7c\xe1\x4a\x93\x5c\x92\xd9\xda\x7e\x46\x48\x48\xdc\x35\x62\x58\xa5\xdc\xdf\x6c\x7a\xa7\xc0\x62\x4a\x4f\xe1\x7c\x2b\x9b\xd4\xfb\x98\x8c\xbc\x48\xe0\x56\x9c\xc1\x54\xd6\x66\xc8\x0d\xe0\x19\xf5\x56\xb8\xb7\x53\xd8\x5a\x6e\xee\x14\x71\xb3\x42\xef\xa7\x30\xf2\x1a\xd4\x7a\xe4\x19\x0f\x71\x84\xf0\xf3\x09\x06\x98\x17\x3c\xf8\x12\x54\x6a\x98\xf7\xa9\x07\x15\x68\xe7\x1e\x0b\x11\xe1\x97\x67\x6d\x82\x48\x65\xd9\x0b\x14\xb6\x8a\x3e\x95\x85\x17\xeb\xe8\x76\x0b\x9f\x0a\x81\xd6\x07\xac\x4c\xf9\x19\x10\x19\x51\xb4\x1e\x02\xe3\x0c\x6b\xa9\xef\x0c\x96\xc7\x0f\x7d\x85\xb0\x77\x0a\x27\x4c\xbb\x77\x1d\xf2\xf4\xcb\x91\xa2\x1c\x0f\xb0\xfb\xcb\x33\xd7\x96\x12\x5a\xc7\xbb\x6c\xa9\x16\xe5\xad\x17\x6b\x71\xde\x7a\xb1\x11\xe9\xa1\xab\xf7\x46\x8c\xfd\x96\x98\x58\xdc\x7e\x54\xd0\xc2\xdc\xaa\xf0\x7b\xe4\x56\x43\x20\xd7\xf2\x57\x82\x47\x58\xa8\xb5\x03\x27\xad\x14\x7d\x27\xb2\x24\xdb\x3b\x1c\xd3\xee\x80\x59\x82\x6c\xcf\x9e\x7d\x19\x62\xf9\xc9\x4d\x27\x3a\x85\x93\x28\x61\x8d\x36\x1f\xab\x3f\x09\xb3\xc8\x6e\x43\xaa\xee\xe2\x60\x01\x59\x9d\x4d\x37\x70\xb0\xc5\xec\x45\x68\x04\x95\x63\x84\xad\xc9\xdf\x19\x51\xb7\xa5\xd6\x5e\x90\xd0\x3a\xd7\x2a\xe0\x6e\x27\xce\xda\x4c\xa9\x06\xe8\x56\x19\xdc\xff\x5e\xf6\x06\x69\x36\x14\xdc\x3d\x8f\xfb\xbb\xe0\x28\xf0\x91\x54\x10\x62\x29\xd1\x1c\x7f\xc9\xe9\x0e\xc9\xe9\x66\x5c\x84\x5b\x55\xef\x05\x17\xe4\x03\x67\x0a\x51\x30\xdf\xa6\x02\xd5\x37\xad\x08\x10\x5c\xa7\x1b\x7a\xd8\x81\x10\xab\x05\x0f\xc6\x4e\xc4\xa5\x72\x80\x04\x63\x67\x9a\x4a\xc2\x01\x5b\x4b\x1b\x3b\x4f\x92\x97\xc2\x14\xe3\x4b\x44\x49\x80\x54\x83\xb5\x6c\x95\xd5\x15\x0e\x21\xaf\xc5\xef\xaa\xc8\xd9\x12\xfb\x8c\x8b\xb1\x43\xf1\x12\x53\x27\x17\x12\x53\x82\xd3\x7e\x52\x83\xb7\x4a\xfb\x30\xd5\xd9\x87\xb5\x2a\x5b\x7e\x32\x65\x7b\x6e\x51\x6b\x05\x83\x91\x8c\x10\xcb\x4c\x04\xbf\x8f\x89\xd0\xfa\xfc\xed\xc8\xd3\x13\xbb\x0c\xbd\xa6\xa8\xd7\xc4\x81\x64\xbf\x8f\xd2\xfd\x3e\x6a\xb5\x5f\x83\xc5\xf4\x24\x92\x2b\xb5\xc2\xd7\xca\xd9\x12\x73\xc2\x16\x2b\xb9\x84\x63\xfa\xd2\x9f\x7d\x44\x14\xf9\x78\xc1\x69\x80\xc5\xd8\x41\x14\x0b\xe5\x80\xb9\xd2\x8e\x1d\x07\xd2\xe3\x16\x0f\xbe\xe3\xc0\x15\xf5\x6c\x33\x75\x67\x3d\xc8\xfc\xc1\x7d\x68\xc2\xaf\x45\x67\xf3\x97\xcb\x5a\xcb\x17\x09\x8c\xea\x44\x0c\x02\x4b\xf2\x41\x87\x82\xcb\x14\x4c\x5b\xf2\x4a\x8e\x9d\xef\x53\x99\x67\xac\xd2\xda\x90\x7e\xe8\x8b\x44\xb2\xe0\x3e\x85\x4b\xd9\xa5\xe4\x94\x04\xb5\xae\xad\x6e\x41\x3b\x45\x68\x62\x6c\x18\xf4\xf9\x6c\x26\xb1\xea\x3f\xdc\xc7\xd5\xad\xee\x8b\x8c\xa7\x21\x51\x95\xee\x8b\x8c\x7d\x1f\x4b\x59\x28\x5d\xbe\xc6\x2c\x48\x63\xd0\xae\xe6\xce\xed\xb8\x37\xf2\xf4\xf1\x6f\x19\xd4\x93\x96\xaf\xc4\x4a\x11\x36\x97\xe5\x96\xf4\x27\x1d\xe7\x7f\x4c\x8a\x8f\x5f\x82\xfb\x81\xc1\xbd\x18\xb4\xcb\x71\x39\xd5\x84\xba\x70\x9e\xcf\x6d\x75\x0b\x93\xe8\xcd\x45\x63\xed\xe6\x63\x55\x89\xe0\x48\x95\xa2\x52\x1b\xe2\xac\x60\xcb\x2f\x23\xcd\xa6\x5b\x15\x73\x5a\xf5\x21\x60\xef\x05\x66\xef\xed\xb3\x7d\x4f\x82\x9c\xc2\x09\x37\x07\x32\xe5\x99\xa2\x17\x68\x46\xde\xb2\x42\x63\xf1\x9a\x6b\xe3\x71\xca\x1c\x79\x2b\x22\x41\x7d\xb1\x8e\x30\x38\x53\xae\x73\x96\x16\xd7\xb7\xad\xf4\xc7\x5f\x60\xff\x9d\x36\xfc\xbc\xed\xd0\x97\x2b\xa2\xfc\x45\x1a\x00\xab\x67\xb0\x46\x50\x37\x9e\x6d\xad\x1b\x09\xc2\x54\x06\x60\x24\xde\x03\x47\x09\x2d\xf8\xcd\x06\x0c\x59\x1c\x64\xf7\x37\xf0\xf6\x9e\xd9\x34\x19\xc8\x0c\xb8\x80\x6e\xe5\xf0\x84\x29\xa7\x57\x3f\xfe\xe8\x7b\xa7\x77\x30\x5b\x58\x1c\x4e\xb1\x68\xc8\x0b\x9b\xf9\x92\x64\x81\x85\xa9\x3f\x92\x56\x47\x33\xcf\x22\xa4\x14\x16\x6c\xec\xfc\xe7\x4d\xff\xbb\xb7\x4f\xde\x0c\xfa\x7f\x7b\xfb\xed\x89\x73\x47\x86\xc4\x4d\x1c\x89\x3f\x2f\x96\xdc\x8a\x23\x95\x43\xcf\x28\x47\xfa\xd4\x9f\xc9\xa1\x73\x3d\xe8\xfe\xdb\xb5\x2f\xbd\x27\x87\x30\x00\xb1\xa0\x46\xf4\xb6\x23\xa8\xb5\x62\x8e\xa1\x4b\x31\x2b\xdb\xe7\xd9\x60\xd0\x4e\x2f\xee\x9a\x54\x1f\xe0\x53\x26\x35\x4c\x6c\x93\x7d\xc3\x61\xad\xc9\xb6\x77\xc2\xa3\xca\xbc\x85\x40\xff\xca\xee\x27\xec\xab\xf3\x35\xd6\xf6\xee\x92\x8a\xe7\x7b\xfa\x7c\x32\xf0\x97\x17\x4f\xbf\xa4\xdf\x47\xed\x97\x1e\xde\x17\xdd\x9d\xeb\x1d\xd6\x17\x7d\xad\x90\x8a\x65\xfb\xe4\xd6\x64\x7a\xfa\xfe\xc8\x15\xba\x14\x31\x63\xb6\x5c\x9f\xe1\x4b\x86\x2c\xc2\xdc\x2b\xe5\xf3\x18\x05\xeb\x7c\xd6\xa8\xff\x5d\xec\xfa\xb0\xd3\xbe\x12\x7c\x2e\xf4\x4d\xfd\x80\xf3\xee\x77\xa8\xc5\x1f\x74\xa7\x04\x5a\x76\x8a\x6a\x96\xf6\xa7\x48\x40\xf1\x23\xd3\x94\x5a\xce\x6f\x41\x12\x36\xe3\xe6\x7e\xb7\xc4\x39\xef\xb7\x20\xec\xef\x43\x2d\xeb\xb5\xbf\xb6\x17\xc3\x14\x64\x8a\x84\x03\x48\x10\xd4\x37\x3e\x9e\xf1\x95\xf1\xe6\x86\xe4\x4a\x10\xa5\x30\x33\x5e\x3e\x07\x09\x09\x1b\x3b\x83\xad\x11\x74\x9d\x2f\x52\x5c\x21\x6a\x96\x6c\xa9\x77\x3a\x9d\xd2\x85\xcd\x66\xd7\x05\xaf\xc2\xb6\x62\xd5\x55\x8a\x3e\x67\x74\x6d\xc2\x67\x69\xa3\x30\x5d\x2b\x2c\x81\xcf\xa0\xbc\x9f\x7d\xf5\xd9\x2d\x6a\xcd\xf5\x99\x03\xc1\x3e\x9e\xa2\xff\xa8\x6f\x1d\x32\x0e\x0f\x34\x6c\xc3\x23\x3f\x59\x7b\x07\xcb\xdc\xdf\x8f\xba\x53\x13\x02\x02\xc1\xa3\x0f\x9c\xe1\x7d\xdd\x08\xae\x50\xb5\xde\x61\x06\x3f\xeb\x0e\xc4\x05\x09\x31\x8f\xd5\x27\x52\x77\x3e\xa0\xc7\xa0\xd2\x8d\xdb\xe4\x32\xfb\xcc\x33\x4a\x6b\xa5\x76\xdc\xf8\x8d\x7b\xee\x32\x1c\x29\x7f\x2b\xfc\xad\xcf\x02\xa3\x20\xbb\x7e\xdd\xdc\x80\x54\x48\x11\xff\xe7\x8b\x5f\x9f\x43\xd7\xbe\xff\xfe\xdb\x73\x70\xbc\x00\xc9\xc5\x94\x23\x11\x78\x48\x4a\xac\xa4\xb7\xc4\x2c\xe0\x42\x7a\x59\x69\x4c\xba\x0c\xab\xfe\x54\x7a\xbe\xb4\xa3\x17\x76\x74\xca\xb9\x92\x4a\xa0\xc8\x0d\x09\x73\x7d\x1d\xce\x66\x88\x4a\xdc\x3b\x22\xd5\xbc\x24\x97\x6e\x20\x1f\xb9\x9f\x0d\x2c\xc8\x7c\x41\x75\xa2\x78\x65\xe9\x29\x1e\x72\x21\xf8\xea\xa8\x87\x4c\xfc\x46\x42\x22\xfd\xac\x23\x61\xe5\x6d\xa2\x56\x41\xb7\xf5\xb2\x82\xfe\x7d\x9d\xd5\xcf\x8c\xbf\x03\x63\x0a\xa7\x95\xe1\xec\x12\x79\xb3\xa5\x67\x36\x16\xea\xbb\xe8\x37\x8f\xb3\x89\x4d\xa2\x68\x86\x72\x93\x8e\x5d\xc9\x23\x6a\x98\x77\x25\xbd\xab\xf7\x31\x16\x6b\xb7\xa0\x64\x9a\x25\x57\xf7\xa1\x59\x53\xa9\x09\x36\xaa\xf3\xbd\xd0\xcc\x75\xb7\x44\xbb\xa0\xd4\x1f\x81\x78\x72\xf6\x46\x4b\x3a\x0e\xf9\xa2\x21\x5d\x15\x3e\xdd\x08\xf9\xef\x8e\x48\x27\x6b\x02\x68\x2a\xd9\xc7\x31\xf9\x98\x5b\xeb\x55\xc9\x58\xab\x04\xda\xfd\x7c\x3c\x33\xb3\x93\x6e\xfa\x4b\xf2\x9e\x6b\xee\x24\xdd\x59\xcc\x4c\x7e\x00\xdd\x5e\xc9\x4e\x9f\xa5\x94\x6d\x51\x43\xea\x18\x05\xe3\x12\x90\x7e\x22\x24\x10\xa5\x98\xfe\x1e\x51\x8e\x02\x39\x84\xb3\xd3\x5a\x98\xf0\x05\x0a\xf1\x10\x9c\x19\x11\xe1\x0a\x09\xec\x54\xc1\x7c\x81\x91\xc2\xbf\x84\x68\x8e\x2f\x16\x71\x38\x65\x88\x50\x39\xb4\x67\xae\x42\x23\xdf\xc7\x91\xc2\xc1\x4f\x84\x62\x39\x84\xad\xe3\x87\xc8\x9f\x12\x86\xc4\xfa\xd4\x9d\x12\xb6\xfd\x37\x51\x9b\xc7\xdb\x89\xd8\x82\x5e\x49\x97\x30\xa2\x7e\x4e\x95\x86\xb0\xf9\x4b\xf6\x9c\xa3\xa0\xdb\x2b\xc1\x66\xfc\x4a\xba\x9f\x3f\x23\x16\x50\x2c\xba\xa2\xcc\x3e\xfd\x90\x19\x74\x85\xb6\xb6\x98\x2a\x18\x8f\xc7\xd0\x99\x21\x42\x71\xd0\xa9\x03\xd6\x0f\xc3\x2b\x78\xf5\x82\x2b\x32\x5b\x77\xeb\x21\xf4\x63\x4a\x16\x43\xe8\xfc\x43\x08\x2e\x3a\x55\xbe\x64\x70\xf8\x5a\x0d\x41\xb8\x49\x57\x7b\x07\xe0\x3a\xd2\xf8\xf0\x1e\x7c\x0b\x12\xe0\x46\x61\xa4\x8f\xf6\xe0\x84\xcd\x87\xd0\xc9\xac\xfc\x61\xa7\x16\x7a\xd3\x7b\x5c\x19\xdf\x64\x25\xce\x6e\xb6\x6f\xf8\x4a\x73\x2e\x66\x36\x0c\x1c\x8b\x79\xaf\xad\xfc\x8e\xc8\x3e\xb9\x17\xe3\xc7\x60\xe0\xb6\xa6\x3f\x28\x85\x5e\x16\xf0\x95\x6b\x6f\x04\x30\xce\x94\xb9\xeb\x87\x41\x1d\x57\x4f\x5c\x74\x85\xae\x1b\x98\x99\x9c\xfa\xd5\xcb\xd7\x17\x0d\x47\x8e\x05\x1d\x42\x27\xbd\x80\x74\xe0\x3b\xf0\xc3\xa0\x1e\x34\xe1\xdd\xb0\x64\x56\xd5\xe3\x95\x8e\x5c\x36\xe5\x93\x6e\xe7\xeb\xec\x17\x57\x9d\x9e\x6b\x7f\xba\x90\x79\xb9\x2e\x5e\x6a\xdf\x57\x3d\x90\x19\x77\x23\x61\xfe\xff\x19\x9e\xa1\x98\xaa\x6e\x0d\x7b\x97\x48\x00\x86\x31\x9c\x74\xd5\x82\xc8\xb2\x73\x38\x16\xcf\xb0\x1b\x09\x1e\x75\x3b\x96\x73\x9d\x5e\x3d\xa8\x0e\xb8\x1a\x56\x62\x41\x10\x25\x1f\x70\xb7\x01\xf0\xd6\xcc\x2d\x9f\xaf\xe6\x47\x14\x5b\x47\xef\x76\xf2\x1c\xd0\x24\x85\x6f\x48\xf0\x76\x2b\x31\x4c\x52\x42\x3d\xde\xe9\xb9\xfe\x02\xb1\x39\xce\xa5\x53\xa7\x83\x4b\x54\xdd\xa8\x7e\x0a\x52\x68\x66\x8f\x8e\x57\x9b\xfa\xe9\x77\x30\xce\xf8\x4c\x82\x4e\xaf\x02\x54\x23\x5b\xed\x96\xd2\x35\x5a\x94\x9d\x9e\xf5\xea\x69\x93\xb5\xd1\x35\xe9\xbd\xbc\x79\xf7\xb6\x40\x32\x69\x90\x76\x7a\xf0\xd5\x18\x3a\x9d\x46\x4f\xb8\x1f\xdf\x12\xd1\x6e\x75\xf7\x9b\xfb\xd1\xcc\xa2\x88\x3b\xbd\x03\xd4\x54\xff\x7b\x7c\xed\x2c\x7e\x17\xee\x0c\xa5\x61\xad\xb3\x26\xb6\xd5\xd5\x0d\xf6\x47\x8d\x36\xe1\xd6\xc6\x0a\xf3\x1f\x7a\x48\x29\x35\x41\xee\x8f\xb5\x7b\xc3\x44\xbb\x10\xd1\xcc\xa0\x32\x44\xf6\x97\x81\xf9\xc5\xeb\xbf\x01\x00\x00\xff\xff\x8a\x01\x28\x6d\xc7\x43\x00\x00"

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

var _localesRuLc_messagesWidgetMo = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\xcf\x6f\x13\xd7\x16\xfe\xb0\xb1\xcd\x33\xef\xf1\x1e\x3c\x4a\x7f\x4b\x17\x55\x40\x69\x35\xa9\x9d\x94\xfe\x70\x62\x02\x24\x20\x50\x09\xa4\xe0\xd2\x4d\xa5\xea\x12\x5f\x9c\x11\xf6\x8c\x75\x67\x06\x9a\x8a\x45\xe2\x54\xa5\x08\xaa\x42\xa5\x0a\x54\x51\xaa\xc0\xa2\x5d\x74\xe1\x24\x98\x18\x9a\x4c\xa4\x76\x53\x75\x75\xee\xae\xdd\xf4\x2f\x40\xed\xa6\x9b\x2e\x2a\x55\x77\xe6\xc6\x04\x48\xbc\xb9\xf7\x9e\x39\xdf\x77\xbe\xf3\x9d\xe3\xdf\x36\xad\xfd\x02\x00\x9e\x07\xf0\x1c\x80\xb5\x6b\x80\x9d\x00\xee\xaf\x41\xf4\xbb\x9c\x00\x52\x00\xae\x24\x80\x7f\x01\xb8\x9a\x00\x36\x02\x98\x4a\x00\xeb\x00\xcc\x26\x80\x34\x80\xf9\x04\xb0\x19\xc0\x8f\x09\x60\x0b\x80\x5f\x4c\xfe\x1f\x26\xff\xaf\x04\x90\x00\xb0\x3e\x19\xf3\x6d\x48\x02\x19\x00\x9b\x93\xc0\x5a\x5d\xdb\xc4\xb7\x26\x63\xbe\x1d\xc9\x98\xbf\xcb\x9c\xaf\x9b\xfc\x3d\x26\xff\xa0\xc9\x1b\x32\xf1\x92\x89\xbf\x67\x78\x78\x12\xd8\x06\xa0\x92\x8c\xfb\xf9\xd4\x7c\x9f\x33\xdf\xe7\x0d\xee\x87\x24\x30\xba\x06\xf8\x39\x09\xfc\x1b\x40\x4f\x0a\xf8\x9f\xae\x93\x02\x76\x01\x78\xdb\xbc\xcf\xa7\x80\xa7\x00\x5c\x4b\x01\x7d\x00\xa6\x53\xc0\x41\x00\xbf\xa7\x80\x27\xb5\xbe\x34\xb0\x03\xc0\xd1\x74\xdc\xe7\xd9\x34\xb0\x01\xc0\x87\x69\x60\x13\x80\x0b\xe9\x58\xef\xf5\x74\xdc\xcf\xad\x34\x90\x05\xf0\x5d\x3a\xe6\x9f\x33\xe7\x4f\x69\xe0\x19\x00\xbf\x1a\xdc\x9f\xe9\x58\x57\x26\x13\xf3\x6d\xce\xc4\x3c\x2f\x64\xe2\xfc\x97\x32\xc0\x30\x80\xde\x0c\xf0\x3e\x80\x4b\x26\x7e\x3f\x13\xe3\xfe\x36\xef\xec\x3a\x40\x8f\x54\x6b\xdb\x10\x8f\x16\x4f\x9b\x73\x3d\x1e\xfe\x3d\x61\x76\x21\x09\xe0\x59\xc4\xbe\x69\x4d\x9a\x47\xd7\xd6\xfe\x69\xed\xba\xef\x2d\x06\xf3\xff\x65\x78\xdd\xdf\x7f\x97\xbd\xf5\x1e\x68\xef\x32\xe6\xbd\xd1\xec\x9b\xd6\xa7\x77\xe6\x3f\x88\x7b\x8d\x7e\x7b\xa5\xe4\x63\xd8\xeb\xfb\xd2\x3e\x19\xf8\x02\xfb\xa4\xcb\xcb\x23\xdc\xf3\x59\x4d\x78\x1e\xaf\x08\x0c\x8c\x8a\x91\xd3\x5e\x50\xc3\x80\xeb\x9c\xb2\x2b\xf1\x21\x6b\x4c\x0a\x4f\xf8\xac\x2c\xce\xd8\x23\x62\x79\xd0\xe7\xb2\x13\x1e\x8c\x0e\xb6\xcd\x5b\xba\xf1\xa5\x42\x1e\x0e\x0d\xe2\xb0\x38\x23\xaa\x18\x32\x85\x8e\xf0\x9a\xc0\x11\xb7\x2c\x3c\x1c\xad\xfb\xb6\xeb\x60\x58\xba\x15\x29\x3c\x4f\x5f\xea\x42\xfa\x63\x38\x16\xf3\xe3\xb8\x70\xca\x38\xee\x73\x3f\xf0\x50\xb2\x6b\xc2\x0d\x7c\x94\xc6\xea\x02\x27\x78\x35\x10\x78\xd7\xae\x56\xd9\x49\xf1\x88\x9c\x2e\xb6\x57\x0a\x36\xe6\x06\xcc\x0b\xa4\xe8\xef\x64\x9d\xb5\xeb\x82\x8d\x44\xdd\xad\x9a\x59\x17\x90\x82\x97\xc7\x20\x03\xc7\xb1\x9d\x0a\x8e\x89\xba\x2b\x7d\x6b\xc8\xab\xd8\x65\x6b\x5f\x50\xf1\xac\x92\x5b\xd0\xf0\x3d\xa7\xed\x51\x5e\x73\xbb\x64\x90\x1d\x3e\x5a\xb2\x06\xa4\xe0\xba\x19\x6b\x90\xfb\xa2\xc0\xba\x73\xf9\x37\xad\x5c\x8f\x95\x7f\x8d\x75\xf7\x14\x76\xed\x7a\x39\xd7\x93\xcb\x65\x0f\x73\xcf\xb7\x4a\x92\x3b\x5e\x95\xfb\xae\x2c\xb0\xb7\x22\x0e\x36\x14\x48\x5e\x73\xcb\x2e\xeb\x7b\x88\x78\x77\xf6\x30\x77\x2a\x01\xaf\x08\xab\x24\x78\xad\xc0\x3a\xef\x02\x3b\x16\x78\x9e\xcd\x9d\xec\xd0\xa1\xa1\xfd\xd6\x09\x21\x3d\xdb\x75\x0a\x2c\xdf\x95\xcb\x0e\xb8\x8e\x2f\x1c\xdf\xd2\x3e\x15\x98\x2f\x3e\xf0\x5f\xa9\x57\xb9\xed\xf4\xb2\x91\x51\x2e\x3d\xe1\x17\xdf\x29\x1d\xb0\xde\x78\x90\xa7\xf5\x9c\x12\xd2\xda\xef\x8c\xb8\x65\xdb\xa9\x14\x58\x76\xb8\x1a\x48\x5e\xb5\x0e\xb8\xb2\xe6\x15\x98\x53\x8f\x9e\x5e\xb1\xa7\x97\xc5\xd7\xa2\xb3\x2d\x9f\x2b\x16\xf3\x6c\xfb\x76\xa6\xaf\xb9\xad\xc5\x7c\x9e\xf5\xb3\x1c\x2b\x44\xef\xdd\xc5\xee\xa5\x4f\x7d\xc5\x57\xf5\xf5\xc5\x28\xad\x2f\x9f\x63\xe7\xce\xc5\x90\xdd\xc5\xee\xdc\x4e\xd6\xcf\xf2\xac\xc0\xba\x7b\x41\xd7\xa9\xa9\x26\xd4\x04\xb5\x69\x06\x74\x59\x35\x54\x43\x8d\x53\x9b\xa6\xd5\xa4\x6a\x80\xbe\xa5\xb6\x1a\xa7\x90\xee\x51\x48\x33\xd4\x52\x17\xa8\xa9\x1a\xd4\xa2\xef\xd5\x25\x5a\xa0\x90\x5a\x4c\x4d\x50\x48\x21\x4d\xab\x0b\xd4\xa2\x05\x6a\x53\x0b\xf4\x0d\xb5\xe8\x9e\x9a\x50\x93\x34\x4f\xf3\xd4\x04\x7d\x49\x21\x2d\xa8\x8f\xa8\x4d\xb3\x6a\x52\x8d\x53\x53\x7d\x4c\x6d\xf5\x19\xe8\x6b\x0a\xe9\xb6\x6a\x44\xdc\xe3\x74\x87\x6e\x2f\x91\x68\xde\x69\x5d\x5a\x4d\x50\x93\xa9\x49\x35\x11\x09\x0b\xe9\xae\xbe\xd1\x8c\x66\x5d\x1d\x4c\x8b\x51\xa8\x45\x73\xd4\xa4\x59\x35\xae\x26\x69\x8e\xee\x51\x7b\x15\xa2\x5b\x8f\x05\x43\xfd\xaf\x7a\xd4\x0e\x75\x71\x65\xfc\xa1\x41\x4d\x31\x1e\x5b\x44\x0b\xea\x12\x68\xea\x71\x53\xae\xd1\xbc\xee\xf8\xab\x48\xf4\x45\xd0\x0d\x5a\xec\xb8\x10\x81\x67\xb5\x62\x3d\x0b\x0d\x9f\x59\x2e\x46\xb7\xfa\x78\x3f\x6d\xd5\xd0\xa5\x6e\xa8\x06\x2d\x6a\x4f\x69\x66\x29\x34\xa5\x1a\x7a\x4e\x5a\x2c\xe8\x26\x35\xe9\xae\x1e\x43\x3c\xd1\x9b\xd4\xa6\x45\xd0\x55\x5a\xa0\xa6\x3a\xff\x40\xde\x4a\x1e\xe8\xb6\xb5\xab\xaa\xb1\x92\xa3\x77\x22\x70\xd8\xc5\xe8\xf3\xc8\x19\x33\x07\xed\xc0\xc5\x7e\xd0\x95\x0e\xb6\x33\xca\x4f\x22\x44\x93\x45\xeb\xb4\xc2\x3e\xac\x68\xef\x6a\xfc\x53\x9d\xfd\x30\x5d\xcf\x52\xa8\x1a\xd1\x10\x42\x44\x32\x17\xd5\xa4\x19\x40\x88\x7f\x02\x00\x00\xff\xff\xc9\xa5\xfa\x87\x1b\x08\x00\x00"

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
