{{ define "componentName" }}secret{{ end }}
{{ define "componentType" }}app{{ end }}

apiVersion: v1
kind: Secret
{{ template "metadata" . }}
data:
  pguser:  {{ .Instance.Spec.Database.User | b64enc }}
  pgpassword:  {{ .Instance.Spec.Database.Password | b64enc }}
