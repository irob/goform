goform
=======

`goform` generate HTML forms dynamically and super easy using Golang/Go.

Status
=======

In development V2.0

Description
=======

`goform` is a super simple form generator, create dynamic forms without having to write HTML.

Templates included:
- bootstrap5
- html

By default `goform` render forms in `Bootstrap 5` style, also its posible to choose `html` format, where the inputs are render in plain html (no divs, no labels, ...)

If anyone need a custom template or custom items, `goform` has the option to choose custom template.

TODO:
- Parse HTML in textlabel element
- Remove Elements Functions
- Enable Set/Group of elements

## Example

	package main

	import(
        "text/template"

        "github.com/irob/goform"
	)

    var res = make(map[string]interface{})
    var tmpl = template.Must(template.ParseGlob("tmpl/*"))

	func main () {

        nInputs := 8

        // CitiesList slice of cities
        var CitiesList = []OptionItem{{Key: "", Value: "Choose your favorite city"}, {Key: "AMS", Value: "Amsterdam"}, {Key: "VEN", Value: "Venice"}, {Key: "KYO", Value: "Kyoto"}, {Key: "PAR", Value: "Paris"}, {Key: "DOH", Value: "Doha"}, {Key: "BAR", Value: "Barcelona"}, {Key: "SMA", Value: "San Miguel de Allende"}, {Key: "BUD", Value: "Budapest"}, {Key: "LIS", Value: "Lisbon"}, {Key: "FLO", Value: "Florence"}, {Key: "HNK", Value: "Hong Kong"}, {Key: "BRU", Value: "Bruges"}}
        // AgeRanges slice of ranges of ages
        var AgeRanges = []OptionItem{{Key: "1", Value: "1 - 9 yo"}, {Key: "2", Value: "10 - 19 yo"}, {Key: "3", Value: "20 - 29 yo"}, {Key: "4", Value: "30 - 39 yo"}, {Key: "5", Value: "40 - 49 yo"}, {Key: "6", Value: ">= 50 yo"}}

        form := Create("profile_form", "POST", "/goform")

        // Label input
        form.NewElement("label", "userdetails", "User profile")

        // Text input
        form.NewElement("text", "text", "")
        form.SetLabel("text", "What's your name")

        // Textlabel input
        form.NewElement("textlabel", "username", "john@bender.com")
        form.SetLabel("username", "Your username:")

        // Password input
        form.NewElement("password", "password", "")

        // Select input
        form.NewElement("select", "select", "VEN")
        form.SetOptions("select", CitiesList)

        // Radio input
        form.NewElement("radio", "radio", "")
        form.SetLabel("radio", "Age range")
        form.SetOptions("radio", AgeRanges)

        // Textarea
        form.NewElement("textarea", "textarea", "")
        form.SetHelpText("textarea", "Error, must write a resume description")

        // Checkbox
        form.NewElement("checkbox", "checkbox", "")

        // File input
        form.NewElement("file", "file", "")

        // Hidden
        form.NewElement("hidden", "hidden", "")

        // Full address init
        form.NewElement("label", "address_info", "Full address")
        form.AddGroupClass("address_info", "col-md-2")
        form.AddGroupClass("address_info", "mb-2")

        form.NewElement("text", "street", "")
        form.SetPlaceHolder("street", "Street")
        form.AddParams("street", "maxlength", "20")
        form.AddGroupClass("street", "col-md-4")
        form.AddGroupClass("street", "mb-2")

        form.NewElement("text", "number", "")
        form.SetPlaceHolder("number", "Number")
        form.AddParams("number", "maxlength", "20")
        form.AddGroupClass("number", "col-md-2")
        form.AddGroupClass("number", "mb-2")

        form.NewElement("select", "city", "VEN")
        form.SetOptions("city", CitiesList)
        form.AddGroupClass("city", "col-md-4")
        form.AddGroupClass("city", "mb-2")
        // Full address end

        form.NewRow("skills")
        // Dyanmic inputs
        for i := 1; i <= nInputs; i++ {
            form.NewElement("text", "skill_"+strconv.Itoa(i), "")
            form.SetPlaceHolder("skill_"+strconv.Itoa(i), "Skill "+strconv.Itoa(i))
            form.AddGroupClass("skill_"+strconv.Itoa(i), "col-md-6")
            form.AddGroupClass("skill_"+strconv.Itoa(i), "mb-2")
        }

        // Buttons
        form.NewElement("submit", "submit", "Update profile")
        form.AddClass("submit", "btn-success")
        form.AddClass("submit", "btn-lg")
        form.AddClass("submit", "btn-xl")
        form.AddClass("submit", "btn-block")

        form.NewElement("button", "button", "Update profile")
        form.AddClass("button", "btn-danger")
        form.AddClass("button", "btn-lg")
        form.AddClass("button", "btn-xl")
        form.AddClass("button", "btn-block")

        // Send to template
        res["Form"] = form

        tmpl.ExecuteTemplate(w, "HTMLTemplate", res)
    }

In your HTML template place
`{{ .Form.Render }}`

### Requests or bugs?
<https://github.com/irob/goform/issues>

## Installation

	go get github.com/irob/goform

## Use you custom templates

	/templates/custom_templates is a templete based on Bootstrap5.

	Step 1.-
	Move/Copy the folder /templates into your application, rename/copy the subdirectory /custom_templates.

	Step 2.-
	Customize the .html files of the elements you want with your own HTML format.

	Step 3.-
	Change the form style template name to you new template form.SetOwnStyleTemplate(YOUR_CUSTOM_THEME_NAME).

## License

The source files are distributed under the
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
