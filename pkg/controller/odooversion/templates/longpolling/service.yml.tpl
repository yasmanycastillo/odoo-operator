{{- define "componentName" }}longpolling{{ end }}
{{- define "componentType" }}app{{ end }}
{{- define "servicePorts" -}}
  - name: longpolling
    protocol: TCP
    port: 80
    targetPort: 8072
{{- end -}}
{{- template "service" . -}}