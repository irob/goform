// Copyright 2019 by Roberto Morales Olivares. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 1.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package goform

//=============================================================================

nInputs := 8

// CitiesList slice of cities
var CitiesList = []OptionItem{{"", "Choose your favorite city"}, {"AMS", "Amsterdam"}, {"VEN", "Venice"}, {"KYO", "Kyoto"}, {"PAR", "Paris"}, {"DOH", "Doha"}, {"BAR", "Barcelona"}, {"SMA", "San Miguel de Allende"}, {"BUD", "Budapest"}, {"LIS", "Lisbon"}, {"FLO", "Florence"}, {"HNK", "Hong Kong"}, {"BRU", "Bruges"}}
// AgeRanges slice of ranges of ages
var AgeRanges = []OptionItem{{"1", "1 - 9 yo"}, {"2", "10 - 19 yo"}, {"3", "20 - 29 yo"}, {"4", "30 - 39 yo"}, {"5", "40 - 49 yo"}, {"6", ">= 50 yo"}}

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


tmpl.ExecuteTemplate(w, "FormHTML", form)