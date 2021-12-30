package main

import (
  "os"
  "strings"
  "path"

  "io/ioutil"

  "github.com/fatih/color"
  "github.com/manifoldco/promptui"
  "github.com/alecthomas/chroma/quick"
  "github.com/inancgumus/screen"
)

func CreateError(error string, isFatal bool) {
  color.Red(error)
  if isFatal {
    os.Exit(1)
  }
}

type Fe struct {
  filepath string
}

func (fe Fe) OpenDirectory(filepath string) {
  files, err := ioutil.ReadDir(filepath)
  if err != nil {
    CreateError(err.Error(), true)
  }

  filenames := []string{".."}
  for _, file := range files {
    name := file.Name()
    filenames = append(filenames, name)
  }
  prompt := promptui.Select{
    Label: "Select file",
    Items: filenames,
  }
  _, result, err := prompt.Run()
  if err != nil {
    CreateError(err.Error(), true)
  }
  fe.Open(path.Join(filepath, result), true)
}

func (fe Fe) OpenFile(filepath string) {
  content, err := ioutil.ReadFile(filepath)
  if err != nil {
    CreateError(err.Error(), true)
  }
  slices := strings.Split(path.Base(filepath), ".")
  extension := slices[len(slices) - 1]
  quick.Highlight(os.Stdout, string(content), extension, "terminal256", "monokai")
  fe.Open(path.Dir(filepath), false)
}

func (fe Fe) Open(filepath string, cls bool) {
  if cls {
    screen.Clear()
    screen.MoveTopLeft()
  }
  stat, err := os.Stat(filepath)
  if err != nil {
    CreateError(err.Error(), true)
  }
  if stat.IsDir() {
    fe.OpenDirectory(filepath)
  } else {
    fe.OpenFile(filepath)
  }
}

func main() {
  args := os.Args[1:]
  path, err := os.Getwd()
  if err != nil {
    CreateError(err.Error(), true)
  }
  if len(args) > 0 {
    path = args[0]
  }
  fe := Fe{path}
  fe.Open(path, false)
}