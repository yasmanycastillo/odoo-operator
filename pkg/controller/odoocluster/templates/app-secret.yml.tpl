{{ define "componentName" }}secret{{ end }}
{{ define "componentType" }}odoo{{ end }}

apiVersion: v1
kind: Secret
{{ template "metadata" . }}
data:
  # adminpasswd only set through Secret Loaning
  pguser:  {{ .Instance.Spec.Database.User }}
  pgpassword:  {{ .Instance.Spec.Database.Password }}
