{{- define "ingressrule" }}
  - host: {{ .Spec.Hostname }}
    http:
      paths:
      - path: /longpolling
        backend:
          serviceName: app-{{ .Spec.Version | replace "." "-" }}-longpolling
          servicePort: 8072
      - path: /
        backend:
          serviceName: app-{{ .Spec.Version | replace "." "-" }}-web
          servicePort: 8069
{{ end -}}