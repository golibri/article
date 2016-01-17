[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/article
Get Text Content from HTML.

Given an HTML string of any "content"-site, this module:

- extracts the relevant paragraphs as plain text
- determines the language of the text
- summarizes the text into a short snippet

# installation
`go get github.com/golibri/article`

# dependencies
`github.com/PuerkitoBio/goquery`
`github.com/endeveit/guesslanguage`
`github.com/JesusIslam/tldr"`
`github.com/mauidude/go-readability`

# usage
````go
a := article.Parse("website-html-string")
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
