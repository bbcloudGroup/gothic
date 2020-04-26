package config

import "github.com/kataras/iris/v12/context"

const (
	StaticRequestPath	=	"StaticRequestPath"
	StaticAssets		=	"StaticAssets"
	StaticFavicon		=	"StaticFavicon"
	StaticIndexName		=	"StaticIndexName"
	StaticGzip			=	"StaticGzip"
	StaticShowList		=	"StaticShowList"
)

type Static struct {
	RequestPath string	`yaml:"RequestPath"`
	Assets		string	`yaml:"Assets"`
	Favicon		string	`yaml:"Favicon"`
	IndexName	string	`yaml:"IndexName"`
	Gzip		bool	`yaml:"Gzip"`
	ShowList	bool	`yaml:"ShowList"`
}

func (s *Static) Register(others *map[string]interface{}) {
	(*others)[StaticRequestPath] = s.RequestPath
	(*others)[StaticAssets] = s.Assets
	(*others)[StaticFavicon] = s.Favicon
	(*others)[StaticIndexName] = s.IndexName
	(*others)[StaticGzip] = s.Gzip
	(*others)[StaticShowList] = s.ShowList
}

func GetStatic(app context.Application) Static {
	return Static {
		RequestPath:	GetString(app, StaticRequestPath),
		Assets:      	GetString(app, StaticAssets),
		Favicon:     	GetString(app, StaticFavicon),
		IndexName:   	GetString(app, StaticIndexName),
		Gzip:			GetBool(app, StaticGzip),
		ShowList:		GetBool(app, StaticShowList),
	}
}

func DefaultStatic() Static {
	return Static{
		RequestPath: "/public",
		Assets:      "./public/",
		Favicon:     "favicon.ico",
		IndexName:   "index.html",
		Gzip:        false,
		ShowList:    false,
	}
}