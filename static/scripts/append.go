package main

import "honnef.co/go/js/dom"

func main() {
	doc := dom.GetWindow().Document()
  

	p := doc.GetElementByID("text").(*dom.HTMLInputElement)
	p.Focus()
	p.AddEventListener("keyup", false, inputKeyUp)
}

func inputKeyUp(event dom.Event) {
	input := event.Target().(*dom.HTMLInputElement)
	span := dom.GetWindow().Document().GetElementByID("value")
	span.SetInnerHTML(input.Value)
}
