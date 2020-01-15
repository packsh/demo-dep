// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// kubepack.dev_applications.yaml
// kubepack.dev_bundles.yaml
// kubepack.dev_orders.yaml
// kubepack.dev_packages.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _kubepackDev_applicationsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3d\x4b\x73\xe3\x36\x93\x77\xff\x8a\x2e\xe7\xe0\x4c\x95\x24\x57\x26\xfb\xed\xc3\x37\xaf\x9d\xc9\x6a\x33\x0f\xaf\xed\xc9\x1e\xbe\xfa\x0e\x10\xd9\x92\xf0\x89\x04\x18\x00\x94\xac\x49\xe5\xbf\x6f\x75\x03\xa0\x28\xf1\x21\xd9\x72\x66\x93\x8a\x70\x99\x11\x08\x36\xfa\x8d\x46\x83\x68\x8b\x42\xfe\x8c\xc6\x4a\xad\xae\x40\x14\x12\x9f\x1c\x2a\xfa\x65\x47\x8b\x7f\xb7\x23\xa9\x2f\x97\xdf\x4d\xd0\x89\xef\xce\x16\x52\xa5\x57\x70\x53\x5a\xa7\xf3\x7b\xb4\xba\x34\x09\xde\xe2\x54\x2a\xe9\xa4\x56\x67\x39\x3a\x91\x0a\x27\xae\xce\x00\x12\x83\x82\x3a\x1f\x65\x8e\xd6\x89\xbc\xb8\x02\x55\x66\xd9\x19\x40\x26\x26\x98\x59\x1a\x03\x20\x8a\xe2\x0a\x2c\x0a\x93\xcc\x33\x39\x9b\xbb\x33\x00\x25\x72\x24\x34\x8a\x4c\x26\x0c\xc1\x8e\x16\xe5\x04\x0b\x91\x2c\x46\x29\x2e\xcf\x6c\x81\x09\xbd\x3c\x33\xba\x2c\xae\x60\xeb\x99\x7f\x3b\xc0\x4e\x84\xc3\x99\x36\x32\xfe\x1e\x56\x63\xc3\x4f\x51\x14\x36\xd1\x29\xf2\x4f\x4f\xda\xf5\x66\x5a\xee\xcd\xa4\x75\x3f\xed\x3e\x79\x2f\xad\xe3\xa7\x45\x56\x1a\x91\x6d\x23\xcb\x0f\xec\x5c\x1b\xf7\x71\x83\x0a\xcf\xe5\x9f\x48\x35\x2b\x33\x61\xb6\x5e\x3a\x03\x28\x0c\x5a\x34\x4b\xfc\xac\x16\x4a\xaf\xd4\x3b\x89\x59\x6a\xaf\x60\x2a\x32\x4b\xf8\xd9\x44\x17\x78\x05\x0c\xb2\x10\x09\xa6\xd4\x57\x4e\x4c\x90\x41\x98\xc6\x3a\xe1\x4a\x7b\x05\xbf\xfe\x76\x06\xb0\x14\x99\x4c\x19\xbc\x7f\xa8\x0b\x54\xd7\x77\xe3\x9f\xbf\x7f\x48\xe6\x98\x0b\xdf\x09\x90\xa2\x4d\x8c\x2c\x78\x5c\x9d\x48\x90\x16\xdc\x1c\xc1\x8f\x86\xa9\x36\xfc\xb3\x4e\x2a\x5c\xdf\x8d\x03\x94\xc2\xe8\x02\x8d\xab\x78\xed\x65\x5b\x69\x55\xd5\xb7\x33\xdf\x05\x21\xe4\xc7\x40\x4a\x7a\x84\x7e\xd2\xa5\xef\xc3\x14\xac\x9f\x5e\x4f\xc1\xcd\xa5\x05\x83\xcc\x28\xe5\x36\x22\x8a\x4d\x4f\x41\x28\xd0\x93\x7f\x62\xe2\x46\xf0\x40\xcc\x34\x96\x24\x51\x66\x29\x24\x5a\x2d\xd1\x38\x30\x98\xe8\x99\x92\x5f\x2a\xc8\x16\x9c\xe6\x29\x33\xe1\x30\x88\x35\x36\xa9\x1c\x1a\x25\x32\x62\x65\x89\x03\x10\x2a\x85\x5c\xac\xc1\x20\xcd\x01\xa5\xaa\x41\xe3\x21\x76\x04\x1f\xb4\x41\x90\x6a\xaa\xaf\x60\xee\x5c\x61\xaf\x2e\x2f\x67\xd2\x45\x3b\x4a\x74\x9e\x97\x4a\xba\xf5\x65\xa2\x95\x33\x72\x52\x3a\x6d\xec\x65\x8a\x4b\xcc\x2e\xad\x9c\x0d\xc9\x10\xa4\xc3\xc4\x95\x06\x2f\x45\x21\x87\x8c\xb8\xf2\x66\x90\xa7\xdf\x54\x02\xbf\xa8\x61\xea\xd6\xa4\x1b\xd6\x19\xa9\x66\x55\x37\xeb\x73\x27\xdf\x49\xa7\x49\xc2\x22\xbc\xe6\xf1\xdf\xb0\x97\xba\x88\x2b\xf7\x3f\x3c\x3c\x42\x9c\x94\x45\xb0\xcd\x73\xe6\xf6\xe6\x35\xbb\x61\x3c\x31\x4a\xaa\x29\x1a\x2f\xb8\xa9\xd1\x39\x43\x44\x95\x16\x5a\x2a\xc7\x3f\x92\x4c\xa2\xda\x66\xba\x2d\x27\xb9\x74\x24\xe9\x5f\x4a\xb4\x8e\xe4\x33\x82\x1b\xa1\x94\x76\x30\x41\x28\x8b\x54\x38\x4c\x47\x30\x56\x70\x23\x72\xcc\x6e\x84\xc5\xdf\x9d\xed\xc4\x61\x3b\x24\x96\xee\x67\x7c\xdd\x09\x6e\x0f\xf4\xdc\xaa\xba\xa3\x23\x6b\x95\x50\xcd\x12\x1f\x0a\x4c\xb6\xac\x83\x5e\x94\xd3\x68\xa7\x64\x99\x42\xd5\x5f\x18\xd5\x80\xb6\x19\x26\x1b\x67\x9a\x7e\x5a\x29\x34\xf7\x38\xdd\x7e\xb0\x8b\xc7\x66\x5c\xc0\xdf\xc2\x10\xa6\x99\x98\x91\xe1\x48\x95\xd2\x9c\x08\x72\x0a\x2b\x04\x85\x98\x52\xb7\x48\xd3\x1d\x98\x00\x11\x0a\x1b\x5c\x2e\x5c\x32\x27\x1d\x8b\x20\x3f\xc4\x0e\x69\x21\xd5\x0a\x61\xb2\x86\x92\x7c\x25\x3c\x60\x86\x89\x23\xef\xa3\x1b\x30\x7f\x29\xd1\xac\x41\x64\x19\xdc\xe8\xbc\xd0\x0a\x95\xfb\x91\x56\x05\xd2\x6e\xbb\x33\xda\x8b\x60\xa2\x75\x86\x62\xdb\x71\x08\x6b\x31\x9f\x64\xeb\xbb\xb9\xb0\xd8\xcf\x8c\xfa\xc8\x9a\xd6\x7b\x5d\x2e\x8d\x41\xe5\xa0\xe0\x87\xec\xb0\xb6\xdc\xe5\xc5\x2e\x46\x9b\x99\x47\x70\xad\x00\xf3\xc2\xad\x83\x1d\x4a\x0b\xf8\x4b\x29\x97\x22\x23\x88\x4e\xc3\xf9\x43\x99\x24\x88\x29\xa6\xe7\xa3\x56\xc2\x76\x94\x90\x5a\x12\x79\xc2\xec\xe8\x25\xac\x85\x7d\xde\x3b\xd0\xfa\x47\xb4\xf8\x2e\x52\xb5\xeb\x5e\x8a\xaa\x39\x2d\x7c\x8b\xa3\xd9\x08\x6e\xb1\xc8\xf4\x3a\xa7\x9e\x01\xdc\xe9\xd4\x0e\xd8\x43\xc8\x04\xed\x00\x6e\xee\x6f\xed\x9b\x11\x8c\x1d\x24\x42\xb1\x75\x5b\x6c\x6a\x8e\x54\xe4\xbe\xff\x59\xaa\x84\xf5\x7d\x25\xdd\x9c\x79\xbb\x85\x49\x5d\x4f\x02\xd6\x06\x56\xa4\x55\x34\xb6\x01\xf3\xba\xbe\x8c\x6d\x90\xde\xe5\xad\x74\x98\x37\x38\xb7\xc3\xbb\x8a\x65\xd1\x2e\x91\x18\xc7\xbd\xbc\x62\x08\x66\xde\x00\x26\xa5\x83\x54\xa3\x05\xa5\x5d\x03\x22\x10\x6f\x13\x04\x11\x97\xbe\x11\xc0\x23\xf9\x4d\x69\x89\x29\xd3\x32\x63\xe6\xcb\x94\x9c\xd2\x74\x4d\x86\x91\x68\x95\x60\xe1\x2c\xa4\x65\x43\xf4\xbe\x65\x5a\x2f\xca\x82\xa2\x82\x19\x5a\xe6\x9b\x2e\x1d\xcc\xc5\x92\x5e\x2f\x84\x71\x52\x64\xd9\xda\x07\x0a\xac\x46\x4d\x79\x76\xb9\x0f\xdf\x7c\x08\xd6\xf2\xa0\x47\x2b\x7d\xdb\x5d\x9d\x0e\x7c\x91\x16\x05\x69\xb0\xe5\xd5\xa1\x47\xa6\xa5\x9f\xa6\x6a\x74\xb7\x3a\xe4\xfa\x23\x61\x8c\x58\x6f\x3d\xa9\x0b\xbd\xcf\x96\x6e\x37\x3f\xc0\x20\x63\x65\x79\x69\x32\xb9\x77\xd8\x1c\x46\x84\x45\x02\xc4\x84\x64\xb2\xe3\x91\xc0\x47\xc6\x6d\xfe\x1c\xf6\x08\xa5\x07\xcb\x5e\x4c\xd9\xda\x27\x46\xe2\x34\x46\x04\xb5\xa1\xd1\x9b\x75\xac\x31\xdb\x9c\xeb\x10\x9d\x4c\xb4\x6a\x55\xa2\x2d\x94\xc6\x34\x8a\x91\x51\xa0\xb9\x4f\x64\x95\x13\x62\x10\x71\xbd\xeb\xe1\x8f\x6f\x04\x6a\x8b\xef\x52\x25\x59\x99\xc6\x35\x94\x23\x9a\x01\x58\xf9\x25\x46\x76\x32\x47\x26\xa1\x0d\x5c\x87\x23\x68\xe2\x9f\x8b\x19\xf2\x82\x4d\x91\x86\x90\x6a\x47\xf4\x41\xda\x20\x69\x5c\xbb\xbb\xf3\x4d\x30\x0b\x88\xe2\x76\xea\xfa\x0d\x13\x78\xa7\xf1\xa5\xb1\x9a\x75\x60\xfd\x6d\x64\xf5\x1b\x78\x24\xde\xc8\x2f\xd5\x02\xe6\x11\x95\x0a\x0a\xf9\x84\x59\xf0\xea\x83\x4e\xb0\x00\x6f\xff\xf6\xf4\xf6\x6f\x6f\xda\x91\x86\xfd\x8e\x81\x31\x37\xc9\x81\x88\x3f\x56\x92\xf4\xee\x91\x91\xad\xd6\x65\x4c\x89\x8d\x28\xdd\x1c\x4d\x9b\x89\x6d\x9a\x98\x58\x9d\x95\x0e\xe1\xf3\xfd\xfb\xb8\x21\xf0\xc0\x48\xd9\xe0\x96\x2c\x95\x1e\x05\xa1\xc6\xe8\x98\x87\x74\x93\x0a\xd5\x8b\x16\x84\xc1\x10\xc3\xa5\xc4\xcd\xfb\x77\x37\xf0\xf6\xfb\xff\xf8\xb7\xa3\xf8\xc4\x43\x5e\x24\xe1\x5c\x2a\xaf\xeb\xdb\x62\xde\x2b\xdc\x73\x1e\x77\x59\xa8\xd9\xf9\x31\x22\xee\x76\xe3\xc0\x2e\xdb\x9a\xa4\x67\x6d\x68\xf5\xda\xd0\xe3\xb9\xa9\x2d\x70\xbd\xd2\xa6\x19\x06\xc1\x2e\xa7\x7e\x0a\x03\xbb\x9c\xd0\x02\xd7\xe0\x07\x08\x6b\x75\x22\x69\x37\xd2\x8a\x6c\x15\xa7\xd4\x3c\x55\x08\x8a\x3e\xac\x1f\xfe\xe7\xfd\x00\xee\x6f\xff\xf3\xc3\xc3\x00\x68\x1d\x98\x08\x8b\xad\x1c\xed\x71\x3c\x7b\xd8\xdc\xc7\x8d\x4c\xaa\xc5\x7e\x56\xbc\xa7\x51\xac\xb9\x9b\x40\xb0\x1a\xb0\x44\xaf\xd8\xb4\x49\x56\xa9\x8f\xfb\x27\xcd\x40\x8b\x1a\xb9\x39\x7a\x6c\x4b\x33\x15\x14\xe4\xa4\xa9\x0c\x5c\x4d\x75\x52\xe6\x71\x3b\x4f\xac\xb0\xf3\x89\x16\x86\xc2\x44\x74\xc9\x31\x9e\x98\x90\xdf\xe3\x84\x83\xb1\x07\xb4\x3a\xd4\xf9\xf9\x18\x1e\xe2\x9c\xf7\x2c\xd3\x2d\xf4\xec\x2c\xd6\xf3\x32\x17\x14\x5f\x88\x54\x4c\x32\x64\x42\x69\xab\x80\x4f\x45\xe6\x1d\x54\x8f\x15\x93\x52\x16\xa5\x29\xf4\x66\x9f\x42\x0a\x71\x94\x2f\x2a\x4d\x76\x20\x21\x9f\x4d\x46\x00\x65\xc2\xc1\x27\xe7\x02\x2c\x08\x07\x02\x56\x38\xb1\xd2\xb1\x7e\x18\xb4\x8d\x60\xfc\x19\xf8\x1c\xe1\x28\x72\x21\x59\x67\xd0\xec\x37\x90\x0f\x9b\xb1\x5d\xee\xa2\x06\x2e\xf0\xba\x15\xe1\x7a\x34\xe3\x3d\x74\x1d\xb4\xf2\x39\x14\x86\x19\x1f\xb0\xd8\x36\x91\x4c\x2b\xd4\x44\xa7\x38\xf0\xbe\xdd\xfa\x20\xa7\x10\xc9\x82\x3c\x7d\x4b\x3a\xef\x18\x63\xbb\x21\x3b\x4b\x1c\x2f\x77\x7b\x02\x1f\x95\xca\xa5\x4c\x4b\x91\x75\x48\x57\x1b\xd0\x66\x26\x94\xfc\xd2\x13\xdb\xed\x37\x30\xcc\x85\x3c\x54\x23\x7f\xa0\xb1\x31\xdf\xc9\x2f\xbe\x82\x0a\x86\x44\xf6\x61\x18\x7c\x14\x39\x46\x04\xea\x1e\x96\x40\x7c\x45\xb3\x4c\x38\x57\xba\x31\xce\x09\xfe\x51\xac\x52\x69\xd7\x2e\xeb\x6d\x36\xd2\xa8\xa8\x80\x20\x76\x9d\xa4\x55\xb2\x28\xd0\xd5\xd6\x2c\xd1\xdc\xf1\x52\x13\xf0\x4b\x29\x93\x05\xed\x9c\x8d\xab\x6c\xa5\xb4\x1b\x23\xde\xda\x0a\xc1\x8d\xce\x73\xad\x3e\x08\xb3\x80\x5c\x98\x45\xaa\x57\xed\xc1\xa6\x5d\x2b\x27\x9e\x38\x35\x1a\x92\x1d\x0c\xdc\xc8\x64\x0e\x0e\x9f\xdc\x4e\x7a\xfb\xd9\x7b\x2c\xbd\x3a\xc8\x6d\x71\x2a\xae\xd3\x63\x11\x75\x1e\x50\x15\x1c\x2a\xeb\x44\x96\x75\xc4\x39\x0d\xd7\xb5\xfd\x72\x3d\xfa\x09\xc9\xf8\x89\x5f\xaf\x44\xd2\x15\x3a\x05\xef\x86\x4b\x5a\xd2\xf4\x14\x04\x68\x93\xa2\xa1\xd8\x59\x1b\x28\x55\xf5\x2b\x95\xd6\x94\x7e\x4d\x14\xd3\x29\x26\xae\x53\xf5\x4e\xbe\xee\xe4\xeb\xfe\x24\xbe\xae\x6b\x5b\xb7\xbd\xf3\xa5\xcd\x5b\xe0\x62\x7d\x23\xd7\xdc\x6c\xfc\xaf\x36\xe9\x1d\x51\xd3\xbe\xb1\x0b\x7b\x91\x1b\x61\xad\x50\xa9\x11\xad\xbb\x90\x5e\x72\x97\xcd\x43\xbe\x56\x94\xe3\x41\xdf\x8e\xdf\x59\xc6\x6e\x7f\x96\xa0\x4d\xe5\x70\x5f\x9e\x6e\xea\xe4\x3e\x1f\x11\xf5\x65\xef\xc6\x6a\xaa\x37\x66\xbd\xb3\x80\x2c\x70\x3d\xf0\x09\xfa\x42\x48\x63\x0f\x41\xf4\x90\x0c\x32\xcd\x39\x76\x98\xfb\x24\x5c\xef\x9c\xb5\x0c\x44\x9b\x0b\xcb\x0b\x6d\x9c\x50\xae\xc5\x1f\xcd\xf5\x8a\x8f\x66\x92\x04\xad\xdd\xcf\xde\x7e\xaf\xd3\x6d\xee\xad\xa6\xde\x20\xca\x49\x97\xc5\x20\x54\x5a\x28\x24\x26\xa4\xc1\x1d\xe6\x56\x23\xa6\xdd\x22\x0f\xd9\x09\xef\xc7\xf6\xb1\x66\x46\x9e\xe3\x15\x82\x51\x42\x2f\x9a\x9e\x61\x1d\x30\xff\xcf\xf1\xe8\xa7\x7d\x6f\xf7\xf2\xb9\xdf\x19\x9d\x1f\x3a\x3f\x8d\xad\x0e\x1b\x05\x18\x9c\xa2\x41\x95\x20\xe9\x4e\x8a\x86\x1c\x75\xc5\x9f\x0e\x71\xf1\x51\xaf\x50\x9a\xd3\x6e\x7e\x77\xf2\xd2\x95\x2d\xd1\x6a\x2a\x67\x1f\x44\xf1\x13\xae\x5b\x4e\x2a\x3b\x08\xf1\x47\x41\x84\xfe\x02\xd7\x3e\x7e\xb8\x89\x80\xba\x5d\xfa\x7e\x6c\xa0\xf3\xa3\x86\x3d\x08\x5d\xdf\x8d\x2b\x37\x17\x14\x2c\x30\xb6\x43\xaa\xb1\x1d\xb0\xea\x51\x9b\x4a\xcc\xd2\x3b\xe1\xe6\xcf\xc0\xe9\x62\x3c\xf5\x48\x70\xb2\x9f\x3c\x43\x65\x86\x9b\xef\x28\x38\xf4\x43\xd1\x95\xa4\x8e\xcd\xbf\x82\xca\x49\x83\xe1\xcd\x81\x37\x9c\x70\x96\xb0\xf9\x0a\x23\x44\xe6\x7c\xe6\xb4\x07\xea\x7f\x3f\x7c\xfa\x78\xf9\xa3\xf6\xd4\x45\xc7\x65\x9d\x70\x98\xa3\x72\x03\xb0\x65\x32\x07\x61\x89\x2c\x69\x30\x7d\xa0\x27\xa3\x5c\x28\x39\x45\xeb\x46\x61\x2e\x34\xf6\xef\x6f\xff\xd1\xcf\x65\x80\x77\xda\x00\x3e\x89\xbc\xc8\x68\xbf\x1c\x62\xe0\xf8\x6d\x43\xb4\x00\x69\x3d\x9b\x2a\xc8\x7b\x80\xae\xa4\x9b\x33\xb1\x85\x4e\x03\x3b\x56\x3e\xe8\x10\x0b\x04\x1d\xd8\x50\x22\x64\x72\x81\x57\x70\x6e\x0b\x4c\x6a\x68\xff\x4a\x4e\xf6\xb7\xf3\x3d\x93\x7c\xbb\x9a\xa3\x41\x38\xa7\xc1\xe7\x1e\xd9\xea\x93\x16\xea\x8b\xfa\x56\xc1\x05\x37\x17\x6d\x11\x49\xbd\x39\x23\x67\x33\x34\x98\x6e\xc2\xf0\x37\x14\xac\xca\x29\x28\x5d\x03\xa5\x82\x7f\x8f\x07\x9f\xfb\x24\xba\x4b\xe2\xdf\xdf\xfe\xe3\x1c\xbe\xdd\xc0\xe3\x94\xa9\x54\x29\x3e\xc1\xdb\x98\xf5\xd8\x03\xb2\xd0\xe9\x9b\x91\x3f\x2a\x0d\xbb\x2b\x69\x21\x99\x6b\x8b\x0a\xb4\xca\xd6\xc4\x8b\xb9\x58\x22\x58\x9d\x23\xac\x30\xcb\x86\x21\x0d\xbf\x4f\x7a\x82\x5d\x47\x14\x3f\x69\xb1\xe0\x43\xd3\x9d\x0f\x8d\x1e\x3f\xdd\x7e\xba\xf2\xd2\x25\x45\x9c\xf5\x1d\x33\x50\x93\x7c\x00\x0c\x53\x49\xb1\x8f\x50\x69\xf8\x3c\x86\x75\x9c\x88\x28\xbd\xda\x39\x0d\xc9\x5c\x28\x3e\x7b\xd9\x27\xad\x39\xc2\xb4\x74\xa5\xc1\xd1\xc5\x6b\x78\x93\x05\xae\x9f\xe1\x47\x68\xc7\x47\x6e\xd6\x69\xb0\xec\x77\x5f\xc5\xa3\x75\x9f\x0f\xb7\x22\xe1\xbf\x67\xda\x75\xad\xff\x4f\x5f\x05\xbd\x98\xe8\xfe\x4d\x54\x93\xe8\x8f\x35\xfb\xee\x25\x7a\x51\x4e\x68\xbf\xec\x90\xe9\x4e\x75\x62\x2f\xe3\xd7\x03\x97\x7a\x89\x66\x29\x71\x75\xb9\xd2\x66\x21\xd5\x6c\x48\x46\x38\x0c\xdf\xe5\x5c\xf2\xe7\x94\x97\xdf\xf0\x3f\xaf\x46\x23\x7f\xc1\xf8\x5c\x42\xf9\xa5\xaf\x41\x2d\xcd\x63\x2f\x5f\x85\xd8\xf8\xe1\xdc\xf3\x23\x86\x8b\x87\xf0\x8d\xd7\x2e\x0c\xb2\xb3\xd5\x9c\x13\x46\xfe\x7b\xc8\xb0\x3a\xed\xf7\x3a\xb9\x48\xfd\xe2\x26\xd4\xfa\x77\x37\x0d\x62\x38\x7f\x0f\x95\xac\x87\x0c\x42\x67\x43\xa1\x52\xfa\xbf\x95\xd6\x51\xff\xab\x70\xb8\x94\xcf\x72\x13\x9f\xc7\xb7\x5f\xc7\x60\x4a\xf9\x2a\x3e\x61\x4f\x16\x01\x78\x73\x34\xa3\xcd\xfd\x73\xa3\x63\x5a\xc2\xc6\xfe\xd5\x53\x48\x7c\x0a\x89\x9b\xed\x14\x12\x6f\xb5\x53\x48\xfc\x57\x0e\x89\xe7\xda\xba\x67\xc6\xc4\x55\x7e\x93\xde\x3d\x45\xc7\xc7\x13\x7d\x8a\x8e\x7b\x08\xfd\x33\x46\xc7\xc5\xf3\x56\xe7\x2d\xa3\xfa\xaf\xc7\xc7\x3b\x06\xf0\x2a\xf6\x74\x8a\xd3\x4f\x71\xfa\x51\x04\x1e\x10\xa7\x5b\x4c\x0c\xba\x63\xf3\xd8\x0f\x0c\xe5\x14\xb1\x9f\x22\xf6\x66\x3b\x45\xec\x5b\xed\x14\xb1\xff\x95\x23\xf6\x53\x12\xfb\x14\xa6\x9f\xc2\xf4\x53\x70\x7c\x0a\x8e\xff\x1c\xc1\x31\x5f\x09\x7e\x7e\x12\x3b\x5e\x26\x3e\x45\xc4\xa7\x88\xb8\xd9\x4e\x11\xf1\x56\x3b\x45\xc4\x7f\xe9\x88\xf8\x14\x8d\xee\x27\xfa\x14\x8d\xfe\xc1\xa3\xd1\x3f\x4e\xd2\xb8\xd0\xe6\xc5\x67\x42\xf4\xee\xa1\x9b\x4d\xff\xed\xf3\x15\x48\xe5\xbe\x7f\x7b\x00\xe2\x52\x39\x9c\xf5\xae\x62\xa7\x88\xfe\x14\xd1\x1f\x45\xe0\x01\x11\xfd\x33\x6a\x13\xc4\xef\xee\xfb\x3e\x15\x3f\x08\xb5\x5e\xb4\x5e\x50\xf1\xc5\x86\x92\x42\xbd\x17\x46\x2e\xaa\xc2\x43\xbe\x60\x92\x98\x60\x16\x0a\x53\x91\x48\x78\xd9\xb5\x1c\x9d\xf9\x92\x88\x98\xc2\x64\xdd\x7a\x0f\x78\xeb\x22\xdd\xd8\x41\x5e\xf2\x75\xdf\x50\xc2\x68\x53\xa5\x28\x16\xcd\xba\xb8\x08\xa5\x14\xeb\xba\xd2\x80\x7a\xbc\xee\xf8\x49\x2e\xbf\xe1\x7f\x87\x91\x29\x0d\x3d\xea\xdb\x68\x31\x15\x3f\x3c\x15\x06\x2d\xd7\x95\xdc\x7b\x57\x68\xf7\x85\xed\x6a\x54\x9e\xc9\x11\x93\x58\x4e\x22\x6f\xab\xe0\xe4\xdb\x23\x5b\xd6\x66\x14\x17\x35\xb8\xfe\x78\x8b\xe9\x31\x37\x01\xaf\x7b\x10\x09\xc5\xf5\xaa\xa2\x54\x2c\xff\x70\xad\xa8\x43\xc3\x7d\x09\xc1\x81\x3f\xee\xf1\xd7\xb5\xf9\x9e\x14\x1a\x51\x81\x30\xc8\x75\x0a\x59\x21\x16\xb8\xe6\x41\xa1\xf4\x60\x2b\xd4\xfd\xbb\xdf\xde\x44\xe9\x16\xb9\x34\x5f\xb8\x71\xe6\xe9\xe6\x8c\x29\x61\xc5\xd7\xd0\x23\xa9\xac\xc6\x2d\xa5\xa4\x36\xcd\xe9\xa3\xee\xf9\x45\x8e\x1c\x88\x76\xc5\xc0\x5a\xbd\x36\x66\xf1\x85\xf5\xec\x24\xfd\x9a\xcb\x82\x77\x81\x3d\x58\x5b\x64\xdd\x8b\x85\x1e\x7f\xe6\xaa\x59\x11\xb8\xd7\xa8\xb1\x1a\xc0\x47\xed\xe8\x9f\x1f\x9e\xa4\xa5\x89\x5a\x6a\x4f\x6d\xda\xad\x46\xfb\x51\x3b\x1e\x7b\x14\x4b\x3c\x52\x07\x32\xc4\x0f\x0e\xd7\xf0\xd8\xeb\xb1\xf7\xad\xd5\x81\xb4\x23\x18\x87\xdd\x72\xa0\xaf\x87\x08\xbe\x1e\x45\xdb\xc8\x40\x79\x75\x3d\xc8\x06\xe0\xec\xc7\x26\x08\x4a\xab\x21\x17\xb9\x8b\xd0\x7b\x80\x56\x42\x93\x36\xb2\x52\x9b\x2d\x7e\x75\x4c\xd4\x03\x73\x82\x10\xa6\xe7\x7d\xa5\x7f\xc7\xd7\x14\xcd\x44\x82\x69\x28\xa4\xe6\x6b\x62\x0a\x87\x33\x99\x40\x8e\x66\xd6\x87\x67\x41\x7e\xaa\x5b\x74\x3d\x9e\xc4\xb7\x67\x2c\xb7\xed\x97\x45\xe1\x80\x72\x3a\x0b\x6c\x7f\x6f\xd8\x2f\xde\xa3\x8a\x68\xb8\x64\xfe\xbe\x56\xe7\x77\xbb\x6d\x0a\xc0\xdc\xed\xf1\x4f\x7b\xf8\xd3\x5c\x33\xfc\xa4\xde\xfb\xe6\xa2\x20\xcd\xfe\xb5\xba\x42\xf9\x9b\xbf\xb7\x39\x82\x6b\xae\xc3\x9b\xb5\x4b\xb6\x3e\x3e\xdc\x44\xaf\x83\x26\xa8\x8d\xfa\x8c\x42\x01\x66\xec\xf8\x5b\x41\x72\x21\x90\xed\x15\x6d\x00\xab\xb9\xb6\xde\x8b\x57\x69\x82\xf3\x05\xae\xcf\x07\x5b\x96\x07\x1d\xb9\x92\xf3\xb1\x3a\x1f\x84\x4c\xc3\x8e\x1d\x54\xd7\x57\x39\x53\x72\xce\xcf\xce\x47\x8d\x45\xb0\x15\x6c\xef\xc2\xb8\x37\x86\x6a\x3c\x6a\x2f\x7c\xea\xcb\x14\x9f\x75\xc8\xb1\x5e\xfa\x94\x47\x56\xf7\x11\x43\xc8\x9e\xa1\xb9\xb0\x21\x99\xc7\x45\x93\x53\x9f\x3f\xdc\xbd\x41\xba\x5b\xcc\xd9\xb7\xae\x25\x71\x53\x05\xb2\x37\xde\xfb\xe4\x13\x39\x9e\x86\xc0\x70\x2e\x49\x97\x65\x8d\xb2\xa6\x3b\x70\x0e\xb9\x0b\xec\xc1\x07\xba\x59\x8b\x67\xa8\xd0\xc8\x24\xce\x38\xd7\x59\x8a\xfe\x8e\x74\xfb\x2c\x47\x14\x6d\x6c\x23\xb4\xbd\xac\x22\x1c\x53\xe2\x71\xbb\xde\x56\x48\x33\xf5\xec\x26\xf6\x4c\x94\x49\xb5\x38\xb4\x20\x94\xd3\x47\x4c\xf4\x9c\x5b\xcf\x47\x51\xd4\x34\x90\x8e\xa9\x2e\xbc\xa2\x8c\xfc\xdd\x5d\x7b\x05\x63\x75\x67\x34\x7f\xd5\x3f\x80\x7b\x14\xe9\x7a\x00\xa1\xa2\x78\xfb\xb6\xaf\x17\x91\x17\x6c\x99\x12\xad\xbc\x6b\xdf\x57\x6f\x36\x0e\xdb\xad\xa2\xeb\xcb\x70\x57\xf6\x5c\x4b\xd9\xbf\xc0\x98\xaa\x69\x42\xf7\x24\x16\x7f\x8c\xd0\x37\x27\x22\x5c\x80\x2a\x41\x43\x8e\xb3\x85\x53\x5c\xa6\xea\xb9\xb7\xe4\x33\x61\xdd\xa3\x11\xca\x32\x0e\x8f\xf2\x20\xed\x79\x2f\xac\x03\xc7\xb5\x28\x7d\x7e\x3f\x50\xe0\x2a\x40\x98\xfa\x2b\xd6\x5a\x61\xd0\x94\xae\x3d\xb8\x8e\xd7\xb0\xdb\x23\x94\x98\x59\x4a\x85\xc3\x21\x4d\xf9\x22\xeb\x13\xd6\x7d\xe6\xa2\xdc\x47\x12\xb8\x12\x96\xf8\x39\xe9\x48\xe1\xbf\x0a\xb2\x39\x5a\x2b\x66\x87\x60\x79\xbd\x7b\x0b\x3f\xbc\x1a\xcb\x54\xf8\x0a\xa9\x4e\xc8\xcc\xfa\x02\x0b\x5d\x32\x98\x63\x4d\x74\x2f\xba\xc3\x6f\x50\xd8\xae\x2c\x5d\x23\xc5\xe8\x07\x57\x95\x29\x2a\xf6\x5e\x58\x96\xd4\xb1\xb8\x1c\xec\x99\xc2\x0a\xb6\x39\xa6\xf2\x68\x0c\x58\x6b\xf5\x14\x1e\x4d\x89\x03\x78\x27\x32\x8b\x95\x8b\xfa\x1a\xc5\x1d\x2a\x4c\x5e\x30\x59\x5f\xf5\xe1\x0e\x43\x1c\x32\xc4\xd7\xf0\xac\x31\xd4\xf9\x91\xc2\x81\xda\x9f\x75\xe8\xa0\xf7\x53\x63\x78\xdc\xbb\xe7\xda\xf2\xdf\x41\xa0\xc8\x75\xb6\x79\x1a\xe1\x37\xf9\x32\x76\x90\x68\x63\xd0\x16\x9a\x53\x59\xfe\x14\xd2\x87\x07\x17\xb6\x06\x63\x10\x52\xc0\xd2\xc6\x32\xfd\xa0\x15\xe4\x65\xcb\x1f\x6b\xa0\xe6\xd3\x60\x7c\x8a\xee\xff\x68\xc0\xee\xd4\xb5\xcc\xf7\xbf\xfe\x4b\x2b\x9b\x9a\xb9\xee\x16\xce\xee\x74\xc5\x72\x35\xb0\xfc\x4e\x64\xc5\x5c\x7c\xb7\xe9\x63\xc5\x1e\x86\xbf\x40\x52\x7b\xec\x3f\x50\xc0\xf4\x0a\x9c\x09\x65\x2f\xac\xd3\x86\x3c\x89\xef\xf9\xbf\x00\x00\x00\xff\xff\x65\x88\x62\x93\x41\x65\x00\x00")

