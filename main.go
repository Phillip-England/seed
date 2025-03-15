package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Phillip-England/mood"
	"github.com/eiannone/keyboard"
)

func main() {

	app := mood.New()

	app.SetDefault(func(app *mood.Mood) error {
		fmt.Println("seed - generate skeleton projects with ease")
		fmt.Println("run 'seed plant' to get started")
		return nil
	})

	app.At("plant", func(app *mood.Mood) error {

		out := app.GetArgOr(1, ".")
		if out != "." {
			if len(out) >= 2 {
				if string(out[0:2]) != "./" {
					out = "./" + out
				}
			}
		}
		menuPosition := 0
		menuLimit := 2
		clear()
		printMenu(menuPosition)
		err := keyboard.Open()
		if err != nil {
			return err
		}
		defer keyboard.Close()

		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				fmt.Println(`error reading key:`, err)
				continue
			}
			switch key {
			case keyboard.KeyArrowDown:
				menuPosition++
				if menuPosition > menuLimit {
					menuPosition = 0
				}
				clear()
				printMenu(menuPosition)
			case keyboard.KeyArrowUp:
				menuPosition--
				if menuPosition < 0 {
					menuPosition = menuLimit
				}
				clear()
				printMenu(menuPosition)
			case keyboard.KeyCtrlZ:
				clear()
				return nil
			case keyboard.KeyCtrlX:
				clear()
				return nil
			case keyboard.KeyCtrlC:
				clear()
				return nil
			case keyboard.KeyEnter:
				clear()
				skeletonType, err := getSkeletonType(menuPosition)
				if err != nil {
					return err
				}
				switch skeletonType {
				case "server":
					err := GenerateGoServer(out)
					if err != nil {
						return err
					}
					return nil
				case "cli":
					err := GenerateGoCli(out)
					if err != nil {
						return err
					}
					return nil
				case "library":
					err := GenerateGoLibrary(out)
					if err != nil {
						return err
					}
					return nil
				}
				return nil
			}
		}
	})

	err := app.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

}

func GenerateGoLibrary(out string) error {
	if out != "." {
		err := makeDir(out)
		if err != nil {
			return fmt.Errorf(`cannot overwrite dir [%s] because it already exists`, out)
		}
	}
	skeleton := NewLibrarySkeleton(out)
	err := makeFile(skeleton.LibGoPath)
	if err != nil {
		return err
	}
	err = makeFile(skeleton.TestGoPath)
	if err != nil {
		return err
	}
	err = writeFile(skeleton.LibGoPath, `package lib

func Add(a, b int) int {
	return a + b
}
`)
	if err != nil {
		return err
	}
	err = writeFile(skeleton.TestGoPath, `package lib

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
`)
	if err != nil {
		return err
	}
	clear()
	if out == "." {
		fmt.Println("to initialize the project, run:\n")
		fmt.Println("1. go mod init github.com/github-name/repo-name\n")
		fmt.Println("thank you for using seed 🌱")
		return nil
	}
	fmt.Println("to initialize the project, run:\n")
	fmt.Println("1. cd " + strings.Replace(out, "./", "", 1))
	fmt.Println("2. go mod init github.com/github-name/repo-name\n")
	fmt.Println("thank you for using seed 🌱")
	return nil
}

func GenerateGoCli(out string) error {
	if out != "." {
		err := makeDir(out)
		if err != nil {
			return fmt.Errorf(`cannot overwrite dir [%s] because it already exists`, out)
		}
	}
	skeleton := NewCliSkeleton(out)
	err := makeFile(skeleton.MainGoPath)
	if err != nil {
		return err
	}
	err = writeFile(skeleton.MainGoPath, `package main

import (
	"fmt"
	"os"
	"github.com/Phillip-England/mood"
)

func main() {

	app := mood.New()

	app.SetDefault(NewDefaultCmd)
	app.At("help", NewHelpCmd)

	err := app.Run()
	if err != nil {
		panic(err)
	}

}

//======================
// DefaultCmd
//======================

type DefaultCmd struct{}

func NewDefaultCmd() (mood.Cmd, error) {
	return DefaultCmd{}, nil
}

func (cmd DefaultCmd) Execute(app *mood.App) error {
	fmt.Println("working..")
	return nil
}

//======================
// HelpCmd
//======================

type HelpCmd struct{}

func NewHelpCmd() (mood.Cmd, error) {
	return HelpCmd{}, nil
}

func (cmd HelpCmd) Execute(app *mood.App) error {
	fmt.Println("helping..")
	return nil
}`)
	if err != nil {
		return err
	}
	clear()
	if out == "." {
		fmt.Println("to install required packages run:\n")
		fmt.Println("1. go mod init github.com/github-name/repo-name")
		fmt.Println("2. go mod tidy\n")
		fmt.Println("thank you for using seed 🌱")
		return nil
	}
	fmt.Println("to install required packages run:\n")
	fmt.Println("1. cd " + strings.Replace(out, "./", "", 1))
	fmt.Println("2. go mod init github.com/github-name/repo-name")
	fmt.Println("3. go mod tidy\n")
	fmt.Println("thank you for using seed 🌱")
	return nil
}

