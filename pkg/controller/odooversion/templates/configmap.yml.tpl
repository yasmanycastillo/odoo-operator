{{- define "componentName" }}config{{ end }}
{{- define "componentType" }}app{{ end }}
apiVersion: v1
kind: ConfigMap
{{- template "metadata" . -}}
data:
  app-config: text
