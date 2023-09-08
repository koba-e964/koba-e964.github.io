package main

type Config struct {
	Nouns []NounConfig `toml:"nouns"`
}
