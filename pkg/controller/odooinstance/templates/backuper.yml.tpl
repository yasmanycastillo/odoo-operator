{{ define "componentName" }}backup-cron{{ end }}
{{ define "componentType" }}backuper{{ end }}

{{ define "jobArgs" }}[dodoo-initializer, "--config", "/run/configs/odoo/", "--from-database", "{{ .Extra.FromDatabase }}", "--new-database", "{{ .Instance.Spec.Hostname }}" {{ if .Instance.Spec.InitModules }}, "--modules", "{{ range $index, $element := .Instance.Spec.InitModules }}{{ if $index }},{{ end }}{{ $element }}{{ end }}"{{ end }}]{{ end }}

{{ template "cronjob" . }}
