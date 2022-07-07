//go:build go1.19

package main

func asmAddPackagePathArg(args []string, packagePath string) []string {
	if packagePath != "" {
		args = append(args, "-p", packagePath)
	}
	return args
}
