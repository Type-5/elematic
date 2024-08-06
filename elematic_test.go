package elematic

import (
	"bytes"
	"testing"
)

func TestCreateElement(t *testing.T) {
	f := Fragment(
		Div(Attrs{"class": "container"},
			H1(nil, Text("Hello, world!")),
		),
	)

	actual := f.ToString()
	expected := `<div class="container"><h1>Hello, world!</h1></div>`

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestDocument(t *testing.T) {
	doc := Document(
		DocType(),
		Html(nil,
			Head(nil,
				Title(nil, "Hello, world!"),
			),
			Body(nil,
				H1(nil, Text("Hello, world!")),
			),
		),
	)

	// render to string
	actual := doc.ToString()
	expected := `<!DOCTYPE html><html><head><title>Hello, world!</title></head><body><h1>Hello, world!</h1></body></html>`

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestIf(t *testing.T) {
	d := Div(nil,
		If(true, Div(nil, Text("true"))),
		If(false, Div(nil, Text("false"))),
	)

	// render to string
	actual := d.ToString()
	expected := `<div><div>true</div></div>`

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestMap(t *testing.T) {
	items := []string{"a", "b", "c"}

	d := Div(nil,
		Map(items, func(item string) Element {
			return Div(nil, Text(item))
		})...,
	)

	actual := d.ToString()
	expected := `<div><div>a</div><div>b</div><div>c</div></div>`

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestRender(t *testing.T) {
	d := Div(nil,
		Text("Hello, world!"),
	)

	s := bytes.NewBufferString("")
	d.Render(s)
	expected := `<div>Hello, world!</div>`

	if s.String() != expected {
		t.Errorf("Expected %s, got %s", expected, s.String())
	}
}

func TestText(t *testing.T) {
	d := Div(nil,
		Text("Hello, world!"),
	)

	actual := d.ToString()
	expected := `<div>Hello, world!</div>`

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
