# Yeppeun

[![Go Reference](https://pkg.go.dev/badge/github.com/prophittcorey/yeppeun.svg)](https://pkg.go.dev/github.com/prophittcorey/yeppeun)

A command line tool with an optional web interface for pretty printing JSON.

![A screenshot demonstrating Yeppeun running in a browser.](.github/screenshot.png)

## Installation

```bash
go install github.com/prophittcorey/yeppeun/cmd/yeppeun@latest
```

## Usage

The command line tool can be use via pipes.

```bash
cat /tmp/dirty.json | yeppeun # outputs pretty json
```

The web interface can be started by running the executable.

```bash
yeppeun --host 0.0.0.0 --port 1234
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
