# NameBeta

[![Release][3]][4] [![MIT licensed][5]][6] [![Build Status][1]][2] [![Go Report Card][7]][8]

[1]: https://travis-ci.org/TimothyYe/namebeta.svg?branch=master
[2]: https://travis-ci.org/TimothyYe/namebeta
[3]: https://img.shields.io/badge/release-v0.2-brightgreen.svg
[4]: https://github.com/TimothyYe/namebeta/releases
[5]: https://img.shields.io/dub/l/vibe-d.svg
[6]: LICENSE
[7]: https://goreportcard.com/badge/github.com/timothyye/namebeta
[8]: https://goreportcard.com/report/github.com/timothyye/namebeta

NameBeta is a command line domain query tool, inspired by [NameBeta.com](https://namebeta.com).

![](https://raw.githubusercontent.com/TimothyYe/namebeta/master/snapshots/namebeta.gif)

## Features

* Domain query
* WHOIS query

## Installation

#### Homebrew

```bash
brew tap timothyye/tap
brew install timothyye/tap/namebeta
```

#### Using Go

```bash
go get github.com/TimothyYe/namebeta
```

#### Manual Installation

Download it from [releases](https://github.com/TimothyYe/namebeta/releases), extract it and install it:

```bash
cp namebeta /usr/local/bin/
```

## Usage

1. Query domain

```text
namebeta <domain>
```

2. Query domain with more results

```text
namebeta -m <domain>
```

3. WHOIS query

```text
namebeta -w <domain>
```

## Help

Just type "namebeta" to get help.
  
## Licence

[MIT License](https://github.com/TimothyYe/namebeta/blob/master/LICENSE)
