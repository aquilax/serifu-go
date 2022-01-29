package serifu

import (
	"bufio"
	"io"
	"strings"
)

const (
	pageSpreadPrefix       = "##"
	pagePrefix             = "#"
	panelPrefix            = "-"
	soundPrefix            = "*"
	sideNotePrefix         = "*"
	textLineSeparator      = ":"
	styleSeparator         = "/"
	preFormattedBlockStart = "/="
	preFormattedBlockEnd   = "=/"
)

// ItemType represents the type of the item in the panel
type ItemType string

const (
	TextLineItemType    = "text"
	SideNoteItemType    = "sideNote"
	SoundEffectItemType = "soundEffect"
)

// TextLine is a text line item
type TextLine struct {
	Type    ItemType
	Source  string
	Style   string
	Content string
}

// SideNote is a side note item used for comments
type SideNote struct {
	Type    ItemType
	Content string
}

// SoundEffect contains sound effect definition
type SoundEffect struct {
	Type            ItemType
	Name            string
	Transliteration string
}

// Items contains list of items per panel
type Items = []interface{}

// Panel contains comics panel items
type Panel struct {
	Id    string
	Items Items
}

// Page is a comic page containing one or more panels
type Page struct { // #
	Title    string
	IsSpread bool     // ##
	Panels   []*Panel // -
}

// Page is the whole script
type Script struct {
	Pages []*Page //Page
}

// Parse parses the input stream and returns script or error
func Parse(io io.Reader) (*Script, error) {
	script := &Script{make([]*Page, 0)}
	scanner := bufio.NewScanner(io)

	var line string
	var trimmedLine string
	var page *Page
	var panel *Panel
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line = scanner.Text()
		trimmedLine = strings.TrimSpace(line)

		if strings.HasPrefix(line, pagePrefix) {
			var title string
			isSpread := false
			if strings.HasPrefix(trimmedLine, pageSpreadPrefix) {
				title = strings.TrimSpace(trimmedLine[2:])
				isSpread = true
			} else {
				title = strings.TrimSpace(trimmedLine[1:])
			}
			page = &Page{
				Title:    title,
				IsSpread: isSpread,
			}
			script.Pages = append(script.Pages, page)
			continue
		}
		if strings.HasPrefix(line, panelPrefix) {
			id := strings.TrimSpace(trimmedLine[1:])
			panel = &Panel{
				Id: id,
			}
			page.Panels = append(page.Panels, panel)
			continue
		}
		if strings.HasPrefix(line, soundPrefix) {
			name := strings.TrimSpace(trimmedLine[1:])
			index := strings.Index(name, "(")
			transliteration := ""
			if index > -1 && strings.HasSuffix(name, ")") {
				transliteration = name[index : len(name)-2]
				name = name[:index]
			}
			sound := &SoundEffect{
				Type:            SoundEffectItemType,
				Name:            name,
				Transliteration: transliteration,
			}
			panel.Items = append(panel.Items, sound)
			continue
		}
		if strings.HasPrefix(line, sideNotePrefix) {
			sideNote := strings.TrimSpace(trimmedLine[1:])
			panel.Items = append(panel.Items, SideNote{
				Type:    SideNoteItemType,
				Content: sideNote,
			})
			continue
		}
		index := strings.Index(line, textLineSeparator)
		if index > -1 {
			source := strings.TrimSpace(trimmedLine[:index])
			content := strings.TrimSpace(trimmedLine[index+1:])
			if strings.HasPrefix(content, preFormattedBlockStart) {
				if strings.HasSuffix(content, preFormattedBlockEnd) {
					content = content[2 : len(content)-3]
					// single line block
				} else {
					// multi-line block
					var b strings.Builder
					b.WriteString(content[2:])
					// ingest block
					for scanner.Scan() {
						line = scanner.Text()
						trimmedLine = strings.TrimSpace(line)
						if strings.HasSuffix(trimmedLine, preFormattedBlockEnd) {
							content = b.String()
							break
						}
						b.WriteString(line)
						b.WriteByte('\n')
					}
				}
			}
			style := ""
			styleIndex := strings.Index(source, styleSeparator)
			if styleIndex > -1 {
				style = source[styleIndex+1:]
				source = source[:styleIndex]
			}
			textLine := TextLine{
				Type:    TextLineItemType,
				Source:  source,
				Style:   style,
				Content: content,
			}
			panel.Items = append(panel.Items, textLine)
		}
	}

	return script, nil
}
