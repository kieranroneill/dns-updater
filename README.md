<div align="center">

[![License: CC0-1.0](https://img.shields.io/github/license/kieranroneill/dns-updater)][license]

</div>

<div align="center">

[![GitHub Pre-release](https://img.shields.io/github/v/release/kieranroneill/dns-updater?include_prereleases&label=pre-release&logo=github)](https://github.com/kieranroneill/dns-updater/releases)
[![GitHub Pre-release Published At](https://img.shields.io/github/release-date-pre/kieranroneill/dns-updater?label=pre-release%20date&logo=github)](https://github.com/kieranroneill/dns-updater/releases)

</div>

<div align="center">

[![GitHub Release](https://img.shields.io/github/v/release/kieranroneill/dns-updater?&logo=github)](https://github.com/kieranroneill/dns-updater/releases/latest)
[![GitHub Release Published At](https://img.shields.io/github/release-date/kieranroneill/dns-updater?logo=github)](https://github.com/kieranroneill/dns-updater/releases/latest)

</div>

<h1 align="center">
   DNS Updater
</h1>

<p align="center">
  Application to update a DigitalOcean DNS record with the current public IP address.
</p>

#### Table of contents

* [1. Overview](#-1-overview)
  - [1.1. Project Structure](#11-project-structure)
* [2. Installation](#-2-installation)
* [3. Usage](#-3-usage)
  - [3.1. `log`](#31-log)
* [4. Development](#-4-development)
  - [4.1. Requirements](#41-requirements)
* [5. Appendix](#-5-appendix)
  - [5.1. Useful commands](#51-useful-commands)
* [6. How to contribute](#-6-how-to-contribute)
* [7. License](#-7-license)

## 🗂️ 1. Overview

### 1.1. Project structure

The project structure is based on the layout outlined in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

<sup>[Back to top ^][table-of-contents]</sup>

## 📦 2. Installation

<sup>[Back to top ^][table-of-contents]</sup>

## 🪄 3. Usage

### 3.1. `log`

#### Flags

| Name          | Shorthand | Description                                                         | Example           |
|---------------|-----------|---------------------------------------------------------------------|-------------------|
| `--lineCount` | `-l`      | Specifies the the amount of lines to be printed in the log command. | `--lineCount 100` |

#### Examples

| Example                  | Description                                                                               |
|--------------------------|-------------------------------------------------------------------------------------------|
| `dns-updater log`        | Prints the last 50 lines of the log file from `$HOME/.dns-updater/logs/dns-updater.log`.  |
| `dns-updater log -l 100` | Prints the last 100 lines of the log file from `$HOME/.dns-updater/logs/dns-updater.log`. |

<sup>[Back to top ^][table-of-contents]</sup>

## 🛠️ 4. Development

### 4.1. Requirements

* [Golang v1.26+](https://go.dev/doc/install)
* [Make](https://www.gnu.org/software/make/)

<sup>[Back to top ^][table-of-contents]</sup>

### 4.2. Setup

To set up the environment, run:

```shell
make install
```

This will install all the project tools and the Go modules. It will also set up Git hooks (see the [`.hooks/`](./.hooks) directory for the hooks that are used).

<sup>[Back to top ^][table-of-contents]</sup>

### 4.3. Usage

Firstly, the binary needs to be built using:

```shell
make build
```

This will build the binary to the `.build/<os>-<arch>/` directory.

> ⚠️ **NOTE**: The subdirectory is based on the OS and architecture and uses the default set at `go env GOOS` and `go env GOARCH`, respectively.

Once the binary is built, it can be run using:

```shell
./.build/<os>-<arch>/dns-updater [COMMAND]
```

<sup>[Back to top ^][table-of-contents]</sup>

## 📑 5. Appendix

### 5.1. Useful commands

| Command                | Description                                                                 |
|------------------------|-----------------------------------------------------------------------------|
| `make build`           | Builds the CLI binary to `.build/<os>-<arch>/dns-updater`.                  |
| `make format`          | Formats Go files.                                                           |
| `make install`         | Installs tools, downloads required Go modules and sets up Git hooks.        |
| `make install_modules` | Installs the required Go modules.                                           |
| `make install_tools`   | Installs tools.                                                             |
| `make package`         | Packages the CLI binary to `dist/dns-updater-<os>-<arch>-<version>.tar.gz`. |
| `make postinstall`     | Sets up Git hooks.                                                          |
| `make test`            | Runs tests.                                                                 |

<sup>[Back to top ^][table-of-contents]</sup>

## 👏 6. How to contribute

Please read the [**contributing guide**](https://github.com/kieranroneill/dns-updater/blob/main/CONTRIBUTING.md) to learn about the development process.

<sup>[Back to top ^][table-of-contents]</sup>

## 📄 7. License

Please refer to the [LICENSE][license] file.

<sup>[Back to top ^][table-of-contents]</sup>

<!-- links -->
[license]: https://github.com/kieranroneill/dns-updater/blob/main/LICENSE
[table-of-contents]: #table-of-contents
