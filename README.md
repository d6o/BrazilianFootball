![BrazilianFootball](http://image.prntscr.com/image/5b8a0445a0ce4360980cb03a819705bd.png)

# BrazilianFootball ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/BrazilianFootball) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

Display data about Brazilian football games that will occur in the next few days.

**Table of Contents**

- [Demo](#demo)
- [Project Status](#project-status)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Get all available currencies](#get-all-available-currencies)
  - [Convert 100 USD to all currencies](#convert-100-USD-to-all-currencies)
- [Contributing](#contributing)
  - [Bug Reports & Feature Requests](#bug-reports--feature-requests)
  - [Developing](#developing)
- [Social Coding](#social-coding)
- [License](#license)

## Demo

[![asciicast](https://asciinema.org/a/a0m1ln1k1u2ibbh1n4mbty1do.png)](https://asciinema.org/a/a0m1ln1k1u2ibbh1n4mbty1do)

## Project Status

BrazilianFootball is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/BrazilianFootball#social-coding)

## Features

- Open source - You can check out our code
- Secure
- Always up-to-date
- Use decimal type
- 100% satisfaction guaranteed
- It's perfect to keep up-to-date with all football matches
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/BrazilianFootball#usage)
- Very fast start up and response time
- Uses native libs
- Colorful

## Installation

### Option 1: Go Get

```bash
$ go get github.com/disiqueira/BrazilianFootball
```

## Usage

### Get 5 matches

```bash
$ BrazilianFootball
```

### Get 10 matches

```bash
$ BrazilianFootball -limit=10
```

### Filter by Team

```bash
$ BrazilianFootball -team="Corinthians"
```

### Filter by Championship

```bash
$ BrazilianFootball -championship="Campeonato Paulista"
```

### Filter by Date

```bash
$ BrazilianFootball -date="02/04/2017"
```

### Filter by Hour

```bash
$ BrazilianFootball -hour="10:00"
```

### Filter by Date and Hour
```bash
$ BrazilianFootball -date="02/04/2017" -hour="10:00"
```

### Filter by Filter by Day of week

```bash
$ BrazilianFootball -dayOfWeek="Dom"
```

### Filter by Location

```bash
$ BrazilianFootball -location="Arena Joinville"
```

### Filter by Round phase

```bash
$ BrazilianFootball -phase="6ª rodada"
```

### Filter by Status

```bash
$ BrazilianFootball -status="Encerrada"
```

### Show only today matches

```bash
$ BrazilianFootball -today=true
```

### Hide Finished matches

```bash
$ BrazilianFootball -upcoming=true
```

### Show Help

```bash
# Show help
$ BrazilianFootball -h
```

## Program help

![ProgramHelp](http://image.prntscr.com/image/fe46f19329144f67afa0f81ceb7f8dc5.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/BrazilianFootball/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone https://github.com/disiqueira/BrazilianFootball.git BrazilianFootball
$ cd BrazilianFootball/
$ go get -v -d 
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it](https://github.com/DiSiqueira/BrazilianFootball/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.