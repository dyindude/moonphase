# MoonPhase
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/janczer/MoonPhase/master/LICENSE)

Package `moonphase` allows calculation of the phase of Moon and other related data. It's an implementation of [php-moon-phase](https://github.com/solarissmoke/php-moon-phase) in go.

## Installation

To install the package on your system, run

```
go get github.com/dyindude/moonphase
```

## Quick Start

```
time := time.Date(2007, 10, 1, 24, 0, 0, 0, time.UTC)
//time := time.Now()
m := MoonPhase.New(time)
```

## License

moonphase is released under the MIT License. It is copyrighted by Ivan Menshykov and
the contributors acknowledged below.

github.com/dyindude - Andrew Missel

## Acknowledgments

This package's code and documentation are very closely derived from the [php-moon-phase](https://github.com/solarissmoke/php-moon-phase)
PHP class for calculating the phase of the Moon created by Samir Shah.
