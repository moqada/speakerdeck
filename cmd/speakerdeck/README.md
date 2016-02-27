# speakerdeck

:ship: Command for Fetch info and Download Slide from Speaker Deck.

## Install

```
$ go get github.com/moqada/speakerdeck/cmd/speakerdeck
```

## Usage

```
NAME:
   speakerdeck - Fetch Slide on Speaker Deck

USAGE:
   speakerdeck [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR(S):
   moqada <moqada@gmail.com>

COMMANDS:
   download, d  Download PDF file to output dir
   info, i      Fetch Slide info to JSON
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
```

### Download

```
$ speakerdeck download https://speakerdeck.com/achiku/pycon-jp-2014-python-plus-hive-on-aws-emrdepin-zhe-falseroguji-ji ./
$ ls
pycon-jp-2014-python-plus-hive-on-aws-emrdepin-zhe-falseroguji-ji.pdf
```

### Info

```
$ speakerdeck info https://speakerdeck.com/achiku/pycon-jp-2014-python-plus-hive-on-aws-emrdepin-zhe-falseroguji-ji | jq .
{
  "slug": "pycon-jp-2014-python-plus-hive-on-aws-emrdepin-zhe-falseroguji-ji",
  "title": "PyCon JP 2014 Python + Hive on AWS EMRで貧者のログ集計",
  "description": "",
  "downloadUrl": "https://speakerd.s3.amazonaws.com/presentations/d93ef1101e17013278dc06e915146373/pycon-2014-aws-emr.pdf",
  "user": {
    "username": "achiku",
    "displayName": "Akira Chiku",
    "url": "https://speakerdeck.com/achiku"
  },
  "stars": 0,
  "category": {
    "name": "Technology",
    "slug": "technology",
    "url": "https://speakerdeck.com/technology"
  },
  "published": "2014-09-14T00:00:00Z",
  "url": "https://speakerdeck.com/achiku/pycon-jp-2014-python-plus-hive-on-aws-emrdepin-zhe-falseroguji-ji"
}
```
