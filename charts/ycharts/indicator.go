package ycharts

import (
	"github.com/wenooij/nuggit"
	"github.com/wenooij/nuggit/graphs"
	"github.com/wenooij/nuggit/v1alpha"
)

func indicatorPage() (b *graphs.Builder, out string) {
	b = &graphs.Builder{}
	basePath := b.Node("string", graphs.Data("indicators"))
	v := b.Node("Var", graphs.Data(v1alpha.Var{Default: "20_year_treasury_rate"}))
	elems := b.Node("Array", graphs.Data(v1alpha.Array{Op: v1alpha.ArrayAgg}),
		graphs.Edge(basePath),
		graphs.Edge(v),
	)
	path := b.Node("String", graphs.Data(v1alpha.String{
		Op: v1alpha.StringURLPathEscape,
	}), graphs.Edge(elems))
	pathJoin := b.Node("String", graphs.Data(v1alpha.String{
		Op: v1alpha.StringURLPathJoin,
	}), graphs.Edge(elems))
	source := b.Node("Source",
		graphs.Data(v1alpha.Source{
			Scheme: "https",
			Host:   "ycharts.com",
		}),
		graphs.Edge(path, graphs.SrcField(pathJoin)),
	)
	return b, source
}

func IndicatorPage() *nuggit.Graph {
	b, source := indicatorPage()
	http := b.Node("HTTP",
		graphs.Data(v1alpha.HTTP{}),
		graphs.Edge(source),
	)
	sink := b.Node("Sink", graphs.Edge(http))
	html := b.Node("HTML", graphs.Edge(sink, graphs.SrcField("sink")))
	indicatorSelect(b, html)
	return b.Build()
}

func IndicatorChromedp() *nuggit.Graph {
	b, source := indicatorPage()
	chromedp := b.Node("Chromedp", graphs.Edge(source, graphs.SrcField("source")))
	html := b.Node("HTML", graphs.Edge(chromedp, graphs.SrcField("bytes")))
	indicatorSelect(b, html)
	return b.Build()
}

func indicatorSelect(b *graphs.Builder, html string) {
	b.Node("Selector", graphs.Key("out"),
		graphs.Data(v1alpha.Selector{Selector: "div.key-stat-title"}),
		graphs.Edge(html, graphs.SrcField("node")),
	)
}
