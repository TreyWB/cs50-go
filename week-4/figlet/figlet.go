package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mbndr/figlet4go"
)

var errFont = fmt.Errorf("invalid font")
var errColor = fmt.Errorf("invalid color")
var errFontFile = fmt.Errorf("cannot access font file")

func main() {
	fontName, fontPath, colors, err := getArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	var input string
	fmt.Print("Input: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()

	options.FontColor = colors
	if fontName != "" {
		options.FontName = fontName
		file_err := ascii.LoadFont(fontPath)
		if file_err != nil {
			fmt.Println(file_err)
			return
		}
	}

	output, err := ascii.RenderOpts(input, options)
	if err != nil {
		fmt.Println("Render Error:", err)
		return
	}

	fmt.Println("Output:\n", output)
}

func getArgs() (string, string, []figlet4go.Color, error) {
	var fontName string
	var fontPath string
	var err error
	var colors []figlet4go.Color

	for index, param := range os.Args {
		if index == 0 {
			continue
		}

		switch param {
		case "-f":
			result := checkFont(os.Args[index+1])
			if !result {
				return "", "", []figlet4go.Color{}, errFont
			}
			fontName, fontPath, err = getFontFile(os.Args[index+1])
			if err != nil {
				return "", "", []figlet4go.Color{}, errors.Join(errFontFile, fmt.Errorf("cannot access program directory"))
			}
		case "-c1":
			figletColor, err := checkColor(os.Args[index+1])
			if err != nil {
				return "", "", []figlet4go.Color{}, errors.Join(errColor, fmt.Errorf("invalid primary color"))
			}
			colors = append(colors, figletColor)
		case "-c2":
			figletColor, err := checkColor(os.Args[index+1])
			if err != nil {
				return "", "", []figlet4go.Color{}, errors.Join(errColor, fmt.Errorf("invalid secondary color"))
			}
			colors = append(colors, figletColor)
		case "-h":
			printHelp()
			os.Exit(0)
		case "--help":
			printHelp()
			os.Exit(0)
		default:
			continue
		}
	}

	return fontName, fontPath, colors, nil
}

func checkFont(input string) bool {
	validFonts := []string{
		"contrast",
		"epic",
		"slant",
		"speed",
		"doom",
	}

	for _, font := range validFonts {
		if font == input {
			return true
		}
	}

	return false
}

func checkColor(input string) (figlet4go.Color, error) {
	validColors := map[string]figlet4go.Color{
		"red":     figlet4go.ColorRed,
		"green":   figlet4go.ColorGreen,
		"yellow":  figlet4go.ColorYellow,
		"blue":    figlet4go.ColorBlue,
		"magenta": figlet4go.ColorMagenta,
		"cyan":    figlet4go.ColorCyan,
		"white":   figlet4go.ColorWhite,
	}

	figletColor, ok := validColors[input]
	if !ok {
		return figlet4go.ColorBlack, errColor
	}

	return figletColor, nil
}

func getFontFile(name string) (string, string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", "", errFontFile
	}

	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "fonts", name+".flf")

	return name, path, nil
}

func printHelp() {
	var help = `
Usage: figlet [OPTIONS]

OPTIONS:
	-f font	Use the specified font
	-c1 color	Use the specified primary color
	-c2 color	Use the specified secondary color

FONTS:
	contrast
	epic
	slant
	speed
	doom

COLORS:
	red
	green
	yellow
	blue
	magenta
	cyan
	white
`

	fmt.Println(help)
}
