{{ define "componentName" }}dbnamespace{{ end }}
{{ define "componentType" }}database{{ end }}

apiVersion: cluster.odoo.io/v1beta1
kind: DBNamespace
{{ template "metadata" . }}
spec:
  host: {{ .Instance.Spec.Database.Host }}
  port: {{ .Instance.Spec.Database.Port }}
  user: {{ .Instance.Spec.Database.User }}
  password: {{ .Instance.Spec.Database.Password }}
  dbAdmin:
    user: {{ .Instance.Spec.Database.Admin.User }}
    password: {{ .Instance.Spec.Database.Admin.Password }}