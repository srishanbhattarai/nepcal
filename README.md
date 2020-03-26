# nepcal &middot; [![Build Status](https://travis-ci.org/srishanbhattarai/nepcal.svg?branch=master)](https://travis-ci.org/srishanbhattarai/nepcal) [![Build status](https://ci.appveyor.com/api/projects/status/6vm0m2ph6usjvdn4/branch/master?svg=true)](https://ci.appveyor.com/project/srishanbhattarai/nepcal-j10el/branch/master) [![Coverage Status](https://coveralls.io/repos/github/srishanbhattarai/nepcal/badge.svg?branch=master&service=github)](https://coveralls.io/github/srishanbhattarai/nepcal?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/srishanbhattarai/nepcal)](https://goreportcard.com/report/github.com/srishanbhattarai/nepcal) [![GoDoc](https://godoc.org/github.com/srishanbhattarai/nepcal?status.svg)](https://godoc.org/github.com/srishanbhattarai/nepcal)

Nepcal is a command line tool and a library that provides several functionalities pertaining to [Bikram Sambat (B.S.)](https://calendars.wikia.org/wiki/Bikram_Samwat) calendars, the official calendar in Nepal.

## Table of Contents

* [Command Line](#command-line)
* [Installation](#installation)
  * [Pre-built binaries](#pre-built-binaries)
  * [Homebrew on Mac](#homebrew-on-mac)
  * [Using go get](#using-go-get)
* [Usage](#usage)
  * [Monthly Calendar](#monthly-calendar)
  * [Today's date and day](#todays-date-and-day)
  * [Convert an A.D. date to B.S.](#convert-an-ad-date-to-bs)
  * [Convert a B.S. date to A.D.](#convert-a-bs-date-to-ad)
* [Contributing](#contributing)
* [Library/Programmatic usage](#library)
* [License](#license)

## Command Line

The `nepcal` CLI was initially inspired from the `cal` command on Linux, but expanded to have much more functionality, namely:

* [x] Show the current Nepali month's calendar (Similar to `cal`)
* [x] Show today's Nepali date and day
* [x] Convert A.D. (gregorian) dates to B.S. dates and vice-versa.

## Installation

There are several ways to install the CLI. Pick the one that best suits you from the options below. In any case, run `nepcal` on your terminal; if you see some formatted output then you are good to go.

### Pre-built binaries

Pre-built tarball executables/binaries are available in the [Releases](https://github.com/srishanbhattarai/nepcal/releases) page for MacOS, Windows, and Linux.
Download and untar the binary for your platform, then move it into your executable path. e.g. `/usr/local/bin` on Mac or Linux.

You might need to give the script execution permissions. On Linux and MacOS this would mean using `chmod` as follows:

```sh
$ chmod +x /usr/local/bin/nepcal
```

### Homebrew on Mac

Tap the repository first.

```sh
$ brew tap srishanbhattarai/nepcal https://github.com/srishanbhattarai/nepcal
```

Then, run:

```sh
$ brew install nepcal
```

### Using `go get`

You can also install `nepcal` manually if you have the Go toolchain installed.

```sh
$ go get -v github.com/srishanbhattarai/nepcal/cmd/nepcal
```

## Usage

Complete details can be found by running `nepcal` without any arguments.

### Monthly Calendar

```sh
$ nepcal cal # or nepcal c
    चैत ११, २०७६
 Su Mo Tu We Th Fr Sa
                   १
 २  ३  ४  ५  ६  ७  ८
 ९  १० ११ १२ १३ १४ १५
 १६ १७ १८ १९ २० २१ २२
 २३ २४ २५ २६ २७ २८ २९
 ३०
```

### Today's Date

```sh
$ nepcal date # or nepcal d

चैत १०, २०७६ सोमबार
```

### Convert A.D. to B.S.

Use the `mm-dd-yyyy` format when converting A.D. to B.S.

```sh
$ nepcal conv tobs 08-21-1994

भदौ ५, २०५१ आइतबार
```

### Convert B.S. to A.D.

Use the `mm-dd-yyyy` format when converting B.S. to A.D.

```sh
$ nepcal conv toad 08-18-2053

December 3, 1996 Tuesday
```

## Library

If you would like to use `nepcal` as a Go library, the best reference is the [Godoc](https://godoc.org/github.com/srishanbhattarai/nepcal/nepcal) documentation for this package which should be fairly easy to navigate. The CLI tool is also built on this library. However, there are additional functionalities provided in the library that are not relevant in the CLI, for example the [`NumDaysSpanned()`](https://godoc.org/github.com/srishanbhattarai/nepcal/nepcal#Time.NumDaysSpanned) method.

## Contributing

Please file an issue if you have any problems with `nepcal` or, have a look at the issues page for contributing on existing issues. Also, read the [code of conduct](https://github.com/srishanbhattarai/nepcal/blob/master/CODE_OF_CONDUCT.md).

## License

[MIT](LICENSE)
