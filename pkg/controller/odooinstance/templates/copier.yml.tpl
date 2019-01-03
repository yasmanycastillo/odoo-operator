{{ define "componentName" }}copier{{ end }}
{{ define "componentType" }}copier{{ end }}
{{ define "command" }}[dodoo-initializer, "--config", "/run/configs/odoo/", "--from-database", "{{ .Extra.FromDatabase }}", "--new-database", "{{ .Instance.Spec.HostName }}" {{ if .Instance.Spec.Modules }}, "--modules", "{{ range $index, $element := .Instance.Spec.Modules }}{{ if $index }},{{ end }}{{ $element }}{{ end }}"{{ end }}]{{ end }}
{{ template "job" . }}
