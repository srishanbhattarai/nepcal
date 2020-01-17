# nepcal &middot; [![Build Status](https://travis-ci.org/srishanbhattarai/nepcal.svg?branch=master)](https://travis-ci.org/srishanbhattarai/nepcal) [![Build status](https://ci.appveyor.com/api/projects/status/6vm0m2ph6usjvdn4/branch/master?svg=true)](https://ci.appveyor.com/project/srishanbhattarai/nepcal-j10el/branch/master) [![Coverage Status](https://coveralls.io/repos/github/srishanbhattarai/nepcal/badge.svg?branch=master)](https://coveralls.io/github/srishanbhattarai/nepcal?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/srishanbhattarai/nepcal)](https://goreportcard.com/report/github.com/srishanbhattarai/nepcal) [![GoDoc](https://godoc.org/github.com/srishanbhattarai/nepcal?status.svg)](https://godoc.org/github.com/srishanbhattarai/nepcal)

> Calendar and conversion utilities for Nepali dates

Inspired from the Linux command line tool `cal`, `nepcal` adds a few nifty features especially for Nepali ([B.S.](https://en.wikipedia.org/wiki/Vikram_Samvat)) dates.

## Feature Rundown

Complete instructions on how to use each of these features are mentioned below.

* [x] Show the current Nepali month's calendar
* [x] Show today's Nepali date and day
* [x] Convert an A.D. date into B.S.
* [x] Convert an B.S. date into A.D.

## Installation

Pre-built tarball binaries are available in the [Releases](https://github.com/srishanbhattarai/nepcal/releases) page. Download and untar the binary for your platform, then move it into your `$PATH` e.g. `/usr/local/bin`.

You might need to give the script execution permissions. On Linux and MacOS this would mean using `chmod` as follows:

```
$ chmod +x /usr/local/bin/nepcal
```

### MacOS via Homebrew

Tap the repository first.

```
$ brew tap srishanbhattarai/nepcal https://github.com/srishanbhattarai/nepcal
```

Then, run:

```
$ brew install nepcal
```

### Manual Installation

You can also install `nepcal` manually if you have Go installed

```
$ go get -v github.com/srishanbhattarai/nepcal
```

Run `nepcal` on your terminal - if you see some formatted output, you are good to go!

## Usage

Complete details can be found by running `nepcal` without any arguments.

### Monthly Calendar

```
$ nepcal cal # or nepcal c

    जेठ 3, 2075
 Su Mo Tu We Th Fr Sa
       1  2  3  4  5
 6  7  8  9  10 11 12
 13 14 15 16 17 18 19
 20 21 22 23 24 25 26
 27 28 29 30 31 32
```

### Today's date and day

```
$ nepcal date # or nepcal d

साउन 29, 2075 मंगलबार
```

### Convert A.D. date to B.S.

Use the `mm-dd-yyyy` format when converting A.D. to B.S.

```
$ nepcal conv adtobs 08-21-1994

भदौ 5, 2051 आइतबार
```

### Convert B.S. date to A.D.

Use the `mm-dd-yyyy` format when converting B.S. to A.D.

```
$ nepcal conv bstoad 18-08-2053

December 3, 1996 Tuesday
```

## Contributing

Please file an issue if you have any problems with `nepcal` or, have a look at the issues page for contributing on existing issues. Also, read the [code of conduct](https://github.com/srishanbhattarai/nepcal/blob/master/CODE_OF_CONDUCT.md).

## License

[MIT](LICENSE)