func GenerateGoServer(out string) error {
	if out != "." {
		err := makeDir(out)
		if err != nil {
			return fmt.Errorf(`cannot overwrite dir [%s] because it already exists`, out)
		}
	}
	skeleton := NewServerSkeleton(out)
	err := makeDir(skeleton.TemplatePath)
	if err != nil {
		return err
	}
	err = makeDir(skeleton.StaticPath)
	if err != nil {
		return err
	}
	err = makeFile(skeleton.IndexHtmlPath)
	if err != nil {
		return err
	}
	err = makeFile(skeleton.IndexJsPath)
	if err != nil {
		return err
	}
	err = makeFile(skeleton.IndexCssPath)
	if err != nil {
		return err
	}
	err = makeFile(skeleton.MainGoPath)
	if err != nil {
		return err
	}
	err = writeFile(skeleton.IndexHtmlPath, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="/static/index.css"/>
  <title>{{ .Title }}</title>
</head>
<body>
  <h1>Hello, World!</h1>
  <script src="/static/index.js"></script>
</body>
</html>`)
	if err != nil {
		return err
	}
	err = writeFile(skeleton.IndexCssPath, `* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html {
  font-size: 16px;
}

body {
  line-height: 1.5;
  background-color: var(--color-white);
  color: var(--color-black);
}

:root {
  --color-black: #000;
  --color-white: #fff
}`)
	err = writeFile(skeleton.IndexJsPath, `console.log('connected!')`)
	if err != nil {
		return err
	}
	err = writeFile(skeleton.MainGoPath, `package main

import (
	"net/http"

	"github.com/Phillip-England/vii"
)

func main() {
	app := vii.NewApp()
	app.Use(vii.MwLogger, vii.MwTimeout(10))
	app.Static("./static")
	app.Favicon()
	err := app.Templates("./templates", nil)
	if err != nil {
		panic(err)
	}
	app.At("GET /", func(w http.ResponseWriter, r *http.Request) {
		vii.ExecuteTemplate(w, r, "index.html", map[string]interface{}{
			"Title": "Some Title",
		})
	})
	app.Serve("8080")
}
`)
	clear()
	if out == "." {
		fmt.Println("to install required packages run:\n")
		fmt.Println("1. go mod init github.com/github-name/repo-name")
		fmt.Println("2. go mod tidy\n")
		fmt.Println("thank you for using seed 🌱")
		return nil
	}
	fmt.Println("to install required packages run:\n")
	fmt.Println("1. cd " + strings.Replace(out, "./", "", 1))
	fmt.Println("2. go mod init github.com/github-name/repo-name")
	fmt.Println("3. go mod tidy\n")
	fmt.Println("thank you for using seed 🌱")
	return nil
}

type LibrarySkeleton struct {
	Root       string
	LibGoPath  string
	TestGoPath string
}

func NewLibrarySkeleton(root string) LibrarySkeleton {
	var fileName string
	if root == "." {
		fileName = "lib.go"
	} else {
		fileName = strings.Replace(root, "./", "", 1) + ".go"
	}
	var testName string
	if root == "." {
		testName = "lib_test.go"
	} else {
		testName = strings.Replace(root, "./", "", 1) + "_test.go"
	}
	return LibrarySkeleton{
		Root:       root,
		LibGoPath:  root + "/" + fileName,
		TestGoPath: root + "/" + testName,
	}
}

type CliSkeleton struct {
	Root       string
	MainGoPath string
}

func NewCliSkeleton(root string) CliSkeleton {
	return CliSkeleton{
		Root:       root,
		MainGoPath: root + "/main.go",
	}
}

type ServerSkeleton struct {
	Root          string
	TemplatePath  string
	StaticPath    string
	IndexHtmlPath string
	IndexCssPath  string
	IndexJsPath   string
	MainGoPath    string
}

func NewServerSkeleton(root string) ServerSkeleton {
	skeleton := ServerSkeleton{
		Root:          root,
		TemplatePath:  root + "/templates",
		IndexHtmlPath: root + "/templates/index.html",
		IndexCssPath:  root + "/static/index.css",
		StaticPath:    root + "/static",
		IndexJsPath:   root + "/static/index.js",
		MainGoPath:    root + "/main.go",
	}
	return skeleton
}

func printMenu(selectPosition int) {
	msg := `select a skeleton:
1. server
2. cli
3. library
	`
	lines := strings.Split(msg, "\n")
	newLines := []string{}
	for i, line := range lines {
		if i == 0 {
			newLines = append(newLines, line)
			continue
		}
		if i == selectPosition+1 {
			newLines = append(newLines, line+" *")
			continue
		}
		newLines = append(newLines, line)
	}
	fmt.Println(strings.Join(newLines, "\n"))
}

func clear() {
	fmt.Print("\033[2J\033[H")
}

func getSkeletonType(menuPosition int) (string, error) {
	switch menuPosition {
	case 0:
		return "server", nil
	case 1:
		return "cli", nil
	case 2:
		return "library", nil

	}
	return "", fmt.Errorf("provided invalid position [%d] to getSkeletonType", menuPosition)
}

func makeDir(path string) error {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return fmt.Errorf("directory '%s' already exists", path)
	}
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error checking directory: %w", err)
	}
	if err := os.Mkdir(path, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

func makeFile(path string) error {
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return fmt.Errorf("file '%s' already exists", path)
	}
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error checking file: %w", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	return nil
}

func writeFile(path, content string) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("file '%s' does not exist", path)
	}
	if err != nil {
		return fmt.Errorf("error checking file: %w", err)
	}
	if info.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a file", path)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
