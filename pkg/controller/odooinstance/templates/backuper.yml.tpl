{{- define "componentName" }}cron{{ end }}
{{- define "componentType" }}backup{{ end }}
{{- define "jobArgs" -}}
        - dodoo-backuper
        - --config
        - /run/configs/odoo/
        - --database
        - {{ .Instance.Spec.Hostname }}
{{- end -}}
{{- template "cronjob" . -}}
