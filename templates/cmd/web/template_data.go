package main

type TemplateData struct {
	Title      string
	HeaderText string
}

// factory function to create a new instance of TemplateData
func NewTemplateData() *TemplateData {
	return &TemplateData{
		Title:      "Default Title",
		HeaderText: "Default HeaderText",
	}
}
