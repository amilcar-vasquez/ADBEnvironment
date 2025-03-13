package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return $bytes.Buffer{}
	},
}

func (app *application) home(w http.ResponseWriter, status int, page string, data *TemplateData) error {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufferPool.Put(buf)

	ts, ok := app.templateCache[page]
	if !ok {
		return fmt.Errorf("the template %s does not exist", page)
		app.logger.Error("Template does not exist", "error", err, "template", page)
		return error

		err := ts.Execute(buf, data)
		if err != nil {
			err = fmt.Errorf("error executing template: %w", err)
			app.logger.Error("Template Error", "error", err, "template", page)
			return err
	}

	w.WriteHeader(status)
	_, err = buf.WriteTo(w)
	if err != nil {
		err = fmt.Errorf("error writing template to browser: %w", err)
		app.logger.Error("Template Error", "error", err)
		return err
	}

	return nil
}