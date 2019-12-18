// Code generated by go-bindata.
// sources:
// ../lib/kubecfg.libsonnet
// DO NOT EDIT!

package utils

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

var _libKubecfgLibsonnet = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x5d\x6f\xdb\xb8\x12\x7d\xf7\xaf\x38\x30\xee\x83\x5d\x28\x76\x5b\x14\xf7\x02\xbe\x28\xb0\x6e\x9b\xa2\xee\xa6\x4e\xd7\x4e\xb6\xc8\x5b\xc6\xd4\x48\x62\x4b\x91\x5a\x92\x8a\x6d\x2c\xf6\xbf\x2f\x48\x49\x96\x3f\x52\xa0\x40\x60\x28\x9c\xe1\xe1\x99\x33\xa3\x43\x4d\xa7\x78\x6f\xaa\xbd\x95\x79\xe1\xf1\xfa\xe5\xab\xff\xe1\xae\x60\xfc\xa8\x37\x2c\xb2\x1c\x54\xfb\xc2\x58\x37\x98\x4e\x9b\x3f\x00\xb8\x91\x82\xb5\xe3\x14\xb5\x4e\xd9\xc2\x17\x8c\x79\x45\xa2\xe0\x2e\x92\xe0\x4f\xb6\x4e\x1a\x8d\xd7\x93\x97\x18\x85\x84\x61\x1b\x1a\x8e\xff\xdf\xa2\xec\x4d\x8d\x92\xf6\xd0\xc6\xa3\x76\x0c\x5f\x48\x87\x4c\x2a\x06\xef\x04\x57\x1e\x52\x43\x98\xb2\x52\x92\xb4\x60\x6c\xa5\x2f\xe2\x51\x2d\xd0\xa4\x85\x79\x68\x61\xcc\xc6\x93\xd4\x20\x08\x53\xed\x61\xb2\xe3\x5c\x90\xef\xd9\x03\x85\xf7\xd5\x6c\x3a\xdd\x6e\xb7\x13\x8a\xbc\x27\xc6\xe6\x53\xd5\xe4\xba\xe9\xcd\xe2\xfd\xf5\x72\x7d\x7d\xf5\x7a\xf2\xb2\xdf\x75\xaf\x15\x3b\x07\xcb\x7f\xd5\xd2\x72\x8a\xcd\x1e\x54\x55\x4a\x0a\xda\x28\x86\xa2\x2d\x8c\x05\xe5\x96\x39\x85\x37\x81\xfb\xd6\x4a\x2f\x75\x9e\xc0\x99\xcc\x6f\xc9\x72\x8b\x94\x4a\xe7\xad\xdc\xd4\xfe\x44\xc0\x8e\xa9\x74\x27\x09\x46\x83\x34\x86\xf3\x35\x16\xeb\x21\xde\xcd\xd7\x8b\x75\xd2\xe2\x7c\x5b\xdc\x7d\xba\xbd\xbf\xc3\xb7\xf9\x6a\x35\x5f\xde\x2d\xae\xd7\xb8\x5d\xe1\xfd\xed\xf2\xc3\xe2\x6e\x71\xbb\x5c\xe3\xf6\x23\xe6\xcb\x07\xfc\xbe\x58\x7e\x48\xc0\xd2\x17\x6c\xc1\xbb\xca\x86\x3a\x8c\x85\x0c\xd2\x72\xda\xe9\xb8\x66\x3e\x21\x92\x99\x86\x98\xab\x58\xc8\x4c\x0a\x28\xd2\x79\x4d\x39\x23\x37\x4f\x6c\xb5\xd4\x39\x2a\xb6\xa5\x74\xa1\xd1\x0e\xa4\xd3\x16\x49\xc9\x52\x7a\xf2\x71\xf5\xa2\xc0\xc9\x60\xf0\xf7\x00\x98\x4e\x51\x91\x75\xfc\xd9\x19\x3d\x4a\xc9\xd3\x78\xd6\x2c\xb8\x98\xfc\x18\x96\x1e\x11\x74\xd0\x39\xc8\x81\xf0\xdd\x19\x8d\xd4\x88\xba\x64\xed\x93\x78\x5c\x84\xb1\xec\x6b\xab\x9b\x6d\x96\x5d\xad\x82\xe8\x31\x5b\xb3\x87\xd9\x7c\x67\xe1\x27\x03\xf4\xc7\xcd\x66\x70\x3e\x9d\x68\xf2\xf2\x89\x47\xc3\xc3\xfa\x70\x9c\x0c\x8e\x98\x3d\x50\xa9\x4e\x98\xfd\x8c\xd8\xc3\xfc\xcb\x4d\x58\x60\x2a\x9f\xa1\x45\x1a\x2f\xc8\x5a\xda\xbf\xe8\x66\xf2\x67\x24\xdd\x04\x98\xc3\x49\x9d\x2b\x6e\x30\x22\x72\x57\x32\xb6\x52\x29\x38\x1f\x7e\x37\xdc\xe2\x73\x1a\x39\x68\xc4\x23\x9a\x77\xc4\xe8\x76\x3b\x2b\x0e\x1b\x0f\xc5\x87\x8a\x9e\x2b\x3e\xac\xf7\xc5\x97\xa4\x65\xc6\xce\xc7\xce\x3c\x91\xaa\x39\x81\xd4\x29\x6b\x3f\x9e\x41\x18\xfd\xc4\xd6\xc7\x3a\x4e\xd9\xe3\x31\xe6\x3e\x36\x20\xde\x80\x3a\x91\x58\x0b\x93\x36\x44\x87\x95\x65\xef\xf7\x43\x8c\xca\x20\xc1\x95\x92\x9a\xc7\xf8\xbc\xbe\x5d\x26\x0d\x77\x26\x51\x34\x08\x9a\x5d\xd4\x48\xf1\x13\xab\x96\x40\xf3\xda\x3d\x36\xff\x3c\xc2\x55\x24\xd8\x85\xf2\x7e\xce\xf9\xed\x9b\xf1\x6c\x86\xd1\x20\xce\xa5\x11\xa4\x90\xe1\xed\x89\x04\xc7\x7b\x83\x39\x85\xcc\xec\xac\xf0\x01\x70\xa1\x4f\x9c\x8f\x98\xf6\x4b\xba\x44\x49\x1a\x84\x4b\x5d\xa8\x6d\xfb\x69\xc7\x8f\x2b\x7b\xa6\x77\xc7\xa1\xbe\x7d\xec\x04\x55\xbc\x8e\x47\xac\x38\xe7\xdd\xc8\x8d\x67\xf8\xa3\x36\x9e\xdb\xe9\xcb\x79\x87\x92\x3d\x89\x82\x2c\x09\xcf\xd6\x21\x33\xb5\x4e\x83\x67\x45\x35\xa7\xd3\x78\x03\x34\x73\x1a\x0c\x89\xda\x5d\xbe\xa0\x76\x0c\x4b\xf2\xa2\x71\x63\x63\x65\x2e\x35\x29\x28\xe9\xd9\x92\x6a\xf6\xf7\xd8\x01\xf0\x82\xd3\x59\x25\x17\xf1\xbe\x1c\xcb\xce\xa8\x27\x5e\x94\x94\xf3\x48\x86\xdf\x33\xb5\x53\x23\x7e\x70\x30\xb3\xe0\x4c\xad\xb2\x99\x35\x65\xb3\x3d\x2e\xcf\x3c\xe5\x90\x3a\xce\x64\x69\xec\x91\xa5\xc5\xf0\x6f\xa9\xcc\xd9\xf9\x04\x29\x57\xac\xd3\x00\x60\x74\x77\xff\xb5\xe5\x98\xb2\x24\x9d\x22\xcc\x2b\x32\x45\x79\x2c\xeb\x98\xdb\x59\x45\xc7\xa1\xe3\x62\x72\xde\x7d\x09\xd2\x8d\xe2\x63\xd2\x12\x1e\xcf\xb0\xea\x5c\xcc\xd6\x0c\x99\xb5\x82\xcb\xbe\x35\xc7\xa3\x33\xc1\xaa\x0b\x93\x8b\x3e\x1e\x5f\x74\x8e\x2d\xcc\x4d\x30\xea\x06\xa0\x42\x45\xe2\x07\xe5\xad\x21\x8c\xaa\xbd\x2f\x8c\xbe\x92\xae\x18\x37\x05\x74\x7c\x2e\xe8\x77\x81\x33\xf2\xeb\x7a\xe3\xfc\x81\xbc\x15\x09\x2c\x57\xea\xc0\xff\xc8\xdf\x82\xdb\x85\x18\x09\xa9\xf3\x23\x84\x38\x65\x56\x34\xaf\x7b\x48\x98\x00\xab\x98\x17\x6b\xe8\x5a\x18\x6e\x74\xa9\x85\xaa\x53\xc6\x7f\x5e\x25\x60\x2f\x0e\xce\x62\x39\x0b\x77\x8a\x81\xab\x37\x71\x10\xd9\x45\x90\x5f\x93\xa4\xb3\xf8\xa8\xcb\xf3\x92\xc4\x2a\x9f\x93\x24\x06\x7a\x49\x72\xd6\x6c\xc9\xf3\x57\x72\x6e\x6b\x6c\x3a\x52\xac\x73\x5f\x24\xd0\x75\xf9\x41\xe6\xd2\xbb\xf8\xb8\xde\x97\x1b\xa3\xc2\xb3\xb9\xaf\x2a\xb6\x09\x48\x29\xb3\x5d\x71\xc5\xe4\x13\x51\x3b\x6f\xba\x9c\xf1\xac\x41\x6e\xe5\xa4\xc3\x11\x29\xaa\xf6\x90\xfe\x43\x28\x63\xf2\xb5\x65\xd7\xcd\x73\x63\x8f\x21\x52\x91\xa5\x92\xbb\xd7\xef\x9c\xe6\x59\x65\xe7\xe1\x23\x2f\xd9\xb1\xf8\x6a\x4d\x6e\xa9\x1c\x69\x2a\x39\x01\xd9\x3c\x5a\x93\x4b\x90\x91\x54\xb7\xfa\xda\x5a\x63\xc7\x33\x5c\xef\x58\xd4\x9e\x83\x57\x54\xcd\x8e\x70\x15\x76\xb7\x60\xdb\xba\x7e\x38\x0e\x26\xd8\x58\x0a\xb9\xf8\xb5\xe4\x59\xc7\xb6\xfa\xd4\xd4\x7e\x82\xf6\x68\x94\xb5\xf3\x68\x20\x36\x1c\xfa\xf9\x75\x7e\xf7\x29\xfa\x4a\xcf\xef\xdc\x51\xfa\x48\x5f\x8e\xf6\x9f\xc8\x15\x1f\xad\x29\x0f\x0d\xeb\x44\x3d\x9d\xe0\xe5\xdd\xe8\xe6\xcb\x18\x05\xb9\x22\x7e\x06\xd1\x41\xfc\x70\xea\x25\xcc\xd9\xe1\x97\x09\x47\x92\x46\xd3\x7f\x47\x8e\xff\xfb\xe6\xde\xaa\x91\xe7\x9d\xef\x0f\xdf\xc4\xf5\xab\xda\xaa\xab\xee\x76\xe8\x2c\x2d\x7c\x8d\xf1\x2e\xde\x09\x67\x18\xe7\xa5\x9f\x46\xc3\xd1\xff\x0c\xfe\x0d\x00\x00\xff\xff\x2e\x7b\xd2\x11\xe4\x0b\x00\x00")

func libKubecfgLibsonnetBytes() ([]byte, error) {
	return bindataRead(
		_libKubecfgLibsonnet,
		"lib/kubecfg.libsonnet",
	)
}

func libKubecfgLibsonnet() (*asset, error) {
	bytes, err := libKubecfgLibsonnetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/kubecfg.libsonnet", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"lib/kubecfg.libsonnet": libKubecfgLibsonnet,
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
	"lib": &bintree{nil, map[string]*bintree{
		"kubecfg.libsonnet": &bintree{libKubecfgLibsonnet, map[string]*bintree{}},
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

