// Code generated by go-bindata. DO NOT EDIT.
// sources:
// tmpl/function.tmpl (258B)
// tmpl/info.tmpl (1.714kB)
// tmpl/object.tmpl (1.902kB)
// tmpl/operation.tmpl (755B)
// tmpl/pair.tmpl (2.214kB)
// tmpl/service.tmpl (7.449kB)
// ../../definitions/infos.hcl (1.253kB)
// ../../definitions/operations.hcl (4.605kB)
// ../../definitions/pairs.hcl (1.893kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _cmdDefinitionsTmplFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\x41\x4a\xc4\x40\x10\x45\xf7\x39\xc5\x67\xc8\x22\x33\x68\x1f\x40\x70\x15\x74\x25\x32\xa8\x17\x28\xda\x1e\x6d\x4c\x55\x9a\x54\x05\x07\xda\xba\xbb\xf4\x44\x04\x5d\x15\xd4\xff\xbc\xff\x6a\xbd\x46\x7f\x92\x0f\xdc\xdc\xa2\x0f\xf7\xab\xc4\xf0\x48\x9c\xf0\x05\x9b\x47\xe2\x34\xc1\xbd\x3b\xad\x12\x31\x28\x0e\xb5\xa2\xbf\xe4\x5a\x28\x6e\xa5\x23\x69\xa4\xd6\xda\xa3\xd6\x46\x72\x1f\xa2\x9d\x11\x67\xb1\x74\xb6\x30\x6e\xf7\xaa\xa5\x1b\xff\x48\x0b\xb1\x86\x97\x25\xf3\x03\xa9\x85\x67\x5b\xb2\xbc\xdd\xc9\xab\x7e\x66\x7b\x1f\x67\x66\x72\xc7\x5c\x0c\x87\x42\x79\xf9\xb7\xe9\xde\x1e\xcd\xf8\xef\xfa\xf0\xcb\x7f\x4a\xba\x4e\xa6\x3f\xdc\x8b\x58\x07\x00\x85\x24\xc7\x61\x97\xb9\x4c\x89\x93\x18\xb2\xed\xf6\x9d\x7f\x07\x00\x00\xff\xff\x5b\xb1\x71\xf7\x02\x01\x00\x00")

func cmdDefinitionsTmplFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplFunctionTmpl,
		"cmd/definitions/tmpl/function.tmpl",
	)
}

func cmdDefinitionsTmplFunctionTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/function.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf3, 0x17, 0x45, 0xc4, 0x8, 0xb1, 0x57, 0xf3, 0x99, 0x7b, 0x1f, 0xcd, 0xe4, 0x30, 0x7e, 0x3, 0x30, 0xba, 0x28, 0xed, 0x36, 0x1c, 0xd3, 0x4a, 0x24, 0xfc, 0x10, 0x73, 0xfe, 0xf7, 0xda, 0x75}}
	return a, nil
}

var _cmdDefinitionsTmplInfoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x53\x5f\x6b\xdb\x3e\x14\x7d\xd7\xa7\x38\xbf\x10\x7e\xc4\x23\x8d\x5b\x18\x7b\xd8\x9a\xa7\xb6\x1b\x65\xb4\x1d\xb4\xec\x61\x63\x0c\xd9\xbe\x0e\x5a\x2c\xc9\x48\xb2\x69\xe6\xfa\xbb\x0f\xc9\x6e\x9c\xa4\x69\x46\xfb\xb4\x37\xfd\xbb\xe7\x9c\x7b\x74\x6e\x1c\xe3\x4c\x67\x84\x05\x29\x32\xdc\x51\x86\x64\x85\x85\x5e\xef\x21\x94\x23\xa3\x78\x11\xa7\x32\xfb\x80\xf3\x1b\x5c\xdf\xdc\xe1\xe2\xfc\xf2\x6e\xc6\x4a\x9e\x2e\xf9\x82\xe0\x56\x25\x59\xc6\x84\x2c\xb5\x71\x98\x30\x00\x18\xe5\xd2\x8d\x58\xc4\x58\xd3\x1c\xc1\x70\xb5\x20\x8c\x97\x53\x8c\x85\xca\xb5\xc5\xfb\x39\x66\x97\x7e\x75\xc5\x4b\xb4\x2d\x6b\x1a\x8c\x2d\x99\x5a\xa4\x74\xcd\x25\xf9\xfb\xf1\x12\x0f\x70\xfa\x8c\x4b\x2a\xfc\x13\x16\xc7\xf8\x28\xa8\xc8\x20\x54\x46\xf7\x10\x0a\x4d\xb3\x59\xd4\xb6\x48\x84\x63\xa9\x56\xd6\x8b\xd8\xe1\xad\x03\x66\xc7\xde\xb6\x41\xe2\x6e\xf9\xa5\xc7\xf5\x4a\xea\x59\x10\xe1\xe9\xbf\x70\x9b\x72\xcf\x8f\x39\x4e\x4e\x4f\xfd\xed\xb2\x13\x7c\x04\x52\x99\x5f\x46\x8c\x79\x03\xb0\xdb\x43\xdb\xc2\x3a\x53\xa5\x0e\x4d\x4f\xb7\x16\xf4\xf3\x39\x41\x9e\xfb\x6e\x55\xae\x01\x86\x93\xe1\xcd\x9a\x38\xec\xe3\xd8\x77\x8d\xca\x52\x06\x6e\xc1\xfd\x4e\xf2\x12\xb9\x36\xd0\xc9\x2f\x4a\x1d\x6a\x5e\x54\x34\xc5\x31\x24\x71\x65\xa1\xb4\x83\x25\x37\xc5\x49\x7f\x60\xc9\x05\xa8\x80\x23\x94\x7b\xf7\x36\x6c\x25\x24\x2f\xbf\x5b\x67\x84\x5a\xfc\x08\x29\xc8\x79\x4a\x4d\xcb\x7a\xe6\xc3\x06\xfb\x5b\x91\x7b\xf5\x17\xf7\x21\x16\x6d\xcb\xf2\x4a\xa5\x98\x48\xbc\x79\x6a\x55\x84\x4f\xe4\xba\x6e\xcf\x85\x2d\x0b\xbe\xea\x2f\x26\xd1\xb6\x09\xbd\x99\x86\x5c\x65\x14\xe4\xec\x89\x67\x5e\xde\x41\xa2\xdb\x67\x88\xea\x6d\xa2\x68\x5f\x75\xcf\xbe\x87\x16\x73\xd4\x5b\xca\x58\x1f\x92\xc2\x06\x55\xaf\x6d\x7e\xb2\x25\x6a\x8a\x44\xeb\x22\xea\x65\x88\x1c\x72\xe6\x7f\xed\xff\x17\x66\xf9\xbf\x39\x8e\x7b\x8c\xc3\x66\x4e\xe1\x4c\x45\xe1\xe1\x10\xbf\xee\x5b\xbf\x91\xd1\x5f\x7d\xb2\x1e\x93\xb9\x81\xd4\xe1\x6c\xbe\x98\x22\xe7\x85\xa5\x21\xc2\xbd\x2d\x7b\x0b\x43\xb3\x0f\xf8\x7d\xa8\xbe\x1b\x81\xbf\xf9\x7a\x55\x59\xf7\xb2\x60\xbd\xda\xd3\xf9\xb6\xa7\x25\x57\x22\x9d\xe4\xd2\xcd\x6e\x4b\x23\x94\xcb\x27\xa3\x5d\x81\xbe\xfe\x33\x25\x3c\x19\x26\xfd\x31\x4c\x62\x3d\xa6\xa3\x28\xda\xb0\xff\xdf\x8d\x7d\x67\xda\xc3\xfc\x65\xae\xed\x9d\x98\xee\x6b\x77\x7e\x7a\x58\xfe\x09\x00\x00\xff\xff\x2e\x94\xa2\xf0\xb2\x06\x00\x00")

func cmdDefinitionsTmplInfoTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplInfoTmpl,
		"cmd/definitions/tmpl/info.tmpl",
	)
}

func cmdDefinitionsTmplInfoTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplInfoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/info.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x85, 0x27, 0x2d, 0x89, 0xde, 0xbb, 0x9, 0x4b, 0x47, 0xf8, 0x4d, 0xa, 0xd, 0x94, 0x37, 0x71, 0xd9, 0x7f, 0xdf, 0xb9, 0x39, 0x86, 0x82, 0x1, 0xc, 0xa4, 0xc6, 0xb2, 0xde, 0xe1, 0x1f, 0x7d}}
	return a, nil
}

var _cmdDefinitionsTmplObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4d\x6f\xe3\x36\x10\x3d\x9b\xbf\xe2\xd5\x08\x0a\xa9\x70\xa5\xa4\x2d\x7a\x68\xe3\x53\x92\x16\x39\x24\x29\x90\xa0\x87\x16\x7b\xa0\xa5\x91\xc3\x8d\x44\x0a\xe4\xc8\x6b\xaf\xa2\xff\xbe\x20\x25\x7f\x25\xde\xc0\xc1\x5e\xf6\x64\x0e\x39\xf3\xe6\xe3\xbd\xb1\xd2\x14\x17\x26\x27\xcc\x49\x93\x95\x4c\x39\x66\x2b\xcc\xcd\xc6\x86\xd2\x4c\x56\xcb\x32\xcd\xaa\xfc\x4f\x5c\xde\xe1\xf6\xee\x01\x57\x97\xd7\x0f\x89\xa8\x65\xf6\x24\xe7\x04\x5e\xd5\xe4\x84\x50\x55\x6d\x2c\x23\x12\x00\x30\x2e\x2a\x1e\xf7\x27\x56\x15\x0d\x47\xb7\xd2\xd9\x58\xc4\x42\xa4\x29\xfe\x52\x54\xe6\x50\x3a\xa7\x25\x94\x86\x99\x7d\xa4\x8c\x31\x53\x2c\x32\xa3\x9d\xc7\x69\xdb\x9f\x61\xa5\x9e\x13\x4e\x9e\x26\x38\x59\xe0\x8f\x29\x92\xbb\xe0\x77\x43\x2c\xd1\x75\x01\xb5\x8f\xbc\xf6\x40\x6d\x8b\x93\x45\x72\x2b\x2b\xc2\x33\xd8\xfc\x23\x5d\x26\x4b\x74\x1d\xa6\x38\x3b\x3f\xf7\xaf\x4f\x3e\xca\x03\x93\xce\xfd\x31\x16\xc2\x97\x8f\x1e\x16\x8e\x6d\x93\x31\xda\x63\x73\x7b\x37\x55\xf8\xa4\x17\xa6\xaa\x48\xf3\xfa\x01\x69\x8a\xbe\x9a\x17\x0f\x3b\xb9\x7b\xd3\xfb\x3c\xac\x6a\x0a\x55\x77\xdd\xce\xcd\x8b\x52\xc5\x28\x4d\x91\x95\xca\x83\x29\x07\x7e\xa4\x8d\xa5\xf1\xe9\x51\x65\x8f\xeb\x2e\x94\x83\x2c\xd5\x82\x12\x31\x1a\x3c\xee\xd9\x58\x39\x27\x1b\x30\x2a\x38\x36\x96\x5c\xf8\xf1\x04\x5a\x2a\x03\xf5\x15\xb1\xcc\x25\xcb\x24\x94\xe6\x2d\x54\xb2\xfe\xdf\xb1\x55\x7a\xfe\x21\x28\xa1\x90\x19\xb5\x9d\x10\x43\x8b\x33\xc5\x68\x1c\xe5\x90\x0e\xd2\x5b\x95\xac\x51\x18\xbb\xa6\x73\x21\xcb\x86\x26\x38\x45\x45\x52\x3b\x68\xc3\x70\xc4\x13\x9c\x0d\x17\x8e\x38\x40\x05\x1c\xa5\xf9\xf7\xdf\xc4\x28\x37\x9a\x82\xf1\xeb\x2f\x62\x54\xf9\x57\x2f\x9b\xe4\xa6\x61\x5a\x8a\x4e\x88\x63\xb8\xd9\xf2\x72\xb5\x0c\xaa\xec\x3a\x51\x34\x3a\x43\x64\xf0\x53\xef\x1a\xe3\x6f\xe2\x7e\xd8\x97\xca\xd5\xa5\x5c\x0d\x0c\x44\xf1\x3e\x07\x68\x43\x89\x96\xb8\xb1\x1a\x26\x79\x45\x99\x2f\xea\x15\xf8\xfd\x57\xc0\x17\xfb\xe0\xf1\x3a\x62\xc8\x72\x00\x1e\x53\x2c\xf6\x2a\x10\x83\x2c\x4a\x17\xb2\xbf\xa7\xb1\x68\x2f\xf9\x04\x33\x63\xca\x78\x93\xda\xb1\xe4\x28\xee\xc9\x55\x05\x4c\xe2\x69\xf9\xf1\xa8\x05\xfb\x61\x8a\xd3\x01\xe7\xed\x61\x4d\xc0\xb6\xa1\xe0\x38\xc8\x68\xcb\xd5\x7f\x64\xcd\xbf\x5e\x31\x9b\x2d\xda\x42\xf5\x40\xbb\x1e\x13\x14\xb2\x74\xb4\xdd\xaa\x61\x1e\x07\x03\x43\xc7\xcf\xf8\xfc\x56\x7c\xbf\x66\x87\x06\x7a\xd3\x38\x7e\x9f\x5a\xbe\x69\x98\xd3\xfd\x61\xd6\x52\xab\x2c\x2a\x2a\x4e\xee\x6b\xab\x34\x17\xd1\x78\x58\xaf\x1d\x84\xae\xf3\x6b\x3f\x2c\xd8\x38\x8e\x87\x09\x7f\x1f\xca\xed\x9b\x7f\x9e\x1e\xd3\xfd\x41\xa9\xf7\xd4\xec\xfe\x19\xbe\x2a\x3b\x2b\x8d\xa6\x68\xb9\x73\xd3\x6e\xa8\x3d\xe6\x2b\x72\xb0\xfc\xe5\xa1\xa1\xbd\x10\xcc\x10\x1e\xfe\x2e\x43\x84\x3f\x89\x9d\xbe\xc3\xa5\xff\xb0\x75\x5f\x02\x00\x00\xff\xff\x89\x2a\x06\x6f\x6e\x07\x00\x00")

func cmdDefinitionsTmplObjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplObjectTmpl,
		"cmd/definitions/tmpl/object.tmpl",
	)
}

func cmdDefinitionsTmplObjectTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplObjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/object.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb5, 0x31, 0x99, 0x29, 0x2, 0xc4, 0x9b, 0xd3, 0x91, 0x35, 0x57, 0x52, 0xe1, 0x99, 0xff, 0xb1, 0xec, 0xc7, 0x4, 0x8a, 0x19, 0xaa, 0xa6, 0x88, 0xca, 0x41, 0x3a, 0x9a, 0x1b, 0x95, 0x45, 0x67}}
	return a, nil
}

var _cmdDefinitionsTmplOperationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\x4b\x4f\xc3\x30\x0c\xbe\xe7\x57\x58\x55\x0f\xad\x34\xba\x3b\x12\x27\x1e\x12\x17\x98\xe0\xc0\x11\x99\xcc\x2b\x16\x6b\x13\x12\x83\x36\x85\xfc\x77\x94\x36\x1d\xdb\x04\x42\x42\xdc\x12\xdb\xdf\xc3\x0f\x8b\xfa\x05\x5b\x02\xd9\x5a\xf2\x4a\x71\x67\x8d\x13\xa8\x14\x00\x40\xa1\x4d\x2f\xb4\x91\x62\xfc\xb1\x29\x54\xad\x54\x08\x27\xe0\xb0\x6f\x09\xca\xc7\x19\x94\x0c\xa7\x67\xd0\x5c\xf7\x42\x6e\x85\x9a\x3c\xc4\xa8\x42\x80\x92\x9b\x0b\xf2\xda\xb1\x15\x36\x7d\x0a\x26\x05\xc8\x19\xf6\x76\x8d\xdb\x1b\xec\x08\x62\x04\x9e\xc0\x10\x06\xa5\xa4\xc0\x2b\x30\x0e\x2a\x7a\x4d\xf5\x43\x61\xe1\xc9\xbd\xb3\x26\x57\xd4\x47\x71\x31\x0e\xdb\x14\x8f\x71\xc0\xdf\x8b\xe3\xbe\xad\x6a\xf0\xc3\x63\xc7\x49\xfd\x32\x19\xd9\xfd\xf7\xba\xa0\xee\x89\x96\xa9\x93\x92\x9b\xcb\xe1\x9d\xb9\xc6\xda\x5c\x70\x64\xfc\x07\xe2\x7d\x5e\x63\x33\xe9\xad\xf5\x13\x62\x3e\x1f\x18\x8d\x1d\xfd\x7f\x80\x98\x05\x7a\x8d\xeb\x34\x8b\x9c\x39\x9a\xdd\xe4\xe2\x5b\x4c\x95\x33\x57\xc6\x75\x28\x0b\x74\xd8\x25\xad\x1a\x0e\x13\x77\xe4\xdf\xd6\xe2\x1f\x58\x9e\x17\xe3\xd2\x0f\xc6\x57\x24\xc8\xaf\xfe\x12\xfa\x7c\xbc\x8a\xbf\x58\xdd\x83\x57\x5a\x36\x90\x0f\xac\xc9\xb1\xd9\x3f\x77\x12\xc2\xb4\x9a\xa8\xbe\xf6\xf4\x19\x00\x00\xff\xff\xad\xfc\xbb\x7f\xf3\x02\x00\x00")

func cmdDefinitionsTmplOperationTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplOperationTmpl,
		"cmd/definitions/tmpl/operation.tmpl",
	)
}

func cmdDefinitionsTmplOperationTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplOperationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/operation.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd0, 0x3a, 0x99, 0xf6, 0x41, 0xfa, 0xdd, 0xc0, 0x2e, 0x3, 0xc1, 0x31, 0x43, 0x68, 0xa5, 0x42, 0xca, 0x90, 0x7f, 0xb7, 0x15, 0x72, 0x87, 0xb, 0x71, 0xa6, 0x44, 0xc8, 0x5a, 0x98, 0x33, 0xa4}}
	return a, nil
}

var _cmdDefinitionsTmplPairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\xcf\x8b\xdb\x3a\x10\xbe\xfb\xaf\x98\x67\xc2\xc3\x06\xc7\x7e\xbc\xa3\x4b\x0e\xa5\x9b\x43\x29\xdd\xec\x21\xb4\x87\x65\x29\x5a\x79\xe2\x88\xc8\x92\x90\x14\x6d\x83\xeb\xff\xbd\x48\x4e\xb2\x4e\x9c\xec\x8f\x43\xa1\x3a\x59\x9a\x6f\x66\xbe\xf9\x46\x23\x17\x05\x7c\x92\x15\x42\x8d\x02\x35\xb1\x58\xc1\xe3\x0e\x6a\x79\xdc\x03\x13\x16\xb5\x20\xbc\xa0\x4d\xf5\x01\x6e\x16\x70\xbb\x58\xc2\xfc\xe6\xf3\x32\x8f\x14\xa1\x1b\x52\x23\x28\xc2\xb4\x89\x22\xd6\x28\xa9\x2d\x24\x11\x00\x40\x4c\xa5\xb0\xf8\xd3\xc6\x51\xbf\xad\x99\x5d\x6f\x1f\x73\x2a\x9b\x82\x48\x33\xad\xd0\x15\xb5\x9c\x1a\x2b\x35\xa9\xb1\x70\xff\x17\x6a\x53\x17\x28\x2a\x25\x99\xb0\xf1\x3b\x7c\xa8\xc6\x0a\x85\x65\x84\xbf\xc7\x6b\x6d\xad\xa2\x9c\xe1\xdb\x73\xd9\x9d\x42\x13\x47\x69\x14\x15\x05\x7c\xe4\x1c\x88\x23\x8c\x93\x47\xbe\xaf\x3f\x8f\xa8\x14\xc6\x97\xdf\xb6\x53\xd0\x44\xd4\x08\x93\x1f\x19\x4c\x1c\x94\x33\xc8\xef\x3c\x06\xba\x2e\x64\xf3\x88\x89\x12\xa4\x41\x6f\x9b\xb8\xfc\xd6\x7f\xfe\x02\x2b\xef\x88\xa1\x84\x1f\x70\x45\x01\x6d\x7b\x40\x76\x1d\x3c\x31\xce\xc3\x89\xcb\x6f\xd0\x50\xcd\x94\x65\x52\x3c\x47\x1d\x40\x67\x10\xf7\xc0\xdb\xfe\x20\x0e\xb4\x50\x54\x1e\x9d\x46\xaf\x90\x7c\x0b\xc1\xa2\x80\xef\xcc\xae\xc7\x04\x89\x52\x7c\x07\x27\xd9\xc1\x11\xbe\x45\xb0\x12\x16\x81\xb2\xf1\xde\xcb\x35\x33\x41\x3b\x60\x06\xb6\x06\x2b\x6f\xbf\x58\xdd\x6a\x2b\xe8\x28\x59\xe2\xf6\xe0\xe5\x4e\xf9\x7d\x0a\xa1\x45\xa1\x08\x68\x83\x22\x1a\xed\x56\x8b\xc1\x79\x7f\xec\xd7\x17\xdc\x95\x27\x8a\x65\x47\xd3\x37\xcf\xb5\x04\xd7\x9f\x74\x51\x37\x94\xce\x13\xbf\x23\xda\x60\x5f\xab\x0a\x9f\x04\x36\x53\x07\x0d\x51\xbe\x82\x70\x1b\xc0\x70\x46\x31\xef\x99\x07\x7c\xd2\x78\xc0\xbd\xb1\x9a\x89\xfa\x21\x4c\xd5\x8a\x50\x6c\xbb\x14\x92\xfb\x87\x67\x8a\x19\xa0\xd6\x52\xa7\xfb\x12\xfa\x68\xe5\x0c\x1a\xb2\xc1\x33\xe0\x7f\x19\x70\x14\x49\x93\xa6\xfd\x90\x39\xa2\xbd\x73\x1f\xa0\x3f\x5a\x49\x0d\x9b\x0c\x42\x7b\xfb\x76\x37\xf0\x2c\x82\x77\x50\x0e\x06\x64\x8e\x26\xf3\xc4\x2c\x5d\xc3\x66\x80\x7e\xfd\x5e\x0f\x91\x6f\xb9\xdf\x87\x45\x89\xc1\x93\x66\x94\x27\xe6\x01\x1f\x1d\xb2\xba\x3c\xf1\x3a\xa4\x03\x72\xa3\x50\xc7\x7b\x31\x8e\x15\x74\x75\x30\x03\xed\x46\x36\x4f\x9d\xad\xbc\x7b\xe8\x9a\x3e\xa7\x7a\xcc\xd1\x37\xf2\x5a\xf0\xd0\x45\x98\x41\xdb\x1e\x23\x75\x5d\xa2\x5d\x7a\x11\xcf\x56\x01\xfe\xcf\x0c\x04\xe3\x17\x8a\x3a\xac\xfd\x7d\x16\x8c\x67\xf0\xef\xdc\x77\xf9\x3a\x76\xa1\x4a\x88\xc3\xf5\x8c\xb3\xab\xa0\xb9\xd6\xa5\x4f\x7d\x1d\xf1\xc2\x9c\x9c\x2f\xaf\x77\x79\x78\x76\xf6\xe2\xbf\x90\xfb\x74\xcc\xce\xd7\x58\xf5\xf1\xc9\x60\x2c\xcf\x4d\x15\xae\xc8\x96\xdb\xcb\xed\xf9\x33\x32\xce\xb5\xf6\xc3\xe0\x0b\xff\xca\x4c\x43\x2c\x5d\xff\xed\xb2\x76\x27\xb3\x7d\x41\xca\xab\x32\xfa\xbf\x3b\x13\x5b\x8c\xc6\xa1\xfa\x17\x6b\xe6\x7f\x03\x28\xaa\x24\x6c\xb3\xe1\x1b\x1c\x6a\xdf\x64\x07\xa6\xca\x75\xe9\xfe\xa5\x1d\x3e\xda\x7b\x3f\xc1\x78\xd4\x45\xbf\x03\x00\x00\xff\xff\xc3\x68\x74\x7f\xa6\x08\x00\x00")

func cmdDefinitionsTmplPairTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplPairTmpl,
		"cmd/definitions/tmpl/pair.tmpl",
	)
}

func cmdDefinitionsTmplPairTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplPairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x38, 0x4c, 0x4b, 0x10, 0xe5, 0xf1, 0x10, 0x79, 0xd6, 0x6d, 0x42, 0x94, 0x3e, 0x38, 0x6a, 0x55, 0x70, 0xd2, 0x39, 0xe3, 0x35, 0x27, 0x83, 0x91, 0xe4, 0xd6, 0xc9, 0xeb, 0x93, 0xf2, 0x88, 0x4b}}
	return a, nil
}

var _cmdDefinitionsTmplServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\x5b\x6f\xdb\x38\x16\x7e\xf7\xaf\x38\x15\xdc\x85\x54\x38\xd2\xee\x3e\xba\xf0\xcb\x26\x6d\x5a\x74\x37\x0d\x36\xd9\x2d\x30\x6d\x10\x30\x12\x65\x13\x96\x48\x95\xa4\x95\x78\x5c\xfd\xf7\x01\x2f\x92\x28\x59\x72\xed\x4e\xda\xf1\x43\x60\x8b\xe7\xfa\x9d\x2b\x95\x28\x82\x73\x96\x60\x58\x62\x8a\x39\x92\x38\x81\x87\x2d\x2c\x59\xf3\x1b\x4a\x82\x80\x50\x89\x39\x45\x59\x14\xe7\x49\x24\x30\x2f\x49\x8c\x5f\xc3\xc5\x47\xb8\xfa\x78\x0b\x6f\x2e\xde\xdf\x86\x93\x02\xc5\x6b\xb4\xc4\xb0\xdb\x41\x78\x85\x72\x0c\x55\x35\x99\x90\xbc\x60\x5c\x82\x3f\x01\x00\xf0\x62\x46\x25\x7e\x92\x9e\xf9\x45\x98\x37\x31\xdf\x96\x44\xae\x36\x0f\x61\xcc\xf2\x08\x31\x71\x96\xe0\x32\x5a\xb2\x33\x21\x19\x47\x4b\x1c\x95\xff\x8c\x8a\xf5\x32\xc2\x34\x29\x18\xa1\x35\xf7\x51\x3c\x31\xc7\x09\xa6\x92\xa0\xec\x14\xae\x95\x94\x45\x9c\x11\x7c\xbc\x2e\x8b\x88\x30\xf4\xe1\x11\x1c\x72\x5b\xd4\xe4\x85\x38\xc6\x2e\x44\xb8\xf0\x26\xc1\x64\x52\x22\x0e\xf7\xd0\x7a\x16\x5e\x73\x56\x92\x04\x73\x7b\x52\xe3\xd4\x7f\x7e\x63\x84\xd5\x3f\x6b\x93\xc3\x1b\xf3\xe5\x0d\xe7\xac\x3e\x6b\x01\x08\x3f\x16\x92\x30\x2a\x26\x93\x28\x82\xdb\x6d\x81\x81\x08\x90\x2b\x0c\xca\x7e\x48\x19\xef\x84\x3b\x66\x54\x48\x43\xb6\x00\xcf\x39\xf1\x34\xbf\xd5\x04\xa8\x44\x24\x43\x0f\x19\x06\xed\x55\x68\xf9\xfc\xc9\x6e\x77\x06\x1c\xd1\x25\x86\xe9\xfd\x0c\xa6\x25\xcc\x17\x10\x5e\x2b\x1a\x25\x5d\x61\xa5\x28\x48\x0a\x94\x49\x98\x96\xe1\x65\xc6\x1e\x50\xe6\x9e\x4d\x0b\xaa\x34\xce\x17\xea\x58\x2b\xff\x06\x92\x5d\x23\x11\xb7\x74\x51\xa4\x8c\xb6\x94\x55\x05\x8f\x24\xcb\xf4\x93\x32\xbc\xc0\x22\xe6\x44\xbb\x5c\x53\x2b\xf5\x1d\x72\xe3\xd9\xb4\x76\xed\xde\x70\x36\x8e\xd6\x96\x60\x9a\x28\x11\xce\xd7\x60\x04\x04\x42\x53\xf6\x3d\x10\xde\x2b\x9a\xe7\x04\xc1\xca\x50\x3e\x13\x51\x64\x68\x5b\xc7\x10\xec\xc7\x11\xb4\x18\x21\x73\x7c\x53\x3f\x95\x8d\x06\x8c\x9b\x98\x15\x3d\xa5\xe6\xe0\x1c\x49\xbc\x64\x7c\xbb\x7f\x36\x06\xef\xd9\x29\xf0\x1e\xce\x9f\x31\xd8\x8e\x81\x2c\x8a\xe0\x13\x91\xab\xfd\xbc\x41\x45\x91\x6d\xa1\x63\x24\x94\x28\xdb\x60\x90\x0c\xea\xe2\x51\xb5\xb3\x22\x42\xa7\xbb\x2a\xa0\x8d\xc0\x89\x3a\x1f\x4c\xba\x74\x43\xe3\x3d\x65\x7e\x69\x89\x75\x71\x55\x55\xa0\xf3\x12\x76\x4d\xb8\x38\x96\x1b\x4e\x4d\xb6\x36\x0f\x3f\xe0\xed\x7c\x2f\x81\x67\xcd\xf1\xff\x95\x9d\x73\x28\xcd\x93\x6a\xd2\x81\xd3\xf9\x3a\x02\xec\x09\x39\xf9\x5c\x29\x79\x62\x46\x2a\x06\xdb\xe7\xae\xac\xfa\x82\x13\x2a\x53\xf0\x5e\x8a\x97\xc2\x03\x7f\x20\x59\x03\xfd\x74\x20\x53\x03\xa7\x81\x5c\x62\xb9\x9f\x0b\x4b\x2c\x07\x33\x21\xe5\x2c\x87\x1c\x4b\x94\x20\x89\x42\x2d\x42\x07\xb9\x27\xc4\xcf\xf7\x0c\x56\x81\xf6\x3b\x81\x9f\xc1\x03\x63\x59\x00\x36\xf4\xe5\x0c\xd8\x5a\xf9\x95\x87\x97\x58\xfa\x75\x09\x76\x65\xb8\x4a\x02\xcd\x46\x52\x78\xc1\xd6\x56\x46\x8b\xfb\x6f\x98\x33\x9d\x14\x2e\xea\x36\xb3\x8c\x15\x2e\xc5\x0c\x52\x94\x09\xdc\x02\x9f\x89\x71\x46\x6d\xfe\x37\xf8\xfd\x10\x7f\x1b\x38\xf3\xd7\x0a\x28\xc3\x2e\x04\xc1\x0c\x24\xdf\x60\x4b\x58\x47\x44\x18\x30\x5b\xec\x75\x48\xc4\x48\x48\x08\x95\x6c\x28\x24\xe2\x88\x90\xcc\x60\xaf\x18\xf7\x89\x2c\xb6\xd6\x85\x3c\xbc\x39\x2a\x3a\x33\x28\x03\x07\x80\x91\x72\x1c\xac\x46\x25\x4d\x14\x28\xc6\x9d\x92\x94\x38\x2f\x32\xb5\xc4\x79\xaa\xf7\x78\x90\xa3\x35\xbe\xc9\xd4\xfc\xf1\x87\x2a\x31\xd0\x0f\xf1\xe3\xb8\x08\xe1\xca\xa8\x25\x4c\xcb\xf0\xed\x86\xc6\x23\x9a\x15\xb0\xba\x0f\x7e\x8f\xb5\xef\x63\x82\x53\x42\x5b\xbd\xdd\x09\xa7\x9c\x26\x34\xc1\x4f\x10\xc2\xdf\x47\xba\xc9\x54\x45\xd7\x25\xfc\x87\x96\x5d\x1f\xf7\x31\x34\xe4\xbd\x19\x38\x8e\xa0\x32\x62\x5a\xee\x77\x1e\xfb\x75\xc8\x8b\xc3\x4e\xb8\x67\x29\x1d\xb7\x7b\x9a\x52\x5d\xf3\xd3\x94\x0e\xf5\xd2\xba\x22\x8a\xa6\xf1\xdb\x2c\x53\x6c\x55\xf5\x1f\x54\xc0\x8a\x65\x89\x00\xa4\xe6\x57\x77\x17\x33\x4d\x05\xf1\x43\xcc\x0b\xc8\x51\xf1\x59\x48\x4e\xe8\xf2\x4e\x48\xbe\x89\xe5\xae\x6a\x87\x4e\x14\xc1\x7f\xf1\xd7\x0d\xe1\x38\x71\x64\x0e\x60\xae\xa7\xa1\x75\xa3\xe1\xe8\xb7\x7b\x44\x78\xf8\x76\x93\x65\xb6\x58\xe6\xd0\x28\x74\x06\x59\xaf\x7b\x58\x2b\xcc\xf4\x45\xd9\xf1\x56\x34\x1c\xcf\x68\xc5\x65\x73\xad\x3a\xda\x8c\x96\xe5\x39\xec\xf8\x6e\x42\xd4\x2b\x7d\x81\xb8\xda\x4b\x8c\x48\xcd\xa3\x97\xfc\x31\x2e\x43\xe7\xac\x20\xda\x3f\xf8\x7c\xa7\x16\x8e\xc9\x4f\x4c\x87\xfd\x94\xef\x1c\xdb\xa6\xfc\x4b\x92\xe3\x1d\x12\x87\xcd\x52\xc3\xfa\x58\x0f\x4e\xf6\xe1\xf9\x52\xab\xdd\xdd\xb4\xfa\x0b\x9c\xa2\x4d\x26\xff\x8c\xa7\x7b\xf6\x3e\x53\xf8\xdc\x6c\xe6\x02\x5f\x8f\x24\xa7\x1e\xfe\x9a\x02\x5e\xe9\x4d\x59\xe8\x96\xad\xc7\xfe\xab\x91\x8c\x6e\xb7\x80\x43\xa2\x7d\x56\xc8\x3a\xcb\x03\xf0\xc7\x84\xcd\x00\xab\x0b\x75\xd0\xec\x01\x42\x01\x3a\x5f\xc0\xdf\x46\x18\x7a\x75\x34\x07\xa5\x67\xe6\xfa\xac\x37\x17\xa1\x97\x3d\xb4\xc6\xbe\xd3\x84\xf5\xdb\x99\x14\xc5\x78\x67\xf7\x3b\x75\x31\xbf\x57\x7b\xca\x7c\x61\x33\x41\x5b\xbd\x1b\x0a\xb9\xbf\x42\xe2\x9a\xe3\x94\x3c\x19\x53\xbc\x2b\xfc\xe8\x05\x6e\x08\x48\xaa\x84\x99\x3d\x73\x7c\x2e\x7c\x2e\xc3\x0f\x78\x7b\xf7\xda\xd9\x2c\xeb\x8f\xdd\x82\x28\xc9\x66\xed\x7b\x87\x2b\xfc\xa8\x30\xfc\x1f\x15\x9b\xa2\x60\x5c\xe2\x44\xbf\x82\xf0\xed\x06\x04\xcd\x16\x34\x92\x50\x06\x0e\xab\x15\x16\x50\x86\x7a\xb1\x74\x11\xb3\x5e\x32\xde\x6d\x2a\x9d\x6a\x1e\xac\x0a\x35\x01\x4b\x70\x60\x6d\x9e\xb2\x75\x9b\xe9\xee\xbe\x62\x93\xf2\x1d\xa2\x49\x86\x81\xef\x77\xbc\x93\xbb\x9d\x5a\x10\xd6\x78\xab\x89\x46\x0a\xa7\xc5\x42\x87\x67\x51\x63\x62\x98\x3b\xc3\xe2\xac\xaa\xee\xdc\x80\x9e\x16\xa4\xda\x42\x13\xa1\x61\xf1\x43\x61\x53\xe0\xef\xeb\x51\xa5\x10\xea\x24\x52\xfe\xe9\x3b\xbf\xdd\xf1\xdd\x0e\xd0\x97\xd7\x4b\x81\x16\x6c\xb6\xdf\xca\x4f\x9e\xf1\x3f\x17\xec\x0e\x06\xd6\xff\xba\xa1\x36\x10\x34\x77\x9a\x9f\x02\xd2\x72\x60\x56\xfc\xc8\x9c\xf8\x65\x30\x1d\x33\x94\x8e\x86\x73\xa0\x7d\xfc\x28\xc2\x8e\x6d\x07\xed\x1a\xa9\xb1\x01\xa5\x8d\xc6\x11\x49\xa3\x5d\xb0\xdf\x7f\x6c\x05\x1b\x25\x33\x55\xc9\xb6\x19\x8e\x5d\xac\x9c\x9b\xd9\x5f\x7d\xb9\xaa\x4d\x39\xe1\x82\x35\xea\xce\x73\xdc\xb2\x8e\xbd\x64\x39\x3b\x87\xf9\x3a\xf4\x0e\x39\x8a\x6a\x72\xfd\x1a\xb0\xb6\xd3\x6c\x29\x31\xc7\x0a\x01\x04\xf6\xbf\x23\xf0\xb0\x55\x1e\xa9\x4c\x70\x5e\x4c\xf8\x02\x5e\xed\x76\xd3\x82\xd6\xaf\x1b\xec\x36\xb2\xdb\x29\x9d\xd7\x88\xa3\x5c\x84\x37\x7a\x1d\xb0\x2f\x8d\xcc\x64\x51\xa9\xe0\x1e\xb4\xf9\x18\xcb\x27\xe5\x9f\xd5\x1a\xfe\x0b\xc5\xeb\x25\x67\x1b\x9a\xf8\x41\xff\x15\x8e\x08\x5b\x8d\x9f\x88\x5c\x9d\x1b\x1e\x3f\x96\x4f\x33\xe8\x58\x70\x8e\xb2\x0c\xf3\xba\x6a\x86\x60\x72\xf8\x0f\x20\x76\xc8\xe7\x9e\x05\x8d\x07\xf6\x59\xcf\xa2\x93\x30\x49\x70\x8a\xb9\x56\xee\x07\xdd\xca\xb5\xed\x4b\xae\xea\x9c\xb0\xe2\xaf\x91\x5c\x19\xa7\xfb\x65\x6b\x9b\x04\xa2\x09\xf8\xf8\xab\x65\xf6\xbc\xc0\xfe\xa2\xe0\xd9\x7f\x05\x74\x97\xad\x96\xdd\xb2\x2c\xc0\x9b\x7d\xf1\xbe\x78\xde\x90\x82\x5e\x53\x53\x1f\xcc\x39\x2c\x40\x84\x29\xe3\x39\x92\x66\x62\x7b\xc6\x75\x95\xc7\x55\xe5\xe9\xdd\xb4\x55\x50\x55\xe0\x34\x39\x27\xfa\x7a\xdf\x29\xe4\xe1\x85\x59\x7d\x58\x21\x67\x56\xef\xc1\xd5\x59\x8f\x9e\xc0\x6d\xf9\x8a\xe9\xc5\x42\xb5\xaa\xc1\x6d\xc4\xe9\xbd\x83\x39\xd9\xad\xce\x73\x94\x63\x55\x9c\x03\x79\x79\xcb\x49\xfe\x6f\x24\xa4\x4d\xd0\x37\x34\x11\x8f\x3a\x8d\xf2\x1c\x55\x95\x72\x20\xd8\xeb\x95\x7f\x04\x00\x00\xff\xff\xce\xd6\x0e\xdc\x19\x1d\x00\x00")

func cmdDefinitionsTmplServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplServiceTmpl,
		"cmd/definitions/tmpl/service.tmpl",
	)
}

func cmdDefinitionsTmplServiceTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/service.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9d, 0x44, 0xdb, 0x3d, 0x2a, 0x2d, 0xbc, 0x23, 0x82, 0x2c, 0x84, 0x21, 0x3, 0x26, 0x5b, 0xac, 0x43, 0xbc, 0xa4, 0xd1, 0x35, 0xae, 0x9e, 0x80, 0xc1, 0xf6, 0xdf, 0x99, 0x89, 0x65, 0xcc, 0x6}}
	return a, nil
}

var _definitionsInfosHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xb1\x72\xd3\x40\x10\xed\xfd\x15\x6f\xae\x49\x93\xa4\x22\x74\x54\x98\x82\x02\x68\x3c\x43\x19\x9d\x75\x6b\x7b\x91\xb4\x2b\xee\xf6\x70\x04\xc3\xbf\x33\x27\xd9\x13\x3b\x89\x95\x21\x6a\x34\xb3\xb7\xef\xed\x7b\xb7\x4f\x5a\xb0\x6c\x14\x4e\xd7\x3f\xa8\x36\x07\xd7\x91\x79\x07\x57\xab\x18\x89\xdd\x74\xe1\xce\xe1\xcf\x02\xb0\xa1\x27\x1c\x9f\x0f\x70\xc9\x22\xcb\xd6\x2d\x80\xc0\xa9\x6f\xfd\x70\x2f\xbe\xa3\x72\xf2\x71\x82\x7e\x59\xde\xb9\xc5\xdf\x57\xe8\x0b\xeb\x09\xff\x09\xef\x25\x24\x99\xdf\xfe\xaf\xa2\x4f\x2b\x3f\xc3\xc8\xe1\x9c\xef\x8c\x8a\x1e\x7a\x8d\x56\x8a\x16\x33\x2d\x80\x5a\xbb\x8e\xc4\x4a\xd7\xe7\x25\x38\xc1\x76\x84\x2c\xfc\x33\x13\x1a\x1a\xc0\x82\x64\x1a\xfd\x96\x6e\x2f\x8f\x2c\xc2\xde\x38\xf4\x6b\xf1\xc4\x09\xc4\xb6\xa3\x38\x4e\xf7\xeb\xa4\x6d\x36\x42\xef\x6d\x07\x9d\x8a\x91\x5a\x6f\xfc\xeb\x50\x34\xdd\xfb\x18\xd2\x51\xda\x55\xc2\x77\x8d\xcd\x92\x23\x02\xf5\x24\x21\x41\x05\x39\x51\xbc\x4a\x60\xe9\xb3\xcd\x88\x4f\xfc\xfb\xc9\xce\x58\xec\xfd\xbb\xcb\x00\xf3\x71\x4b\x36\xe3\xf7\xc4\xde\x6a\xec\x3d\xde\x6b\x1a\xba\x96\xa5\xc1\xc4\x80\xcd\xe8\x8d\x13\xa6\x09\xd7\x50\x69\x07\xd0\x03\x27\xc3\x7e\x47\x72\xa8\x4f\x43\x38\xa1\x60\x67\x8c\x3c\x09\xdf\xa4\xea\xdb\xd8\xb4\x2a\x47\xaf\x6c\xa2\xf4\xa0\xd6\xdc\x06\xac\x09\x2a\x04\xdd\xa0\xda\x70\x4b\xd5\x35\xaa\xc0\xb1\xbc\x8a\x84\xaa\xec\xa4\xca\xd2\x88\xee\xa5\x9a\x11\x94\xfb\xe0\x8d\xc2\xbd\xb7\xf3\xfb\x35\xee\xe8\x76\xc5\x1d\x3d\x42\x0f\x8b\x7c\xc4\xb6\x5a\x7b\x63\x95\xf9\xaf\xe9\x19\xec\x59\x12\x5f\x0a\xe2\xc1\xfd\x45\x92\xbd\xc6\xe6\x26\x70\x7c\x3b\x51\x32\x6f\x9c\x8c\xeb\xf1\xdf\x90\xc5\x66\x03\xf6\x32\x6e\x26\x97\xff\x02\x00\x00\xff\xff\x6d\x6a\x6e\x6a\xe5\x04\x00\x00")

