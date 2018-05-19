# nepcal &middot; [![Build Status](https://travis-ci.org/nepcal/nepcal.svg?branch=master)](https://travis-ci.org/nepcal/nepcal) [![Build status](https://ci.appveyor.com/api/projects/status/2kg411mphyaor4ki?svg=true)](https://ci.appveyor.com/api/projects/status/2kg411mphyaor4ki?svg=true) [![Coverage Status](https://coveralls.io/repos/github/nepcal/nepcal/badge.svg?branch=master)](https://coveralls.io/github/nepcal/nepcal?branch=master)
Print the current Nepali date in the terminal.

Nepcal is a minimal version of the Linux command line tool `cal` for [Vikram Samvat](https://en.wikipedia.org/wiki/Vikram_Samvat) dates. It only implements the default command, which shows the current date, and month in a neat calendar.
```
    जेठ 3, 2075
 Su Mo Tu We Th Fr Sa
       1  2  3  4  5
 6  7  8  9  10 11 12
 13 14 15 16 17 18 19
 20 21 22 23 24 25 26
 27 28 29 30 31 32
```

## Usage
Run `nepcal` on your terminal to see today's date and calendar for the month. For viewing only the date,
use the `-d` flag as follows

```
$ nepcal -d
```

## Installation
Pre-built binaries are available in the [Releases](https://github.com/nepcal/nepcal/releases) page.

You can also install `nepcal` manually if you have Go installed
```
go get -v github.com/nepcal/nepcal
```

## Contributing
Please file an issue if you have any problems with `nepcal` or, have a look at the issues page for contributing on existing issues. Also, read the [code of conduct](https://github.com/nepcal/nepcal/blob/master/CODE_OF_CONDUCT.md).

## License
MIT
