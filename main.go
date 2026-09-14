package main

import (
	"fmt"
	"os"

	"xmeta/xmeta"
)

func main() {
	if len(os.Args) < 3 {
		help()
	}
	cmd, path := os.Args[1], os.Args[2]

	switch cmd {
	case "add": // xmeta add ./photo.jpg fingerprint hciow
		key, val := os.Args[3], os.Args[4]
		if err := xmeta.Set(path, key, []byte(val)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "get": // xmeta get ./photo.jpg fingerprint
		val, err := xmeta.Get(path, os.Args[3])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(string(val))

	case "list": // xmeta list ./photo.jpg
		keys, err := xmeta.List(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, k := range keys {
			v, _ := xmeta.Get(path, k)
			fmt.Printf("%s = %s\n", k, string(v))
		}

	case "remove": // xmeta remove ./photo.jpg fingerprint
		if err := xmeta.Remove(path, os.Args[3]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:

	}
}

// This just prints the help text
func help() {
	fmt.Println("usage: xmeta <add|get|list|remove> <file> [key] [value]")
	os.Exit(1)
}
