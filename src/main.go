package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"flag"
	"strings"
	"path/filepath"
	"os"
)

type handle struct {
	reverseProxy string
}

func substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start : end])
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host

	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.URL.Path + " " + r.Proto + " " + r.UserAgent())

	path := r.URL.Path
	pos := strings.Index(path, specUrlHead)
	if pos != -1 {
		realFilePath := substring(path, pos + len(specUrlHead), len(path))
		if len(realFilePath) > 0 {
			log.Println("start serve local file " + realFilePath)

			localPath := filepath.Join(".", localDir)
			os.MkdirAll(localPath, os.ModePerm)

			http.ServeFile(w, r, localPath + "/" + realFilePath)
		}
	} else {
		proxy.ServeHTTP(w, r)
	}
}




var remoteHttpAddr = "http://114.55.5.207:82"
var localDir = "spa"
var specUrlHead = "/spa/"

func main() {
	bind := flag.String("l", "0.0.0.0:8888", "listen on ip:port")
	remote := flag.String("r", remoteHttpAddr, "reverse proxy addr")
	flag.Parse()
	log.Printf("Listening on %s, forwarding to %s", *bind, *remote)
	h := &handle{reverseProxy: *remote}
	err := http.ListenAndServe(*bind, h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}