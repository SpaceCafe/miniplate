package functions

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

type XMLWriter interface {
	io.ByteWriter
	io.Writer
}

func (DataFuncs) ToXML(obj any) (string, error) {
	return toXMLHelper(obj, "")
}

func (DataFuncs) ToXMLPretty(indent string, obj any) (string, error) {
	return toXMLHelper(obj, indent)
}

func (DataFuncs) XML(in any) (obj any, err error) {
	switch in := in.(type) {
	case []byte:
		err = xml.Unmarshal(in, &obj)
	case string:
		err = xml.Unmarshal([]byte(in), &obj)
	}

	return
}

func toXMLHelper(obj any, indent string) (string, error) {
	var builder strings.Builder

	err := buildXML(&builder, obj, nil, 0, []byte(indent))

	return builder.String(), err
}

func buildXML(w XMLWriter, obj any, key []byte, depth int, indent []byte) error {
	var (
		newline   []byte
		curIndent []byte
	)

	if len(indent) != 0 {
		newline = []byte("\n")
		curIndent = bytes.Repeat(indent, depth)
	}

	switch value := obj.(type) {
	case map[string]any:
		return buildXMLMap(w, value, depth, key, newline, indent, curIndent)
	case []any:
		return buildXMLSlice(w, value, depth, key, indent)
	default:
		return buildXMLScalar(w, value, key, newline, curIndent)
	}
}

func buildXMLMap(
	w XMLWriter,
	value map[string]any,
	depth int,
	key, newline, indent, curIndent []byte,
) error {
	if len(key) == 0 {
		for nestedKey, nestedValue := range value {
			err := buildXML(w, nestedValue, []byte(nestedKey), depth, indent)
			if err != nil {
				return err
			}
		}

		return nil
	}

	attrs, content, err := extractAttributesAndContent(value)
	if err != nil {
		return err
	}

	return xmlCreateNode(w, key, newline, curIndent, attrs, func(w XMLWriter) error {
		if content != nil {
			return xmlEscapeContent(w, content)
		}

		_, _ = w.Write(newline)

		for nestedKey, nestedValue := range value {
			err := buildXML(w, nestedValue, []byte(nestedKey), depth+1, indent)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func buildXMLSlice(w XMLWriter, value []any, depth int, key, indent []byte) error {
	for _, item := range value {
		err := buildXML(w, item, key, depth, indent)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildXMLScalar(w XMLWriter, value any, key, newline, curIndent []byte) error {
	if len(key) == 0 {
		return nil
	}

	return xmlCreateNode(w, key, newline, curIndent, nil, func(w XMLWriter) error {
		return xmlEscapeContent(w, value)
	})
}

func xmlCreateNode(
	w XMLWriter,
	key, newline, indent []byte,
	attrs *bytes.Buffer,
	callback func(XMLWriter) error,
) error {
	_, _ = w.Write(indent)
	_ = w.WriteByte('<')

	attrKey, err := xmlEscape(key)
	if err != nil {
		return err
	}

	_, _ = w.Write(attrKey)

	if attrs != nil {
		_, _ = attrs.WriteTo(w)
	}

	_ = w.WriteByte('>')

	err = callback(w)
	if err != nil {
		return err
	}

	_, _ = w.Write([]byte("</"))
	_, _ = w.Write(attrKey)
	_ = w.WriteByte('>')
	_, _ = w.Write(newline)

	return nil
}

// extractAttributesAndContent checks if a map represents an element with attributes
// Returns attributes string and content (or nil if it's a regular nested structure).
func extractAttributesAndContent(m map[string]any) (*bytes.Buffer, any, error) {
	var (
		buf        = &bytes.Buffer{}
		content    any
		hasContent = false
	)

	for key, value := range m {
		// Convention: keys starting with @ are attributes, #text is the text content
		if attrName, ok := bytes.CutPrefix([]byte(key), []byte("@")); ok {
			buf.WriteByte(' ')

			attrKey, err := xmlEscape(attrName)
			if err != nil {
				return nil, nil, err
			}

			buf.Write(attrKey)
			buf.WriteString(`="`)

			err = xmlEscapeAttr(buf, value)
			if err != nil {
				return nil, nil, err
			}

			buf.WriteByte('"')
		} else if key == "#text" {
			content = value
			hasContent = true
		}
	}

	if hasContent || buf.Len() > 0 {
		return buf, content, nil
	}

	return nil, nil, nil
}

func xmlEscape(in []byte) ([]byte, error) {
	buf := &bytes.Buffer{}

	err := xml.EscapeText(buf, in)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func xmlEscapeContent(w io.Writer, in any) error {
	return xml.EscapeText(w, fmt.Appendf([]byte(""), "%v", in))
}

func xmlEscapeAttr(w io.Writer, in any) error {
	// Use a temporary element to leverage xml.Marshal's attribute escaping
	type TempAttr struct {
		XMLName xml.Name `xml:"-"`
		Value   string   `xml:"v,attr"`
	}

	temp := TempAttr{Value: fmt.Sprintf("%v", in)}

	buf, err := xml.Marshal(temp)
	if err != nil {
		return err
	}

	// Extract the escaped value from: <TempAttr v="escaped_value"></TempAttr>
	start := bytes.Index(buf, []byte(`v="`))
	if start == -1 {
		return ErrInvalidXMLAttr
	}

	start += 3
	end := bytes.LastIndexByte(buf[start:], '"')

	if end == -1 {
		return ErrInvalidXMLAttr
	}

	_, err = w.Write(buf[start : start+end])

	return err
}
