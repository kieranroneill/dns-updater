#!/usr/bin/env bash

SCRIPT_DIR=$(dirname "${0}")

source "${SCRIPT_DIR}/utilities/_set_vars.sh"

# Public: Injects the version and builds the Go app with watch.
#
# Examples
#
#   ./scripts/dev.sh
#
# Returns exit code 0.
function main() {
  local arch
  local build_subdir
  local os
  local version

  _set_vars

  arch=$(go env GOARCH)
  os=$(go env GOOS)
  build_subdir="${os}-${arch}"

  # get the version in the version file
  version=$(<VERSION)

  # load env variables
  set -a
  source "${PWD}/.env.dev"
  set +a

  printf "%b starting app in watch mode...\n" "${INFO_PREFIX}"
  CompileDaemon \
    -build="go build -ldflags=-X=main.Version=${version} -o ${PWD}/${BUILD_DIR}/${build_subdir}/${EXECUTABLE} ${PWD}/cmd/${EXECUTABLE}/main.go" \
    -command="${PWD}/${BUILD_DIR}/${build_subdir}/${EXECUTABLE}"

  printf "%b done!\n" "${INFO_PREFIX}"

  exit 0
}

# and so, it begins...
main "$@"
