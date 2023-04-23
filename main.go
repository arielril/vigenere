package main

import (
	"errors"
	"os"
	"strings"

	"github.com/arielril/vigenere/internal/vigenere"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

type Options struct {
	Encode      bool
	Decode      bool
	Crack       bool
	CipherKey   string
	Message     string
	MessagePath string

	Verbose bool
	Silent  bool
}

var options *Options = &Options{}

func init() {
	set := goflags.NewFlagSet()
	set.SetDescription("Vigenere cipher")

	set.BoolVarP(&options.Decode, "decode", "d", false, "decode")
	set.BoolVarP(&options.Encode, "encode", "e", false, "encode")
	set.BoolVarP(&options.Crack, "crack", "c", false, "crack an encrypted message")
	set.StringVarP(&options.CipherKey, "key", "k", "", "cipher key")
	set.StringVarP(&options.Message, "message", "m", "", "message or encrypted message")
	set.StringVarP(&options.MessagePath, "message-path", "mp", "", "message path")

	set.BoolVarP(&options.Verbose, "verbose", "v", false, "verbose output")
	set.BoolVarP(&options.Silent, "silent", "s", false, "silent output")

	if err := set.Parse(); err != nil {
		gologger.Fatal().Msgf("could not parse program flags: %s\n", err)
	}
}

func main() {
	configureOutput(options)
	gologger.Info().Msg("Vigenere cipher is amazing :P")

	err := validateOptions(options)
	if err != nil {
		gologger.Fatal().Msgf("could not run vigenere: %s\n", err)
	}

	readMessage(options)

	var vigenereResult string
	switch {
	case options.Encode:
		vigenereResult, err = vigenere.Encode(options.Message, options.CipherKey)
	case options.Decode:
		vigenereResult, err = vigenere.Decode(options.Message, options.CipherKey)
	case options.Crack:
		vigenereResult, err = vigenere.Crack(options.Message)
	default:
		gologger.Warning().Msg("no runner option selected")
	}

	if err != nil {
		gologger.Fatal().Msgf("could not run vigenere: %s\n", err)
	}

	gologger.Silent().Msgf("Vigenere result: %s\n", vigenereResult)
}

func readMessage(opts *Options) {
	if opts.MessagePath != "" {
		data, err := os.ReadFile(opts.MessagePath)
		if err != nil {
			gologger.Fatal().Msgf("could not read message file: %s\n", err)
		}

		opts.Message = string(data)
	}

	/* deal only with lower case */
	opts.Message = strings.TrimSpace(strings.ToLower(opts.Message))
	opts.CipherKey = strings.TrimSpace(strings.ToLower(opts.CipherKey))
	/* deal only with lower case */
}

func validateOptions(opts *Options) error {
	switch {
	case !opts.Encode && !opts.Decode && !opts.Crack:
		return errors.New("please, select one option: `encode` or `decode`")
	case opts.Encode && opts.CipherKey == "":
		return errors.New("to encode a message you must provide a cipher key")
	case opts.Decode && opts.CipherKey == "":
		return errors.New("to decode a message you must provide a cipher key")
	case opts.Message == "" && opts.MessagePath == "":
		return errors.New("no message was provided")
	}

	return nil
}

func configureOutput(opts *Options) {
	if opts.Verbose && opts.Silent {
		gologger.Fatal().Msg("silent and verbose output can not set at the same time")
	}

	if opts.Verbose {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	}

	if opts.Silent {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	}
}
