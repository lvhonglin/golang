package main

import (
	"fmt"
	version "github.com/hashicorp/go-version"
)

func main() {
	oneVersion, err := version.NewVersion("v1.1.1")
	if err != nil {
		return
	}
	twoVersion, err := version.NewVersion("v1.1.2")
	if err != nil {
		return
	}
	fmt.Println(oneVersion.Equal(twoVersion))
	fmt.Println(oneVersion.LessThan(twoVersion))
	fmt.Println(oneVersion.GreaterThan(twoVersion))
	fmt.Println(oneVersion.Prerelease())
}
