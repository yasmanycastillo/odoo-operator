{{ define "componentName" }}backup{{ end }}
{{ define "componentType" }}storage{{ end }}

{{ template "pvc" . }}
