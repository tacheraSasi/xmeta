package xmeta

import "github.com/pkg/xattr"

// Set attaches attrVal to filepath under attrKey.
func Set(filepath, attrKey string, attrVal []byte) error {
	err := xattr.Set(filepath, attrKey, attrVal)
	return err
}

// Get retrieves the value stored under attrKey on filepath.
func Get(filepath, attrKey string) ([]byte, error) {
	val, err := xattr.Get(filepath, attrKey)
	return val, err
}

// Remove deletes attrKey from filepath.
func Remove(filepath, attrKey string) error {
	err := xattr.Remove(filepath, attrKey)
	return err
}

// List returns all attribute keys set on filepath.
func List(filepath string) ([]string, error) {
	attrs, err := xattr.List(filepath)
	return attrs, err
}
