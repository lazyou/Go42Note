module ModulesTest

go 1.12

replace (
	golang.org/x/net v0.0.0-20180218175443-cbe0f9307d01 => github.com/golang/net v0.0.0-20180218175443-cbe0f9307d01
	golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a
)

require (
	github.com/PuerkitoBio/goquery v1.5.0 // indirect
	github.com/antchfx/htmlquery v1.0.0 // indirect
	github.com/antchfx/xmlquery v1.0.0 // indirect
	github.com/antchfx/xpath v0.0.0-20190129040759-c8489ed3251e // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/temoto/robotstxt v0.0.0-20180810133444-97ee4a9ee6ea // indirect
)
