// Code generated by go-bindata.
// sources:
// data/aws-vpc-public-private/build/template.json
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package goapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"reflect"
	"strings"
	"unsafe"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
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
	name string
	size int64
	mode os.FileMode
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

var _dataAwsVpcPublicPrivateBuildTemplateJson = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x00\x02\xa5\xdc\xcc\xbc\xf8\x82\xc4\xe4\xec\xd4\xa2\xf8\xb2\xd4\xa2\xe2\xcc\xfc\x3c\x25\x2b\x05\x25\x03\x3d\x0b\x3d\x03\x25\xae\x5a\x2e\x40\x00\x00\x00\xff\xff\x1b\x82\xd0\x08\x26\x00\x00\x00"

func dataAwsVpcPublicPrivateBuildTemplateJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsVpcPublicPrivateBuildTemplateJson,
		"data/aws-vpc-public-private/build/template.json",
	)
}

func dataAwsVpcPublicPrivateBuildTemplateJson() (*asset, error) {
	bytes, err := dataAwsVpcPublicPrivateBuildTemplateJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-vpc-public-private/build/template.json", size: 38, mode: os.FileMode(420), modTime: time.Unix(1435862031, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x94\x4d\x6f\xf3\x36\x0c\xc7\xef\xfe\x14\x9c\xf3\x6c\xcf\x06\xd4\x32\x5a\x14\x3d\x04\x6d\x80\xa2\x03\xba\x1e\x86\x76\x6b\xd6\xcb\x30\x34\x8a\x4d\xdb\x42\x65\x51\x90\xe4\xbc\xf4\xe5\xbb\x8f\xb2\x9d\x36\x29\xd0\x61\x40\x80\xd8\x14\xf9\x23\xf9\x27\xad\x09\x5c\xa3\x41\x27\x03\x96\xb0\xdc\xc2\x6d\x08\x74\x04\x25\x81\xa1\x00\x58\xaa\xf0\x43\x32\x49\x26\x30\x6f\x94\x07\xfe\x85\x06\xe1\x41\xd6\x4e\x9a\x50\x29\x8d\x50\x7f\x8e\x85\x8a\x5c\xef\x55\xe2\x0a\x35\xd9\x16\x4d\x00\xaa\x18\x11\x22\x42\x5a\xab\x55\x21\x83\x22\x93\x7b\x74\x2b\x55\xa0\x80\x9b\x00\xbe\xa1\x4e\x97\x7d\xd2\x25\x42\x23\x4d\x99\xc5\xe4\x58\x0a\x98\x13\xb4\x54\xaa\x6a\x1b\xb1\xcc\xd9\x4b\x7f\x04\x9d\xc7\x3e\xdb\xa5\xb5\xd1\x20\x92\x64\x3c\x16\x05\x99\x4a\xd5\x9d\xc3\x9f\xd3\x93\xf4\x97\xd8\xd1\xeb\x60\x7a\x4d\x00\x86\x27\xb1\x6a\xc5\x92\x36\x70\x01\x69\x23\x7d\xa3\x0a\x72\x36\xb7\x0e\x0b\xe5\xf1\xec\x34\x4d\xd8\x71\x02\xf7\x18\x3a\x0b\x12\xfc\xd6\x14\xdc\x66\x45\xba\x44\x07\x95\xa3\x16\xa8\x73\xb0\x26\xf7\xa4\x4c\x0d\xa5\xe2\xb8\x40\x8e\xab\x24\xc8\x57\x43\x11\x07\x99\x06\xc0\xe3\x08\x48\x5f\x5e\xc0\xca\xd0\x88\x1d\xe0\xed\x2d\x3d\x82\x74\x17\x39\x26\xbf\x31\x3e\x48\xad\xe1\x9a\x60\xd9\x29\x16\x08\xcd\x4a\x39\x32\x51\xd5\x03\xb8\x75\xb4\x52\x9e\x55\x85\xd4\x37\xa8\x35\xb3\x94\xd1\xca\xe0\x14\xbe\xf9\xc2\x29\x1b\x1e\x6b\xd2\xd2\xd4\x03\xf7\x77\xf9\x84\xa0\x58\x76\x62\xf5\x64\x80\xc5\x98\x16\xbc\x6f\x16\x50\x13\xfa\xb1\x21\xdd\xf7\x13\x15\x66\x71\xa2\x21\xda\xff\x67\x66\x76\x03\xf8\xf1\x8f\xbf\xb1\x68\x08\xd2\xa2\x7c\x97\x25\x85\xd9\x0c\xf2\x86\x5a\xdc\x59\x72\xb1\xe4\x01\xb8\xe2\x9f\x04\x4d\x99\x24\x87\x25\xf3\x7c\xce\xcf\xef\xaf\xfe\xbc\xb9\x9b\x27\x1e\x03\x64\x98\x24\x03\xf3\x57\x5a\x1b\x4d\xb2\x8c\xfa\x5d\x93\x10\x22\x4d\xd6\x75\xf4\xb8\xfd\x84\xaf\x49\x04\xe9\x44\xfd\x0c\x4d\x08\xd6\x4f\xf3\xdc\xf3\xac\x64\x8d\xa2\x26\xaa\x35\x4a\xab\x3c\x6f\x4c\x9b\x0f\x19\xf9\xef\x58\x9c\x8a\x13\xc1\x7d\x74\x9b\x4c\xb6\xe5\xd9\xe9\x08\xd8\xa5\xfe\xcb\xf0\xbb\xdb\x4b\xec\x3b\x5e\x31\xb6\x41\x76\x05\x79\xe7\x5d\xae\xa9\x90\x1a\xb2\xcd\x73\xf5\x55\x31\x3b\x16\x4f\xa3\x07\xdd\xde\x5d\xce\x7f\xfb\x80\xb5\x4f\xac\x35\x64\x16\x72\xb2\x31\x2a\xae\xcb\x70\xc2\x51\x6b\x03\x8b\x75\x43\xb2\x55\x8b\xe9\xee\xe1\xc0\x71\x64\xf3\xfa\x86\x08\xe7\x1d\x7e\xa7\xf7\x27\xdf\x71\x63\xc9\x85\xde\x7a\xb1\x17\x98\x2f\x95\x99\x7e\x34\xc0\xd6\xde\xf2\x2d\xfa\x7d\xff\x72\x70\x87\xcc\xa1\x93\x7d\xea\x7f\x44\x8e\x85\x8e\xab\x1e\x6b\x7d\xb8\xba\xf7\xfd\x3d\x52\x13\x5f\x30\xe1\x43\x11\x69\x43\x16\xe7\xdb\xd9\x92\xef\x1c\xc8\xb6\x30\xcb\xf9\x9e\xc9\x4d\xc7\xdf\xc8\xc9\xec\xa7\xe3\x43\x37\x35\x7e\x3d\xec\x57\xf3\xb2\x2f\x9f\x1d\xb4\xe8\x8a\xce\x29\xa9\x93\x71\xa1\xfe\x0d\x00\x00\xff\xff\x01\xa6\x23\x4f\xff\x04\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 1279, mode: os.FileMode(420), modTime: time.Unix(1439846697, 0)}
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
	"data/aws-vpc-public-private/build/template.json": dataAwsVpcPublicPrivateBuildTemplateJson,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-vpc-public-private": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"template.json": &bintree{dataAwsVpcPublicPrivateBuildTemplateJson, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
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

