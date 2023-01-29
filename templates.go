package goform

import (
	"html/template"
	"log"
)

var themes = map[string]map[string]string{}

func init() {

	// Initialisize maps
	themes["html"] = make(map[string]string)
	themes["bootstrap5"] = make(map[string]string)

	// HTML plain inputs

	themes["html"]["form"] = `<form{{if .Name}} name="{{.Name}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}} method="{{ .Method }}" action="{{ .Action }}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}: {{$v}}; {{end}}"{{end}}{{ if eq .MultipartFormData "enabled" }} enctype="multipart/form-data"{{end}}{{ if .Classes }} class="{{range .Classes}}{{.}} {{end}}"{{end}}>
			{{ .RenderElements }}
	</form>`

	themes["html"]["label"] = `
	<label{{if .Name}} name="{{.Name}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}} class="control-label{{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}">{{.Value}}</label>`

	themes["html"]["text"] = `<input type="text" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="{{ if .Classes }}{{range .Classes}}{{.}} {{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}{{ if .Value}} value="{{.Value}}"{{end}}{{if .PlaceHolder}} placeholder="{{.PlaceHolder}}"{{end}}>`

	themes["html"]["password"] = `<input type="password" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="{{ if .Classes }}{{range .Classes}}{{.}} {{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}{{if .PlaceHolder}} placeholder="{{.PlaceHolder}}"{{end}}>`

	themes["html"]["select"] = `<select name="{{.Name}}" class="{{ if .Classes }}{{range .Classes}}{{.}} {{end}}{{end}}"{{if .ID}} id="{{.ID}}"{{end}}{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}: {{$v}}; {{end}}"{{end}}>
	{{ $p := . }}
	{{range $option := .Options}}
	<option value="{{$option.Key}}"{{ if eq $p.Value $option.Key}} selected{{end}}>{{$option.Value}}</option>
	{{end}}
	</select>`

	themes["html"]["radio"] = `{{ $p := . }}
	{{range $option := .Options}}
	<input type="radio" value="{{$option.Key}}"{{ if eq $p.Value $option.Key}} selected{{end}}> {{$option.Value}}<br />
	{{end}}`

	themes["html"]["textarea"] = `<textarea name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}: {{$v}}; {{end}}"{{end}}{{if .PlaceHolder}} id="{{.PlaceHolder}}"{{end}} rows="6">{{.Value}}</textarea>`

	themes["html"]["checkbox"] = `<input type="checkbox" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} value="{{.Value}}" class="{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}>`

	themes["html"]["file"] = `<input type="file" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="{{ if .Classes }}{{range .Classes}}{{.}} {{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}{{if .PlaceHolder}} placeholder="{{.PlaceHolder}}"{{end}}>`

	themes["html"]["hidden"] = `<input type="hidden" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} {{ if .Value}} value="{{.Value}}"{{end}}>`

	themes["html"]["button"] = `<button type="button" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="{{ if .Classes }}{{range .Classes}}{{.}} {{end}}{{end}}">{{.Value}}</button>`

	themes["html"]["submit"] = `<button type="submit" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="{{ if .Classes }}{{range .Classes}}{{.}} {{end}}{{end}}">{{.Value}}</button>`

	themes["html"]["row"] = `<br />`

	// Bootstrap5 inputs

	themes["bootstrap5"]["form"] = `
	<form{{if .Name}} name="{{.Name}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}} method="{{ .Method }}" action="{{ .Action }}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}: {{$v}}; {{end}}"{{end}}{{ if eq .MultipartFormData "enabled" }} enctype="multipart/form-data"{{end}}{{ if .Classes }} class="{{range .Classes}} {{.}}{{end}}"{{end}}>
	<div class="row" name="row_main" id="row_main">
		{{ .RenderElements }}
	</div>
	</form>`

	themes["bootstrap5"]["label"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	<label{{if .ID}} name="{{.ID}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}} class="control-label{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}>{{.Value}}</label>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["textlabel"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}row">
	<label{{if .ID}} name="{{.ID}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}} class="col-sm-2 col-form-label{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}>{{.Label}}</label>
	<label class="col-sm-10">
		<input type="text" readonly class="form-control-plaintext" {{if .ID}} name="static_{{.ID}}"{{end}}{{if .ID}} id="static_{{.ID}}"{{end}} value="{{.Value}}">
	</label>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["text"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{ if .Label }}<label class="control-label{{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Label}}</label>{{end}}
	<input type="text" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="form-control{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}{{ if .Value}} value="{{.Value}}"{{end}}{{if .Params}}{{range $k, $v := .Params}} {{$k}}="{{$v}}"{{end}}{{end}}{{if .PlaceHolder}} placeholder="{{.PlaceHolder}}"{{end}}>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["password"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{ if .Label }}<label class="control-label{{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Label}}</label>{{end}}
	<input type="password" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="form-control{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}{{if .PlaceHolder}} placeholder="{{.PlaceHolder}}"{{end}}>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["select"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{if .Label}}<label class="control-label{{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Label}}</label>{{end}}
	<select name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="form-control{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}: {{$v}}; {{end}}"{{end}}>
	{{ $p := . }}{{range $option := .Options}}
	<option value="{{$option.Key}}"{{ if eq $p.Value $option.Key}} selected{{end}}>{{$option.Value}}</option>{{end}}
	</select>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["radio"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{if .Label}}<label class="control-label{{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Label}}</label>{{end}}
	{{ $p := . }}
	{{range $option := .Options}}
	<div class="form-check">
	<input class="form-check-input" type="radio" name="{{$p.Name}}" id="{{$p.ID}}_{{$option.Key}}" value="{{$option.Key}}"{{ if eq $p.Value $option.Key}} checked{{end}}>
	<label class="form-check-label" for="{{$p.ID}}_{{$option.Key}}">
	{{$option.Value}}
	</label>
	</div>
	{{end}}
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["textarea"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{ if .Label }}<label class="control-label{{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Label}}</label>{{end}}
	<textarea name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="form-control{{ if .Classes }} {{range .Classes}}{{.}} {{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}: {{$v}}; {{end}}"{{end}}{{if .PlaceHolder}} id="{{.PlaceHolder}}"{{end}} rows="6">{{.Value}}</textarea>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["checkbox"] = `
	<div id="group_{{.Name}}" class="form-check{{ if .GroupClass }}{{range .GroupClass}} {{.}}{{end}}{{end}}">
	<input type="checkbox" name="{{.Name}}"{{ if .ID }}{{if .ID}} id="{{.ID}}"{{end}} value="{{.Value}}" class="{{range .Classes}} {{.}}{{end}}"{{end}}{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}>
	{{ if .Label }}
	<label class="form-check-label" for="{{.Name}}">
	{{.Label}}
	</label>
	{{end}}
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["file"] = `
	<div id="group_{{.Name}}" name="group_{{.Name}}"{{if .ID}} id="group_{{.ID}}"{{end}} class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{ if .Value }}<label class="control-label{{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Value}}</label>{{end}}
	<div class="custom-file">
	<input type="file" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="{{ if .Classes }}{{range .Classes}}{{.}} {{end}}{{end}}"{{if .CSS}} style="{{range $k, $v := .CSS}}{{$k}}:{{$v}};{{end}}"{{end}}{{if .PlaceHolder}} placeholder="{{.PlaceHolder}}"{{end}}>
	{{ if .Label }}<label class="custom-file-label" for="{{.Name}}">{{.Label}}</label>{{end}}
	</div>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["hidden"] = `
	<input type="hidden" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} {{ if .Value}} value="{{.Value}}"{{end}}>`

	themes["bootstrap5"]["button"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{ if .Label }}<label class="control-label {{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Label}}</label>{{end}}
	<button type="button" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="btn{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}">{{.Value}}</button>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["submit"] = `
	<div id="group_{{.Name}}" class="{{ if .GroupClass }}{{range .GroupClass}}{{.}} {{end}}{{end}}">
	{{ if .Label }}<label class="control-label {{ if .LabelClass }}{{range .LabelClass}} {{.}}{{end}}{{end}}"{{if .ID}} for="{{.ID}}"{{end}}>{{.Label}}</label>{{end}}
	<button type="submit" name="{{.Name}}"{{if .ID}} id="{{.ID}}"{{end}} class="btn{{ if .Classes }}{{range .Classes}} {{.}}{{end}}{{end}}">{{.Value}}</button>
	{{ if .HelpText }}<small id="{{.Name}}Help" class="form-text text-muted">{{.HelpText}}</small>{{end}}
	</div>`

	themes["bootstrap5"]["row"] = `
	</div>
	<div class="row" name="row_{{.Name}}"{{if .ID}} id="row_{{.ID}}"{{end}}>`

}

// HTMLTemplate parse html template
func HTMLTemplate(theme string, input string) *template.Template {

	t, err := template.New("tmpl").Parse(themes[theme][input])
	if err != nil {
		log.Println(err)
	}

	return t
}
