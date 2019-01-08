{{ define "componentName" }}config{{ end }}
{{ define "componentType" }}odoo{{ end }}
apiVersion: v1
kind: ConfigMap
{{ template "metadata" . }}
data:
  app-config: |
    {{ .Extra.ConfigFile }}