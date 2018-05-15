package colors

// Solarized colors
const (
	// Base03 is dark a background tone.
	Base03 = "#002b36"
	// Base02 is a dark background tone.
	Base02 = "#073642"
	// Base01 is content tone.
	Base01 = "#586e75"
	// Base00 is content tone.
	Base00 = "#657b83"
	// Base0 is content tone.
	Base0 = "#839496"
	// Base1 is content tone.
	Base1 = "#93a1a1"
	// Base2 is a light background tone.
	Base2 = "#eee8d5"
	// Base3 is a light background tone.
	Base3 = "#fdf6e3"

	Yellow  = "#b58900"
	Orange  = "#cb4b16"
	Red     = "#dc322f"
	Magenta = "#d33682"
	Violet  = "#6c71c4"
	Blue    = "#268bd2"
	Cyan    = "#2aa198"
	Green   = "#859900"
)

const (
	// Dark palette.
	Dark = "dark"
	// Light palette.
	Light = "light"
)

// Palette type
type Palette map[string]string

func baseDarkPalette() Palette {
	base := Palette{
		"EmphasizedContent":    Base1,
		"BodyText":             Base0,
		"Comment":              Base01,
		"SecondaryContent":     Base01,
		"BackgroundHighlights": Base02,
		"Background":           Base03,
	}

	full := make(map[string]string)
	for k, v := range base {
		full[k] = v
	}
	for k, v := range base {
		full["Inverted"+k], _ = Invert(v)
	}

	return full
}

func addColors(palette Palette) Palette {
	colors := Palette{
		"Yellow":  Yellow,
		"Orange":  Orange,
		"Red":     Red,
		"Magenta": Magenta,
		"Violet":  Violet,
		"Blue":    Blue,
		"Green":   Green,
		"Cyan":    Cyan,

		"Base03": Base03,
		"Base02": Base02,
		"Base01": Base01,
		"Base00": Base00,
		"Base0":  Base0,
		"Base1":  Base1,
		"Base2":  Base2,
		"Base3":  Base3,

		"BrBlack":   Base03,
		"Black":     Base02,
		"BrGreen":   Base01,
		"BrYellow":  Base00,
		"BrBlue":    Base0,
		"BrCyan":    Base1,
		"White":     Base2,
		"BrWhite":   Base3,
		"BrRed":     Orange,
		"BrMagenta": Violet,
	}

	for k, v := range palette {
		colors[k] = v
	}

	return colors
}

// DarkPalette is the solarized dark palette.
func DarkPalette() Palette {
	return addColors(Palette{
		"EmphasizedContent":    Base1,
		"BodyText":             Base0,
		"Comment":              Base01,
		"SecondaryContent":     Base01,
		"BackgroundHighlights": Base02,
		"Background":           Base03,

		"InvertedEmphasizedContent":    Base01,
		"InvertedBodyText":             Base00,
		"InvertedComment":              Base1,
		"InvertedSecondaryContent":     Base1,
		"InvertedBackgroundHighlights": Base2,
		"InvertedBackground":           Base3,
	})
}

// LightPalette is the solarized light palette.
func LightPalette() Palette {
	return addColors(InvertPalette(DarkPalette()))
}

// Invert inverts color between light and dark color scheme.
func Invert(color string) (string, bool) {
	switch color {
	case Base03:
		return Base3, true
	case Base02:
		return Base2, true
	case Base01:
		return Base1, true
	case Base00:
		return Base0, true
	case Base0:
		return Base00, true
	case Base1:
		return Base01, true
	case Base2:
		return Base02, true
	case Base3:
		return Base03, true
	}

	return color, false
}

// InvertPalette the given palette.
func InvertPalette(palette Palette) Palette {
	invertedPalette := Palette{}

	for key, value := range palette {
		color, _ := Invert(value)
		invertedPalette[key] = color
	}

	return invertedPalette
}
