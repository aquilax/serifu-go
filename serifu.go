// Package serifu contains parser for the serfu markup language
package serifu

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type parseState = int

const (
	inScriptState parseState = iota
	inPageState
	inPanelState
)

const (
	pageSpreadPrefix       = "##"
	pagePrefix             = "#"
	panelPrefix            = "-"
	soundPrefix            = "*"
	sideNotePrefix         = "!"
	textLineSeparator      = ":"
	styleSeparator         = "/"
	preFormattedBlockStart = "/="
	preFormattedBlockEnd   = "=/"
)

// ItemType represents the type of the item in the panel
type ItemType string

const (
	// TextLineItemType is the type for a TextLine
	TextLineItemType = "text"
	// SideNoteItemType is the type for a SideNote
	SideNoteItemType = "sideNote"
	// SoundEffectItemType is the type for a SoundEffect
	SoundEffectItemType = "soundEffect"
)

// TextLine is a text line item
type TextLine struct {
	Type           ItemType `json:"type"`
	Source         string   `json:"source"`
	Style          string   `json:"style"`
	IsPreFormatted bool     `json:"is_pre_formatted"`
	Content        string   `json:"content"`
}

// SideNote is a side note item used for comments
type SideNote struct {
	Type    ItemType `json:"type"`
	Content string   `json:"content"`
}

// SoundEffect contains sound effect definition
type SoundEffect struct {
	Type            ItemType `json:"type"`
	Name            string   `json:"name"`
	Transliteration string   `json:"transliteration"`
}

// Items contains list of items per panel
type Items = []interface{}

// Panel contains comics panel items
type Panel struct {
	ID    string `json:"id"`
	Items Items  `json:"items"`
}

// Page is a comic page containing one or more panels
type Page struct {
	Title    string   `json:"title"`
	IsSpread bool     `json:"is_spread"`
	Panels   []*Panel `json:"panels"`
}

// Script is the whole script
type Script struct {
	Pages []*Page `json:"pages"`
}

// Parse parses the input stream and returns script or error
func Parse(io io.Reader) (*Script, error) {
	script := &Script{make([]*Page, 0)}
	scanner := bufio.NewScanner(io)

	state := inScriptState
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
			state = inPageState
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
			if state != inPageState && state != inPanelState {
				return nil, fmt.Errorf("line %d: unexpected panel definition outside of page", lineNumber)
			}
			state = inPanelState
			id := strings.TrimSpace(trimmedLine[1:])
			panel = &Panel{
				ID: id,
			}
			page.Panels = append(page.Panels, panel)
			continue
		}
		if strings.HasPrefix(line, soundPrefix) {
			if state != inPanelState {
				return nil, fmt.Errorf("line %d: unexpected sound definition outside of panel", lineNumber)
			}
			name := strings.TrimSpace(trimmedLine[1:])
			index := strings.Index(name, "(")
			transliteration := ""
			if index > -1 && strings.HasSuffix(name, ")") {
				transliteration = name[index+1 : len(name)-2]
				name = strings.TrimSpace(name[:index])
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
			if state != inPanelState {
				return nil, fmt.Errorf("line %d: unexpected side note definition outside of panel", lineNumber)
			}
			sideNote := strings.TrimSpace(trimmedLine[1:])
			panel.Items = append(panel.Items, SideNote{
				Type:    SideNoteItemType,
				Content: sideNote,
			})
			continue
		}
		index := strings.Index(line, textLineSeparator)
		if index > -1 {
			if state != inPanelState {
				return nil, fmt.Errorf("line %d: unexpected text line definition outside of panel", lineNumber)
			}
			isPreFormatted := false
			source := strings.TrimSpace(trimmedLine[:index])
			content := strings.TrimSpace(trimmedLine[index+1:])
			if strings.HasPrefix(content, preFormattedBlockStart) {
				isPreFormatted = true
				if strings.HasSuffix(content, preFormattedBlockEnd) {
					content = content[2 : len(content)-3]
					// single line block
				} else {
					// multi-line block
					var b strings.Builder
					b.WriteString(content[2:])
					// ingest block
					for scanner.Scan() {
						lineNumber++
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
				Type:           TextLineItemType,
				Source:         source,
				Style:          style,
				Content:        content,
				IsPreFormatted: isPreFormatted,
			}
			panel.Items = append(panel.Items, textLine)
			continue
		}
		if trimmedLine != "" {
			return nil, fmt.Errorf("line %d: unexpected markup: `%s`", lineNumber, line)
		}

	}

	return script, nil
}

func (s Script) String() string {
	var b strings.Builder
	for _, p := range s.Pages {
		b.WriteString(p.String())
		b.WriteByte('\n')
	}
	return b.String()
}

func (p Page) String() string {
	var b strings.Builder
	if p.IsSpread {
		b.WriteString(fmt.Sprintf("%s %s", pageSpreadPrefix, p.Title))
	} else {
		b.WriteString(fmt.Sprintf("%s %s", pagePrefix, p.Title))
	}
	b.WriteByte('\n')
	for _, pn := range p.Panels {
		b.WriteString(pn.String())
		b.WriteByte('\n')
	}
	return b.String()
}

func (pn Panel) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("    %s %s", panelPrefix, pn.ID))
	b.WriteByte('\n')
	for _, i := range pn.Items {
		st := i.(fmt.Stringer)
		b.WriteString(st.String())
		b.WriteByte('\n')
	}
	return b.String()
}

func (t TextLine) String() string {
	heading := t.Source
	if t.Style != "" {
		heading = fmt.Sprintf("%s%s%s", t.Source, styleSeparator, t.Source)
	}
	content := " " + t.Content
	if t.IsPreFormatted {
		content = fmt.Sprintf("%s%s%s", preFormattedBlockStart, t.Content, preFormattedBlockEnd)
	}
	return fmt.Sprintf("    %s%s%s", heading, textLineSeparator, content)
}

func (se SoundEffect) String() string {
	if se.Transliteration != "" {
		return fmt.Sprintf("    %s %s (%s)", soundPrefix, se.Name, se.Transliteration)
	}
	return fmt.Sprintf("    %s %s", soundPrefix, se.Name)
}

func (sn SideNote) String() string {
	return fmt.Sprintf("    %s %s", sideNotePrefix, sn.Content)
}
