// Code generated by go-bindata.
// sources:
// ../../../../../bamboo-scripts/ack.lua
// ../../../../../bamboo-scripts/close.lua
// ../../../../../bamboo-scripts/consume.lua
// ../../../../../bamboo-scripts/enqueue.lua
// ../../../../../bamboo-scripts/fail.lua
// ../../../../../bamboo-scripts/maxfailed.lua
// ../../../../../bamboo-scripts/maxjobs.lua
// ../../../../../bamboo-scripts/recover.lua
// ../../../../../bamboo-scripts/test.lua
// DO NOT EDIT!

package bamboo

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
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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

var _bambooScriptsAckLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\xef\x6f\xda\x30\x10\xfd\xee\xbf\xe2\x94\x7d\x28\x95\x20\x5b\xbf\x32\xa8\xd4\x96\x88\xd1\x76\xb0\x85\xad\x68\xab\x2a\x14\xc8\x05\x32\x82\xcd\x6c\xa7\xac\xff\xfd\xce\x3f\x02\x01\xa1\x6a\x1d\x52\x05\xb9\xf8\xde\x7b\x7e\xf7\xec\xb6\x5a\x8f\x8f\x8c\x25\xf3\x15\x74\xb8\xba\x84\x26\x74\x7e\x89\x59\x9e\x5e\x32\xc6\x55\x1b\x86\xc9\x1a\xd5\x26\x99\x23\x94\x3c\x45\x09\xdb\x65\x3e\x5f\xc2\xef\x12\x4b\x84\x34\xd1\x09\xe0\x9f\x5c\x69\x15\x32\xdb\xd4\x86\x5b\x31\x83\x3c\x45\xae\xf3\x2c\x47\x19\x32\x16\x49\x29\x24\x48\xd4\xa5\xe4\xa0\xb4\xcc\xf9\x42\xb5\x19\x03\xfa\x7c\x1f\xde\x0d\x47\x93\xe1\xf4\x76\x74\x3d\x1d\xf4\x18\x6b\xb5\x9e\x9e\x18\x2b\xc4\x3c\x29\x20\xe3\x44\x0c\x5d\x08\x48\x58\xe0\x6b\x0a\x37\xa6\xd2\xae\x9e\x25\xaa\xb2\xd0\x55\x47\x21\x16\xd3\x6d\x42\x24\x5d\xc8\x4a\x3e\xd7\xb9\xe0\xd0\x20\xf1\x2a\x59\xe0\xb9\xe5\x93\x98\xe6\x2a\xa4\x75\x0d\xf7\xeb\x7e\xd4\x9f\x4e\xae\xe2\xe1\x60\xd8\x6f\x42\xd0\x09\x20\x0c\x3d\x2f\xfd\x08\x2e\xed\x73\x00\xf6\x6b\x87\x83\x3c\xad\x13\x3e\xa3\x9c\x09\x85\x6f\xe3\x7c\x88\xe2\xeb\xd1\x38\xfa\x1f\xce\x5c\x4d\xd1\x1a\xba\x27\x6c\x38\x1b\x2a\x3a\xeb\xb3\x7e\xd9\x60\x55\x87\x6e\x17\xce\x74\x32\x2b\xf0\x0c\x12\x9e\x7a\xd7\x42\x82\xa9\x03\x73\x45\x90\x77\xd1\x8f\xf1\xe3\xc5\x93\x2f\xd9\x89\x52\xf5\x2a\xee\x3f\x98\xaa\x2f\xaf\xb6\x42\xae\x68\x8a\xf4\x86\x9a\x48\xa7\x19\x8b\x51\x3d\x19\xc5\x77\x64\x65\x00\xd0\x6a\x99\x1c\x28\xd0\xcb\x44\xc3\x32\x79\x46\x98\x21\x72\x98\x0b\xae\xca\x35\xa6\x75\x20\x94\xea\x24\x50\x14\x8f\x1d\xd0\xc4\x2e\x82\x41\x4f\x31\xe6\xa4\xd3\x7a\xe7\xe5\x86\x50\x8a\x46\xf0\x73\x7c\x33\x8a\xa3\xa0\xb9\x53\xd6\x74\xd2\xcf\x59\x9e\x41\xd5\x42\x1c\x79\x01\x64\x5c\xe5\xe0\xce\x1e\xbd\x44\x6e\xbd\xab\x02\xd4\x08\xbe\x48\xf1\x4c\x21\x4e\x0d\x0e\x70\xa1\x21\x13\x94\xfe\xd0\x86\x7b\xd0\x6b\xbb\xf1\x38\x7b\x8c\xde\xaf\xe6\x34\xf8\x6a\xa5\xe1\x60\x1c\x4e\xae\xa5\x9d\x4a\xdc\x14\x2f\x8d\xe0\x30\xfb\x7e\xec\xb7\x75\x3a\xc8\xb9\x3b\x67\xe1\x11\x70\x6d\x6a\x2b\x23\xf0\xd8\x3e\x82\x1c\x07\xb5\x82\x15\x6a\x4e\x17\xbc\x3b\xf9\x31\x6f\x62\x5c\x0b\x1a\x93\x81\xcb\xa4\x58\x43\x35\x63\x2b\xe0\x95\x56\xb7\x33\x3f\x87\x38\xfa\x7c\x6a\x0a\x6f\x67\x46\x79\xa6\x48\xbd\x7e\xa5\xd3\xed\x7f\x5e\xe4\x74\xd5\x4c\xfd\x5d\x51\x17\xf3\xa9\x1f\x7d\x33\x62\x08\x97\xce\x99\xd8\x72\x94\x81\x0d\xc4\x41\x0b\xdd\x26\x81\x09\xc5\x51\x31\x4b\x0a\x3a\xd3\x27\x82\x31\x14\x5e\x1f\x71\x2d\xe8\xde\x43\xa2\x7c\xef\xa6\x95\x11\x0c\x91\x85\xc4\x82\xd4\xed\x1b\x6b\x39\x27\x81\xbb\xc4\xef\x87\x53\x63\xf6\x89\x39\x15\xf1\x71\xcd\x5a\x94\x3b\x67\x4d\x03\x59\x34\xc8\x8c\x56\x89\x90\xd0\x1f\x17\xb0\x16\xf4\x2d\x4a\xad\x34\x9d\x77\x33\x45\x5a\xae\x9a\x84\x68\x9d\xa6\xa5\xd5\x26\xac\xe1\xe6\x99\xbc\x06\x91\x55\x78\x5e\x65\x58\xdb\x03\x37\x10\x47\x16\x8f\x6f\xae\xe2\xde\x5e\x95\x93\x43\x0e\xbb\xb5\x9d\x2e\x7c\xd8\x5b\xb8\xbf\x0a\x0f\x36\xc4\x49\x56\xcd\x01\x07\x61\xf2\x6d\x33\xfe\x2f\xb9\x31\x47\xa6\x47\xff\x87\x5e\xcd\xe8\x29\x4b\x7b\xd1\xbd\xcf\xc7\xf9\x47\x73\xb3\xd8\x83\x7a\xc1\xfe\x06\x00\x00\xff\xff\x41\xad\xee\x33\x0e\x07\x00\x00")

func bambooScriptsAckLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsAckLua,
		"bamboo-scripts/ack.lua",
	)
}

