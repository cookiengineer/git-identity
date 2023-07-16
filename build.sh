#!/bin/bash

GO="$(which go)";
ROOT="$(pwd)";
SUDO="$(which sudo)";


build() {

	local os="${1}";
	local arch="${2}";
	local folder="${ROOT}/build";

	if [[ ! -d "${folder}" ]]; then
		mkdir -p "${folder}";
	fi;

	cd "${ROOT}";

	if [[ "${os}" == "windows" ]]; then
		env CGO_ENABLED=0 ${GO} build -o "${folder}/git-identity_${os}_${arch}.exe" "${ROOT}/main.go";
	else
		env CGO_ENABLED=0 ${GO} build -o "${folder}/git-identity_${os}_${arch}" "${ROOT}/main.go";
	fi;

}

build "windows" "amd64";

build "linux" "amd64";
build "linux" "arm";
build "linux" "arm64";

build "darwin" "amd64";
build "darwin" "arm64";

build "freebsd" "amd64";
build "freebsd" "arm64";

build "openbsd" "amd64";
build "openbsd" "arm64";

