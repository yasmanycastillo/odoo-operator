{{- define "componentName" }}web{{ end }}
{{- define "componentType" }}app{{ end }}
{{- define "servicePorts" -}}
  - name: web
    protocol: TCP
    port: 80
    targetPort: 8069
{{- end -}}
{{- template "service" . -}}