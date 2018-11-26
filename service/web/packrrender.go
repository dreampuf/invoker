package web

import (
	"github.com/gobuffalo/packr/v2"
	"net/http"
)

type PackrRender struct {
	Packr *packr.Box
	FilePath string
}

func (r *PackrRender) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := r.Packr.Find(r.FilePath)
	if err != nil {
		return err
	}

	w.Write(bytes)
	return nil
}
func (r *PackrRender) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}
}
