package domain

type WebsitesYaml struct {
	Websites []struct {
		URL  string `yaml:"url"`
		Type string `yaml:"type"`
		Ram  int    `yaml:"ram"`
		Shop string `yaml:"shop"`
	} `yaml:"websites"`
}