func kubepackDev_applicationsYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubepackDev_applicationsYaml,
		"kubepack.dev_applications.yaml",
	)
}

func kubepackDev_applicationsYaml() (*asset, error) {
	bytes, err := kubepackDev_applicationsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubepack.dev_applications.yaml", size: 25921, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubepackDev_bundlesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x24\x35\x0c\xbd\xf7\xaf\xb0\xc4\x61\x2e\xdb\xd5\x1a\x81\x10\xaa\xdb\x32\x7c\x68\x04\x88\xd5\xce\xb2\x77\x57\xe2\xae\x0a\x93\x4a\x42\xec\xf4\x32\x20\xfe\x3b\x4a\xea\xa3\xbb\xb7\xab\x6b\x86\x45\x9b\x5b\x9e\x63\xc7\x79\x7e\x49\x8c\xc1\xbc\xa7\xc8\xc6\xbb\x1a\x30\x18\xfa\x53\xc8\xe5\x19\x57\x8f\xdf\x70\x65\xfc\xee\x70\xdb\x90\xe0\xed\xe6\xd1\x38\x5d\xc3\x5d\x62\xf1\xfd\x5b\x62\x9f\xa2\xa2\xef\x68\x6f\x9c\x11\xe3\xdd\xa6\x27\x41\x8d\x82\xf5\x06\x40\x45\xc2\x0c\xbe\x33\x3d\xb1\x60\x1f\x6a\x70\xc9\xda\x0d\x80\xc5\x86\x2c\xe7\x35\x00\x18\x42\x0d\x4c\x18\x55\x67\x4d\xdb\xc9\x06\xc0\x61\x4f\x35\x34\xc9\x69\x4b\x5c\x3d\xa6\x86\x02\xaa\xc7\x4a\xd3\x61\xc3\x81\x54\xf6\x6b\xa3\x4f\xa1\x86\x33\xdb\xe0\x38\x86\x55\x28\xd4\xfa\x68\xa6\xf9\x76\x5e\x3b\x4e\x31\x04\x56\x5e\x53\x99\x0e\xa7\xfa\xb6\xec\x58\x00\x6b\x58\x7e\x3a\x01\x7f\x36\x2c\xc5\x10\x6c\x8a\x68\xe7\xec\x0a\xc6\xc6\xb5\xc9\x62\x9c\xd0\x0d\x40\x88\xc4\x14\x0f\xf4\x9b\x7b\x74\xfe\x83\xfb\xc1\x90\xd5\x5c\xc3\x1e\x2d\x67\x33\x2b\x1f\xa8\x86\x3b\x9b\x58\x28\x66\x20\x35\x71\x64\x73\x4c\x98\x05\x25\x71\x0d\x7f\xff\xb3\x01\x38\xa0\x35\xba\x90\x39\x18\x7d\x20\xf7\xfa\xcd\xfd\xfb\x2f\x1f\x54\x47\x3d\x0e\x60\xde\xd5\x07\x8a\x32\x1f\x7a\xe0\x77\xae\xec\x8c\x01\x68\x62\x15\x4d\x28\x11\xe1\x26\x87\x1a\xd6\x80\xce\xb5\x24\x06\xe9\x08\x0e\x03\x46\x1a\xb8\x6c\x03\x7e\x0f\xd2\x19\x86\x48\xe5\x7c\x4e\x4a\x4a\x27\x61\x21\x2f\x41\x07\xbe\xf9\x9d\x94\x54\xf0\x90\x39\x88\x0c\xdc\xf9\x64\x35\x28\xef\x0e\x14\x05\x22\x29\xdf\x3a\xf3\xd7\x1c\x99\x41\x7c\xd9\xd2\xa2\xd0\xc8\xf4\x34\x8c\x13\x8a\x0e\x6d\x26\x21\xd1\x2b\x40\xa7\xa1\xc7\x27\x88\x94\xf7\x80\xe4\x4e\xa2\x95\x25\x5c\xc1\x2f\x3e\x12\x18\xb7\xf7\x35\x74\x22\x81\xeb\xdd\xae\x35\x32\x69\x59\xf9\xbe\x4f\xce\xc8\xd3\x4e\x79\x27\xd1\x34\x49\x7c\xe4\x9d\xa6\x03\xd9\x1d\x9b\x76\x9b\xc5\x68\x84\x94\xa4\x48\x3b\x0c\x66\x5b\x12\x77\x52\x2e\x44\xaf\xbf\x98\x4b\x75\x73\x92\xa9\x3c\xe5\x92\xb2\x44\xe3\xda\x19\x2e\xc2\xba\xca\x7b\x56\x18\x18\x06\x1c\xdd\x86\xfc\x8f\xf4\x66\x28\xb3\xf2\xf6\xfb\x87\x77\x30\x6d\x5a\x4a\x70\xce\x79\x61\xfb\xe8\xc6\x47\xe2\x33\x51\xc6\xed\x29\x0e\x85\xdb\x47\xdf\x97\x88\xe4\x74\xf0\xc6\x49\x99\x28\x6b\xc8\x9d\x93\xce\xa9\xe9\x8d\xe4\x4a\xff\x91\x88\x25\xd7\xa7\x82\x3b\x74\xce\x0b\x34\x04\x29\x68\x14\xd2\x15\xdc\x3b\xb8\xc3\x9e\xec\x1d\x32\x7d\x76\xda\x33\xc3\xbc\xcd\x94\x3e\x4f\xfc\xe9\x43\x74\xbe\x70\x60\x6b\x86\xa7\x17\x65\x1a\x4b\x77\xa8\xdc\x23\xad\xbd\xfb\x08\x03\x30\x42\xfd\x05\x78\x3d\xc8\x30\xca\xfb\xb6\x80\x5f\x39\xca\x71\xa4\x68\x3f\xc9\x6f\xbc\xc8\x8b\xc9\x5c\x3d\xc3\x8b\x22\x4f\x0b\x30\x46\x7c\xba\xb0\x67\xf5\x98\x48\xfa\x32\xf8\xb6\x90\xb0\x00\xa7\x68\x17\xd0\xe9\x00\x17\xa6\xc5\x92\xae\x27\x96\xff\x00\x6c\x2f\x2b\xf3\x49\xa5\x54\x1d\x46\x59\xe6\x6e\xdd\x11\x56\x75\x00\x2f\x61\x7e\x45\x0f\x2f\xf4\x5f\xd7\x05\x3c\xa7\x8d\x17\xee\xb2\xae\x11\x58\xd5\x09\x5c\xd7\x0a\x5c\xd5\x0b\xac\x6b\x06\xd6\x75\xf3\x7c\x42\x83\x73\xe3\xbd\x25\x74\xff\x49\xf4\x45\x2e\x0b\xf8\xe4\xf3\xff\xf5\xbd\xb4\xfb\x76\x7c\xbc\xce\xa0\xe9\x1a\x3c\xfb\x3e\x0e\x7d\xc8\x0b\x5e\x48\xdf\x94\x96\x47\xff\x48\x8e\xe2\x49\xb3\x72\x1c\x67\xbf\xdf\xaf\x17\xcb\xf3\x57\x98\xff\xa3\xde\x73\xe9\x11\xc8\x09\xb4\x47\xeb\x14\xff\x82\xa4\xbd\x8f\x53\x5f\x32\x7c\x91\x15\xdc\x0b\x28\x1f\x23\x71\xf0\x4e\xcf\xdd\xc5\x64\xbf\xe1\x93\xb8\xaf\xe0\x43\x67\x54\x77\x11\xd5\xf0\xf4\xcd\x81\x77\xd0\xa7\xa1\xd9\x81\xe6\xa9\x84\x7a\xfd\xe6\x7e\xfc\x64\xab\xcd\x45\x36\x3d\x4a\x9d\xfb\x96\xaf\xbf\x5a\xac\x59\xee\x68\xda\xd2\xf6\xad\x30\xff\x11\x74\x98\xba\xf3\xc3\x2d\xda\xd0\xe1\xed\xe6\xfc\xfa\x6e\xc7\xae\xf9\xc4\x0c\x30\xf0\x55\x83\xc4\x44\x63\x53\xe9\x23\xb6\x34\x22\xff\x06\x00\x00\xff\xff\x1d\x25\xfd\x6b\xf5\x0b\x00\x00")

