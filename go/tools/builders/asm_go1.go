//go:build !go1.19

package main

func asmAddPackagePathArg(args []string, packagePath string) []string {
	// go tool asm does neither support nor require the -p arg in Go < 1.19.
	return args
}
