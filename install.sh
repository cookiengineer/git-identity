#!/bin/bash

GO="$(which go)";
ROOT="$(pwd)";
SUDO="$(which sudo)";



build_and_install() {

	local os="${1}";
	local arch="${2}";

	cd "${ROOT}";

	env CGO_ENABLED=0 ${GO} build -o "/tmp/git-identity" "${ROOT}/main.go";

	if [[ "$?" == "0" ]]; then
		${SUDO} mv "/tmp/git-identity" "/usr/bin/git-identity";
		${SUDO} chmod +x "/usr/bin/git-identity";
	fi;

}

build_and_install "linux" "amd64";

