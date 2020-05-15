package gui

// Bitmap font loader (.fnt)
// Format: http://www.angelcode.com/products/bmfont/doc/file_format.html
// Exporter: https://github.com/libgdx/libgdx/wiki/Hiero

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/maxfish/GoNativeUI-Core/utils"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// BitmapGlyph holds information about a single glyph/character
type BitmapGlyph struct {
	Id             int
	X              int
	Y              int
	Width          int
	Height         int
	OffsetX        int
	OffsetY        int
	AdvanceX       int
	PageIndex      int
	TextureChannel int
	kerning        map[int32]int
}

func (c *BitmapGlyph) Kerning(previousGlyph int32) int {
	k, ok := c.kerning[previousGlyph]
	if !ok {
		return 0
	}
	return k
}

// BitmapFont holds information about the Bitmap font
type BitmapFont struct {
	path       string
	pageNames  []string
	pageImages []image.Image
	glyphs     map[int32]*BitmapGlyph

	// Info
	face          string
	size          int
	bold          bool
	italic        bool
	charset       string
	unicode       bool
	stretchH      int
	smooth        bool
	superSampling int
	padding       [4]int // top, right, bottom, left
	spacing       [2]int // X, Y
	// Common
	lineHeight int
	base       int
	pageWidth  int
	pageHeight int
	packed     bool
	numPages   int
	// Chars
	charactersCount int
}

// NewFontFromFile parse the font data out of a file
func NewFontFromFile(fileName string) *BitmapFont {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("Error loading font -> %s", err)
	}

	f := &BitmapFont{}
	f.path, _ = path.Split(fileName)
	f.pageNames = make([]string, 0)
	f.pageImages = make([]image.Image, 0)
	f.glyphs = make(map[int32]*BitmapGlyph)
	f.loadDefinitionFromString(string(fileContent))
	f.loadPageImagesFiles()

	return f
}

func NewFontFromData(definition string, base64Images []string) *BitmapFont {
	f := &BitmapFont{}
	f.path = "n/a"
	f.pageNames = make([]string, 0)
	f.glyphs = make(map[int32]*BitmapGlyph)
	f.loadDefinitionFromString(definition)

	for _, base64String := range base64Images {
		data, err := base64.StdEncoding.DecodeString(base64String)
		if err != nil {
			log.Panicf("Cannot decode b64 -> '%s'\n", err)
		}

		r := bytes.NewReader(data)
		decodedImage, err := png.Decode(r)
		if err != nil {
			panic("Bad png")
		}

		f.pageImages = append(f.pageImages, decodedImage)
	}

	return f
}

func (f *BitmapFont) loadDefinitionFromString(definition string) {
	lines := strings.Split(definition, "\n")
	for _, line := range lines {
		section, keyValues := f.tokenizeLine(line)
		switch section {
		case "info":
			f.parseInfoSection(keyValues)
		case "common":
			f.parseCommonSection(keyValues)
		case "page":
			f.parsePageSection(keyValues)
		case "char":
			f.parseCharSection(keyValues)
		case "kerning":
			f.parseKerningSection(keyValues)
		}
	}
}

func (f *BitmapFont) loadPageImagesFiles() {
	for _, imageFileName := range f.pageNames {
		file, err := os.Open(path.Join(f.path, imageFileName))
		if err != nil {
			log.Panicf("Error loading image '%s' -> '%s'\n", imageFileName, err)
		}

		decodedImage, format, err := image.Decode(file)
		if err != nil {
			log.Panicf("Cannot decode image:'%s' format:'%s' -> '%s'\n", imageFileName, format, err)
		}
		file.Close()
		f.pageImages = append(f.pageImages, decodedImage)
	}
}

func (f *BitmapFont) FaceName() string {
	return f.face
}

func (f *BitmapFont) PageImages() []image.Image {
	return f.pageImages
}

func (f *BitmapFont) PageSize() utils.Size {
	return utils.Size{f.pageWidth, f.pageHeight}
}

func (f *BitmapFont) Size() int {
	return f.size
}

// TODO: this doesn't support multiline text
func (f *BitmapFont) TextSize(size int, text string) utils.Size {
	scale := float32(size) / float32(f.Size())
	h := f.lineHeight
	var previousGlyph int32
	var w = 0
	for _, glyph := range text {
		bmc := f.glyphs[glyph]
		w += bmc.AdvanceX + bmc.Kerning(previousGlyph)
		previousGlyph = glyph
	}
	return utils.Size{int(float32(w) * scale), int(float32(h) * scale)}
}

