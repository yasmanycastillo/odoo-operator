{{ define "componentName" }}secret{{ end }}
{{ define "componentType" }}app{{ end }}

apiVersion: v1
kind: Secret
{{ template "metadata" . }}
data:
  # adminpasswd only set through Secret Loaning
  pguser:  {{ .Instance.Spec.Database.User }}
  pgpassword:  {{ .Instance.Spec.Database.Password }}
