package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "clother"
	app.Usage = "Create template directory for Competitive Programming."
	app.Version = "1.0.0"

	app.Action = clother
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func clother(context *cli.Context) error {

	if context.Args().Get(0) == "" || context.Args().Get(1) == "" {
		fmt.Printf("Usage\n      clother [Directory Name] [Language Name]\ne.g\n      clother ABC024 python\n")
		return nil
	}

	var ext = ""
	switch strings.ToLower(context.Args().Get(1)) {
	case "python":
		ext = ".py"
	case "rust":
		ext = ".rs"
	case "c++":
		ext = ".cpp"
	case "nim":
		ext = ".nim"
	default:
		fmt.Println("please choose one of the following.")
		fmt.Println("-> Python, Rust, C++, Nim")
		return nil
	}

	if err := os.Mkdir(context.Args().Get(0), 0777); err != nil {
		fmt.Println(err)
		return nil
	}

	homePath, err := homedir.Dir()
	if err != nil {
		return err
	}

	tempPath := homePath + "/.clother/" + "template" + ext
	tempStr := ""
	if _, err := os.Stat(tempPath); err == nil {
		tmp, err := ioutil.ReadFile(tempPath)
		if err == nil {
			tempStr = string(tmp)
		}
	}

	aText := context.Args().Get(0) + "/" + context.Args().Get(0) + "_A" + ext
	bText := context.Args().Get(0) + "/" + context.Args().Get(0) + "_B" + ext
	cText := context.Args().Get(0) + "/" + context.Args().Get(0) + "_C" + ext
	dText := context.Args().Get(0) + "/" + context.Args().Get(0) + "_D" + ext
	eText := context.Args().Get(0) + "/" + context.Args().Get(0) + "_E" + ext
	fText := context.Args().Get(0) + "/" + context.Args().Get(0) + "_F" + ext

	files := []string{aText, bText, cText, dText, eText, fText}

	for _, i := range files {
		_, err := os.Create(i)
		if err != nil {
			return err
		}
		insertErr := insertTemplate(i, tempStr)
		if insertErr != nil {
			return err
		}
	}
	return nil
}

func insertTemplate(f, t string) error {
	file, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, ferr := fmt.Fprint(file, t)
	if ferr != nil {
		return ferr
	}
	return nil
}
