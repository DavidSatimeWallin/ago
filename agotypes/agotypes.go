package agotypes

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
		Domain      string `fig:"domain" default:"ago.ofnir.xyz"`
		Protocol    string `fig:"protocol" default:"https"`
		Author      string `fig:"author" default:"Joane Doe"`
		Email       string `fig:"email" default:"joane.doe@ago.ofnir.xyz"`
		Title       string `fig:"title" default:"an Ago Blog!"`
		Description string `fig:"description" default:"This is an awesome Ago Blog!"`
		Tags        string `fig:"tags" default:"ago,blog,awesome"`
		Intro       string `fig:"intro" default:"You should have a small intro here to describe a little bit about yourself and the purpose of the blog"`
	}

	// Hard is for hard coded config values
	Hard struct {
		PostsFolder   string
		SiteFolder    string
		EntriesFolder string
	}
)
