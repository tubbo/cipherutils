unrot(1) - decode rotational ciphers
======================================

## SYNOPSIS

`unrot` FILE

or

STDIN | `unrot`

## DESCRIPTION

unrot decodes rotational ciphers

## ARGUMENTS

*FILE*
  Path to a file on disk.

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
