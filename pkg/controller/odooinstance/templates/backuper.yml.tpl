{{- define "componentName" }}cron{{ end }}
{{- define "componentType" }}backup{{ end }}
{{- define "jobArgs" -}}
        - dodoo-backuper
        - --config
        - /run/configs/odoo/
        - --database
        - {{ .Instance.Spec.Hostname }}
{{- end -}}
apiVersion: v1beta1
kind: CronJob
{{- template "metadata" . -}}
spec:
  schedule:
  concurrencyPolicy: Forbid
  jobTemplate:
{{- template "metadata" . | indent 4 -}}
{{- template "jobspec" . | indent 4 -}}
