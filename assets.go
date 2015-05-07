package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _dockerfile_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x80\x80\xe2\xe4\xa2\xc4\x92\xe4\x0c\x2e\x47\x17\x17\xa8\x88\x82\x9e\x7e\x75\xb5\x5e\x40\x51\x7e\x56\x6a\x72\x49\x6d\x2d\x1a\x97\xcb\xd5\x2f\x24\x28\x32\xc0\xdf\xd3\x2f\x44\x21\x5a\x09\x55\x4e\x29\x96\x0b\x10\x00\x00\xff\xff\xbc\x53\x4a\x49\x5a\x00\x00\x00")

func dockerfile_template_bytes() ([]byte, error) {
	return bindata_read(
		_dockerfile_template,
		"Dockerfile.template",
	)
}

func dockerfile_template() (*asset, error) {
	bytes, err := dockerfile_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "Dockerfile.template", size: 90, mode: os.FileMode(420), modTime: time.Unix(1428456742, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dockerignore_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcf\x2c\xd1\xe7\x4a\x2a\xcd\xcc\x49\xd1\xe7\x72\xcf\x4f\x49\x2d\x28\xd6\xe7\xf2\x4d\xcc\x4e\x4d\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x49\x54\x7b\x12\x1e\x00\x00\x00")

func dockerignore_template_bytes() ([]byte, error) {
	return bindata_read(
		_dockerignore_template,
		".dockerignore.template",
	)
}

func dockerignore_template() (*asset, error) {
	bytes, err := dockerignore_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: ".dockerignore.template", size: 30, mode: os.FileMode(420), modTime: time.Unix(1428456742, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _gitignore_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\x2a\xcd\xcc\x49\xd1\x4f\xce\xcf\x2b\x49\xcc\xcc\x4b\x2d\xe2\xd2\xd2\x2b\x2e\x2f\x00\x92\x25\xb9\x05\x5c\xd5\xd5\x7a\x01\x45\xf9\x59\xa9\xc9\x25\xb5\xb5\x5c\x80\x00\x00\x00\xff\xff\xda\xe9\xe4\xc5\x29\x00\x00\x00")

func gitignore_template_bytes() ([]byte, error) {
	return bindata_read(
		_gitignore_template,
		".gitignore.template",
	)
}

func gitignore_template() (*asset, error) {
	bytes, err := gitignore_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: ".gitignore.template", size: 41, mode: os.FileMode(420), modTime: time.Unix(1428456742, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _main_go_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\xd2\xd7\x57\x28\x4e\x2d\x51\x48\xaa\x54\x28\xc9\x48\x55\x48\xce\xcf\x2d\xc8\xcc\x49\x2d\xe2\x2a\x4b\x2c\x52\x70\x2a\xcd\xcc\x49\x09\xf6\x70\x54\x28\x2e\x29\xca\xcc\x4b\xe7\xe2\x4a\x2b\xcd\x4b\x06\x6b\xd3\xd0\x54\xa8\xe6\xaa\xe5\x02\x04\x00\x00\xff\xff\xd1\xe3\x56\xcc\x4a\x00\x00\x00")

func main_go_template_bytes() ([]byte, error) {
	return bindata_read(
		_main_go_template,
		"main.go.template",
	)
}

func main_go_template() (*asset, error) {
	bytes, err := main_go_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "main.go.template", size: 74, mode: os.FileMode(420), modTime: time.Unix(1428456742, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _makefile_template = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\x7f\x6f\xe3\xb8\x11\xfd\xdb\xfc\x14\x53\xc5\x40\x92\x85\x65\xdd\x6d\x8a\xed\x41\x81\x0f\xe7\x4b\xdc\x24\x68\x10\x07\x76\x82\x76\xb1\x77\x48\x68\x89\xb6\xd9\xc8\xa2\x2a\x52\xf6\xba\xae\xbf\x7b\x67\x48\xc9\x96\x9d\x64\xb1\xd9\x16\x8b\xcb\x1f\x81\x45\x91\x8f\x6f\x7e\xbd\x19\x1d\x40\x3f\x89\x7d\xfd\xa4\x54\x02\xa3\x42\x26\x31\x18\xfc\xa9\xdb\xec\x80\x1d\xc0\x1d\xcf\x27\xc2\x68\x38\xd2\x42\x80\xe0\xd1\x14\x8c\x5d\x81\xb1\xca\x61\xa6\x72\x01\x32\xc5\x9f\x33\x6e\xa4\x4a\x8f\x43\x3c\x01\x0e\x24\x88\x54\x6a\xb8\x4c\x45\x1e\x36\xec\x82\x06\x33\x15\x70\xae\xa2\x27\x91\x83\x9c\xf1\x89\x80\x42\x0b\xba\x0c\x22\x35\xcb\x64\x22\x80\x27\x09\x4c\x54\xc2\xd3\x09\x2e\xc5\x62\x0b\x16\x42\xf9\xd7\xa8\xb0\x46\x32\xe5\xb9\x14\xba\x05\x59\xc2\x23\x89\x27\x2c\xb9\x54\x7c\x36\x84\x28\xcd\xa1\x86\x5c\xe8\x4c\x44\x46\xce\x05\xcc\x90\x09\x64\x4f\x13\x0b\x69\x84\x36\x21\x42\x39\xc0\xbc\x48\xb5\x5d\xd2\xf6\x65\x22\x53\x13\xba\x37\xb4\x83\x1e\x1d\x73\xad\x8a\x3c\x12\x60\x72\xe1\x88\xc9\x54\x1b\x24\x1c\xda\x8d\x25\xad\x96\x03\x6a\xd1\x81\x14\x6d\xc8\x90\xa1\x3d\x8c\x54\x8a\xc4\x10\x4d\x4b\x7c\x49\x1c\x9b\x17\xfd\xdb\xee\xdd\x65\x80\x2b\x81\x45\x8c\xad\x6f\xe4\xbf\x45\x58\x37\x75\x17\x73\xc6\x9f\x10\x92\xef\xfa\x91\x62\x61\xcd\x77\xe0\xec\x00\x89\x47\x89\xe0\x69\x68\x6d\x40\x1b\xc5\x4c\xcd\xf1\x9c\x0b\x2f\xcf\x8d\x1c\xf3\x88\xa2\xca\x9f\xf8\xc6\x95\xc7\xdb\x73\x3e\x19\x66\x4f\xda\xc7\xea\xf2\x0a\x86\x2c\xaa\x40\xec\xe5\x7b\x11\x67\x6c\x78\xd9\xbb\xbe\x86\xb0\x03\xd6\xba\x11\xd7\x53\x86\x26\x6a\x35\x43\x07\x8a\x7c\x86\x37\x26\xe8\x9e\x84\x78\xeb\x88\x67\xc2\xc6\x5b\xb3\xeb\xab\x8b\xcb\xbb\x87\x8b\x41\xaf\x77\x43\x87\x9b\x47\x7a\x2a\x30\x27\x44\x34\x55\xe0\x0b\xf0\x7e\xfb\xe1\xe4\xe4\xd3\x8f\xa7\x27\xef\x67\xde\x31\xbb\x39\x7b\x7d\xcf\x0f\xb8\x01\x0e\xe0\x46\xc1\x19\xdd\x42\x97\xab\x11\x91\x73\xe4\x13\x4e\x3e\x85\x89\x34\x30\xbc\xec\x5a\x13\x68\x39\x2a\xf2\x5c\xa4\x06\xed\xcc\x54\x0b\x16\x53\x89\x2e\x95\x7a\x93\xa6\x86\x4f\xca\x18\x21\x9c\xf5\x3c\x46\x22\x8d\x41\x63\x39\x64\xb9\xca\x04\xba\x04\x97\x30\xd3\x08\xac\x4c\xe4\xca\xbb\xec\xe2\xea\xee\x01\x2f\xab\x51\xa6\xeb\x73\x31\xf7\x33\x9e\x6b\x01\x97\xbd\xee\x39\xbc\xff\x39\x88\xc5\x3c\x48\x0b\x7c\xfd\x1f\xe4\x63\xc0\x8f\xe0\x47\xff\x2f\xc7\x64\xc1\x2d\xd2\xa6\x5a\x0b\x78\x8e\xc4\x28\xb7\xe5\x58\x46\x10\xe5\x45\x4c\x85\xe7\x43\x3f\xa5\x7c\x2d\x3e\xb7\x00\x93\xda\x59\x54\x85\x04\x16\xd2\x4c\x77\x8c\x2c\x64\xdc\xc2\x88\xc0\x18\x2b\x4f\x13\xff\xb8\x88\xd0\xce\x71\xae\x66\x36\x1b\xe9\x80\xdc\x47\xe1\x58\xf0\x6a\x91\xe2\xbe\xd1\x72\x17\x4d\x8b\x1c\xaf\xe5\xb8\x46\xbe\xe4\x98\x2c\x4a\x99\x76\x45\xab\x3f\xfc\x47\x0b\x7d\x97\x1e\x1a\xfc\x8f\x4e\x5b\x12\xf6\xa4\x46\x6a\x73\x83\x03\x22\xb7\x26\xe8\xd6\x11\x82\xbc\x77\x3e\xc7\xd4\x4f\xd1\xe5\x96\x5b\x46\x39\xa4\x35\x4a\x8e\x06\xe5\x38\x06\xf7\x78\x50\xa3\x24\x15\xc8\x06\x09\x18\xd4\x00\x8a\x0c\x86\x2f\xb3\x2c\xa2\x5c\x69\xed\x57\x52\x73\x24\xc7\x28\x16\x91\xd0\x1a\x0b\xe6\x18\xf6\x0d\x6d\x61\x4a\x1b\x78\x12\x22\x73\xb5\xcf\x31\x71\xcb\xca\x4d\xf1\x37\x3b\xef\x9f\xfd\xad\x37\x78\xb8\x1f\xf6\x06\xf5\x24\x44\xd0\x4f\x9f\xc0\x6b\x36\xfb\xc3\xbb\x8f\xb7\x3d\x0f\xfe\xd4\x01\x2f\xe6\xf9\x42\xa6\xde\x3b\xf8\xfd\xf7\x53\x57\x47\x74\xec\xa1\x3b\xb8\xe8\x78\xbe\x4f\xf6\x76\x1e\x65\x0c\x7e\xf1\xe8\x9d\x62\x34\x4e\x5d\x2a\x23\x48\xb5\x0d\x73\xfd\xa2\xdf\x1f\xa2\xdd\x5f\xba\xaa\xf3\xe2\x55\x0e\xab\x5c\x47\xe8\x04\x33\xcd\xad\xd9\x4c\xb1\x37\x12\x7c\x77\x70\x76\xf9\x1c\xfe\xb1\x20\x6b\xc1\xe7\x8f\x04\xff\xce\xfb\xfc\xd3\x87\x87\x0f\x7f\x7e\x8e\xcf\x67\x31\x2e\xef\xc0\x9f\xfc\xf4\xa1\x04\x47\xef\x93\x0f\x3d\x4a\x08\x8f\x44\x78\xe3\xe6\xd8\xf9\x5d\xda\xf4\xfb\x27\x4a\x75\x1b\xe0\x6a\x6c\x77\x97\x0b\xd5\x56\x0c\x2c\xe9\x67\x96\x50\xfc\xb7\xda\x5f\x13\x3d\xd4\xfb\x4a\xe4\x31\xad\xd0\x82\x11\xb5\x26\x54\x4a\x5d\x8c\x62\x99\x83\x1a\xc3\x70\x70\xf6\x30\xe8\xf7\xef\x58\xf5\xa3\xb3\x5a\xb5\x07\x58\xeb\x5a\x1a\x95\x2f\xd7\xeb\x00\x9f\x6f\xd0\x60\x9d\xf1\x48\xb8\xc7\x5b\xc7\x03\x1f\x18\x6b\x9f\xf7\xfe\xda\xbd\xbf\x46\x81\xea\x77\xad\xb8\x59\xd9\x23\xfb\x7e\xdd\x36\xb8\xb8\x2e\xcc\x36\x0f\x17\xe2\x10\xe9\x60\x9c\xf7\x5b\x1d\x32\xd3\xe2\x5f\x05\x95\xcf\x6e\xd7\x33\xaa\x88\xa6\x21\x4c\x8d\xc9\xc2\x20\x58\x2c\x16\xed\x49\x5a\xb4\x55\x3e\x09\xb4\x1a\x9b\x05\x16\x61\x40\x7d\x00\xff\xa5\x05\x4f\xec\xef\xf6\xd4\xcc\x92\x83\xde\x2c\x33\x4b\xbf\xec\xd9\x6c\xbf\x0f\x97\x32\xed\x5a\x07\x95\x3d\x6b\xfc\x52\xe6\xda\xaa\xa6\xbc\x6b\xbb\x8d\xca\xb3\xde\x64\xf0\xb4\xc2\x23\x79\x46\xda\xb9\xe7\x28\xbf\xee\xa8\xaa\xc4\xda\xed\x76\x73\x75\x73\xb6\xf6\xf0\x96\xd2\x29\xae\xfd\xf8\x7e\xaa\xfc\x08\x23\x87\x89\x65\xde\x8a\x5a\xda\x00\x3f\xc3\x46\x28\x59\xc3\xfa\x0b\x9a\xbf\x30\xe6\xba\x1e\x6b\xe4\x33\xf0\xc7\x50\x07\x28\xdf\xd9\xce\xe6\x9a\x5c\xb5\xeb\x59\xf3\x3a\x00\x3b\x13\x70\x78\x9c\x60\x3c\x32\xac\xfe\xb9\x78\x74\xc9\x54\x13\x08\x55\x98\xac\x30\xb6\xa7\x6f\x3b\xbc\x86\xb9\xe4\xf6\x79\xae\x92\x02\x4b\xc7\xea\x11\xb3\x38\xe1\xb3\x9b\x36\x8e\x21\xb9\xf6\x7d\xa4\xf3\x1b\x6b\x34\xfc\x39\x95\xf5\xed\xdf\xcf\xbd\xd0\x0b\x26\x2a\xd0\x79\x14\x34\x57\x55\xce\xae\x3d\xb7\x69\x01\xaf\xbf\x6c\xae\x6a\xfa\xb4\xa6\x25\xf8\x16\x57\x13\xd4\xd6\x03\xac\x7d\x7b\xd9\xbf\xf9\x18\x82\x5d\x62\x6c\xb5\xba\x53\xf7\x19\x2a\x31\x6c\x8f\x3e\x20\x91\x21\x6c\x55\x64\x2c\x51\xc7\xdb\x78\xf7\x12\x5b\xfc\x18\x7c\xab\x26\x87\xef\xda\x13\x75\x68\x75\xa1\x36\x18\x96\xea\x6a\xbd\x5c\xe6\x5d\xad\xed\x20\x4c\x39\x4d\x49\x3b\xdf\xf1\x1d\xff\xc2\x51\xe0\xa2\x11\x1c\xb3\xba\x25\x21\x12\x79\x95\xe6\xf1\x0b\xe1\xf8\x62\x35\x34\x57\xa4\xc2\xce\x3d\x38\xee\x8e\xd0\x43\xb4\x44\xca\xb9\xae\xe8\xd3\x24\x51\x67\xf0\x42\x11\xfc\x9f\x63\xed\x10\x56\x08\xb1\x0e\xbc\xb0\xf4\xc3\x2b\x59\xd0\xa0\xf9\x88\x8c\xe8\x94\xb6\x78\x65\x6e\xd8\x65\x32\xa4\xb3\xb1\xc8\xdb\x1c\xf8\xf5\xea\xa6\x3b\xf8\xd8\xa9\x9b\xb5\x7d\x59\x4e\x35\x74\xce\xfd\xaa\xde\xbd\x39\xdd\x18\x2b\x47\xfd\xdd\xaa\xad\x8a\x31\x5d\xba\x39\x98\xe6\x6e\x19\x8b\x17\xb2\x84\xb9\xb9\xde\x89\xf2\xcb\xb1\x44\xa8\xd4\x16\xac\x45\xfa\xfe\xc1\xfa\xf6\x0a\x9c\x28\x4b\x1a\x90\x43\x3b\x40\xa2\x9b\x6a\xa4\x55\xc6\x36\x1f\x23\xe5\x07\x1c\x2d\xbe\xec\x02\x2c\xa3\xe5\xde\x77\xc8\xca\x7d\x88\xac\xed\xac\x5e\x73\x42\x94\x55\xf6\xee\x74\x42\x68\xee\x9e\xd8\x50\x29\x49\x50\xcc\xb6\x9f\x4d\xd8\xdd\x24\x0e\x8d\xe8\x0e\xfb\xe9\xf4\x7a\xf8\x5a\xd8\x22\x89\x18\xb6\x42\x3c\xcc\xdc\x77\xd8\xd7\x95\x28\xed\x95\x65\xfb\x7c\x6b\x10\x5f\x0d\x60\xed\x45\xf0\xbf\x44\x8f\xc8\xed\xc5\xcc\x1a\x58\x8d\x0d\xe8\x0a\x3b\x8a\xd5\x75\xa5\x0c\x0e\xc9\x1e\x7f\x61\xa2\xc0\x59\x62\x47\x34\xb9\x76\x1f\xbc\xe8\xeb\x7c\x99\x29\x42\xaf\x7d\x4b\xbe\x51\xe6\x76\x42\x5d\x93\x35\x99\x46\x49\x41\x23\x37\x69\x74\x7d\x32\xf8\x03\x6a\xdc\x77\x93\xac\x37\x0c\x50\x70\xf8\x85\xe9\x32\xdc\x92\x39\xfc\x9a\xc1\xa9\xe4\xeb\x7d\x1d\xa4\x07\xce\x79\x9b\x04\xdc\x64\x07\x63\xff\x0d\x00\x00\xff\xff\x94\x91\x58\x74\xfe\x11\x00\x00")

func makefile_template_bytes() ([]byte, error) {
	return bindata_read(
		_makefile_template,
		"Makefile.template",
	)
}

func makefile_template() (*asset, error) {
	bytes, err := makefile_template_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "Makefile.template", size: 4606, mode: os.FileMode(420), modTime: time.Unix(1429911995, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"Dockerfile.template": dockerfile_template,
	".dockerignore.template": dockerignore_template,
	".gitignore.template": gitignore_template,
	"main.go.template": main_go_template,
	"Makefile.template": makefile_template,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	".dockerignore.template": &_bintree_t{dockerignore_template, map[string]*_bintree_t{
	}},
	".gitignore.template": &_bintree_t{gitignore_template, map[string]*_bintree_t{
	}},
	"Dockerfile.template": &_bintree_t{dockerfile_template, map[string]*_bintree_t{
	}},
	"Makefile.template": &_bintree_t{makefile_template, map[string]*_bintree_t{
	}},
	"main.go.template": &_bintree_t{main_go_template, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