func definitionsInfosHclBytes() ([]byte, error) {
	return bindataRead(
		_definitionsInfosHcl,
		"definitions/infos.hcl",
	)
}

func definitionsInfosHcl() (*asset, error) {
	bytes, err := definitionsInfosHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/infos.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd5, 0xc2, 0xfb, 0x71, 0x3e, 0x4e, 0x9f, 0x8c, 0xdf, 0x49, 0x4, 0xb4, 0x97, 0xef, 0x61, 0x2, 0x10, 0xb0, 0xb8, 0xb8, 0x46, 0x8b, 0x7b, 0xc6, 0x83, 0x42, 0x20, 0x63, 0xb7, 0x84, 0xe9, 0x1c}}
	return a, nil
}

var _definitionsOperationsHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x3d\x8b\xe4\x46\x10\xcd\xe7\x57\x14\x4a\x0e\xcc\x30\x91\x71\xb6\xd1\x81\xc1\xc1\x61\x73\x1b\x38\x30\xc7\xd2\x2b\x95\x46\x65\xa4\x6e\x51\x5d\x9a\xb9\x3d\xb3\xff\xdd\x54\xeb\xfb\x73\xa4\x63\x37\x58\x06\xf5\xab\x8f\xf7\xba\xfa\xa9\x75\x22\x2b\xc8\xa9\x89\x11\xa2\xd8\x95\x84\x1c\xc1\x7f\x27\x80\x04\x7d\xcc\x54\x0a\x39\x0b\x4f\x10\x91\x07\xc9\x10\x7a\x70\xea\x18\x3e\xbb\xf2\xed\x12\x9d\x4e\x00\xae\x0c\xd1\x6f\x75\xec\x2c\xfa\x4e\x79\x0e\xba\x0e\xc6\xc2\x9f\xaf\xff\x62\x2c\xe0\x18\x8a\x2a\x17\x2a\x73\x04\x57\x3f\x22\x1b\x8a\x78\xe4\x1b\xc5\x78\x89\x42\xaa\xd2\xb0\x29\x3c\x84\xbf\x27\xf8\x27\xf2\x1c\x47\x67\x88\x12\x2f\xd1\xb7\x13\xc0\xfb\xe9\x7d\x48\x21\x21\x7e\xc9\xc9\xcb\x3a\x8d\xca\x63\x12\xba\x4f\x88\x31\x16\xc7\x6f\xf0\x6a\xf4\x99\x17\xc7\xe6\xda\x95\x07\x71\xa0\x99\x9a\xe6\x3c\x54\x36\x41\x06\xa3\x71\x3d\x69\x45\xbc\x24\xc4\x5b\xc4\x19\xa5\x62\x5b\x27\x33\xe0\x4b\x8c\x29\xa5\xb8\xc9\xb3\x44\x51\xf3\x7d\x0b\x2b\x8c\xbe\xca\xc5\x77\x2b\x8e\xd6\x48\x7b\xbc\x16\x68\xc5\x7f\x34\xfb\x36\xef\x84\x3e\x00\x16\xaf\x98\xc0\x60\x5f\x6a\x20\x6a\xeb\x13\x71\xba\xe6\xb6\x54\x1a\x57\xbb\x91\xf9\x29\x81\xfc\xa2\x40\x64\x13\xfc\xfe\xd2\x77\xb8\x7b\xc2\x43\x60\xab\x50\x1d\xbe\x97\x3c\x59\x92\x97\x51\xe1\x2d\xf6\x8a\xd6\xc3\xb1\x56\x70\xae\x41\x69\x24\x5b\x13\x01\xaf\x8d\x0a\x75\x2b\x77\x26\xc1\xfd\xbd\x04\x38\x18\xad\xa8\x67\x52\xdc\xb1\xc6\xb4\xfa\x19\x22\xd6\x7f\x21\x4a\x7f\x78\xfa\x81\x4b\x3b\x53\xb8\xdb\x91\xfd\xf8\xe2\x6e\xd8\x1f\x3e\x0d\xde\x22\xa2\xeb\xda\xfb\xc7\xd9\x4b\xc9\x98\xd2\xf7\xbd\x67\xac\x46\x1f\xb4\x97\x3a\x68\xe2\x30\xf5\xc3\x0f\x35\x99\x26\xe5\x11\x9f\x69\xd8\x1f\xb5\x9a\x5d\x32\xcc\x7c\xa6\xd3\x61\xbf\xd5\x4c\xfa\xdb\xed\x36\x3f\xa3\xd0\xb2\xd1\x30\x9a\x38\x3b\x32\xd0\x5f\x35\xa0\xdf\xec\x10\xbf\xd5\x77\xc9\xee\x46\x89\x9e\xce\xbb\x79\x3b\xc3\x3d\xa3\x38\x83\xd8\x58\x08\x91\xa1\x40\x3d\x50\xc7\x5d\xa3\xe2\x7c\x89\xd2\xc4\x35\xc3\x8a\x35\x39\x3c\x81\x70\x85\x6d\xe3\xe6\xd5\xb1\xec\x31\x97\x00\xd4\x09\x7d\xe8\x21\x43\x07\x8b\x5d\x51\xe6\x28\xb8\xa7\x42\x8b\xed\x8b\x80\xb1\x09\x14\xc8\x57\x54\x81\x8a\xc6\xd3\xe0\x77\xca\x57\x5d\xa0\xab\x3f\xd1\x22\x0c\xed\xf2\xfe\xea\x2e\x14\x86\xac\x18\xb2\xcd\xd5\x46\x0d\x74\x32\xf0\x7e\x70\x61\x62\x34\xb2\x69\x60\x35\x02\x0c\x58\xbc\xb7\x89\xf4\xa5\xe4\xc5\xd8\x55\x07\xb3\xa6\xc0\xb5\xa1\x15\xc7\x38\x12\x36\x41\x95\x6a\xab\x87\x1a\xa1\x62\x1e\xad\xdf\xd6\xb8\xe2\xe6\x7e\x5d\x51\xe7\xe1\x66\x72\x4a\xe6\x35\xc2\x21\xd9\x76\xec\x43\x7c\xf5\xdc\x3f\xf4\x05\x93\xe7\xf3\x4e\x5a\x67\x92\x8c\xfc\xa4\xa3\x79\xd9\x45\x73\xf0\x62\x84\xbc\x50\x4c\xc6\xee\x77\x88\xe7\x36\xca\xe4\xfd\xe8\xf8\xfe\xe1\xf6\x51\xa8\xac\xb4\xcd\x7e\xf2\xd0\x85\xf9\x33\xf8\x2a\xce\xc0\x78\x78\xa6\x1f\x78\x86\xcf\x8a\x5c\xa5\xd3\x44\x2d\x93\xaa\x75\xda\x4f\x68\x72\x20\x7a\x52\xfb\x67\xb1\xfb\x88\x48\xd9\x15\x0f\xc6\xa3\x73\xbc\x76\x04\x0a\x14\x93\x18\x31\x3b\xde\xa6\x71\xc5\xac\xfe\xd1\x92\xfc\xe4\xa1\x8d\x5e\xd1\x4a\x97\x47\xd5\x18\x4d\xb2\x5d\xc9\x24\x41\xa4\x94\x72\xdd\xa2\x41\xee\x45\x26\x67\x88\xee\x2b\xd3\x6e\x47\x85\x75\xd7\xb6\x0a\xeb\x7a\xb8\xe4\x49\xa6\x2f\x61\x3d\x85\x64\x53\x07\x2e\xed\xaf\x4c\xc7\x5f\x22\x6e\x7e\xf1\x7c\x7c\xd5\x54\xce\xad\x29\xa7\xeb\xa6\xdc\xf2\x5f\xbb\xfb\xdb\x6e\x3e\x4f\x29\x61\x9e\x84\x0f\xa3\xba\xba\xbc\x95\xa8\x45\xbd\x30\xd9\x6b\x74\x7a\xef\x10\xad\x1d\xac\x22\x90\x27\x39\x90\xd9\xf1\x00\x50\xdf\x71\x47\x10\xb2\x32\x00\x84\x99\x18\xad\x3f\xd7\xf3\xf4\x45\x17\x7a\x9c\x9d\x25\xf9\xed\xd7\xe1\xb2\xfa\xdc\x76\xaf\x6e\xbc\xfe\x4b\x7d\x4a\x86\x00\x5a\x44\xfc\x21\xc8\x46\x46\xac\x4a\x43\xec\xc7\xe0\xcb\xe5\xf2\x97\xa1\x31\x48\xb2\x07\x2d\x0d\x6f\xae\xab\xa0\x89\xc2\xe4\x2e\x5f\xd1\x24\x38\x2c\xa5\x6f\xe4\xb1\x86\xcd\x4d\x60\x00\x99\x92\x6b\x20\x0b\xec\xc2\xe7\xc8\xa6\xda\xfa\x1d\xb0\xdd\x74\x6f\x8b\x4b\x7b\xdb\xd9\xf6\x28\x62\xd6\x61\x63\x2c\x4b\x2d\x86\x77\xd7\x52\xea\x21\x4a\x2f\x6d\xdb\x6d\xde\x67\xda\xfe\xad\x87\x2e\x24\xf9\x3f\x00\x00\xff\xff\x04\xee\x56\x42\xfd\x11\x00\x00")

