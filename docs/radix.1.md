---
path: /radix
---

radix(1) - decode baseNN ciphers
======================================

## SYNOPSIS

`radix` [-hv] FILE

## DESCRIPTION

radix decodes a cipher written in base64, base32, or any applicable
baseNN obfuscation method, and returns the winning result to the
terminal.

Part of the cipherutils: <https://tubbo.github.io/cipherutils>

## ARGUMENTS

*FILE*
  Path to a file on disk. Data can also be supplied via `STDIN`

## OPTIONS

  * `-h`: Show usage information
  * `-v`: Verbose mode, show all solutions even if they don't match in
    the dictionary

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
