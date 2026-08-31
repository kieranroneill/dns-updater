#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}/utilities/_set_vars.sh"

# Public: Builds the Go application with the version from the VERSION file injected to the `./.build/<os>-<arch>/dns-updater`.
#
# See https://go.dev/doc/install/source#environment for a list of supported OS/arch combos.
#
# $1 - [optional] The OS and architecture to build to. It must be formatted in as "<OS>-<ARCH>".
#
# Examples
#
#   ./scripts/build.sh # uses the default set at `go env GOOS` & `go env GOARCH`
#   ./scripts/build.sh "darwin-amd64"
#   ./scripts/build.sh "darwin-arm64"
#   ./scripts/build.sh "linux-amd64"
#   ./scripts/build.sh "linux-arm64"
#   ./scripts/build.sh "windows-amd64"
#   ./scripts/build.sh "windows-arm64"
#
# Returns exit code 0.
function main() {
  local arch
  local build_subdir
  local os
  local version

  _set_vars

  # get the version in the version file
  version=$(<VERSION)

  if [[ -n "${1}" ]]; then
    IFS='-' read -r os arch <<< "${1}"
  fi

  if [[ -z "${os}" ]]; then
    os=$(go env GOOS)
  fi

  if [[ -z "${arch}" ]]; then
    arch=$(go env GOARCH)
  fi

  build_subdir="${os}-${arch}"

  # build the server
  GOOS="${os}" GOARCH="${arch}" go build \
    -ldflags="-X main.Version=${version}" \
    -o "${PWD}/${BUILD_DIR}/${build_subdir}/${EXECUTABLE}" \
    "${PWD}/cmd/${EXECUTABLE}/main.go"

  printf "%b build at \"%b\"\n" "${INFO_PREFIX}" "${PWD}/${BUILD_DIR}/${build_subdir}/${EXECUTABLE}"

  exit 0
}

# and so, it begins...
main "$@"