func definitionsOperationsHclBytes() ([]byte, error) {
	return bindataRead(
		_definitionsOperationsHcl,
		"definitions/operations.hcl",
	)
}

func definitionsOperationsHcl() (*asset, error) {
	bytes, err := definitionsOperationsHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/operations.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa8, 0xf2, 0xba, 0x1d, 0xac, 0x71, 0x3b, 0xd9, 0xbf, 0x70, 0x8a, 0x5d, 0x3b, 0xaf, 0x19, 0x6f, 0x28, 0x36, 0x6, 0xdf, 0x45, 0x8f, 0x76, 0x79, 0x87, 0xa7, 0x14, 0xc0, 0x11, 0x3f, 0x2, 0xd0}}
	return a, nil
}

var _definitionsPairsHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x41\x6f\xdb\x3a\x0c\xc7\xef\xf9\x14\x44\x4e\xed\x43\xe1\x5e\x1e\xde\xad\x97\x37\x60\x40\x0f\xc3\x06\x6c\x3b\x6d\x43\xa0\x48\x54\x4d\x44\x11\x35\x8a\x8e\x9b\x0d\xfb\xee\x83\x64\x3b\x31\x3a\x37\x2b\x9a\x4b\x62\x90\xfc\xff\x7f\x21\x69\x69\x95\x0c\x09\xac\x2d\x47\xc5\x47\x5d\xc3\xcf\x15\x80\x1e\x13\xc2\xf8\xb9\x3b\xc5\x9a\x37\x63\xce\x0a\xc0\x61\xb6\x42\x49\x89\xe3\x2c\x03\x28\x82\x09\x01\x04\xbf\x77\x98\xc7\x44\x6f\xba\xa0\x4f\xa5\xfe\x37\x76\xf7\x20\xdc\x45\x77\x75\xbd\x5e\xfd\x9a\x41\x50\xec\x4c\x91\xdd\x28\xef\x30\x2e\xf2\x64\x15\x8a\x0f\x0b\x18\x39\xa1\x25\x7f\x04\x6d\x11\xe6\x5a\x50\xb5\xc0\xb3\x40\xa0\xac\x1b\x47\x02\xd3\xef\x24\xe8\xe9\xb1\x99\x41\x08\x3a\x8c\x4a\x26\x2c\x9a\xff\x73\x8e\x37\x1f\x84\x0f\xe4\x50\x2e\xa0\xb4\xdc\x83\x32\xa4\x21\x13\xce\xc5\x95\x26\xa3\x1c\xc8\x62\x81\xc9\xca\x62\x1e\xb0\x48\x25\x23\x19\xe5\xdc\xfe\x99\x61\x89\x9c\x51\x31\xba\xc4\x14\x97\xa7\x36\x05\x5f\x41\x39\x95\xbe\x9c\xf1\x6c\xf6\x84\xf0\x31\x91\xe0\x22\x5f\xe1\x7e\x9e\xa8\x6f\x31\xd6\x39\x76\x52\xf6\x49\x3b\x89\xe8\x60\x7b\x04\x41\x63\x5b\xe8\x29\x04\x18\xc5\xff\xa4\xa9\x8f\xf7\x45\x7f\xe2\x68\x55\xd3\xc6\x06\xc2\xa8\x1b\xae\x66\x79\x79\xba\x25\x71\xc8\x6b\xde\x8f\x79\x0b\x90\x98\x4e\x7b\x36\xaa\xd5\x56\x95\xe7\x22\x00\x83\xc2\xd9\x9e\xa2\xa2\x58\x4c\xca\x32\xb3\xbd\x83\x75\xf9\xce\xcd\xfd\x2c\x7c\xaa\x09\x6c\xeb\xf2\xbe\xfe\x15\x98\x14\x9e\x1d\xe3\x64\x15\xcd\x7e\x79\x46\x2f\xb2\x19\xe5\xa0\xaa\x9c\x34\xd9\xfb\x8c\xcb\x9b\x49\x51\xff\xfb\xf7\x82\xe8\x50\x3a\x76\x94\xf2\x74\x9c\xdc\x9c\x9c\xea\xf4\x33\xe2\xae\x2c\x6d\x4d\x19\x4b\xb6\xe8\x59\xb0\xec\x88\xbb\xb0\x16\xc5\x7c\xc2\x2c\xa9\x1b\x6b\x42\xd8\x1a\xbb\xdb\xf8\x2e\xda\x45\xe4\x12\xb8\xfa\xf2\x6d\x7b\x54\xbc\xbe\xb8\xb5\x46\x41\xd9\x31\xe0\x01\xe5\x08\x4a\x7b\x84\x7e\x00\x02\x67\xd4\x80\x17\xde\x43\xe6\x4e\xec\xac\x57\x99\x7e\x3c\xfb\x8e\x5c\xec\x54\x29\xfc\x5b\x9f\x38\x86\xe3\x00\x10\x68\x4f\x8a\xae\x1e\x8d\x18\xb5\x02\xbd\xb0\x4d\x3d\xcb\xae\x1c\x9b\xaf\xdf\x92\xa2\x00\xe5\xe0\x5d\x5e\xc6\x9b\xb1\x63\x9c\x50\x86\xa5\xad\xf0\xdb\xd2\xbb\x60\x94\x0e\x78\x9a\xb5\x23\x69\x60\x02\x82\x77\x9f\x3f\x7e\x82\xac\x46\x14\x7a\xd2\x16\x6e\xab\xc1\x20\x36\xf5\x61\xb4\xcb\xb3\xb2\x49\x7d\xba\x9f\x94\xe1\x16\xc8\x43\x64\x85\x8c\xda\xc0\xd7\x08\x6f\x59\xc0\xe7\xa7\x2a\x50\xd9\xa2\xe3\x3e\x43\x0a\x46\x3d\xcb\xfe\xa6\xfe\xc5\x2d\xb6\xe6\x40\x2c\x40\x19\xba\xe8\xd0\x53\x44\x57\x6f\x96\xdf\x01\x00\x00\xff\xff\x6d\x3b\xd7\x5b\x65\x07\x00\x00")

