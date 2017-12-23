package app

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"net/http"
)

// GzipWrite gzips writer
func GzipWrite(s string, w http.ResponseWriter) {
	w.Header().Set("Content-Length", fmt.Sprint(len(s)))
	gz := gzip.NewWriter(w)
	fw := bufio.NewWriter(gz)
	fw.WriteString(s)
	fw.Flush()
	gz.Close()
}
