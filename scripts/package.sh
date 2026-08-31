#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}/utilities/_set_vars.sh"

# Public: Packages the binary.
#
# See https://go.dev/doc/install/source#environment for a list of supported OS/arch combos.
#
# $1 - [optional] The OS and architecture to build to. It must be formatted in as "<OS>-<ARCH>".
#
# Examples
#
#   ./scripts/package.sh # uses the default set at `go env GOOS` & `go env GOARCH`
#   ./scripts/package.sh "darwin-amd64"
#   ./scripts/package.sh "darwin-arm64"
#   ./scripts/package.sh "linux-amd64"
#   ./scripts/package.sh "linux-arm64"
#   ./scripts/package.sh "windows-amd64"
#   ./scripts/package.sh "windows-arm64"
#
# Returns exit code 0. If no binary file exists at the OS/arch subdirectory, exit 1 is returned.
function main() {
  local arch
  local binary_file_path
  local os
  local version

  _set_vars

  if [[ -n "${1}" ]]; then
    IFS='-' read -r os arch <<< "${1}"
  fi

  if [[ -z "${os}" ]]; then
    os=$(go env GOOS)
  fi

  if [[ -z "${arch}" ]]; then
    arch=$(go env GOARCH)
  fi


  binary_file_path="${PWD}/${BUILD_DIR}/${os}-${arch}"

  if [ ! -f "${binary_file_path}/${EXECUTABLE}" ]; then
    printf "%b \"%b\" file not found \n" "${INFO_PREFIX}" "${binary_file_path}/${EXECUTABLE}"

    exit 1
  fi

  # create the dist directory if it does not exist
  if [ ! -d "${DIST_DIR}" ]; then
    printf "%b no \"%b\" directory found, creating a new one \n" "${INFO_PREFIX}" "${DIST_DIR}"

    mkdir -p "${DIST_DIR}"
  fi

  # get the version in the version file
  version=$(<VERSION)

  tar -czf "${PWD}/${DIST_DIR}/${EXECUTABLE}-${os}-${arch}-${version}.tar.gz" -C "${binary_file_path}" "${EXECUTABLE}"

  printf "%b package at \"%b\" \n" "${INFO_PREFIX}" "${PWD}/${DIST_DIR}/${EXECUTABLE}-${os}-${arch}-${version}.tar.gz"

  exit 0
}

# and so, it begins...
main "$@"
