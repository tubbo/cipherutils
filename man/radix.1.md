radix(1) - decode baseNN ciphers
======================================

## SYNOPSIS

`radix` FILE

or

STDIN | `radix`

## DESCRIPTION

radix decodes a cipher written in base64, base32, or any applicable
baseNN obfuscation method, and returns the winning result to the
terminal.

Part of the cipherutils: <https://tubbo.github.io/cipherutils>

## ARGUMENTS

*FILE*
  Path to a file on disk.

## EXAMPLES

Read from a file:

```
radix ./file.txt
```

Read from STDIN:

```
echo ./file.txt | radix
```

## AUTHOR

Tom Scott <http://psychedeli.ca>

## SEE ALSO

byteme(1)
unrot(1)