func kubepackDev_bundlesYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubepackDev_bundlesYaml,
		"kubepack.dev_bundles.yaml",
	)
}

func kubepackDev_bundlesYaml() (*asset, error) {
	bytes, err := kubepackDev_bundlesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubepack.dev_bundles.yaml", size: 3061, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubepackDev_ordersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x56\x4b\x8f\xe4\x34\x10\xbe\xf7\xaf\x28\x89\xc3\x5c\x36\xdd\x1a\x81\x10\xca\x6d\x35\x3c\x34\xe2\xb1\xab\x9d\x65\xef\x95\xb8\x3a\x31\xe3\xd8\xa6\xaa\xdc\xbb\x0d\xe2\xbf\x23\x3b\x49\x3f\x76\x12\x06\x84\xb8\xa0\xf5\xcd\xf5\xf6\x57\x65\xfb\xc3\x68\xdf\x11\x8b\x0d\xbe\x06\x8c\x96\x3e\x28\xf9\xbc\x93\xed\xe3\x57\xb2\xb5\x61\x77\xb8\x6d\x48\xf1\x76\xf3\x68\xbd\xa9\xe1\x2e\x89\x86\xe1\x0d\x49\x48\xdc\xd2\xd7\xb4\xb7\xde\xaa\x0d\x7e\x33\x90\xa2\x41\xc5\x7a\x03\xd0\x32\x61\x16\xbe\xb5\x03\x89\xe2\x10\x6b\xf0\xc9\xb9\x0d\x80\xc3\x86\x9c\x64\x1b\x00\x8c\xb1\x06\x21\xe4\xb6\x77\xb6\xeb\x75\x03\xe0\x71\xa0\x1a\x02\x1b\x62\xd9\x3e\xa6\x86\x22\xb6\x8f\x5b\x43\x87\x8d\x44\x6a\xb3\x5b\xc7\x21\xc5\x1a\xae\x74\xa3\xdf\x14\xb5\x45\xa5\x2e\xb0\x9d\xf7\xd5\xc9\x76\xda\x62\x8c\xd2\x06\x43\x65\x3b\x1e\xea\x55\x4e\x58\xf6\xce\x8a\x7e\x7f\x96\xfd\x60\x45\x8b\x3c\xba\xc4\xe8\xe6\xd2\x8a\x48\xac\xef\x92\x43\x9e\x84\x1b\x80\xc8\x24\xc4\x07\xfa\xd9\x3f\xfa\xf0\xde\x7f\x6b\xc9\x19\xa9\x61\x8f\x4e\x72\x36\x69\x43\xa4\x1a\x7e\xca\xb5\x46\x6c\xc9\x64\x59\x6a\x78\xc2\x72\xaa\x57\x14\x35\x49\x0d\xbf\xff\xb1\x01\x38\xa0\xb3\xa6\x40\x39\x2a\x43\x24\xff\xf2\xf5\xfd\xbb\xcf\x1f\xda\x9e\x06\x1c\x85\x39\x71\x88\xc4\x7a\x3a\xf3\x88\xee\xa9\xaf\x27\x19\x80\x21\x69\xd9\xc6\x12\x11\x6e\x72\xa8\xd1\x06\x4c\xee\x24\x09\x68\x4f\x70\x18\x65\x64\x40\x4a\x1a\x08\x7b\xd0\xde\x0a\x30\x95\x23\x7a\x2d\x25\x5d\x84\x85\x6c\x82\x1e\x42\xf3\x0b\xb5\xba\x85\x87\x0c\x03\x0b\x48\x1f\x92\x33\xd0\x06\x7f\x20\x56\x60\x6a\x43\xe7\xed\x6f\xa7\xc8\x02\x1a\x4a\x4a\x87\x4a\x13\xd4\xf3\xb2\x5e\x89\x3d\xba\x0c\x42\xa2\x17\x80\xde\xc0\x80\x47\x60\xca\x39\x20\xf9\x8b\x68\xc5\x44\xb6\xf0\x63\x60\x02\xeb\xf7\xa1\x86\x5e\x35\x4a\xbd\xdb\x75\x56\xe7\x49\x6e\xc3\x30\x24\x6f\xf5\xb8\x6b\x83\x57\xb6\x4d\xd2\xc0\xb2\x33\x74\x20\xb7\x13\xdb\x55\x79\x14\xad\x52\xab\x89\x69\x87\xd1\x56\xa5\x70\xaf\xe5\x3a\x0c\xe6\xb3\x53\xab\x6e\x2e\x2a\xd5\x63\xee\xaa\x28\x5b\xdf\x9d\xc4\x65\xae\x56\x71\xcf\x13\x06\x56\x00\x27\xb7\xb1\xfe\x33\xbc\x59\x94\x51\x79\xf3\xcd\xc3\x5b\x98\x93\x96\x16\x5c\x63\x5e\xd0\x3e\xbb\xc9\x19\xf8\x0c\x94\xf5\x7b\xe2\xb1\x71\x7b\x0e\x43\x89\x48\xde\xc4\x60\xbd\x96\x4d\xeb\x2c\xf9\x6b\xd0\x25\x35\x83\xd5\xdc\xe9\x5f\x13\x89\xe6\xfe\x6c\xe1\x0e\xbd\x0f\x0a\x0d\x41\x8a\x06\x95\xcc\x16\xee\x3d\xdc\xe1\x40\xee\x0e\x85\xfe\x73\xd8\x33\xc2\x52\x65\x48\x9f\x07\xfe\xf2\x19\xba\x36\x1c\xd1\x3a\x89\xe7\x07\x65\x5e\x4b\x77\xa8\x8c\xa1\xd2\xf0\x91\x68\x45\xb8\x1e\x63\x5c\x4d\xf2\xc6\xd1\x92\xe6\x39\xcf\xbc\xca\xd3\xb8\xa2\x5b\xc1\xe2\x7a\x25\x76\xff\xca\x7f\x7a\x15\x56\x0b\x5c\x45\xe5\x1f\x65\x99\x8d\x90\x19\x8f\x8b\x36\x79\x34\x2d\x93\x59\x4e\x54\x15\xa0\x56\x54\x89\xdd\x8a\x66\x3e\xdc\xa2\x7a\x71\x7e\xce\xab\xed\x91\xf5\x53\x5b\xff\x6f\x6d\x8d\xc8\x38\x90\x12\xaf\x9c\xfd\x19\x77\x80\x0f\x55\x66\x1e\xec\x49\x49\x2a\x1a\x1a\x32\x86\x4c\x35\x3f\xe7\x35\x28\xa7\xe5\x13\x5d\x39\xce\xa4\xa2\x4a\x23\xab\xa8\xf6\x13\xad\x58\x74\xff\x8b\xa2\x96\x1b\xb0\x04\x7b\x35\xb6\xfc\xd9\x17\x74\x64\x2a\x7f\xe3\x0d\x0d\x4d\x39\x82\xf9\x8e\x3c\xf1\x05\x9d\x39\xaf\xab\xff\xf1\xd5\x13\xf3\xfc\x59\xe6\x1f\x6b\x08\x52\x58\x04\x79\x85\xee\xac\x9d\xe3\x3f\x81\x63\x1f\x78\x66\x2e\x23\xea\x5b\xb8\x57\x68\x03\x33\x49\x0c\xde\x9c\xf8\xc7\xac\xbf\x91\x8b\xb8\x2f\xe0\x7d\x6f\xdb\xfe\x49\x54\x2b\xf3\x47\x08\xc1\xc3\x90\x46\x3a\x04\xcd\xb1\x84\x7a\xf9\xfa\x7e\xfa\x86\xb7\x9b\x27\xd5\x0c\xa8\x75\x66\x36\x5f\x7e\xb1\xd8\x9d\xcc\x79\xba\x89\x8a\xae\x22\xff\x91\xe8\x30\xb3\xf7\xc3\x2d\xba\xd8\xe3\xed\xe6\xfa\x52\x57\x13\xab\xbe\x50\x03\x8c\x78\x5d\x0c\x91\x68\x60\xec\xe6\xa9\xfc\x33\x00\x00\xff\xff\xc3\x36\x1d\x71\x15\x0c\x00\x00")