func (f *BitmapFont) Glyph(index int32) BitmapGlyph {
	char, ok := f.glyphs[index]
	if !ok {
		// Glyph not available in the current font, returns 'space'
		return *f.glyphs[32]
	}
	return *char
}

func (f *BitmapFont) parseInfoSection(keyValues map[string]string) {
	f.face = keyValues["face"]
	f.size, _ = strconv.Atoi(keyValues["size"])
	f.bold, _ = strconv.ParseBool(keyValues["bold"])
	f.italic, _ = strconv.ParseBool(keyValues["italic"])
	f.unicode, _ = strconv.ParseBool(keyValues["unicode"])
	f.stretchH, _ = strconv.Atoi(keyValues["stretchH"])
	f.smooth, _ = strconv.ParseBool(keyValues["smooth"])
	f.superSampling, _ = strconv.Atoi(keyValues["aa"])

	// Padding
	paddingStrings := strings.Split(keyValues["padding"], ",")
	for i := 0; i < 4; i++ {
		f.padding[i], _ = strconv.Atoi(paddingStrings[i])
	}
	// Spacing
	spacingStrings := strings.Split(keyValues["spacing"], ",")
	for i := 0; i < 2; i++ {
		f.spacing[i], _ = strconv.Atoi(spacingStrings[i])
	}
}

func (f *BitmapFont) parseCommonSection(keyValues map[string]string) {
	f.lineHeight, _ = strconv.Atoi(keyValues["lineHeight"])
	f.base, _ = strconv.Atoi(keyValues["base"])
	f.pageWidth, _ = strconv.Atoi(keyValues["scaleW"])
	f.pageHeight, _ = strconv.Atoi(keyValues["scaleH"])
	f.packed, _ = strconv.ParseBool(keyValues["packed"])
	f.numPages, _ = strconv.Atoi(keyValues["pages"])
}

func (f *BitmapFont) parsePageSection(keyValues map[string]string) {
	// TODO: This assumes the files are in order
	//Id, _ := strconv.Atoi(keyValues["Id"])
	f.pageNames = append(f.pageNames, keyValues["file"])
}

func (f *BitmapFont) parseCharSection(keyValues map[string]string) {
	c := &BitmapGlyph{}
	id, _ := strconv.ParseInt(keyValues["id"], 10, 32)
	c.Id = int(id)
	c.X, _ = strconv.Atoi(keyValues["x"])
	c.Y, _ = strconv.Atoi(keyValues["y"])
	c.Width, _ = strconv.Atoi(keyValues["width"])
	c.Height, _ = strconv.Atoi(keyValues["height"])
	c.OffsetX, _ = strconv.Atoi(keyValues["xoffset"])
	c.OffsetY, _ = strconv.Atoi(keyValues["yoffset"])
	c.AdvanceX, _ = strconv.Atoi(keyValues["xadvance"])
	c.PageIndex, _ = strconv.Atoi(keyValues["page"])
	c.TextureChannel, _ = strconv.Atoi(keyValues["chnl"])
	c.kerning = make(map[int32]int)
	f.glyphs[int32(c.Id)] = c
}

func (f *BitmapFont) parseKerningSection(keyValues map[string]string) {
	first, _ := strconv.ParseInt(keyValues["first"], 10, 32)
	second, _ := strconv.ParseInt(keyValues["second"], 10, 32)
	amount, _ := strconv.Atoi(keyValues["amount"])

	glyph, ok := f.glyphs[int32(second)]
	if !ok {
		fmt.Printf("Kerning parse error: glyph %v not found", first)
	}
	glyph.kerning[int32(first)] = amount
}

var bmSectionRex = regexp.MustCompile("^(\\w+) ")
var bmKeyValueRex = regexp.MustCompile("(\\w+)=\"?([\\w\\s ,._\\-]*)\"?( |$|\")")

func (f *BitmapFont) tokenizeLine(line string) (string, map[string]string) {
	sectionMatches := bmSectionRex.FindStringSubmatch(line)
	if sectionMatches == nil {
		return "", nil
	}
	sectionName := sectionMatches[1]
	data := bmKeyValueRex.FindAllStringSubmatch(line, -1)

	keyValues := make(map[string]string)
	for _, kv := range data {
		k := kv[1]
		v := strings.Trim(kv[2], " ")
		keyValues[k] = v
	}

	return sectionName, keyValues
}
