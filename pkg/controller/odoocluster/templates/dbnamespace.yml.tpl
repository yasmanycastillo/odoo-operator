{{- define "componentName" }}namespace{{ end }}
{{- define "componentType" }}database{{ end }}
apiVersion: cluster.odoo.io/v1beta1
kind: DBNamespace
{{- template "metadata" . -}}
spec:
  user: {{ .Instance.Spec.Database.User }}
  password: {{ .Instance.Spec.Database.Password }}
  dbAdmin:
    host: {{ .Instance.Spec.Database.Admin.Host }}
    port: {{ .Instance.Spec.Database.Admin.Port }}
    user: {{ .Instance.Spec.Database.Admin.User }}
    password: {{ .Instance.Spec.Database.Admin.Password }}