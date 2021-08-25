package ioc

import "fmt"

type Widget struct {
	X, Y int
}

type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

type Button struct {
	Label // Embedding (delegation)
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}

func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}

func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}

func (box ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", box.Texts)
}

func (box ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", box.Texts)
}
