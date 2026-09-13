package xmeta

import "github.com/pkg/xattr"

func Set(filepath, attrKey string, attrVal []byte) error {
	err := xattr.Set(filepath, attrKey, attrVal)
	return err
}

func Get() {}

func Remove() {}

func List() {}
