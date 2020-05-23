package types

type (
	// Placeholder is for replacing tags in html code with dynamic valyes
	Placeholder struct {
		Tag   string
		Value string
	}

	// Placeholders are just the slice of Placeholder
	Placeholders []Placeholder

	// Config contains everything needed to run the blog
	Config struct {
		Domain        string `fig:"domain" default:"ago.ofnir.xyz"`
		Author        string `fig:"author" default:"Joane Doe"`
		Email         string `fig:"email" default:"joane.doe@ago.ofnir.xyz"`
		WebsiteName   string `fig:"website_name" default:"an Ago Blog!"`
		GithubAccount string `fig:"github_account" default:"dvwallin"`
		Title         string `fig:"title" default:"an Ago Blog!"`
		Description   string `fig:"description" default:"This is an awesome Ago Blog!"`
		Keywords      string `fig:"keywords" default:"ago,blog,awesome"`
	}
)
