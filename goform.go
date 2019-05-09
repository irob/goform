// Package goform Generate html forms dynamically and super simple using Golang/Go.
package goform

import (
	"bytes"
	"html/template"
	"os"
	"path"
	"strings"
)

var (
	formText string
)

// Form structure.
type Form struct {
	Name          string
	ID            string
	Method        string
	Action        string
	StyleTemplate string
	StyleOrigin   string
	FormTypes     map[string]int
	FormElements  []Field
	Classes       []string
	CSS           map[string]string
	FormText      string
	FormTemplates map[string]*template.Template
	GroupClass    []string
}

// Field structure.
type Field struct {
	FieldType   string
	Name        string
	ID          string
	Classes     []string
	CSS         map[string]string
	Label       string
	LabelClass  []string
	Value       string
	Options     []OptionItem
	PlaceHolder string
	HelpText    string
	Params      map[string]string
	Set         string
	GroupClass  []string
}

// OptionItem structure.
type OptionItem struct {
	Key   string
	Value string
}

//=============================================================================

// Create configure a new Form structure
func Create(name string, method string, action string) *Form {

	return &Form{
		name,
		name,
		method,
		action,
		"bootstrap4",
		"",
		make(map[string]int),
		[]Field{},
		[]string{},
		make(map[string]string),
		"",
		make(map[string]*template.Template),
		[]string{},
	}
}

// SetStyleTemplate set style format, (html or bootstrap4: default option)
func (f *Form) SetStyleTemplate(style string) {

	f.StyleTemplate = style
}

// SetOwnStyle set style format, target different templates folder
// In case if any one need to use a custom templates
func (f *Form) SetOwnStyleTemplate(style string) {

	f.StyleTemplate = style
	f.StyleOrigin = "OWN"
}

// DefaultGroupClass set default group classes for all the elements
func (f *Form) DefaultGroupClass(width string) {

	f.GroupClass = append(f.GroupClass, width)
}

// Render returns form generated in plain format
func (f *Form) Render() template.HTML {

	var tmpl *template.Template
	buf := new(bytes.Buffer)
	cwd := ""
	var err error

	if f.StyleOrigin == "OWN" {
		cwd, _ = os.Getwd()
		cwd += path.Join("/templates/", f.StyleTemplate, ".html")

		tmpl, err = template.ParseFiles(cwd)
		if err != nil {
			panic(err)
		}

	} else {
		tmpl = HTMLTemplate(f.StyleTemplate, "form")
	}

	tmpl.Execute(buf, f)

	return template.HTML(buf.String())
}

// RenderElements returns form elements generated in plain format
func (f *Form) RenderElements() template.HTML {

	var tmpl *template.Template
	buf := new(bytes.Buffer)
	cwd := ""
	var err error

	// Load ONCE all the necesary templates depending the input types
	for keyTemplate := range f.FormTypes {

		if f.StyleOrigin == "OWN" {
			cwd, _ = os.Getwd()
			cwd += path.Join("templates", f.StyleTemplate, keyTemplate, ".html")

			tmpl, err = template.ParseFiles(cwd)
			if err != nil {
				panic(err)
			}
		} else {
			tmpl = HTMLTemplate(f.StyleTemplate, keyTemplate)
		}

		f.FormTemplates[keyTemplate] = tmpl
	}

	// Apply the template to each item of the form
	for _, itemForm := range f.FormElements {

		// Apply the default classes if exists, only if the element have not own group classes
		if len(itemForm.GroupClass) == 0 && len(f.GroupClass) > 0 {
			itemForm.GroupClass = f.GroupClass
		}

		f.FormTemplates[itemForm.FieldType].Execute(buf, itemForm)
		f.FormText += buf.String()

		// Clear buffer
		buf = new(bytes.Buffer)
	}

	return template.HTML(f.FormText)
}

