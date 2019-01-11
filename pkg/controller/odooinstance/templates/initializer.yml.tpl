{{- define "componentName" }}initializer{{ end }}
{{- define "componentType" }}app{{ end }}
{{- define "jobArgs" -}}
        - dodoo-initializer
        - --config
        - /run/configs/odoo/
        - --new-database
        - {{ .Instance.Spec.Hostname }}
	{{- if .Instance.Spec.InitModules }}
        - --modules
        - {{ .Instance.Spec.InitModules | join "," }}
	{{- end -}}
{{- end -}}
{{- template "job" . -}}
