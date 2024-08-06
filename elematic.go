package elematic

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

// An alias for readability.
type Attrs map[string]string

// The base type for all elements.
// It wraps an HTML node.
type Element struct {
	Node *html.Node
}

// Appends the provided attributes to the node.
func appendAttributes(n *html.Node, attrs Attrs) {
	for key, value := range attrs {
		n.Attr = append(n.Attr, html.Attribute{Key: key, Val: value})
	}
}

// Returns an HTML element with the provided attributes and children.
func CreateElement(name string, attrs Attrs, children ...Element) Element {
	e := Element{&html.Node{Type: html.ElementNode, Data: name}}
	appendAttributes(e.Node, attrs)
	for _, child := range children {
		e.Node.AppendChild(child.Node)
	}
	return e
}

// Renders the element to the provided writer.
func (e Element) Render(w io.Writer) {
	html.Render(w, e.Node)
}

// Renders the element to a string.
func (e Element) ToString() string {
	var b bytes.Buffer
	e.Render(&b)
	return b.String()
}

/*
HELPER FUNCTIONS
*/

// Returns an empty document node.
func Document(children ...Element) Element {
	n := &html.Node{Type: html.DocumentNode, Data: "html"}
	for _, child := range children {
		n.AppendChild(child.Node)
	}
	return Element{n}
}

// returns <!DOCTYPE html>
func DocType() Element {
	return Element{&html.Node{Type: html.DoctypeNode, Data: "html"}}
}

// Takes and simply returns the passed in node.
// Use this for readibility when you want to return an element in a `return` statement but preserve indentation.
func Fragment(element Element) Element {
	return element
}

// Returns an element if the condition is true, otherwise an empty text node.
func If(condition bool, element Element) Element {
	if condition {
		return element
	}
	return Text("")
}

// Iterates over the items and builds a list of elements.
func Map[T any](items []T, f func(T) Element) []Element {
	var elements []Element
	for _, item := range items {
		elements = append(elements, f(item))
	}
	return elements
}

// Returns a text node with the provided text.
func Text(text string) Element {
	return Element{&html.Node{Type: html.TextNode, Data: text}}
}
