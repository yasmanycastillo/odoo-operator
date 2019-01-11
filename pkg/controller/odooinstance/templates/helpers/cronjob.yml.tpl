{{- define "cronjob" }}
apiVersion: v1beta1
kind: CronJob
{{- template "metadata" . -}}
spec:
  schedule:
  concurrencyPolicy: Forbid
  jobTemplate:
{{- template "metadata" . | indent 4 -}}
{{- template "jobspec" . | indent 4 -}}
{{ end -}}