func bambooScriptsAckLua() (*asset, error) {
	bytes, err := bambooScriptsAckLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/ack.lua", size: 1806, mode: os.FileMode(420), modTime: time.Unix(1448246920, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsCloseLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x8d\x8e\xe6\xe2\x72\x4b\xcc\xcc\x51\x48\xcc\xab\x54\xc8\x4f\x53\x28\xc9\xc8\x2c\x56\x28\xcf\x2f\xca\x4e\x2d\x52\x2f\x56\x48\x2e\x2d\x2a\x4a\xcd\x2b\x51\xc8\xca\x4f\x2a\xd6\xe3\x0a\x4a\xcd\xcd\x2f\x4b\x85\xca\x2a\x00\xc5\x8b\x32\x53\x81\xc2\x5c\xba\xba\xb1\xb1\x5c\x80\x00\x00\x00\xff\xff\x53\xaf\x4e\xfd\x4b\x00\x00\x00")

func bambooScriptsCloseLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsCloseLua,
		"bamboo-scripts/close.lua",
	)
}

func bambooScriptsCloseLua() (*asset, error) {
	bytes, err := bambooScriptsCloseLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/close.lua", size: 75, mode: os.FileMode(420), modTime: time.Unix(1448250332, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsConsumeLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x58\xdd\x73\xda\x48\x12\x7f\xd7\x5f\xd1\xa5\x7b\x30\x54\x61\x25\xd9\xbb\xba\x07\x5f\xec\x2a\x02\x2a\x87\xd8\x86\x2c\xd8\xf1\xde\xa5\x52\xec\x80\x06\x98\x58\x8c\x58\xcd\x28\x98\xfb\xeb\xb7\xbb\x67\x06\x84\x3f\xb0\x13\xca\x55\x46\xa3\xfe\xee\x5f\x7f\x0c\xc7\xc7\x5f\xbf\x46\xd1\xb4\xd0\xa6\x5a\x4a\x78\xaf\xcd\x19\xb4\xe0\xfd\x34\x57\x52\xdb\xb1\x16\x4b\x79\x06\xef\x33\x61\xa5\x55\xfc\xf5\x7b\x31\x51\x19\xfe\x97\xf7\x2b\x55\x4a\x73\x16\x45\x17\x72\x63\x4e\x98\x31\x02\xfc\x68\x7c\xe8\x23\x9b\x59\x89\xa9\x84\x4a\x67\xb2\x84\xf5\x42\x4d\x17\xf0\x57\x25\x2b\x09\x28\x4b\x80\xbc\x57\xc6\x9a\x24\x6a\x97\x73\xe2\x7d\xbd\x36\xd2\x50\xa3\x3e\x81\xeb\x85\x04\xfa\xf6\x46\x65\x50\xcc\xc0\xe2\xe3\xba\x28\xef\x64\x99\x30\x6d\x90\x75\x02\x9d\xaa\x2c\x91\x6d\x7b\x02\xc2\xc0\x8d\x56\xf7\x70\x73\xdd\x01\x23\x31\x00\x99\x01\xa3\x34\x1a\x2d\x57\xc5\x74\x91\xc0\x50\xfe\x55\xa1\xd6\x0c\x6c\x01\x13\xc9\xe2\xc2\x67\x25\x8c\xc1\x17\x4a\x43\x86\x2e\xe1\xfb\x5c\x2d\x95\x15\x56\x61\x18\xc9\x8c\xcb\x0a\x7d\xd4\x73\xa5\x25\xd2\x18\x95\x49\x14\x96\x29\xe3\x6c\x62\xa7\x9c\xe5\x9f\x8a\x09\xf4\xba\x24\xc0\x27\x20\x81\xc1\x8a\xa4\x88\x3c\x81\xae\x9c\x89\x2a\xb7\x86\x5e\xc7\x71\x02\xbd\x19\xe8\xc2\x82\x59\xc9\xa9\x9a\x29\x99\xed\x59\x44\x8e\x6b\x79\x6f\x61\xa1\xe6\x0b\x69\x2c\xac\x4a\x55\x94\xca\x6e\x48\x1d\x28\x13\x14\x64\xce\x06\x1f\xd0\x13\x18\x79\xcf\x2b\x6d\x55\xce\xf6\xf8\x57\x20\x74\xc6\x7c\xf8\x5f\x51\xd0\xc8\x8c\x39\x46\x42\x4c\xef\xa0\xd0\xac\x90\x33\x9a\xec\xd9\x81\x46\xc6\x71\x0b\xb2\x9a\xed\xff\x7e\x1b\xe2\x9b\x44\xd1\x50\xda\xaa\x24\x8c\x90\x2a\xc2\x42\x94\x96\x65\x51\xe2\xc1\x55\xfb\x8f\x4f\x83\x0f\xa3\xf1\x30\x6d\x77\x3e\xa6\xdd\x16\xf4\xfa\x5f\xda\x97\xbd\xee\xb8\x73\xd9\x4b\xfb\xd7\xe3\x7e\xfb\x2a\x6d\x41\x7f\x30\xee\x5d\xa7\x57\xa3\x16\xdc\xf4\x2f\xfa\x83\xdb\xfe\x18\x99\xc6\xbd\x6e\x14\x1d\x1f\xc3\xb7\x6f\x51\x94\x17\x53\x91\xc3\x8c\x30\x01\xa7\x10\x7b\xb7\x63\x7f\x6e\xe4\x8a\x4e\x4f\xe2\x40\x98\x17\xf3\x31\x46\x55\x4d\x89\x7a\x56\xe9\x29\x45\x1f\x1a\x88\x5f\x23\xe6\xb2\xc9\xbe\x95\x9c\x3b\xa4\x6c\xb8\x6f\x97\x83\xf3\x71\x7f\x70\xdd\xeb\xa0\x3d\xf1\xfb\x18\x92\xc4\xeb\xc3\x2f\xf1\x19\x3f\xc7\xc0\xff\xb6\x62\xa4\xce\xea\x1a\xd7\xa2\xd4\x3f\xa7\xef\xb6\x3d\xec\xf7\xfa\xe7\xbf\xaa\xf0\x87\x2c\x27\x85\xf9\x49\x1f\xbf\xa4\xc3\x0f\x83\xd1\x2f\x39\xa9\xcc\x58\x52\x5e\x6b\x0a\x51\xb0\x41\x48\x04\x75\x04\x03\xb0\x9b\x95\x0c\xe7\x70\x7a\x0a\x47\x56\x4c\x72\x79\xc4\xd8\x73\xc7\x09\x8a\xa9\x0b\xc6\xfa\x3a\x85\x8b\xf4\xbf\xa3\xaf\xef\xbe\xf9\xa3\x3b\xd7\x58\x00\x5f\xe0\x5b\x34\x88\x92\x4c\xe6\xfd\x7e\x93\xde\xa4\xdd\x18\xdf\x20\x36\x3e\x97\x6a\x29\xca\x0d\xe3\x8e\x19\x02\x37\xb5\x0b\xa5\xe7\x8f\xb8\x6f\x07\xc3\x0b\x8c\x78\xcc\xdc\xc8\x85\x58\x5e\x08\xac\x2f\xf1\x43\x62\x3b\x90\x7a\x5b\x51\x41\xd0\x52\xdc\x7f\x27\xb2\x87\x82\x3c\xb0\x9d\xa0\x2b\x71\x0f\xba\x5a\x4e\xb0\x2b\x62\x9f\x60\x7a\x91\xe7\xc5\x7a\x27\xc6\xb5\xaf\xc7\x62\xc8\x9e\x74\xe8\xc5\xdc\x32\x11\xf6\x0e\x13\x02\x53\xeb\x8a\xc8\xda\x1e\x9e\x7f\xd9\x45\x28\xb3\x95\x9d\xea\x62\x0d\x1c\x24\x5b\x38\x03\x1a\x4c\xf4\xdb\xb7\xa6\xa7\xe2\xbe\xe4\xaa\xd8\x0b\xf8\x67\x10\x10\x7a\xc2\x53\x02\xfe\x85\x02\xd4\x6c\x4b\x82\x59\x8c\x63\xc0\xcc\xd7\x0e\x34\xb6\x16\xec\x17\x7a\x77\x46\x4d\xa1\x96\x55\x0c\xdd\xd8\xc7\x6e\x2b\xdc\xe1\x70\x85\xaf\xf3\x46\x7c\x9e\x5e\x63\x57\x09\x21\x6e\x06\x93\xf5\x41\xa6\xff\x75\xda\xc3\x2e\xb1\xf9\x14\x6f\xd9\x1c\xb4\xb8\x65\xfc\xe3\xc9\x0f\xbd\xe9\x16\xfa\xc8\x86\x24\xc3\xb2\x28\x25\x01\xc0\x35\xbe\xe5\xa3\x2c\x1e\x90\x85\xc1\x71\x76\x12\xaa\xb7\x9e\xd2\x43\xc3\x9d\x9f\x9d\x6e\x8f\x9b\x1c\xa7\x7a\x89\x38\x8f\xb8\x98\xc6\xa5\x5c\xe5\x9b\x46\xfc\xa0\x53\xc6\xbe\xf2\x0e\x7a\x43\xd3\xc6\x41\x84\xc7\x25\x2c\x2b\x9c\x11\x13\x9c\x19\x85\x3e\x9e\xe4\x42\xdf\x25\x44\xf5\x81\xbe\xe1\xc8\xca\xd4\x14\x3b\xbe\xc1\x6f\xca\x2a\x07\x0d\x30\x96\x86\xc0\x5a\xd9\x05\x32\x41\xb1\xd6\x34\x63\x0f\x7a\xbd\x07\x49\x86\xc5\xd6\xb9\xd0\x04\x1b\x71\xa7\x66\xd4\xde\x78\x4b\xe2\xe6\x0b\x71\x78\x62\x40\xbc\x2a\x16\x3c\x48\x3d\xdc\xb7\xea\x5a\x0e\xa1\x56\xdc\x49\x4e\xf1\x4c\x95\x18\x21\x65\xe5\x92\xc6\x7c\x6d\xda\x39\x01\xc2\xf3\xaf\x71\x8d\xa8\xc9\x98\x61\xec\x5c\x9f\xe0\xb1\xfb\x90\xf1\x19\x9b\xfc\x6c\x9a\x22\xc6\x28\x6e\x4e\xf2\x83\x88\x21\xf7\x4a\xca\x3b\x40\xd1\xef\x9c\x59\xb3\xb2\x58\xee\xe4\xfb\x58\x11\xb0\xb1\x20\x5c\xb4\x7c\x19\x0c\xdb\xfd\xf3\x94\xea\x80\x09\x5b\xf0\x96\xff\x8e\x28\x95\xac\xd3\x1c\x35\x23\x66\x47\xdd\xdc\x80\x93\xb9\xb4\xba\xde\x97\xdf\xee\xec\x08\xd9\xf3\x13\xa5\x11\xf7\x0b\x36\xc7\xec\x2d\x05\x27\x6e\x34\x38\x8d\xcd\x2d\xe3\xf3\xb9\x0c\x73\xdd\x27\x9d\x93\xe8\xdd\xa6\x86\x81\x8b\x88\x70\xb6\x51\xc1\xe5\xb8\x5b\x21\x0c\x7f\x73\x68\x24\xa5\x8f\xf2\xc4\xf5\x85\x4f\x41\x08\x3b\xfa\x26\x2c\x45\xbb\x3d\x8c\x63\x45\x7e\x52\xb7\xa4\x53\x26\xdc\x6f\x29\xfc\x9a\xfa\xa4\xcc\xcd\x93\x71\x0e\xfd\x66\xd4\x19\x0c\xeb\x81\x66\x0d\xcd\x10\xda\x30\x13\xb7\x81\x7d\x14\x53\xf7\x9a\xc7\x2c\x0e\xdd\x8e\xd0\x54\x10\xc5\xc4\x0a\xf4\xcc\xd9\x35\xc3\xc6\x8a\x52\x7d\x78\x9d\x07\xc9\xfe\x06\x46\x9f\x98\xb1\xf1\x7f\x23\xed\x7e\x22\xfe\xf3\x8a\x4c\xec\xef\x55\xb5\x7c\x1c\x88\xce\xae\xea\x6e\xa5\x9b\x90\xa1\x3e\x26\x1b\xac\xb5\x75\x82\x73\xef\x8e\xb2\x74\x64\x9c\x4a\xb8\x93\x9b\x24\x4c\x3c\x2a\x95\x87\xd3\xce\x4d\xcc\xdd\x01\x4b\x7b\xa1\xae\x3f\x17\x2b\x1e\xef\xec\xfc\x15\x05\xed\x77\xae\x8c\x97\x0a\x4f\x97\x72\x59\xfc\x90\xd9\xa3\xba\x49\xaf\x1e\x27\xf3\xb0\x09\xed\x2c\x63\x13\x70\xed\xbd\xf5\x7b\xc5\x48\xda\x03\x3c\xcf\x20\xa9\xdd\xad\x0f\xae\x96\x8b\xfb\x2b\x6d\xb8\x59\xf1\xba\xfe\x29\xf4\xec\x9f\xd6\xfe\xf1\x6a\xe4\xc6\x2d\xaa\x6b\x3d\xc2\x96\x03\x18\xf7\x7f\xa4\xa9\xf5\xf8\xe7\x48\xd9\x08\x24\x8d\xbd\x37\xf1\x73\x84\x61\xa1\xa2\xfb\x83\xdf\x58\x10\xb1\x87\x9d\x1d\xca\x39\x5e\x25\x71\x0e\xdb\x05\xa2\xca\xed\x4e\x74\xb5\x13\x30\xf5\x17\xbe\x70\xc6\x2b\xe5\x96\x58\x86\xfe\xed\x3a\x88\xe2\x01\xee\x6f\x8e\x00\x29\xef\x29\xb5\xdb\x24\x88\x19\xf1\xa1\x54\x59\x72\x39\x8a\x65\x81\x57\x26\xbe\x72\xe2\x65\xf2\x50\x73\x7f\x3a\xc6\xa3\x5a\x86\x71\xdb\xdb\x0b\x64\x93\x1a\x16\xc2\x86\xa4\x7b\xf5\x74\x6c\xf6\x37\x44\x94\xb8\xdd\x15\x77\x75\x52\x13\xf3\x2a\xcd\x01\x53\x35\x95\xdf\xb7\xbb\xae\x08\xea\x17\x82\xe3\xc3\x45\x2c\xa8\x76\xb9\x11\xd5\xc2\xc3\xe4\x7e\xbb\x4b\xf6\xed\x1c\x0b\x5c\xff\x7f\xc8\x17\xcd\xad\x57\x7f\xbb\x73\xdd\xfb\x92\xc6\x51\xbd\x1c\x3d\x26\xf7\x84\xb6\xe0\x1d\x5b\x7e\x81\x16\x61\xc9\x19\x35\xd7\x6a\xb6\x71\xd6\xd4\xac\xa3\xe9\xc1\xf4\xc9\x9e\xc4\xf4\x8f\xcf\x3d\xd7\xaf\x1f\x08\xf5\x8e\xb8\x42\xc3\x9b\x5e\x7a\xe2\x11\x91\xf9\x8d\x9d\xb1\x41\xf7\x29\x7e\xf4\x73\x27\x5c\x24\xdc\xfc\x69\x20\xe0\xde\x14\xf4\xdb\x07\xd1\xfd\xc9\x45\xf0\x27\xc9\x43\x63\xb0\x2f\xf3\xa5\x3e\x14\x44\x93\xd1\xb9\x20\xd8\x6a\xb7\x59\x41\xc3\x99\xd4\x0c\x97\x0e\xe3\x8d\xca\x10\x9d\x6d\xef\x17\x43\x96\xc5\x3f\x08\x35\x25\xa8\xc6\x12\xae\xff\x2a\x93\xf4\x84\xbb\x4f\x5e\xe8\x39\xc1\xd9\x05\x05\xda\x7a\x43\xb2\x8a\xca\xa2\x99\xb8\xfc\xa1\x13\x7c\xe1\x99\x48\x22\xa4\x47\xb4\x96\x0d\xf1\x01\x35\x8b\xa2\xca\x33\xda\x1f\x4b\x29\x56\x28\xd3\x55\xd7\x31\xbb\x9e\x25\x2f\xd6\xac\xbb\xfc\xf9\x1f\x40\xf8\x57\x80\x43\xf5\x53\x9b\x54\x9c\xb8\xa3\x8f\x78\x1d\x68\x5f\x5e\x1e\xb9\x16\xd5\x8c\xfe\x0e\x00\x00\xff\xff\xf7\x27\x5f\xf4\xc4\x12\x00\x00")

func bambooScriptsConsumeLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsConsumeLua,
		"bamboo-scripts/consume.lua",
	)
}

func bambooScriptsConsumeLua() (*asset, error) {
	bytes, err := bambooScriptsConsumeLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/consume.lua", size: 4804, mode: os.FileMode(420), modTime: time.Unix(1448248180, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsEnqueueLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\x51\x73\x1a\x37\x10\x7e\xe7\x57\xec\xdc\x4c\x27\xc7\x14\x5f\x43\xfa\x66\x8c\x67\x48\xb9\x71\x48\x6c\x9c\x02\x71\xdb\x78\x3c\x8c\xe0\x16\x90\x7d\xe8\xae\x92\x70\x4a\x3b\xfd\xef\xdd\x95\x4e\xc7\x41\xd2\x76\x12\x5e\x38\xad\x56\xfb\xad\xbe\xfd\x76\x75\x76\x76\x7f\xdf\x3a\x3b\x03\x54\xbf\xef\x70\x87\xfc\xf9\x0e\xf7\xe6\x1c\x2e\x94\xb9\xe4\x15\x28\x5a\x8c\xc5\x16\x4d\x29\x96\x08\x3b\x95\xa1\x86\x4f\x1b\xb9\xdc\x80\x3b\x01\x99\xb0\x02\xf0\x0f\x69\xac\x49\xf8\xc0\x40\xaf\xf9\x78\xa9\x65\xa1\xa5\xdd\x5f\xc2\xc5\x63\xb1\x90\x19\xfd\x3f\x21\xaf\x9e\x45\x7e\x09\xf7\xcd\x45\x92\x24\x0f\x0e\xca\xfd\x68\xa3\x03\x64\x87\x52\x48\x6d\x40\x68\xe4\xd5\x0e\x0d\x68\x2c\x35\x1a\x54\x56\xaa\x35\x08\x78\x5b\x2c\xa0\x58\x3c\xe2\xd2\x32\x2e\x07\x98\xa0\xdd\x69\xce\xf7\x25\xc8\x15\x10\x2c\x88\x5c\xa3\xc8\xf6\x3e\x3f\xcc\x12\xe8\xf2\x8e\xc8\x32\xfa\xe6\x13\xa9\xd6\x85\xa6\x03\xa3\xf1\xdd\xe0\x7a\x34\x9c\xbf\x1f\x4c\x06\x37\xe9\x2c\x9d\x4c\x79\xf7\xe1\xa1\xd5\xca\x8b\x25\xe5\xb2\x52\xc4\x00\xf4\x21\xaa\x78\x8a\x2a\xbb\xc1\x92\xad\xe7\x51\x70\xcc\x8b\xf5\xfc\x93\xd0\x8a\xac\xab\x9d\x5a\x5a\x59\x28\x88\x89\x3c\x23\xd6\xd8\x6e\xf1\xfd\x34\x66\xd2\x24\xe4\x17\xfb\xaf\xeb\xdb\xab\xf9\x2f\x83\xc9\x78\x34\xbe\xea\x40\x74\x11\x11\x1d\x15\x1c\x7d\x44\x97\x6e\x1d\x81\xfb\xab\xe3\xa0\xca\x9a\x80\xaa\xb0\x72\x89\x5f\x07\x39\xbe\x9d\x8d\x7e\x4a\xbf\x15\xf1\x19\xf5\xa2\x30\x5f\x09\x79\x97\x4e\x5e\xdf\x4e\xbf\x09\x53\x9a\x39\x72\xa9\x1a\x80\x14\xd8\xec\x72\x1b\xe0\xb8\xf2\x60\xf7\x25\x06\x3b\xf4\xfb\xf0\xc2\x8a\x45\x8e\x2f\x40\xa8\x0c\xbc\x39\xa1\x30\xcd\xc0\xca\x50\xc8\x77\xe9\x6f\xd3\xfb\x6e\x5d\xec\x20\x5d\xda\xb1\x85\xb1\x9a\xe4\x16\x0f\x26\x57\x77\xe4\xd2\x0e\xf9\x1c\xf9\xa8\xdd\x76\x81\xfa\xd4\xc7\xc9\x9e\x93\xeb\x83\xdb\x79\x45\x00\x24\xaa\xf7\x5a\x6e\x85\xde\x3b\xf9\xfa\x0e\x8a\x3f\x4e\xd1\x86\x53\x4f\xde\x46\x87\x28\x35\x62\x83\x15\xc6\xdc\xfc\xfc\x21\xfd\x90\x0e\xa3\xaa\x41\xa1\x58\x81\xdd\xa0\x0b\xe2\x1a\x30\x7e\x73\x23\xca\x3a\x06\x4b\xff\x34\xc0\xdb\xdb\xd7\xd3\xa8\x61\x70\xe9\x85\x2b\x6f\xcd\xba\x53\x31\xd4\xab\x4c\x42\xaf\x97\xcd\xdb\x39\x2a\x93\x35\x5a\xe5\x2e\xda\x6e\xf7\xdc\x75\x6e\xc4\x13\x52\x33\xba\x5d\xce\x4a\xe4\x79\xa3\x33\xa9\x89\x35\x15\xd9\xa2\x36\x07\x56\xe6\x2e\xe3\x3e\xfc\xf5\x77\xc0\xe2\x76\x79\xd9\x6b\xad\xa8\xc0\x92\x3e\x7f\xec\x78\xf4\xac\x70\xc5\xf5\xc8\x52\x19\xd4\x36\x0e\xe7\x3b\x9e\x54\xf9\x40\x79\xb0\x13\x87\x50\xdf\x77\x7d\x6d\x29\xaf\x19\xb1\xe3\x33\xe7\xac\x68\xa8\xfc\xe0\xa6\x88\x4f\xc8\xb0\xc7\x01\xfa\x4b\x77\x0c\x38\xee\x9e\x34\x33\x62\x05\xdf\xc1\x2b\x16\x55\xb7\xcd\xd4\x2b\x87\x4a\xbc\x71\xff\x8f\x14\x05\xa7\x62\x1f\x00\x1f\xbf\x44\xc1\xb9\x17\x78\x2d\x2a\xe5\xc5\x1b\x66\x46\x4c\xd1\x8e\xe4\xec\x5b\xc7\x09\x7f\x4e\xc3\x2f\xdf\xc7\xd1\xe7\x83\x2a\x3a\x6a\x94\xe5\x4e\x6b\x9a\x90\x73\xb3\x2c\x34\x36\xaf\xe6\x63\x91\x4b\x1e\x47\x7f\xba\xdd\xa8\x53\x89\xad\xe3\xb5\xe0\x6e\xea\xa3\xf8\x71\x4e\xc7\x9b\xa7\xd2\x5f\x47\xd3\xd9\x94\x4f\x91\x7b\x9b\x39\xa9\x83\x7b\x7f\xd7\x72\x5d\xd7\x6c\xc7\x69\xb0\xe5\xd0\x31\x97\xfd\x93\xed\x53\x36\x61\x5c\x58\xff\x20\xf1\xa0\x97\x16\xb7\x24\xdf\x16\xd4\x3f\x56\xf3\x40\xf9\x24\x83\x07\x6c\x84\x71\x1d\x61\x78\xa6\x90\x8e\x72\x9a\x22\x54\x0a\x87\x90\x44\xbd\xff\x26\x9a\xb4\x17\x74\xe3\xf4\x1c\x5a\xab\xaa\x20\x85\x16\x0c\xb0\x81\xad\x28\x93\x56\x93\x94\x37\x37\xd3\x74\x56\x71\xd2\xa1\xd7\x91\x1e\xc9\xa7\x23\xed\xfc\xbb\x73\x14\x18\x21\x4b\xf8\x24\xab\xb1\xc2\x72\x6d\xc2\x53\x93\x45\x55\xa7\x0d\xb2\xcc\x25\x35\x1a\x12\xf1\x7e\x74\xb4\x7c\xcb\xd6\x85\x2a\x3d\xce\xc7\xc1\x70\xd8\xa8\xae\x3c\x44\xf7\x85\xee\x71\xf1\xc2\x48\xad\xe7\x65\x5d\x86\x66\xca\xc3\xf4\x3a\x54\x1c\xdc\x0b\xbb\x2d\x9e\x0f\xfc\xc4\xba\xc8\xf3\x05\x5d\xb9\xdd\xac\x1f\xc3\x43\x51\xa2\x16\xee\x61\x58\x09\x99\xf3\xd3\xeb\x71\x4e\x7b\xa0\x42\xff\x9f\x02\x7d\xde\x09\xce\xc3\x15\xad\x72\xe9\xf6\x5a\xff\x04\x00\x00\xff\xff\xc5\xd9\xa1\x88\xcd\x08\x00\x00")

func bambooScriptsEnqueueLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsEnqueueLua,
		"bamboo-scripts/enqueue.lua",
	)
}

func bambooScriptsEnqueueLua() (*asset, error) {
	bytes, err := bambooScriptsEnqueueLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/enqueue.lua", size: 2253, mode: os.FileMode(420), modTime: time.Unix(1448244065, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsFailLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\x6d\x53\x1b\x37\x10\xfe\x7e\xbf\x62\xe7\xfa\x21\x76\x63\xdc\x90\x74\xfa\xc1\xc5\xcc\x38\xf8\x42\x1c\xc0\xb4\x76\x08\x6d\x33\xd4\x23\xdf\xe9\x40\xc9\x59\x72\x24\x5d\x0c\xff\xbe\xab\xb7\x7b\xc1\x06\x42\x3d\xc3\xd8\x48\xbb\xcf\xbe\x68\x9f\xdd\xdd\xdb\xfb\xfc\x39\x8a\x72\xc2\x0a\x38\xe0\xea\x10\x7a\x70\xf0\x45\x2c\x59\x76\x08\x07\x19\xd1\x54\xb3\x15\xc5\x9f\x1b\xc2\xf4\x62\x49\x54\xf5\x9b\xde\xae\x0f\xa3\xe8\x84\xde\xa9\x41\x04\xf8\xe1\x6a\x00\x53\xb2\xa2\x6a\x4d\x52\x0a\x25\xcf\xa8\x84\xcd\x0d\x4b\x6f\xe0\x5b\x49\x4b\x0a\x88\x45\x80\xde\x32\xa5\x55\x3f\x8a\x46\xf2\xda\xeb\x59\x5b\x03\xf8\x20\x96\xc0\x32\xca\x35\xcb\x19\x95\x7d\x7b\x15\xcc\x0f\xe0\xa8\x94\x12\xef\xaa\x13\x20\x0a\x2e\x38\xbb\x85\x8b\x8f\x47\xa0\x68\x2a\x78\xa6\x40\x31\x8e\x96\xe9\x5a\xa4\x37\x7d\x98\xd1\x6f\x25\x93\x34\x03\x2d\x60\x49\x2d\x5c\xf8\xac\x89\x52\x78\xc1\x38\x64\xe8\x17\xde\x17\x6c\xc5\x34\xd1\x4c\x70\x05\x22\x87\xd3\x12\x1d\xe5\xd7\x8c\x53\x94\x51\xe8\x14\x82\x65\x4c\x39\x9f\xaa\x3c\xf4\x20\xa4\x61\x60\x8d\x99\x18\xd7\x92\x09\xc9\xf4\x1d\x74\xe6\x47\xef\x93\xf1\xc5\x69\x32\x76\xd1\x77\x61\xc3\x8a\x02\x1d\x01\x94\x24\x05\x1a\x6d\x05\x08\x2f\xa1\x53\x01\xc3\xcf\x60\x1e\xa3\x94\x54\xfd\x1b\x4c\x74\xa3\x68\x46\x75\x29\x4d\x92\xf7\x81\xe5\x56\x02\x63\x48\x05\xe6\x25\xd5\xc5\x5d\x3f\x4a\xa4\x14\x12\xaf\x2f\xa6\x27\xd3\xf3\xcb\xe9\xe2\xc3\xf9\xdb\xc5\x64\x1c\x45\x7b\x7b\x57\x57\x51\x54\x88\x14\xcd\xe6\x1c\x1f\x08\x86\x10\x1b\xf5\xd8\x1f\x2a\xba\x36\x47\x83\x38\x48\x15\xe2\x7a\xb1\x21\x92\xe3\x69\x5e\xf2\xd4\xe4\x05\x3a\xf8\xb0\x8a\x5c\xd3\xae\xf5\x5b\xda\x7c\xa0\x5c\xc7\xfd\x3a\x3d\x3f\x5e\x5c\x8e\x66\xd3\xc9\xf4\xb8\x07\xf1\x41\x0c\xfd\xbe\xb7\x85\x3f\xe2\x43\xfb\x7f\x0c\xf6\xab\xc2\xa1\x3c\x6b\x1a\xfc\x4e\xe5\x52\x28\xfa\x3c\x9b\x9f\x92\xd9\xdb\xf3\x79\xf2\x7f\x6c\x32\xb5\xa0\x26\x63\x0d\x83\x08\xac\xca\x42\x07\x73\x26\xdd\xa0\xef\xd6\x34\x9c\xc3\x70\x08\x2f\x34\x59\x16\xf4\x05\x10\x9e\x81\x3b\xee\x23\x4c\x13\x18\x6b\x68\x08\x27\xc9\xdf\xf3\xcf\xfb\x57\xfe\xc8\x96\x38\x9e\x8e\x66\xc7\x9f\xea\xd3\x4c\x97\x3a\xe5\x62\x83\x17\x5a\xf0\x72\xb5\xa4\xb2\x63\x25\x5e\x5f\x75\xbd\x48\x5d\x13\xf7\x65\xde\xb4\x65\xb0\x44\xb6\x44\x7e\x45\x11\xac\x94\x06\xc6\x10\x38\x92\x5c\xdf\x50\xde\x42\x7e\xf3\xdb\xab\x57\x60\x22\x08\xd2\x16\xed\xbe\xb0\x33\xf1\x1a\x1a\xa1\x7e\xdd\x08\xf9\x95\xf1\x6b\x4c\xd7\xd0\xc4\x8d\xa9\x36\xc5\x64\x12\x7f\x79\x3e\x3b\xc1\x6a\x88\xf1\x6a\x6f\xcf\x90\x5b\x21\x14\xd1\x70\x43\xbe\x53\xa4\x01\x82\x22\x67\x55\xb9\xa2\x59\x13\x8b\x4a\xf5\x00\x56\x32\x9b\x7b\xac\x4b\x2b\x07\x93\xb1\x0a\x9a\x9e\x0b\x3b\x34\xdf\x8d\x26\x48\x41\xa3\x68\x34\xdf\x39\xb9\x3f\x0d\x21\x9f\xd6\xad\x18\x1c\x1b\xdd\x79\x7a\x43\xb3\x72\x4b\x7d\x45\x6e\xbf\x98\xd8\xb6\xd5\xcf\x46\x7f\x21\x03\x83\xd3\x67\xe4\x16\xdc\xdb\x98\x06\x63\x55\x48\x51\x88\x4d\x1d\x3e\x22\x79\x5f\x76\x20\x85\x38\xb6\x90\x42\xa3\xa8\xd0\x02\x1c\x9a\xd8\x02\x72\xfe\xd4\x07\xb6\x2c\xbd\x82\xab\xe5\x28\x72\xdf\xa8\xea\x48\xb6\xc6\xbb\xa2\x13\xff\x33\x3f\x3a\x9f\x25\x71\xaf\x7a\xf1\x9e\x53\xb6\x05\x16\x54\x5c\xc1\x20\xa3\x02\xb5\x2a\xde\x98\x22\xb2\xa4\x0a\x9d\xa5\x13\xff\x21\xc5\x77\xec\xac\x99\xc1\x01\x2e\x34\xe4\x02\x47\x46\xdf\x8e\x81\xc9\x78\xe0\x78\xeb\x78\x63\x5c\xb7\x49\xf7\xa7\xc1\x87\x16\x4f\x9d\xbb\xd6\xec\x42\xd2\x75\x71\xd7\x89\xdb\x8d\xd0\xf7\x83\x0f\x4d\x73\x66\x04\xd8\xf6\xdc\xbf\x07\x6c\x6b\x1c\x93\xfd\xd3\xce\x8f\xb9\x99\xd1\x95\xc0\x4a\x36\xde\xe7\x52\xac\xc0\xd7\xbb\x83\x7b\x44\xd5\xf9\xe9\xb3\x3a\x4b\xce\x76\xe5\xf4\x71\xcb\x13\x9e\x4a\xba\x32\xc3\xd0\x3f\x3e\x32\xa9\xe4\xda\x5c\xcd\x69\x7d\x68\x46\xcb\x23\x40\x7e\x1e\x84\xf2\x69\xf4\x8e\xd6\xc3\xbf\x3f\x4e\x3e\x1a\x17\xd1\xb3\x9e\x9b\x1a\x46\x3c\xee\xda\x87\xaf\xb5\x1b\xbd\xa2\x01\xe9\x9a\x4a\x33\xe2\xf7\x67\xf3\x1a\xaf\x39\x93\x6b\xe8\x1e\x7a\xa2\xb4\xc4\x7c\x74\xc2\xd9\xcb\xfd\x6e\x5b\xd8\xc4\xe6\xc8\xd2\x14\x0f\xfd\xb4\xfb\x54\x06\x13\x86\x9e\x4a\x20\x99\x5d\x0e\xea\x49\x8d\xa5\xeb\x88\xe6\x5e\x51\x41\x46\xd7\x18\x81\x69\x70\x38\x8d\x90\x86\x0b\x77\xfd\x64\x5a\x9b\x5c\xde\x9d\x57\x9f\xd6\x4a\xd0\x25\xb4\xa1\x57\x91\xa9\xca\xe7\xe1\xb0\x71\x5f\x31\xea\x61\x57\xfc\xed\x99\x29\x53\x0c\xb3\x19\x59\x0f\xe9\x62\xcb\xd7\xb0\xc1\xec\x65\x3f\x00\xf5\x40\x63\x18\x8d\xc7\x26\x10\xe7\x56\xaf\x9a\x69\xa1\x96\x7f\x7f\x44\x77\x9c\x9c\xfa\x52\x40\x31\x5a\x28\xfa\xbc\x80\xee\x6d\x58\x88\x44\xb1\xad\x3d\x23\x22\xbc\x7d\x4b\xd2\xaf\x22\xcf\x91\x38\x88\x81\xcb\xa3\x6f\x52\x76\x8c\xd3\x5b\x6d\x57\xb3\x61\x3d\xa8\xef\x6d\x69\x9d\xed\x35\xad\xfb\x03\xb9\x52\x61\x94\xf4\x2a\x23\x75\xba\x9e\xdf\x79\xdc\xdc\x7c\x81\xeb\x2f\xd5\x4f\x56\x66\x5a\x30\xec\x1c\x0b\xbf\x06\xb6\xa8\xd9\x62\xba\xd8\x70\x2a\x63\x5b\x94\x2d\x15\xdc\x13\x63\x53\x95\xf7\x0e\x73\x82\xcf\xb7\xab\xcd\x4f\x85\xf7\x0f\x6d\x5d\xe3\xea\x4f\xd1\xe4\x2f\xae\xf7\xe6\x08\x83\xc6\xfa\x68\xa5\x7a\xfc\xd6\x2a\x80\x0e\x56\x4b\x41\x3d\xb4\x1a\x96\x1f\xc9\xf5\xbc\xd1\x5a\xa9\xac\x3a\xab\x7f\xf6\x49\x6e\x7c\xc5\x26\x49\xf0\x8f\x0b\x58\x09\xfc\x16\xa5\x56\x9a\x38\xba\x9b\x01\x5d\x91\x04\x45\x43\x10\x36\xe1\xe6\x7f\xcc\x35\x8e\xdf\x80\xe7\xbd\xec\x37\xcb\xc7\xce\xf8\x76\x8a\xe7\x47\xa3\xd9\xb8\xf6\xca\xb9\x83\x19\x76\xb2\x07\xa6\x61\x56\x29\xac\x37\xde\x56\x40\x1c\xdd\x6a\x64\xc0\x41\x98\x9a\xb1\x75\xe3\x67\xe1\x7e\xf4\x5f\x00\x00\x00\xff\xff\xcb\x85\x27\x09\xd3\x0d\x00\x00")

func bambooScriptsFailLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsFailLua,
		"bamboo-scripts/fail.lua",
	)
}

func bambooScriptsFailLua() (*asset, error) {
	bytes, err := bambooScriptsFailLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/fail.lua", size: 3539, mode: os.FileMode(420), modTime: time.Unix(1448249569, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsMaxfailedLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x52\x4d\x8f\xda\x30\x10\xbd\xe7\x57\x8c\xdc\x4b\x90\x48\xd4\x5e\x57\x80\x14\x69\x53\x84\x0a\xb4\x82\xd5\xb6\x15\x42\xc8\x24\x93\xac\xb5\x8e\x8d\x6c\xa7\xb4\xff\xbe\x9e\x7c\xc1\x76\x29\x73\x71\x32\x1e\xbf\x79\xf3\xe6\x45\xd1\x6e\x07\x15\xff\x5d\x70\x21\x31\x87\x89\xb2\x33\x18\xc3\x6e\x32\xa4\x66\xfb\x20\xd8\xa2\x03\x6d\x60\xee\x0f\xf7\x82\x54\x0e\xaa\xae\x8e\x68\x40\x17\x40\x65\xb5\x41\x0b\x47\x2c\xb4\x41\x50\x1a\x2a\x3a\x0d\x3a\x23\xd0\xc6\x41\x90\x1a\xa3\x8d\x7d\x80\xc5\xfa\x39\x59\x2e\x1e\x0f\xdf\x92\x4d\xb2\x4a\x9f\xd2\x4d\x10\x44\xd1\xde\xe3\x4b\x9d\x71\x09\x85\xe2\x15\xc2\x14\xd8\xd0\x9b\x75\x37\x16\x4f\x94\x7f\x60\x7d\xa9\xd4\xe5\xe1\xcc\x8d\xf2\xd9\xa2\x56\x99\x13\x5a\x41\x58\xa1\xb5\xbc\xc4\x51\x00\x3e\x0c\xe6\xc2\xc6\xbe\x2e\x6c\xbf\x96\x5f\xe7\x87\xef\xc9\x66\xbd\x58\xcf\xc7\xc0\x26\x0c\xe2\xb8\x6b\xe8\x3f\xd8\xac\xf9\x67\xd0\x1c\x03\x0e\xaa\xbc\x6f\xa8\xac\x6f\xf5\x25\xfd\xb9\xdd\x7d\xda\x77\xa9\x8b\x68\x53\x50\x42\xf6\x95\xaf\x6f\xf2\x96\x00\x89\x3e\xc1\xaf\x92\x1f\x9f\x93\xc5\x32\x7d\x64\x10\x45\xb0\xba\xad\x22\x97\x52\x9f\x31\x27\x69\xe0\xc3\xcd\xa0\x9b\x45\xd1\x6f\xa2\xeb\x75\xe2\xc6\x0f\xe3\x3c\xd8\x99\x5b\x38\x19\xfd\x4b\xe4\x98\x8f\x7d\x6f\x07\xc2\xc5\x77\xd0\x84\x87\xe2\x47\x89\x71\x89\x4e\x85\xc9\x66\xfe\x3c\x82\x19\x7c\x24\x7c\xd5\x48\x79\x3d\x10\x5d\x93\x04\x94\xff\x2f\x8d\xf7\x2c\xb2\x17\xcc\x5e\x7d\x29\x27\x32\x20\xfc\x94\xdd\xe8\x71\x83\x44\x14\x74\x9b\x08\x07\xac\x11\x4c\x1b\x61\x2f\x44\x28\xfa\xcd\x87\xec\x9d\x9b\x20\x6a\xf7\xe7\xb4\xf5\xce\x53\xe5\x35\x54\x1c\x0f\x08\x43\x30\x22\xa2\xb4\x03\xae\x40\x28\x87\xa5\xa7\xc3\x46\x43\x9d\xf7\x6f\xed\x1d\xd6\xda\x07\xc9\xc2\x07\x83\x27\xf9\xe7\x46\xe7\xee\x15\xf9\xe5\xe2\x3d\x6f\x06\x19\xb2\x6d\xfa\xc4\xc6\x57\xa6\x18\x5f\xc4\xea\x0c\x76\x77\xcf\x9b\x96\x04\x89\x9c\xd5\xc6\xa0\x72\x9d\x70\x77\x9e\xbd\x21\xde\xb2\x98\xff\xc3\x62\x14\xfc\x0d\x00\x00\xff\xff\xaf\xaf\x91\xab\xf8\x03\x00\x00")

func bambooScriptsMaxfailedLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsMaxfailedLua,
		"bamboo-scripts/maxfailed.lua",
	)
}

func bambooScriptsMaxfailedLua() (*asset, error) {
	bytes, err := bambooScriptsMaxfailedLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/maxfailed.lua", size: 1016, mode: os.FileMode(420), modTime: time.Unix(1448290779, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsMaxjobsLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x52\xd1\x8e\x9b\x30\x10\x7c\xe7\x2b\x56\xee\x0b\x91\x00\xb5\xaf\xa7\x24\x12\x55\x11\x4a\x7b\xc9\x55\xe4\x74\x6d\x15\x45\x91\x03\x0b\xe7\x9e\xb1\x23\xdb\x34\xed\xdf\xd7\x06\x43\x1a\xdd\x35\x7e\x01\x96\xf5\xcc\xec\xec\xc4\xf1\x6e\x07\x2d\xfd\xfd\x53\x1e\x35\xcc\x85\x5e\x42\x04\xbb\xb9\x2f\x2c\xf7\x41\xb0\x45\x03\x52\x41\x6e\x1f\xe6\x19\x5d\x2b\x88\xae\x3d\xa2\x02\x59\x83\x66\x6d\xc7\x0d\x15\x28\x3b\x0d\xee\x46\x12\x04\x99\x52\x52\xe9\x3b\x58\x6d\x9e\xd2\xfb\xd5\xa7\xc3\xd7\xb4\x48\xd7\xd9\x63\x56\x04\x41\x1c\xef\x2d\x22\x97\x25\xe5\x50\x0b\xda\x22\x2c\x80\x78\x2e\xe2\xeb\x1a\x4f\xae\x7a\x47\xc6\x46\x2e\x9b\xc3\x99\x2a\x61\xab\x75\x27\x4a\xc3\xa4\x80\xb0\x45\xad\x69\x83\xb3\x00\xec\x51\x58\x31\x9d\xd8\xbe\x70\x78\xbb\x7f\xc8\x0f\xdf\xd2\x62\xb3\xda\xe4\x11\x90\x39\x81\x24\xf1\x74\xf6\x85\x2c\xfb\x6f\x02\xfd\x63\xc2\x41\x51\x8d\x84\x42\x5b\xaa\x2f\xd9\x8f\xed\xee\xc3\xde\x97\x46\x83\x16\x20\x18\x1f\xfb\x5e\xfe\xa9\x6a\x07\xe6\xa4\x3b\xe8\x75\xfa\xfd\xf3\xc3\xc7\x2d\x81\x38\x86\xf5\x95\x5f\x35\x65\xbc\x53\xa8\x81\x72\x2e\xcf\x58\x39\x4b\xe0\xdd\x9b\xc7\xfd\x59\xd5\xa3\xe7\x3d\xcf\x89\x2a\x3b\x84\xb1\x50\x67\x6a\xbf\x94\xfc\xc5\x2a\xac\x22\xcb\x6b\x80\x99\xe4\x06\x16\xb3\x40\xf4\xc8\x31\x69\xd0\x88\x30\x2d\xf2\xa7\x19\x2c\xe1\xbd\x43\x17\xbd\x85\x97\x51\xdc\x4f\x37\xb8\xab\xfe\x47\xc2\x6b\x05\xe5\x33\x96\x2f\xb6\x91\x3a\x21\xc0\xec\x7c\x7e\xe8\xa4\xc7\x71\xf4\x72\x28\x84\x1e\x69\x06\x8b\xde\xcc\x8b\x04\x77\xc6\x5d\x87\xe4\x55\x7a\x20\x1e\x36\x66\xa4\x36\x8a\x89\xe6\x02\x94\x24\xd3\xfd\xe9\x10\x27\x42\x48\x03\x54\x00\x13\x06\x1b\x2b\x85\xcc\xa6\x3e\x85\xa6\xb3\x89\x1a\xe2\x82\x2e\xb0\x07\x85\x27\xfe\xe7\x0d\x5e\x7f\xcb\xe5\xe3\x92\x35\xbb\x7e\x1e\x92\x6d\xf6\x48\xa2\x29\x06\xd1\x68\x92\x0f\xd3\xcd\xcd\x16\x83\x00\x67\x6d\xd9\x29\x85\xc2\x78\xc3\x6e\x5c\xbb\x12\x3d\x28\xc8\xaf\x14\xcc\x82\xbf\x01\x00\x00\xff\xff\x0b\x12\xe7\x70\xce\x03\x00\x00")

func bambooScriptsMaxjobsLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsMaxjobsLua,
		"bamboo-scripts/maxjobs.lua",
	)
}

func bambooScriptsMaxjobsLua() (*asset, error) {
	bytes, err := bambooScriptsMaxjobsLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/maxjobs.lua", size: 974, mode: os.FileMode(420), modTime: time.Unix(1448290824, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsRecoverLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x04\xc0\xcd\x09\xc3\x30\x0c\x05\xe0\xbb\xa6\x78\x0b\x38\x4b\x14\x7a\xe8\x0a\x21\x07\x19\x3f\x88\x69\x79\x02\x59\xfd\xdb\x3e\x5f\x6b\xfb\x6e\x76\x9f\x1a\x70\xfd\xf1\x88\xbe\x50\xa7\x17\x4e\xff\x10\x9d\x14\xbc\xbb\x46\x88\x63\xb3\xdb\x8b\x2e\xf0\x57\xe9\x62\xbc\x17\xbe\x91\x4f\x26\xa8\xca\xc9\xb5\x99\xb5\x76\x1c\x76\x05\x00\x00\xff\xff\x0e\xef\x90\xf3\x55\x00\x00\x00")

func bambooScriptsRecoverLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsRecoverLua,
		"bamboo-scripts/recover.lua",
	)
}

