#!/usr/bin/env sh

set -eu

repo="$(mktemp -d)"
trap "rm -rf ${repo}" 0

git clone --depth 1 --branch main --quiet git@github.com:zellij-org/zellij.git "${repo}" 

cp "${repo}"/zellij-utils/src/plugin_api/*.proto "${PWD}/proto/"
