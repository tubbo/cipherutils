# cipherutils

A suite of programs for solving trivial ciphers, typically used in
alternate reality games.

## Installation

The easiest way to install **cipherutils** is by using our handy-dandy
[installation script](install.sh):

```bash
$ curl -sL https://tubbo.github.io/cipherutils/install.sh | bash
```

This will find the latest release from GitHub, detect your platform, and
download/install the correct tarball. You can [view its source code][]
for more information.

### From Source

Installing from source is as easy as cloning the repo and running
`make`. You'll need [Go][] installed in order for this to work.

```bash
$ git clone https://github.com/tubbo/cipherutils.git
$ cd cipherutils
$ make
$ sudo make install
```

### Using Go Get

You can also `go get` any of the individual CLI tools:

```bash
$ go get github.com/tubbo/cipherutils/unrot
$ go get github.com/tubbo/cipherutils/byteme
$ go get github.com/tubbo/cipherutils/radix
```

## Usage

Each of the cipherutils has their own man pages, so running `man` on
any of the provided commands will show detailed usage information. All
of the command provided by cipherutils take input in a similar way, by
either reading from STDIN or a provided filename. They also only respond
back with information when a match is found in the built-in dictionary
of over 500,000 words.

The cipherutils are:

- `unrot` for decoding rotational (ROTN) ciphers
- `byteme` for decoding byte-based ciphers such as hexadecimal or binary
- `radix` for decoding numerical ciphers such as base64, base32, etc.

For example, to decode the ROTN-encoded `"uryyb jbeyq"`, one might call:

```bash
$ echo "uryyb jbeyq" | unrot
```

This will produce the output:

```
ROT14:  HELLO WORLD
```

However, if you were to decode the String `"nfsy"`, which is just a
bunch of garbled letters, you wouldn't receive any response. This is
because a match could not be found in the dictionary. To bypass the
dictionary lookup, you can pass `-v` to any of the cipherutils commands:

```bash
$ echo "uryyb jbeyq" | unrot -v
ROT25:  SPWWZ HZCWO
ROT4:   XUBBE MEHBT
ROT8:   BYFFI QILFX
ROT9:   CZGGJ RJMGY
ROT10:  DAHHK SKNHZ
ROT14:  HELLO WORLD
ROT12:  FCJJM UMPJB
ROT16:  JGNNQ YQTNF
ROT15:  IFMMP XPSME
ROT21:  OLSSV DVYSK
ROT6:   ZWDDG OGJDV
ROT11:  EBIIL TLOIA
ROT1:   URYYB JBEYQ
ROT23:  QNUUX FXAUM
ROT18:  LIPPS ASVPH
ROT17:  KHOOR ZRUOG
ROT13:  GDKKN VNQKC
ROT5:   YVCCF NFICU
ROT24:  ROVVY GYBVN
ROT19:  MJQQT BTWQI
ROT2:   VSZZC KCFZR
ROT7:   AXEEH PHKEW
ROT22:  PMTTW EWZTL
ROT20:  NKRRU CUXRJ
ROT3:   WTAAD LDGAS
```

You can also pass `-h` to any command to view its manual page.

## Running Tests

To run tests for all programs:

```bash
$ make check
```

## Contributing

Got an idea for a tool not included in the ciphertools, or a feature
request/bug report? Head on over to the [issue tracker][] to file a new
issue and we'll get right to work on it. Or, you can [submit a pull request][]
to the project if you have a patch that you'd like included.

[Go]: https://golang.org
[issue tracker]: https://github.com/tubbo/cipherutils/issues
[submit a pull request]: https://github.com/tubbo/cipherutils/pulls
[view its source code]: https://github.com/tubbo/cipherutils/blob/master/docs/install.sh
