# slgo

[![Go Reference](
  <https://pkg.go.dev/badge/github.com/eggplants/slgo.svg>
)](
  <https://pkg.go.dev/github.com/eggplants/slgo>
) [![ci](
  <https://github.com/eggplants/slgo/actions/workflows/ci.yml/badge.svg>
)](
  <https://github.com/eggplants/slgo/actions/workflows/ci.yml>
) [![release](
  <https://github.com/eggplants/slgo/actions/workflows/release.yml/badge.svg>
)](
  <https://github.com/eggplants/slgo/actions/workflows/release.yml>
)

A Go port of [sl](https://github.com/mtoyoda/sl).

## Install

```sh
go install github.com/eggplants/slgo/cmd/sl@latest
```

## Usage

```sh
sl [-a] [-F] [-l] [-c] [-p]
```

| Option | Effect                       |
| ------ | ---------------------------- |
| `-a`   | An accident occurs           |
| `-F`   | The train flies              |
| `-l`   | Show a small train (SL logo) |
| `-c`   | Show C51 instead of D51      |
| `-p`   | [sl5-1.patch](https://www.izumix.xyz/sl/sl5-1.patch) mode |

Options can be combined, for example `-aF`. As in the original, you cannot stop the train with Ctrl-C.

## License

[MIT](LICENSE)

The ASCII art and animation logic come from sl by [TOYODA Masashi](https://github.com/mtoyoda/).

The `-p` mode is based on [sl5-1.patch](https://www.izumix.xyz/sl/sl5-1.patch) by [IZUMI Tomonori](https://www.izumix.xyz/).
