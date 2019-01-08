{{ define "ingressrule" }}
  - host: {{ .Spec.Hostname }}
    http:
      paths:
      - path: /longpolling
        backend:
          serviceName: {{ .Spec.Hostname }}-longpolling
          servicePort: 8072
      - path: /
        backend:
          serviceName: {{ .Spec.Hostname }}-web
          servicePort: 8069
{{ end }}