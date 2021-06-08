package h

import (
  "io"
  "strings"
)

var voidTags = map[string]bool{
  "area":   true,
  "base":   true,
  "br":     true,
  "col":    true,
  "embed":  true,
  "hr":     true,
  "img":    true,
  "input":  true,
  "keygen": true,
  "link":   true,
  "meta":   true,
  "param":  true,
  "source": true,
  "track":  true,
  "wbr":    true,
}

func encodeForHtml(raw string) (encoded string) {
  encoded = raw
  encoded = strings.ReplaceAll(encoded, "&", "&amp;")
  encoded = strings.ReplaceAll(encoded, "'", "&#39;")
  encoded = strings.ReplaceAll(encoded, "\"", "&quot;")
  encoded = strings.ReplaceAll(encoded, "<", "&lt;")
  encoded = strings.ReplaceAll(encoded, ">", "&gt;")
  return
}

type parsedSel struct {
  tag       string
  id        string
  className string
}

func parseSel(sel string) parsedSel {
  var tag string
  var id string
  var classes []string
  lastSep := '\xff'
  for len(sel) > 0 {
    i := strings.IndexAny(sel, ".#")
    var part string
    var nextSep rune
    if i == -1 {
      part = sel
      nextSep = '\xff'
      sel = ""
    } else {
      part = sel[:i]
      nextSep = int32(sel[i])
      sel = sel[i+1:]
    }
    switch lastSep {
    case '\xff':
      tag = strings.ToLower(part)
    case '.':
      classes = append(classes, part)
    case '#':
      id = part
    default:
      panic("Unreachable")
    }
    lastSep = nextSep
  }
  return parsedSel{
    tag:       tag,
    id:        id,
    className: strings.Join(classes, " "),
  }
}

type Rendered struct {
  Html string
}

func (r Rendered) Reader() io.Reader {
  return strings.NewReader(r.Html)
}

type C = []interface{}

type A = map[string]string

func H(selRaw string, args ...interface{}) Rendered {
  var buffer strings.Builder
  sel := parseSel(selRaw)
  var attrs A
  var children []interface{}
  switch v := args[0].(type) {
  case A:
    attrs = v
  case string:
    children = C{v}
  case C:
    children = v
  default:
    panic("Unrecognised argument")
  }
  if len(args) > 1 {
    switch v := args[1].(type) {
    case string:
      children = C{v}
    case C:
      children = v
    }
  }
  buffer.WriteString("<")
  buffer.WriteString(sel.tag)
  for n, v := range attrs {
    buffer.WriteString(" ")
    buffer.WriteString(n)
    buffer.WriteString("=\"")
    buffer.WriteString(encodeForHtml(v))
    buffer.WriteString("\"")
  }
  buffer.WriteString(">")
  for _, c := range children {
    switch c := c.(type) {
    case string:
      buffer.WriteString(encodeForHtml(c))
    case Rendered:
      buffer.WriteString(c.Html)
    default:
      panic("Unreachable")
    }
  }
  if _, ok := voidTags[sel.tag]; !ok {
    buffer.WriteString("</")
    buffer.WriteString(sel.tag)
    buffer.WriteString(">")
  }
  return Rendered{Html: buffer.String()}
}
