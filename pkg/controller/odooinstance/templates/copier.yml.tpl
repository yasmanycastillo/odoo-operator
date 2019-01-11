{{- define "componentName" }}copier{{ end }}
{{- define "componentType" }}app{{ end }}
{{- define "jobArgs" -}}
        - dodoo-initializer
        - --config
        - /run/configs/odoo/
        - --from-database
        - {{ .Extra.FromDatabase }}
        - --new-database
        - {{ .Instance.Spec.Hostname }}
	{{- if .Instance.Spec.InitModules }}
        - --modules
        - {{ .Instance.Spec.InitModules | join "," }}
	{{- end -}}
{{- end -}}
{{- template "job" . -}}
