package article

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestArticleParser(t *testing.T) {
	fileBytes, _ := ioutil.ReadFile("./data/bbc.html")
	a := Parse(string(fileBytes))
	fmt.Println(a)
}
