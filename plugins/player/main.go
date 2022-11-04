package player

import (
	"embed"
	"fmt"
	. "m7s.live/engine/v4"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

//go:embed ui
var f embed.FS

var plugin = InstallPlugin(new(PlayerConfig))

type PlayerConfig struct {
}

func (c *PlayerConfig) OnEvent(event interface{}) {
	switch event.(type) {
	case FirstConfig:

	}
}

func (c *PlayerConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/preview/" {
		Streams.Range(func(streamPath string, s *Stream) {
			w.Write([]byte(fmt.Sprintf("<a href='%s'>%s</a><br>", streamPath, streamPath)))
		})
		return
	}
	ss := strings.Split(r.URL.Path, "/")
	if b, err := f.ReadFile("ui/" + ss[len(ss)-1]); err == nil {
		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(ss[len(ss)-1])))
		w.Write(b)
	} else {
		//w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		//w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		b, err = f.ReadFile("ui/demo.html")
		w.Write(b)
	}
}
