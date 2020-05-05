---
path: /byteme
---

byteme(1) - decode binary and hexadecimal
======================================

## SYNOPSIS

`byteme` [-hv] FILE

## DESCRIPTION

byteme decodes binary and hexadecimal ciphers. Pass either a binary or
hex cipher and it will output whichever one is solved.

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
byteme ./file.txt
```

Read from STDIN:

```
echo ./file.txt | byteme
```

## AUTHOR

Tom Scott <http://psychedeli.ca>

## SEE ALSO

unrot(1)
radix(1)
