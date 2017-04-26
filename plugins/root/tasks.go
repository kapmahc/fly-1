package root

import "github.com/astaxie/beego/toolbox"

// GenerateSitemapTask generate sitemap.xml.gz robots.txt googleXXX.html baidu_verify_XXX.html
func GenerateSitemapTask() error {
	return nil
}

func init() {
	toolbox.AddTask(
		"generate sitemap", toolbox.NewTask("sitemap", "0 0 3 * * *", GenerateSitemapTask),
	)
}
