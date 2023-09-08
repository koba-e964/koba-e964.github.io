package main

type NounEntry struct {
	Noun        Noun
	Filename    string
	Translation string
}

type NounConfig struct {
	Title          string `toml:"title"`
	Synopsis       string `toml:"synopsis"`
	Genitive       string `toml:"genitive"`
	PluralGenitive string `toml:"plural_genitive"`
	Declension     int    `toml:"declension"`
	Gender         string `toml:"gender"`
	Translation    string `toml:"translation"`
}

func (conf NounConfig) ToNounEntry() NounEntry {
	var noun Noun
	var gender Gender
	if conf.Gender == "masculine" {
		gender = Masculine
	} else if conf.Gender == "feminine" {
		gender = Feminine
	} else if conf.Gender == "neuter" {
		gender = Neuter
	} else {
		panic(conf.Gender)
	}
	if conf.Declension != 3 {
		noun = CreateNoun(gender, conf.Synopsis, conf.Synopsis, 2)
	} else {
		noun = CreateThirdDeclensionNoun(gender, conf.Synopsis, conf.Genitive, conf.PluralGenitive)
	}
	return NounEntry{
		Noun:        noun,
		Filename:    conf.Title,
		Translation: conf.Translation,
	}
}
