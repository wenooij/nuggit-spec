package ycharts

import (
	"github.com/wenooij/nuggit"
	"github.com/wenooij/nuggit/graphs"
	"github.com/wenooij/nuggit/v1alpha"
)

func indicatorPage() (b *graphs.Builder, out string) {
	b = &graphs.Builder{}
	v := b.Node("Var", graphs.Data(v1alpha.Var{
		Default: &v1alpha.Const{Type: nuggit.TypeString, Value: "20_year_treasury_rate"},
	}))
	path := b.Node("String", graphs.Data(v1alpha.String{
		Op: v1alpha.StringURLPathEscape,
	}), graphs.Edge(v))
	source := b.Node("Source",
		graphs.Data(v1alpha.Source{
			Scheme: "https",
			Host:   "ycharts.com",
			Path:   "/indicators/",
		}),
		graphs.Edge(path, graphs.SrcField("path"), graphs.Glom(nuggit.GlomAppend)),
	)
	return b, source
}

func IndicatorPage() *nuggit.Graph {
	b, source := indicatorPage()
	http := b.Node("HTTP",
		graphs.Data(v1alpha.HTTP{}),
		graphs.Edge(source),
	)
	b.Node("Sink", graphs.Key("out"), graphs.Edge(http))
	return b.Build()
}

func IndicatorChromedp() *nuggit.Graph {
	b, source := indicatorPage()
	b.Node("Chromedp",
		graphs.Key("out"),
		graphs.Edge(source, graphs.SrcField("source")),
	)
	return b.Build()
}
