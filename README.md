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
  Small application to update the DigitalOcean DNS records with the current public IP address.
</p>

#### Table of contents

* [1. Overview](#-1-overview)
* [2. Development](#-2-development)
  - [1.1. Project Structure](#21-project-structure)
  - [2.2. Requirements](#22-requirements)
* [3. Appendix](#-3-appendix)
  - [3.1. Useful commands](#31-useful-commands)
* [4. How to contribute](#-4-how-to-contribute)
* [5. License](#-5-license)

## 🗂️ 1. Overview

TBC...

<sup>[Back to top ^][table-of-contents]</sup>

## 🛠️ 2. Development

### 2.1. Project structure

The project structure is based on the layout outlined in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

<sup>[Back to top ^][table-of-contents]</sup>

### 2.2. Requirements

* [Golang v1.26+](https://go.dev/doc/install)
* [Make](https://www.gnu.org/software/make/)

<sup>[Back to top ^][table-of-contents]</sup>

## 📑 3. Appendix

### 3.1. Useful commands

| Command        | Description                                                                             |
|----------------|-----------------------------------------------------------------------------------------|
| `make build`   | Builds and packages the application to `dist/dns-updater-<os>-<arch>-<version>.tar.gz`. |
| `make dev`     | Runs in development mode - builds and watches.                                          |
| `make format`  | Formats Go files.                                                                       |
| `make install` | Installs the development tools and downloads the Go modules.                            |
| `make test`    | Runs tests.                                                                             |

<sup>[Back to top ^][table-of-contents]</sup>

## 👏 4. How to contribute

Please read the [**contributing guide**](https://github.com/kieranroneill/dns-updater/blob/main/CONTRIBUTING.md) to learn about the development process.

<sup>[Back to top ^][table-of-contents]</sup>

## 📄 5. License

Please refer to the [LICENSE][license] file.

<sup>[Back to top ^][table-of-contents]</sup>

<!-- links -->
[license]: https://github.com/kieranroneill/dns-updater/blob/main/LICENSE
[table-of-contents]: #table-of-contents
