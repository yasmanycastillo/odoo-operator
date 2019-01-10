{{ define "metadata" }}
metadata:
  name: {{ .Instance.Name }}.{{ block "componentType" . }}{{ end }}.{{ block "componentName" . }}{{ end }}
  namespace: {{ .Instance.Namespace }}
{{ template "metadatalabels" . | 2 }}
{{ end }}