goform
=======

`goform` generate HTML forms dynamically and super easy using Golang/Go.

Description
=======

`goform` is a super simple form generator, create dynamic forms without having to write HTML.

Templates included:
- bootstrap4
- html

By default `goform` render forms in `Bootstrap 4` style, also its posible to choose `html` format, where the inputs are render in plain html (no divs, no labels, ...)

If anyone need a custom template or custom items, `goform` has the option to choose custom template.

TODO:
- Parse HTML in textlabel element
- Remove Elements Functions
- Enable Set/Group of elements

## Example

	package main

	import(
		"github.com/irob/goform"
	)

	var res = make(map[string]interface{})

	func main () {

		nInputs := 8

		// CitiesList slice of cities
		var CitiesList = []goform.OptionItem{{Key: "", Value: "Choose your favorite city"}, {Key: "AMS", Value: "Amsterdam"}, {Key: "VEN", Value: "Venice"}, {Key: "KYO", Value: "Kyoto"}, {Key: "PAR", Value: "Paris"}, {Key: "DOH", Value: "Doha"}, {Key: "BAR", Value: "Barcelona"}, {Key: "SMA", Value: "San Miguel de Allende"}, {Key: "BUD", Value: "Budapest"}, {Key: "LIS", Value: "Lisbon"}, {Key: "FLO", Value: "Florence"}, {Key: "HNK", Value: "Hong Kong"}, {Key: "BRU", Value: "Bruges"}}
		// AgeRanges slice of ranges of ages
		var AgeRanges = []goform.OptionItem{{Key: "1", Value: "1 - 9 yo"}, {Key: "2", Value: "10 - 19 yo"}, {Key: "3", Value: "20 - 29 yo"}, {Key: "4", Value: "30 - 39 yo"}, {Key: "5", Value: "40 - 49 yo"}, {Key: "6", Value: ">= 50 yo"}}

		form := goform.Create("profile_form", "POST", "/goform")
		//form.SetStyleTemplate("html")
		//form.SetOwnStyleTemplate("local_custom_template") // Local template files
		form.DefaultGroupClass("col-md-12")
		form.DefaultGroupClass("mb-2")
		form.NewElement("label", "userdetails", "User profile").AddCSS("font-size", "2em").AddCSS("font-weight", "bold").AddCSS("font-weight", "bold")
		form.NewElement("textlabel", "username", "john@bender.com").SetLabel("Your username:").AddCSS("font-weight", "bold").AddCSS("font-weight", "bold")
		form.NewElement("text", "name", "").SetLabel("What's your name").SetID("name").SetPlaceHolder("What's your name").AddCSS("color", "red")

		form.NewElement("radio", "age_range", "").SetLabel("Age range").SetOptions(AgeRanges)

		form.NewElement("label", "address_info", "Full address").AddGroupClass("col-md-2").AddGroupClass("mb-2")
		form.NewElement("text", "street", "").SetPlaceHolder("Street").AddParams("maxlength", "20").AddGroupClass("col-md-4").AddGroupClass("mb-2")
		form.NewElement("text", "number", "").SetPlaceHolder("Number").AddParams("maxlength", "20").AddGroupClass("col-md-2").AddGroupClass("mb-2")
		form.NewElement("select", "city", "VEN").SetOptions(CitiesList).AddGroupClass("col-md-4").AddGroupClass("mb-2")

		form.NewRow("skills")
		// Dyanmic inputs
		for i := 1; i <= nInputs; i++ {
			form.NewElement("text", "skill_"+strconv.Itoa(i), "").SetPlaceHolder("Skill " + strconv.Itoa(i)).AddGroupClass("col-md-6").AddGroupClass("mb-2")
		}

		form.NewElement("textarea", "resume", "Resume").AddCSS("font-weight", "bold").SetHelpText("Error, must write a resume description")
		form.NewElement("password", "password", "").SetLabel("Set new password").SetPlaceHolder("Set new password").SetHelpText("Use upper and lower case, numbers and special characters")
		form.NewElement("file", "pic", "Attach your photo")
		form.NewElement("checkbox", "legal", "").SetLabel(" Must read and accept Legal/Privacy")
		form.NewElement("hidden", "id", "1")

		form.NewButton("submit", "Update profile").AddClass("btn-danger").AddClass("btn-lg").AddClass("btn-xl").AddClass("btn-block")

		// Send to template
		res["Form"] = form
	}

In your HTML template place
`{{ .Form.Render }}`

### Requests or bugs?
<https://github.com/irob/goform/issues>

## Installation

	go get github.com/irob/goform

## Use you custom templates

	/templates/template_for_customize is a templete based on Bootstrap4.

	Step 1.-
	Move/Copy the folder /templates into your application, rename/copy the subdirectory /template_for_customize.

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
