package ycharts

import (
	"github.com/wenooij/nuggit"
	"github.com/wenooij/nuggit/graphs"
	"github.com/wenooij/nuggit/v1alpha"
)

func IndicatorPage() *nuggit.Graph {
	var b graphs.Builder

	v := b.Node("Var", graphs.Data(v1alpha.Var{
		Default: &v1alpha.Const{Type: nuggit.TypeString},
	}))
	enc := b.Node("String", graphs.Data(v1alpha.String{
		Op: v1alpha.StringURLPathEscape,
	}), graphs.Edge(v))
	host := b.Node("Const", graphs.Data(v1alpha.Const{
		Type:  nuggit.TypeString,
		Value: "https://ycharts.com/indicators/",
	}))
	http := b.Node("HTTP",
		graphs.Edge(host, graphs.SrcField("source.host")),
		graphs.Edge(enc, graphs.SrcField("source.path")),
	)
	b.Node("Sink", graphs.Key("out"), graphs.Edge(http))
	return b.Build()
}

