# speakerdeck

:ship: Fetch info and Download Slide from Speaker Deck for Go.

> Unofficial and Implemented by dirty scraping...

## Install

```
$ go get github.com/moqada/speakerdeck
```

## Usage

```go
import (
	"fmt"
	"io/ioutil"

	"github.com/moqada/speakerdeck"
)

// Fetch Slide
slide, _ := speakerdeck.GetSlide("https://speakerdeck.com/achiku/pycon-jp-2014-python-plus-hive-on-aws-emrdepin-zhe-falseroguji-ji")
fmt.Println(slide.Title, slide.User.Username, slide.Category.Name)
// PyCon JP 2014 Python + Hive on AWS EMRで貧者のログ集計 achiku Technology

// Download pdf file as byte[]
pdf, _ := slide.DownloadPDF()

// Save pdf file to current directory
ioutil.WriteFile("./" + slide.Slug + ".pdf", pdf, 0644)
```

The documentation is on [GoDoc](https://godoc.org/github.com/moqada/speakerdeck)


## CLI

- [cmd/speakerdeck](https://github.com/moqada/speakerdeck/tree/master/cmd/speakerdeck)


[godoc-url]: https://godoc.org/github.com/moqada/speakerdeck
[godoc-image]: https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square
[travis-url]: https://travis-ci.org/moqada/speakerdeck
[travis-image]: https://img.shields.io/travis/moqada/speakerdeck.svg?style=flat-square
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/github/license/moqada/speakerdeck.svg?style=flat-square
