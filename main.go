/*
Copyright 2019 Special Brands Holding BV
Copyright 2014 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// outyet is a web server that announces whether or not a particular Go version
// has been tagged.
package main

import (
	roman "fantastic-lamp/task2"
	"flag"
	"html/template"
	"log"
	"net/http"
)

// Command-line flags.
var (
	httpAddr = flag.String("http", ":8080", "Listen address")
)

func main() {
	flag.Parse()
	http.Handle("/", NewServer())
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}

// Server implements the outyet server.
// It serves the user interface (it's an http.Handler)
// and polls the remote repository for changes.
type Server struct {
	Expr string
}

// NewServer returns an initialized outyet server.
func NewServer() *Server {
	s := &Server{}
	return s
}

// ServeHTTP implements the HTTP user interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Expr = "XI+I"
	data := struct {
		Expr  string
		Value string
	}{
		s.Expr,
		//	"XI",
		roman.RomanNumeralCalculate(s.Expr),
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
	}
}

// tmpl is the HTML template that drives the user interface.
var tmpl = template.Must(template.New("tmpl").Parse(`
<!DOCTYPE html><html><body><center>
	<h2>{{.Expr}}</h2>
	<h1>
	{{.Value}}
	</h1>
</center></body></html>
`))
