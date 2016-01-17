// Extract Text Content from a HTML string
package article

import (
	"bytes"
	"github.com/JesusIslam/tldr"
	"github.com/PuerkitoBio/goquery"
	"github.com/endeveit/guesslanguage"
	"github.com/mauidude/go-readability"
	"strings"
)

type Article struct {
	Language    string
	Description string
	Fulltext    string
}

func Parse(s string) Article {
	a := Article{}
	a.Fulltext = fulltextFromBody(s)
	a.Description = summaryFromFulltext(a.Fulltext)
	a.Language = detectLanguage(a.Fulltext)
	return a
}

func fulltextFromBody(str string) string {
	doc, err := readability.NewDocument(str)
	if err != nil {
		return ""
	}
	text := doc.Content()
	dom := docFromString(text)
	matches := dom.Find("p")
	if matches.Length() > 0 {
		paragraphs := make([]string, matches.Length())
		matches.Each(func(i int, s *goquery.Selection) {
			txt := strings.Trim(s.Text(), " ")
			if len(txt) > 60 {
				paragraphs[i] = txt
			} else {
				paragraphs[i] = ""
			}
		})
		return strings.Join(paragraphs, "\n\n")
	}
	return ""
}

func summaryFromFulltext(str string) string {
	bag := tldr.New()
	result, err := bag.Summarize(str, 3)
	if err != nil {
		return ""
	}
	return result
}

func detectLanguage(str string) string {
	lang, err := guesslanguage.Guess(str)
	if err != nil {
		lang = "en"
	}
	return lang
}

func docFromString(str string) goquery.Document {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(str)
	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		return goquery.Document{}
	}
	return *doc
}
