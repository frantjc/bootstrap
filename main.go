package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer client.Close()

	odamex := client.Container().
		From("ubuntu:24.04").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y",
			"build-essential",
			"cmake",
			"deutex",
			"libcurl4-gnutls-dev",
			"libasound2-dev",
			"libegl1-mesa-dev",
			"libcairo2-dev",
			"libgl1-mesa-dev",
			"libglu1-mesa-dev",
			"libjsoncpp-dev",
			"libpango1.0-dev",
			"libpng-dev",
			"libportmidi-dev",
			"libsdl2-dev",
			"libsdl2-mixer-dev",
			"libwayland-dev",
			"libx11-dev",
			"libxft-dev",
			"libzstd-dev",
			"zlib1g-dev",
		}).
		WithWorkdir("/odamex").
		WithMountedDirectory("/odamex", client.Host().Directory(".")).
		WithExec([]string{"mkdir", "-p", "build"}).
		WithWorkdir("/odamex/build").
		WithExec([]string{"cmake", "..",
			"-DBUILD_SERVER=0",
			"-DBUILD_CLIENT=1",
			"-DBUILD_LAUNCHER=0",
			"-DBUILD_MASTER=0",
			"-DBUILD_OR_FAIL=1",
			"-DCMAKE_BUILD_TYPE=MinSizeRel",
			"-DCMAKE_C_COMPILER=gcc",
			"-DCMAKE_C_FLAGS=-O2 -w",
			"-DCMAKE_CXX_COMPILER=g++",
			"-DCMAKE_CXX_FLAGS=-O2 -w",
		}).
		WithExec([]string{"make", "-j", "$(nproc)", "odamex", "install"})

	binDir := client.Host().Directory("/usr/local/bin")
	shareDir := client.Host().Directory("/usr/local/share/odamex")

	_, err = odamex.Directory("/usr/local/bin").Export(ctx, binDir.Path())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = odamex.Directory("/usr/local/share/odamex").Export(ctx, shareDir.Path())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}