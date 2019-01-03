{{ define "componentName" }}initializer{{ end }}
{{ define "componentType" }}initializer{{ end }}
{{ define "command" }}[dodoo-initializer, "--config", "/run/configs/odoo/", "--new-database", "{{ .Instance.Spec.HostName }}" {{ if .Instance.Spec.Modules }}, "--modules", "{{ range $index, $element := .Instance.Spec.Modules }}{{ if $index }},{{ end }}{{ $element }}{{ end }}"{{ end }}]{{ end }}
{{ template "job" . }}
