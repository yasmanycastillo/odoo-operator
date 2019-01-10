{{ define "metadata" }}
metadata:
  name: {{ .Instance.Name }}-{{ block "componentName" . }}{{ end }}
  namespace: {{ .Instance.Namespace }}
  labels:
    cluster.odoo.io/name: {{ .Instance.Spec.Cluster }}
    cluster.odoo.io/track: {{ .Instance.Spec.Track }}
    app.kubernetes.io/name: {{ block "componentName" . }}{{ end }}
    app.kubernetes.io/instance: {{ .Instance.Name }}-{{ block "componentName" . }}{{ end }}
    app.kubernetes.io/component: {{ block "componentType" . }}{{ end }}
    app.kubernetes.io/managed-by: odoo-operator
    app.kubernetes.io/part-of: {{ .Instance.Name }}
    app.kubernetes.io/version: {{if .Instance.Spec.Version }}{{ .Instance.Spec.Version }}{{ else }}n/a{{ end}}
{{ end }}