func kubepackDev_ordersYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubepackDev_ordersYaml,
		"kubepack.dev_orders.yaml",
	)
}

func kubepackDev_ordersYaml() (*asset, error) {
	bytes, err := kubepackDev_ordersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubepack.dev_orders.yaml", size: 3093, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubepackDev_packagesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x8f\xe3\xb8\x0d\x7f\xcf\xa7\x20\xa6\x0f\xf3\x32\xc9\x60\xd1\xa2\x28\xf2\x36\x98\xdd\x2b\x82\x5e\xb7\x8b\xdb\xed\xbe\x33\x36\x63\xab\x91\x25\x9f\x44\x67\x36\x3d\xdc\x77\x2f\x28\x59\x8e\x13\xff\x49\xe6\xb6\x7f\x50\x60\xf4\x16\x4a\xa2\xc8\x1f\xa5\x1f\x29\x39\x58\xab\xaf\xe4\xbc\xb2\x66\x0d\x58\x2b\xfa\xc6\x64\xe4\x97\x5f\xed\xff\xe4\x57\xca\x3e\x1e\xde\x6d\x89\xf1\xdd\x62\xaf\x4c\xbe\x86\xe7\xc6\xb3\xad\x7e\x22\x6f\x1b\x97\xd1\x7b\xda\x29\xa3\x58\x59\xb3\xa8\x88\x31\x47\xc6\xf5\x02\x20\x73\x84\x22\xfc\xa2\x2a\xf2\x8c\x55\xbd\x06\xd3\x68\xbd\x00\xd0\xb8\x25\xed\x65\x0c\x00\xd6\xf5\x1a\x3c\xa1\xcb\x4a\xad\x8a\x92\x17\x00\x06\x2b\x5a\x43\x8d\xd9\x1e\x0b\xf2\xab\x7d\xb3\x25\xf9\xb1\xca\xe9\xb0\xf0\x35\x65\x32\xb1\x70\xb6\xa9\xd7\x70\xd6\x17\x67\xb6\x7a\x33\x64\x2a\xac\x53\xe9\xf7\xb2\x1b\xdb\xfe\xc4\xba\xf6\x99\xcd\x29\xfc\x8c\x6e\x7d\x4a\xbd\x5a\x79\xfe\x4b\x27\xfa\x51\x79\x0e\xe2\x5a\x37\x0e\xf5\xc9\xb4\x20\xf4\xa5\x75\xfc\xf1\xb4\xf0\x12\xea\x7d\x11\x7b\x94\x29\x1a\x8d\xae\x9b\xb0\x00\xa8\x1d\x79\x72\x07\xfa\xbb\xd9\x1b\xfb\x62\x7e\x50\xa4\x73\xbf\x86\x1d\x6a\x2f\xdd\x3e\xb3\x35\xad\xe1\x59\x37\x9e\xc9\x89\xa0\xd9\xba\x16\xe6\x56\xbf\x67\xe4\xc6\xaf\xe1\x97\x5f\x17\x00\x07\xd4\x2a\x0f\x28\xc7\x4e\x5b\x93\x79\xfa\xb4\xf9\xfa\xfb\xcf\x59\x49\x15\x46\xa1\xac\x6a\x6b\x72\xdc\x81\x11\x81\xef\x42\xde\xc9\x00\x72\xf2\x99\x53\x75\xd0\x08\xf7\xa2\x2a\x8e\x81\x5c\x82\x4c\x1e\xb8\x24\x38\x44\x19\xe5\xe0\xc3\x32\x60\x77\xc0\xa5\xf2\xe0\x28\xf8\x67\x38\x98\xd4\x53\x0b\x32\x04\x0d\xd8\xed\x3f\x28\xe3\x15\x7c\x16\x0c\x9c\x17\xf0\x1a\x9d\x43\x66\xcd\x81\x1c\x83\xa3\xcc\x16\x46\xfd\xb3\xd3\xec\x81\x6d\x58\x52\x23\x53\x1b\x85\xd4\x94\x61\x72\x06\xb5\x80\xd0\xd0\x03\xa0\xc9\xa1\xc2\x23\x38\x92\x35\xa0\x31\x3d\x6d\x61\x88\x5f\xc1\x5f\xad\x23\x50\x66\x67\xd7\x50\x32\xd7\x7e\xfd\xf8\x58\x28\x4e\x9b\x3c\xb3\x55\xd5\x18\xc5\xc7\xc7\xcc\x1a\x76\x6a\xdb\xb0\x75\xfe\x31\xa7\x03\xe9\x47\xaf\x8a\xa5\xec\x52\xc5\x94\x71\xe3\xe8\x11\x6b\xb5\x0c\x86\x1b\x0e\x27\xa5\xca\x7f\xd7\x85\xea\xbe\x67\x29\x1f\x25\xa4\x9e\x9d\x32\x45\x27\x0e\x1b\x6e\x12\x77\xd9\x7b\xa0\x3c\x60\x3b\x2d\xda\x7f\x82\x57\x44\x82\xca\x4f\x1f\x3e\x7f\x81\xb4\x68\x08\xc1\x39\xe6\x01\xed\xd3\x34\x7f\x02\x5e\x80\x52\x66\x47\x2e\x06\x6e\xe7\x6c\x15\x34\x92\xc9\x6b\xab\x0c\x87\x1f\x99\x56\x64\xce\x41\xf7\xcd\xb6\x52\x2c\x91\xfe\xb9\x21\xcf\x12\x9f\x15\x3c\xa3\x31\x96\x61\x4b\xd0\xd4\x39\x32\xe5\x2b\xd8\x18\x78\xc6\x8a\xf4\x33\x7a\xfa\x8f\xc3\x2e\x08\xfb\xa5\x40\x7a\x1d\xf8\x3e\x43\x9d\x0f\x8c\x68\x75\xe2\xc4\x34\xa9\x8d\x9d\xa1\x40\x34\x25\x3a\x3e\x17\x4d\x0f\x96\x16\xe8\x6d\x20\x9d\x30\x37\xb5\xc6\xe9\x57\xce\x91\x08\x29\x47\xf9\xe5\xb4\x65\x30\x60\x20\x6c\x9c\x5e\x8c\xe9\xbe\x40\x25\x6a\x3e\x23\xa4\xdb\xbc\xb6\x2f\x66\x68\x8b\x34\xc5\x54\x8d\x8c\x87\xcb\x43\x91\xb2\xcd\xe6\x3d\xa8\x5c\x82\xbf\x53\x24\x27\x24\x59\x33\xaa\x61\xce\xa2\xd8\x62\x22\x99\xe8\xbc\x12\x94\xd8\x2e\x8f\xf2\x8c\x0f\xe9\x5c\xcb\xe1\xf2\xe4\x14\xea\xc0\x4e\xa2\x21\x72\x28\x75\xee\xac\x00\x36\x3c\xa9\x15\x44\x8b\xb1\xae\x42\xad\x8f\xa7\x93\x16\x38\x30\xa5\x9d\xd5\xf7\x38\x35\xb5\x49\x47\x9c\xba\x97\xe4\x97\xbc\x8a\x29\x32\x4c\xbf\xf4\x48\xa8\x3c\xe4\xbe\x6b\xae\x55\x8d\x67\xa8\x90\xb3\x32\xcc\xef\xeb\x9a\x2a\x3c\x96\x8e\x0a\xe5\xd9\x0d\xf3\xce\x85\xe7\xd6\xae\x5b\x1b\x57\x21\xf2\x01\x31\xc5\x71\xcd\x2d\x01\x6a\x0d\xda\xbe\x90\xcb\x84\xb8\xee\xbf\x07\xc2\x98\xcb\x6f\xc3\x30\x79\xf4\x59\xe6\x04\xe6\x37\x40\xa6\xa9\x62\xda\x4d\x74\x9f\xab\xdd\x8e\xdc\x25\x25\x8f\xac\xea\x01\x0f\xa8\x34\x6e\x75\x40\x1d\x21\x0b\xb8\xcd\x9f\x95\x1b\xdd\x3a\x0c\x8b\x86\x57\xea\x98\x62\xa6\xd8\x96\xf1\x48\x4e\xf4\xc9\x51\x99\xe8\x1a\xa1\xb5\xd4\x15\x40\x99\xe8\x6b\xfd\x19\xed\x9d\xa4\xc0\x7e\x37\x3a\x87\xc7\x41\xef\x9c\x93\x6f\x94\xf7\x46\x79\xe7\xed\x8d\xf2\xde\x28\xef\xff\x9c\xf2\xa6\xeb\xcd\x50\xfa\x0d\xa4\x69\xfc\xad\x55\xe7\x0b\x2a\xfe\xc1\xba\x01\xab\x4d\x90\xe9\x3c\x0d\xa2\x1e\x2d\xa7\xd3\xf2\x5b\x6b\x35\xe1\x18\x40\x3b\xeb\xe6\x26\x4e\x6e\x81\xfe\x93\xc7\x65\x3b\x3b\x16\x4f\x71\x28\x78\xd2\x94\xb1\x75\xf1\x1a\x18\x65\x3f\x37\xe4\x8e\x60\x0f\xe4\xe4\x66\x48\x2c\x14\xd1\x15\xe3\x53\xdc\xf7\x25\xf2\x51\xa3\xc3\xf0\xc0\x31\x3f\x06\x63\xda\x0b\x33\x67\xe5\x87\x6f\x72\x41\x0c\xef\x3d\x80\x8e\xe0\xe9\xe3\x7b\xb9\xc3\x3d\x4d\xd1\x0a\x55\x35\x1f\x2f\xed\x0c\x9a\xe4\x18\x6a\xdd\x86\xcf\xaf\xe0\x29\x3c\xf9\x5c\x0c\x9d\xd0\x9a\x14\x18\xdb\xcd\xff\x8d\x09\xee\xd2\xa9\x1b\x19\x69\x80\x45\x84\x5e\xf9\x80\xdc\x4d\x3e\xc0\xe9\x20\x54\xf1\xc2\x1d\xe1\x3f\x49\x7a\x00\x4f\xea\x98\xa9\x0f\x46\xcc\x1e\xec\x98\xde\x72\xed\x23\xc2\x75\xa3\x01\xb8\x44\x06\xb9\x80\xa3\x32\xbe\x7d\x30\x79\x00\x84\x3d\x1d\xe3\xdb\x0a\x1a\x10\xe0\x51\x96\x08\x83\x1d\x85\x57\x99\x2b\x5a\x49\x34\x04\x05\xed\x23\xcc\xcc\xf8\xeb\xa1\x8d\x6d\x4f\xc7\xf9\x01\x17\x10\x89\x05\x6d\xb6\x8e\x58\x89\x20\xf8\x10\xcb\x92\x1b\xe0\x81\xf8\x56\xa9\xa5\x02\x63\x3b\xe7\x04\xdc\x96\x54\x62\x4b\x88\xbe\xca\x9d\x2e\x0c\xa7\x97\x9d\x18\xa8\x7b\x1f\x83\x22\xbb\xb7\x54\xe3\x09\xe5\xcc\x4c\x7b\x22\x92\xf4\x44\xf6\x15\xb5\xca\xbb\x25\xe2\x7e\xdd\x98\x07\xf8\x68\x79\x63\x1e\xae\xaa\xfc\xf0\x4d\x79\x8e\xdc\xf2\xde\x92\xff\x68\x39\x48\xfe\x6d\x80\x45\x33\x5f\x05\x57\x9c\xd2\x96\x18\x21\x75\x89\xbf\xfd\x97\x35\xbf\x82\xcd\xee\x3a\x5a\x25\x9d\xa0\x57\x1e\x36\x06\xac\x6b\x71\x89\xef\xa2\x71\xa1\xb8\x84\x94\x57\x57\x55\x6e\x09\x8c\x35\xcb\x40\xa8\x62\xc3\x60\x8d\x16\x4e\xeb\xce\xd0\xbc\x1e\x86\x51\x73\x64\xb9\x76\xa9\x2f\xa5\x4a\x3d\xf1\xdd\x56\x63\x36\xc8\xc5\x23\xb8\x36\x01\xb4\xf0\x2e\x89\x4c\x85\xca\xa0\x22\x57\x10\xd4\xc2\x9d\xd7\x82\x7c\x95\xd7\x5a\xdb\x6f\xdd\x0b\xf3\x37\xb0\x53\x9b\xaf\xbe\x62\x5b\xca\xf9\x99\xed\x4f\x61\x99\xab\xb5\xe7\x8b\xa7\x5b\x6d\xee\x25\xe9\x69\x93\x31\xcf\xc3\x3d\x00\xf5\xa7\x9b\x58\xf3\x26\x54\x87\xf9\xb0\xad\x15\x42\x1e\xa9\xb0\x96\x93\xf3\x8b\xa4\x84\xb0\xb9\x7e\x85\x1a\x95\x93\x3c\x3f\x57\x9e\x2b\x53\x68\x3a\x9b\xa5\x4c\xd8\xa0\xfd\x05\x44\xb7\xf2\x20\x91\x3a\xa0\x9e\x2f\xf8\x85\xb6\x0c\x90\x8e\x29\x2e\x55\x35\xbd\xcc\xfd\x00\x2f\xa5\xf5\x31\xf3\xec\x14\xe9\x70\x03\xbd\xdb\xd3\xf1\x6e\xee\xe4\x5c\x9e\xbd\xbb\x8d\xb9\x8b\xa9\x6f\x70\x9a\xba\x3c\x69\x8d\x9e\xdb\x35\x77\x61\xd6\xdd\x6f\x2b\x03\xae\xee\xa6\x2b\x03\xa6\x2f\xb4\x57\xf6\x42\xaa\x29\xc7\x27\xff\x37\x1e\x18\xe6\x2d\xb8\x51\xc9\xff\xf6\xda\x36\x7b\xf1\xfc\xbe\x3b\x98\xaa\xc8\x36\x83\x6f\x0e\x37\xf8\x34\xed\xcf\x52\xea\xf6\x11\xe9\x6e\x84\xef\x96\xed\x6d\x66\xa4\x63\xf4\x42\x3a\x83\xc5\x32\x39\x33\xe8\x99\xc1\x60\x9c\x40\xc7\x7c\x5b\xc6\xaf\x33\x8b\x31\x5b\xfc\x99\x34\xdd\x2e\x17\x57\x2c\x68\xbf\xbe\xf6\x86\x4d\x1d\x07\xbb\x0d\x2f\x3f\xf9\x9f\xc9\x90\xeb\x7d\xa2\x3d\xb5\x33\xb2\xfd\xdb\x60\x78\xaa\x58\x2b\xeb\xc3\x97\x51\xe1\xba\xe2\xd4\x9b\xf4\x0f\x90\xdb\x59\x97\xbe\xc6\xa6\x77\xb5\x8d\x54\xf6\xce\x91\xaf\xad\xc9\xbb\x6f\xaa\xa9\xff\xde\xf7\xf4\x0a\x75\xaa\xac\x1c\x68\x55\x3e\x7d\xdc\x03\x6b\xa0\x6a\xe2\x27\x5e\xd8\x1e\x83\xaa\xa7\x4f\x9b\xf6\xd3\xe2\x25\x9d\xed\xac\xab\x90\xd7\xa0\x0c\xff\xf1\x0f\xa3\x81\x54\x86\xa9\x20\x37\x8f\xfc\x85\x28\x1d\x6d\x38\xbc\x43\x5d\x97\xf8\xee\x24\x0b\x61\x58\xb6\x7f\x22\xe8\x75\x43\x7c\x89\xcb\xd7\xc0\xae\x89\x3b\xd1\xb3\x75\x58\x50\x2b\xf9\x57\x00\x00\x00\xff\xff\x14\x52\xc3\xe9\x04\x21\x00\x00")

func kubepackDev_packagesYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubepackDev_packagesYaml,
		"kubepack.dev_packages.yaml",
	)
}

func kubepackDev_packagesYaml() (*asset, error) {
	bytes, err := kubepackDev_packagesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubepack.dev_packages.yaml", size: 8452, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"kubepack.dev_applications.yaml": kubepackDev_applicationsYaml,
	"kubepack.dev_bundles.yaml":      kubepackDev_bundlesYaml,
	"kubepack.dev_orders.yaml":       kubepackDev_ordersYaml,
	"kubepack.dev_packages.yaml":     kubepackDev_packagesYaml,
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
	"kubepack.dev_applications.yaml": {kubepackDev_applicationsYaml, map[string]*bintree{}},
	"kubepack.dev_bundles.yaml":      {kubepackDev_bundlesYaml, map[string]*bintree{}},
	"kubepack.dev_orders.yaml":       {kubepackDev_ordersYaml, map[string]*bintree{}},
	"kubepack.dev_packages.yaml":     {kubepackDev_packagesYaml, map[string]*bintree{}},
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
