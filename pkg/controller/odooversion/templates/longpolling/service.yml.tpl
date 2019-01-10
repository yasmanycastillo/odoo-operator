{{ define "componentName" }}longpolling{{ end }}
{{ define "componentType" }}app{{ end }}

{{ define "servicePorts" }}[{name: longpolling-port, protocol: TCP, port: 8069, targetPort: longpolling-port}]{{ end }}
{{ template "service" . }}