func definitionsPairsHclBytes() ([]byte, error) {
	return bindataRead(
		_definitionsPairsHcl,
		"definitions/pairs.hcl",
	)
}

func definitionsPairsHcl() (*asset, error) {
	bytes, err := definitionsPairsHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/pairs.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0xdf, 0x9e, 0x97, 0x5b, 0x28, 0xfa, 0x2f, 0x73, 0xf5, 0x37, 0x7f, 0x3a, 0x55, 0xb0, 0x25, 0x7d, 0x8f, 0xe9, 0x1b, 0x50, 0xdb, 0x12, 0x7a, 0xdf, 0x3b, 0x69, 0x5b, 0x3a, 0x10, 0xc6, 0xe9}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"cmd/definitions/tmpl/function.tmpl":  cmdDefinitionsTmplFunctionTmpl,
	"cmd/definitions/tmpl/info.tmpl":      cmdDefinitionsTmplInfoTmpl,
	"cmd/definitions/tmpl/object.tmpl":    cmdDefinitionsTmplObjectTmpl,
	"cmd/definitions/tmpl/operation.tmpl": cmdDefinitionsTmplOperationTmpl,
	"cmd/definitions/tmpl/pair.tmpl":      cmdDefinitionsTmplPairTmpl,
	"cmd/definitions/tmpl/service.tmpl":   cmdDefinitionsTmplServiceTmpl,
	"definitions/infos.hcl":               definitionsInfosHcl,
	"definitions/operations.hcl":          definitionsOperationsHcl,
	"definitions/pairs.hcl":               definitionsPairsHcl,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"cmd": &bintree{nil, map[string]*bintree{
		"definitions": &bintree{nil, map[string]*bintree{
			"tmpl": &bintree{nil, map[string]*bintree{
				"function.tmpl":  &bintree{cmdDefinitionsTmplFunctionTmpl, map[string]*bintree{}},
				"info.tmpl":      &bintree{cmdDefinitionsTmplInfoTmpl, map[string]*bintree{}},
				"object.tmpl":    &bintree{cmdDefinitionsTmplObjectTmpl, map[string]*bintree{}},
				"operation.tmpl": &bintree{cmdDefinitionsTmplOperationTmpl, map[string]*bintree{}},
				"pair.tmpl":      &bintree{cmdDefinitionsTmplPairTmpl, map[string]*bintree{}},
				"service.tmpl":   &bintree{cmdDefinitionsTmplServiceTmpl, map[string]*bintree{}},
			}},
		}},
	}},
	"definitions": &bintree{nil, map[string]*bintree{
		"infos.hcl":      &bintree{definitionsInfosHcl, map[string]*bintree{}},
		"operations.hcl": &bintree{definitionsOperationsHcl, map[string]*bintree{}},
		"pairs.hcl":      &bintree{definitionsPairsHcl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