// EmptyField create and return empty form field
func EmptyField() Field {

	field := Field{}
	field.FieldType = ""
	field.Name = ""
	field.ID = ""
	field.Classes = []string{}
	field.CSS = map[string]string{}
	field.Label = ""
	field.LabelClass = []string{}
	field.Value = ""
	field.Options = []OptionItem{}
	field.PlaceHolder = ""
	field.HelpText = ""
	field.Params = map[string]string{}
	field.GroupClass = []string{}

	return field
}

// NewElement insert new form element
// Returns the last item appended to the slice of elements
func (f *Form) NewElement(fieldType string, fieldName string, fieldValue string) *Field {

	// fieldName remove spaces and convert un lowercase
	fieldName = strings.ToLower(fieldName)
	fieldName = strings.Replace(fieldName, " ", "", -1)

	field := EmptyField()
	field.FieldType = fieldType
	field.Name = fieldName
	field.ID = fieldName
	field.Value = fieldValue

	// Apped/Or Increase the input-type counter of FormTypes map
	// This map will is used in the RenderElements function
	f.FormTypes[fieldType]++

	f.FormElements = append(f.FormElements, field)
	return &f.FormElements[len(f.FormElements)-1]
}

// NewButton insert new button element.
// Returns the last item appended to the slice of elements
func (f *Form) NewButton(buttonType string, fieldValue string) *Field {

	// fieldName remove spaces and convert un lowercase
	fieldName := strings.ToLower(buttonType)
	fieldName = strings.Replace(fieldName, " ", "", -1)

	field := EmptyField()
	field.FieldType = buttonType
	field.Name = fieldName
	field.ID = fieldName
	field.Value = fieldValue

	// Apped/Or Increase the input-type counter of FormTypes map
	// This map will is used in the RenderElements function
	f.FormTypes[buttonType]++

	f.FormElements = append(f.FormElements, field)
	return &f.FormElements[len(f.FormElements)-1]
}

// NewRow insert a new row.
func (f *Form) NewRow(rowName string) {

	// fieldName remove spaces and convert un lowercase
	fieldName := strings.ToLower(rowName)
	fieldName = strings.Replace(fieldName, " ", "", -1)

	field := EmptyField()
	field.FieldType = "row"
	field.Name = fieldName
	field.ID = fieldName

	// Apped/Or Increase the input-type counter of FormTypes map
	// This map will is used in the RenderElements function
	f.FormTypes["row"]++

	f.FormElements = append(f.FormElements, field)
}

// SetID set/change the ID to the field.
func (f *Field) SetID(id string) *Field {

	// fieldID remove spaces and convert un lowercase
	id = strings.ToLower(id)
	id = strings.Replace(id, " ", "", -1)

	f.ID = id
	return f
}

// SetLabel set/change the text label to the field.
func (f *Field) SetLabel(label string) *Field {
	f.Label = label
	return f
}

// ADD/SET Functions

// AddClass adds a class to the input.
func (f *Field) AddClass(class string) *Field {
	f.Classes = append(f.Classes, class)
	return f
}

// AddCSS add a CSS value (in the form of option-value - e.g.: color - red).
func (f *Field) AddCSS(key, value string) *Field {
	f.CSS[key] = value
	return f
}

// AddLabelClass adds a class to the label of the input.
func (f *Field) AddLabelClass(class string) *Field {
	f.LabelClass = append(f.LabelClass, class)
	return f
}

// SetOptions set/change the Options of the dropdown.
func (f *Field) SetOptions(options []OptionItem) *Field {
	f.Options = options
	return f
}

// SetPlaceHolder set the placeholder text to the input.
func (f *Field) SetPlaceHolder(placeholder string) *Field {
	f.PlaceHolder = placeholder
	return f
}

// SetHelpText set the help-text to the input.
func (f *Field) SetHelpText(helptext string) *Field {
	f.HelpText = helptext
	return f
}

// AddParams add a Param value (in the form of option-value - e.g.: maxlength - 15).
func (f *Field) AddParams(key, value string) *Field {
	f.Params[key] = value
	return f
}

// AddGroupClass adds a class to the group input.
func (f *Field) AddGroupClass(class string) *Field {
	f.GroupClass = append(f.GroupClass, class)
	return f
}
