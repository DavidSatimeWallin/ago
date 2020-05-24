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
		Style       string `fig:"style" default:"body{max-width:650px;margin:40px auto;padding:0 10px;font:18px/1.5 -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji';color:#444}h1,h2,h3{line-height:1.2}@media (prefers-color-scheme: dark){body{color:white;background:#444}a:link{color:#5bf}a:visited{color:#ccf}}"`
	}

	// Hard is for hard coded config values
	Hard struct {
		PostsFolder   string
		SiteFolder    string
		EntriesFolder string
	}
)
