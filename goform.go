// Package goform Generate html forms dynamically and super simple using Golang/Go.
package goform

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

var (
	logForm    []ErrorItem
	fieldTypes = map[string]string{
		"label":     "label",
		"text":      "text",
		"textlabel": "textlabel",
		"password":  "password",
		"select":    "select",
		"radio":     "radio",
		"textarea":  "textarea",
		"checkbox":  "checkbox",
		"file":      "file",
		"hidden":    "hidden",
		"button":    "button",
		"submit":    "submit",
		"row":       "row",
	}
)

// Form structure.
type Form struct {
	Name              string
	ID                string
	Method            string
	Action            string
	MultipartFormData string
	TemplateStyle     string
	TemplateSource    string
	FormTypes         map[string]int
	Elements          map[string]Field
	Classes           []string
	CSS               map[string]string
	FormText          string
	FormTemplates     map[string]*template.Template
	GroupClass        []string
}

// Element structure.
type Element struct {
	ID   int
	Name string
}

// Field structure.
type Field struct {
	Position    int
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

type SortField struct {
	Index int
	Field Field
}

type ErrorItem struct {
	RelatedTo string
	Message   string
}

//=============================================================================

// Create configure a new Form structure
func Create(name string, method string, action string) *Form {

	return &Form{
		name,
		name,
		method,
		action,
		"disabled",
		"bootstrap5",
		"",
		make(map[string]int),
		make(map[string]Field),
		[]string{},
		make(map[string]string),
		"",
		make(map[string]*template.Template),
		[]string{},
	}
}

// SetMultipartFormData set style format, (html or bootstrap5: default option)
func (f *Form) SetMultipartFormData(status string) {

	f.MultipartFormData = status
}

// SetTemplateStyle set style format, (html or bootstrap5: default option)
func (f *Form) SetTemplateStyle(style string) {

	f.TemplateStyle = style
}

// SetOwnTemplateStyle set style format, target different templates folder
// In case if any one need to use a custom templates
func (f *Form) SetOwnTemplateStyle(style string) {

	f.TemplateStyle = style
	f.TemplateSource = "OWN"
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

	if f.TemplateSource == "OWN" {
		cwd, _ = os.Getwd()
		cwd += path.Join("/templates/", f.TemplateStyle, ".html")

		tmpl, err = template.ParseFiles(cwd)
		if err != nil {
			panic(err)
		}

	} else {
		tmpl = HTMLTemplate(f.TemplateStyle, "form")
	}

	tmpl.Execute(buf, f)

	return template.HTML(buf.String())
}

// RenderElements returns form elements generated in plain format
func (f *Form) RenderElements() template.HTML {

	elementsSort := f.SortElements()

	var tmpl *template.Template
	buf := new(bytes.Buffer)
	cwd := ""
	var err error

	// Load ONCE all the necesary templates depending the input types
	for keyTemplate := range f.FormTypes {

		if f.TemplateSource == "OWN" {
			cwd, _ = os.Getwd()
			cwd += path.Join("templates", f.TemplateStyle, keyTemplate, ".html")

			tmpl, err = template.ParseFiles(cwd)
			if err != nil {
				panic(err)
			}
		} else {
			tmpl = HTMLTemplate(f.TemplateStyle, keyTemplate)
		}

		f.FormTemplates[keyTemplate] = tmpl
	}

	// Apply the template to each item of the form
	for _, itemForm := range elementsSort {

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
	field.Position = 0
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

// SortElements create and return empty form field
func (f *Form) SortElements() []Field {

	var elements []SortField
	for _, v := range f.Elements {
		elements = append(elements, SortField{v.Position, v})
	}

	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Index < elements[j].Index
	})

	// Elements sorted
	formElements := []Field{}

	for _, sortField := range elements {
		formElements = append(formElements, sortField.Field)
	}

	return formElements
}

// NewElement insert new form element
func (f *Form) NewElement(fieldType string, fieldName string, fieldValue string) string {

	// fieldName remove spaces and convert un lowercase
	fieldName = strings.ToLower(fieldName)
	fieldName = strings.Replace(fieldName, " ", "", -1)

	_, typeOk := fieldTypes[fieldType]
	// If the key exists
	if typeOk {

		field := EmptyField()
		field.Position = len(f.Elements) + 1
		field.FieldType = fieldType
		field.Name = fieldName
		field.ID = fieldName
		field.Value = fieldValue

		// Apped/Or Increase the input-type counter of FormTypes map
		// This map will is used in the RenderElements function
		f.FormTypes[fieldType]++

		_, fieldOk := f.Elements[fieldName]
		if !fieldOk {
			f.Elements[fieldName] = field
		} else {
			logForm = append(logForm, ErrorItem{RelatedTo: fieldName, Message: "Field Already Exists"})
		}

	} else {
		logForm = append(logForm, ErrorItem{RelatedTo: fieldType, Message: "Type Do Not Exists"})
	}

	return fieldName
}

// NewRow insert a new row shortcut.
func (f *Form) NewRow(rowName string) {
	f.NewElement("row", rowName, "")
}

// NewButton insert a new button shortcut.
func (f *Form) NewButton(buttonName string) {
	f.NewElement("button", buttonName, "")
}

// SetID set/change the ID to the field.
func (f *Form) SetID(fieldName string, id string) {

	// fieldID remove spaces and convert un lowercase
	id = strings.ToLower(id)
	id = strings.Replace(id, " ", "", -1)

	field := f.Elements[fieldName]
	field.ID = id
	f.Elements[fieldName] = field
}

// SetLabel set/change the text label to the field.
func (f *Form) SetLabel(fieldName string, label string) {
	field := f.Elements[fieldName]
	field.Label = label
	f.Elements[fieldName] = field
}

// AddClass adds a class to the input.
func (f *Form) AddClass(fieldName string, class string) {
	field := f.Elements[fieldName]
	field.Classes = append(field.Classes, class)
	f.Elements[fieldName] = field
}

// AddCSS add a CSS value (in the form of option-value - e.g.: color - red).
func (f *Form) AddCSS(fieldName string, key, value string) {
	f.Elements[fieldName].CSS[key] = value
}

// AddLabelClass adds a class to the label of the input.
func (f *Form) AddLabelClass(fieldName string, class string) {
	field := f.Elements[fieldName]
	field.LabelClass = append(field.LabelClass, class)
	f.Elements[fieldName] = field
}

// SetOptions set/change the Options of the dropdown.
func (f *Form) SetOptions(fieldName string, options []OptionItem) {
	field := f.Elements[fieldName]
	field.Options = options
	f.Elements[fieldName] = field
}

// SetPlaceHolder set the placeholder text to the input.
func (f *Form) SetPlaceHolder(fieldName string, placeholder string) {
	field := f.Elements[fieldName]
	field.PlaceHolder = placeholder
	f.Elements[fieldName] = field
}

// SetHelpText set the help-text to the input.
func (f *Form) SetHelpText(fieldName string, helptext string) {
	field := f.Elements[fieldName]
	field.HelpText = helptext
	f.Elements[fieldName] = field
}

// AddParams add a Param value (in the form of option-value - e.g.: maxlength - 15).
func (f *Form) AddParams(fieldName string, key, value string) {
	f.Elements[fieldName].Params[key] = value
}

// AddGroupClass adds a class to the group input.
func (f *Form) AddGroupClass(fieldName string, class string) {
	field := f.Elements[fieldName]
	field.GroupClass = append(field.GroupClass, class)
	f.Elements[fieldName] = field
}

// LogOutput display the Log
func LogOutput(format string) string {
	text := ""
	for _, logField := range logForm {
		switch format {
		case "return":
			text += logField.RelatedTo + " : " + logField.Message + "\n\r"
		default:
			log.Println(logField.RelatedTo, logField.Message)
		}
	}
	return text
}
