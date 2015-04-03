`hashpipe` helps you venture into the unknown. It reads from stdin, checks the hash of the content, and outputs it IF AND ONLY IF it matches the provided hash checksum. This makes executing things a teensy bit safer, as it requires compromising more communication channels.

`<>` with `<3` by [@juanbenet](https://twitter.com/juanbenet)

<iframe src="https://ghbtns.com/github-btn.html?user=twbs&repo=bootstrap&type=star&count=true&size=large" frameborder="0" scrolling="0" width="160px" height="30px"></iframe>

## Example

```sh
> echo "hello" | multihash
QmUJPTFZnR2CPGAzmfdYPghgrFtYFB6pf1BqMvqfiPDam8
> echo "hello" | hashpipe QmUJPTFZnR2CPGAzmfdYPghgrFtYFB6pf1BqMvqfiPDam8
hello
> echo "goodbye" | hashpipe QmUJPTFZnR2CPGAzmfdYPghgrFtYFB6pf1BqMvqfiPDam8
error: multihash checksums did not match
```

## Use Case

If you tell people to pipe things directly into the shell... Don't do that. If you're going to, _at least_ provide them a hash to ensure that man-in-the-middle attacks or compromised CDNs do not hurt your users.

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
    go get https://github.com/jbenet/hashpipe
    ```

- Pre-built binaries: https://gobuilder.me/github.com/jbenet/hashpipe

## Usage

```
> hashpipe -h
usage: hashpipe [MULTIHASH] <[FILE] >[FILE]

    cat untrustedFile | hashpipe <expected-checksum> | trustedContext

hashpipe - boldly journey into the unknown.

It reads from stdin, checks the hash of the content, and outputs it IF AND
ONLY IF it matches the provided hash checksum. This makes executing things
a bit safer, as it requires compromising more communication channels. On
error, hashpipe returns a non-zero error code, failing pipelines.

OPTIONS
  -a="sha2-256": one of: sha1, sha2-256, sha2-512, sha3 (shorthand)
  -algorithm="sha2-256": one of: sha1, sha2-256, sha2-512, sha3
  -e="base58": one of: raw, hex, base58, base64 (shorthand)
  -encoding="base58": one of: raw, hex, base58, base64
  -l=-1: checksums length in bits (truncate). -1 is default (shorthand)
  -length=-1: checksums length in bits (truncate). -1 is default
  -q=false: quiet output (no newline on checksum, no error text) (shorthand)
  -quiet=false: quiet output (no newline on checksum, no error text)
```
