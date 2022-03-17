package static

import (
	"os"
	"sumbur/utils"
	"sync"

	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
)

var (
	Debug bool
	cache sync.Map
)

func TimeStamp(path string) int64 {
	stamp_cache, ok := cache.Load(path)
	if ok {
		return stamp_cache.(int64)
	}

	stat, err := os.Stat("static/" + path)
	if err != nil {
		return 0
	}

	stamp := stat.ModTime().Unix()

	if !Debug {
		cache.Store(path, stamp)
	}

	return stamp
}

func StaticStampGET(ctx *atreugo.RequestCtx) error {
	fasthttp.ServeFileUncompressed(ctx.RequestCtx, "static/"+*utils.StrPathValue(ctx, "path"))
	return nil
}
