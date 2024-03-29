---
path: /unrot
---

unrot(1) - decode rotational ciphers
======================================

## SYNOPSIS

`unrot` [-hv] FILE

## DESCRIPTION

unrot decodes rotational "Caesar" ciphers, which are achieved by
rotating letters of the alphabet `N` times (giving them names like ROT14
and ROT10 ciphers).

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
unrot ./file.txt
```

Read from STDIN:

```
echo ./file.txt | unrot
```

## AUTHOR

Tom Scott <http://psychedeli.ca>

## SEE ALSO

byteme(1)
radix(1)
