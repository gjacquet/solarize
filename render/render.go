package render

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/gjacquet/solarize/colors"
)

// Renderer renders template based on a color palette.
type Renderer interface {
	Render(colors.Palette, io.Writer) error
	RenderToFile(colors.Palette, string) error
}

type renderer struct {
	template *template.Template
}

// FromFile create a renderer from a template at a given path.
func FromFile(path string) (Renderer, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}

	return New(tmpl), nil
}

// New creates a new template.
func New(tmpl *template.Template) Renderer {
	return &renderer{
		template: tmpl,
	}
}

func (r *renderer) Render(palette colors.Palette, output io.Writer) error {
	return r.template.Execute(output, palette)
}

func (r *renderer) RenderToFile(palette colors.Palette, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	writer := bufio.NewWriter(file)

	err = r.Render(palette, writer)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
