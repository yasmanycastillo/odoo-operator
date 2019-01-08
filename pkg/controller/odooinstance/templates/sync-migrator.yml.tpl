{{ define "componentName" }}simple-migrator{{ end }}
{{ define "componentType" }}migrator{{ end }}
{{ define "jobArgs" }}[dodoo-migrator, "--config", "/run/configs/odoo/", --database", "{{ .Instance.Spec.Hostname }}", "--file", "/opt/odoo/.migration.yml" ]{{ end }}
{{ template "job" . }}