func bambooScriptsRecoverLua() (*asset, error) {
	bytes, err := bambooScriptsRecoverLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/recover.lua", size: 85, mode: os.FileMode(420), modTime: time.Unix(1448250260, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bambooScriptsTestLua = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\x28\x4a\x2d\x29\x2d\xca\x03\x52\x29\x99\xc5\x7a\xa9\x45\x45\xf9\x45\xf1\x45\xa9\x05\x39\x95\x1a\x4a\xbe\x8e\x11\x5e\xfe\x4e\xc1\x4a\x9a\x5c\xe8\xca\x8a\x4b\x12\x4b\x4a\x8b\x61\xea\x6c\x4a\x52\x8b\x4b\xec\x14\x20\x82\x28\xaa\xab\x0d\x6a\x91\x79\x4a\x60\xe3\x95\x74\x14\x94\x92\x4b\x8b\x4b\xf2\x73\x75\xc1\x7c\xdd\xfc\xa4\xac\xd4\xe4\x12\xa5\x5a\x2e\x88\xe1\xc9\x89\x39\x39\x1a\x4a\x51\x8e\x2e\x2e\x20\x95\x21\x56\x81\xa1\xae\xa1\xae\x40\xa6\x21\x90\x97\x59\x92\x9a\x0b\xb4\x01\xc5\x31\x50\xf5\xc1\xce\xfe\x41\xae\xa8\x3a\x60\xca\x01\x01\x00\x00\xff\xff\xb1\xb3\xdb\x89\xe6\x00\x00\x00")

func bambooScriptsTestLuaBytes() ([]byte, error) {
	return bindataRead(
		_bambooScriptsTestLua,
		"bamboo-scripts/test.lua",
	)
}

func bambooScriptsTestLua() (*asset, error) {
	bytes, err := bambooScriptsTestLuaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bamboo-scripts/test.lua", size: 230, mode: os.FileMode(420), modTime: time.Unix(1448227483, 0)}
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
	"bamboo-scripts/ack.lua": bambooScriptsAckLua,
	"bamboo-scripts/close.lua": bambooScriptsCloseLua,
	"bamboo-scripts/consume.lua": bambooScriptsConsumeLua,
	"bamboo-scripts/enqueue.lua": bambooScriptsEnqueueLua,
	"bamboo-scripts/fail.lua": bambooScriptsFailLua,
	"bamboo-scripts/maxfailed.lua": bambooScriptsMaxfailedLua,
	"bamboo-scripts/maxjobs.lua": bambooScriptsMaxjobsLua,
	"bamboo-scripts/recover.lua": bambooScriptsRecoverLua,
	"bamboo-scripts/test.lua": bambooScriptsTestLua,
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
	"bamboo-scripts": &bintree{nil, map[string]*bintree{
		"ack.lua": &bintree{bambooScriptsAckLua, map[string]*bintree{}},
		"close.lua": &bintree{bambooScriptsCloseLua, map[string]*bintree{}},
		"consume.lua": &bintree{bambooScriptsConsumeLua, map[string]*bintree{}},
		"enqueue.lua": &bintree{bambooScriptsEnqueueLua, map[string]*bintree{}},
		"fail.lua": &bintree{bambooScriptsFailLua, map[string]*bintree{}},
		"maxfailed.lua": &bintree{bambooScriptsMaxfailedLua, map[string]*bintree{}},
		"maxjobs.lua": &bintree{bambooScriptsMaxjobsLua, map[string]*bintree{}},
		"recover.lua": &bintree{bambooScriptsRecoverLua, map[string]*bintree{}},
		"test.lua": &bintree{bambooScriptsTestLua, map[string]*bintree{}},
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

