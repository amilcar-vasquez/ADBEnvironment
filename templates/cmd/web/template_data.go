package main

type TemplateData struct {
	Title      string
	HeaderText string
	FormErrors map[string]string
	FormData   map[string]string
}

// factory function to create a new instance of TemplateData
func NewTemplateData() *TemplateData {
	return &TemplateData{
		Title:      "Default Title",
		HeaderText: "Default HeaderText",
		FormErrors: make(map[string]string),
		FormData:   make(map[string]string),
	}
}
