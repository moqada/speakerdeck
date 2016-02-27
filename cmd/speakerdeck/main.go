package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/moqada/speakerdeck"
)

func download(c *cli.Context) error {
	args := c.Args()
	if len(args) < 2 {
		return fmt.Errorf("few args: speakerdeck download [url] [outputdir]")
	}
	url := args[0]
	output := args[1]
	fi, err := os.Stat(output)
	if err != nil {
		return fmt.Errorf("outputDir does not exists: %s", output)
	}
	if !fi.IsDir() {
		return fmt.Errorf("outputDir is not directory: %s", output)
	}
	slide, err := speakerdeck.GetSlide(url)
	if err != nil {
		return err
	}
	pdf, err := slide.DownloadPDF()
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(path.Join(output, slide.Slug+".pdf"), pdf, 0644); err != nil {
		return err
	}
	return nil
}

func info(c *cli.Context) error {
	args := c.Args()
	if len(args) < 1 {
		return fmt.Errorf("few args: speakerdeck info [url]")
	}
	url := args[0]
	slide, err := speakerdeck.GetSlide(url)
	if err != nil {
		return err
	}
	info, err := json.Marshal(slide)
	if err != nil {
		return err
	}
	fmt.Println(string(info))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "speakerdeck"
	app.Usage = "Fetch Slide on Speakerdeck"
	app.Version = "0.1.0"
	app.Author = "moqada <moqada@gmail.com>"
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "Download PDF file to output dir",
			Action: func(c *cli.Context) {
				if err := download(c); err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err.Error())
					os.Exit(1)
				}
			},
		},
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "Fetch Slide info to JSON",
			Action: func(c *cli.Context) {
				if err := info(c); err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err.Error())
					os.Exit(1)
				}
			},
		},
	}
	app.Run(os.Args)
}
