package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	mh "github.com/jbenet/go-multihash"
	mhopts "github.com/jbenet/go-multihash/opts"
)

var usage = `usage: %s [MULTIHASH] <[FILE] >[FILE]

    cat untrustedFile | hashpipe <expected-checksum> | trustedContext

hashpipe - boldly journey into the unknown.

It reads from stdin, checks the hash of the content, and outputs it IF AND
ONLY IF it matches the provided hash checksum. This makes executing things
a bit safer, as it requires compromising more communication channels. On
error, hashpipe returns a non-zero error code, failing pipelines.

OPTIONS
`

// flags
var opts *mhopts.Options
var quiet bool

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}

	opts = mhopts.SetupFlags(flag.CommandLine)

	quietStr := "quiet output (no newline on checksum, no error text)"
	flag.BoolVar(&quiet, "quiet", false, quietStr)
	flag.BoolVar(&quiet, "q", false, quietStr+" (shorthand)")
}

func parseFlags(o *mhopts.Options) error {
	flag.Parse()
	if err := o.ParseError(); err != nil {
		return err
	}
	return nil
}

func getInput(o *mhopts.Options) (mh.Multihash, error) {
	args := flag.Args()
	if len(args) < 1 {
		return nil, fmt.Errorf("multihash is a required argument")
	}
	raw := args[0]

	h, err := mhopts.Decode(o.Encoding, raw)
	if err != nil {
		return nil, fmt.Errorf("fail to decode multihash '%s': %s", raw, err)
	}
	return h, nil
}

func run() error {
	if err := parseFlags(opts); err != nil {
		return err
	}

	// parse the given checksum
	expect, err := getInput(opts)
	if err != nil {
		return err
	}

	// have to read all of it before we output any of it :/
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	// calculate the checksum of the input
	actual, err := mh.Sum(input, opts.AlgorithmCode, opts.Length)
	if err != nil {
		return err
	}

	// ensure checksums match
	if !bytes.Equal(expect, actual) {
		return mhopts.ErrMatch
	}

	// ok, checksums matched, write it out
	if !quiet {
		_, err = os.Stdout.Write(input)
	}
	return err
}

func main() {
	if err := run(); err != nil {
		if !quiet {
			fmt.Fprintf(os.Stderr, "error: %s (-q for no output)\n", err)
		}
		os.Exit(1)
	}
}
