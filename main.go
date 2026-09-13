package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/xattr"
)

func main() {
	filePath := "test.txt"
	err := os.WriteFile(filePath, []byte("hello from xmeta"), 0o644)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer os.Remove(filePath)

	attrKey := "user.department"
	attrVal := []byte("engineering")

	err = xattr.Set(filePath, attrKey, attrVal)
	if err != nil {
		log.Fatalf("Failed to get xattr: %v", err)
	}
	fmt.Printf("Successfully set %s = %s\n", attrKey, attrVal)

	val, err := xattr.Get(filePath, attrKey)
	if err != nil {
		log.Fatalf("Failed to get xattr: %v", err)
	}
	fmt.Printf("Retrieved %s = %s\n", attrKey, string(val))

	attrs, err := xattr.List(filePath)
	if err != nil {
		log.Fatalf("Failed to list xattr: %v", err)
	}
	fmt.Printf("All metadata keys on file: %v\n", attrs)
}
