package unrolled_render

import (
	"html/template"
	"io"
	"net/http"
)

// Engine is the generic interface for all responses.
type Engine interface {
	Render(io.Writer, interface{}) error
}

// Head defines the basic ContentType and Status fields.
type Head struct {
	ContentType string
	Status      int
}

// Data built-in renderer.
type Data struct {
	Head
}

// HTML built-in renderer.
type HTML struct {
	Head
	Name      string
	Templates *template.Template
}

// Write outputs the header content.
func (h Head) Write(w http.ResponseWriter) {
	w.Header().Set(ContentType, h.ContentType)
	w.WriteHeader(h.Status)
}

// Render a data response.
func (d Data) Render(w io.Writer, v interface{}) error {
	if hw, ok := w.(http.ResponseWriter); ok {
		c := hw.Header().Get(ContentType)
		if c != "" {
			d.Head.ContentType = c
		}
		d.Head.Write(hw)
	}

	w.Write(v.([]byte))
	return nil
}

// Render a HTML response.
func (h HTML) Render(w io.Writer, binding interface{}) error {
	// Retrieve a buffer from the pool to write to.
	out := bufPool.Get()
	err := h.Templates.ExecuteTemplate(out, h.Name, binding)
	if err != nil {
		return err
	}

	if hw, ok := w.(http.ResponseWriter); ok {
		h.Head.Write(hw)
	}
	out.WriteTo(w)

	// Return the buffer to the pool.
	bufPool.Put(out)
	return nil
}
