package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type PageView struct {
	vecty.Core
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Text("Hello Vecty!"),
	)
}

func main() {
	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(&PageView{})
}
