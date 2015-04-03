# hashpipe - pipe iff the hash matches

`hashpipe` helps you venture into the unknown. It reads from stdin, checks the hash of the content, and outputs it IF AND ONLY IF it matches the provided hash checksum. This makes executing things a teensy bit safer, as it requires compromising more communication channels.

## Example

```sh
> echo "hello" | multihash
QmUJPTFZnR2CPGAzmfdYPghgrFtYFB6pf1BqMvqfiPDam8
> echo "hello" | hashpipe QmUJPTFZnR2CPGAzmfdYPghgrFtYFB6pf1BqMvqfiPDam8
hello
> echo "goodbye" | hashpipe QmUJPTFZnR2CPGAzmfdYPghgrFtYFB6pf1BqMvqfiPDam8
error: checksums did not match
```

## Use Case

If you tell people to pipe things directly into the shell... Don't do that. If you're going to _at least_ provide them a hash to ensure that man-in-the-middle attacks or compromised CDNs do not hurt your users.

```
curl http://you.shouldnt.be/doing/this | hashpipe Qmepk1VCHpjyCmWeh61vyDKsWfKymyrBQLcmUpXUdUd3yM | sh
```

## Hashes Supported

`hashpipe` uses [multihash](https://github.com/jbenet/multihash/), a self-describing hash function. It supports all the hashes in [go-multihash](https://github.com/jbenet/go-multihash/):

- sha1
- sha2-256
- sha2-512
- sha3-256
- sha3-512
- [add the one you need](https://github.com/jbenet/go-multihash/pulls)

## Install

- From source:

    ```
    go get https://github.com/jbenet/go-hashpipe
    ```

- Pre-built binaries: https://gobuilder.me/github.com/jbenet/go-hashpipe

## Usage

```
cat untrustedFile | hashpipe <expected-multihash-checksum> | trustedContext

hashpipe - boldly journey into the unknown.

It reads from stdin, checks the hash of the content, and outputs it IF AND
ONLY IF it matches the provided hash checksum. This makes executing things
a bit safer, as it requires compromising more communication channels. On
error, hashpipe returns a non-zero error code, failing pipelines.

OPTIONS

```