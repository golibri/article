[![Build Status](https://travis-ci.org/golibri/article.svg?branch=master)](https://travis-ci.org/golibri/article)
[![GoDoc](https://godoc.org/github.com/golibri/article?status.svg)](https://godoc.org/github.com/golibri/article)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/article

Get Text Content from HTML. An **Article** gets constructed through processing a
HTML page. The relevant content is stripped from all the useless junk and markup
and stored as `Fulltext`. Works best with blog posts or news articles, but even
a tweet should suffice.

Given an HTML string of any "content"-site, this module:

- `Fulltext`: extracts the relevant paragraphs as plain text
- `Language`: determines the language of the text (fallback: `en`)
- `Description`: summarizes the text into a short snippet of upto 3 sentences

# installation
`go get -u github.com/golibri/article`

# usage
````go
import "github.com/golibri/article"

func main() {
    // ...get HTML string somewhere, e.g.: with golibri/fetch
    a := article.Parse("website-html-string")
    // a is an Article object, see below
}
````

# data fields
````go
type Article struct {
    Language    string
    Description string
    Fulltext    string
}
````

# license
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
