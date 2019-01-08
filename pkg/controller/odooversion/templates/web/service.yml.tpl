{{ define "componentName" }}web{{ end }}
{{ define "componentType" }}odoo{{ end }}

{{ define "servicePorts" }}[{name: server-port, protocol: TCP, port: 8072, targetPort: server-port}]{{ end }}
{{ template "service" . }}