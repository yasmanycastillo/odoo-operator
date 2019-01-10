{{ define "metadatalabels" }}
labels:
  cluster.odoo.io/name: {{ .Instance.Spec.Cluster }}
  cluster.odoo.io/track: {{ .Extra.Track }}
  app.kubernetes.io/name: {{ block "componentName" . }}{{ end }}
  app.kubernetes.io/instance: {{ .Instance.Name }}-{{ block "componentName" . }}{{ end }}
  app.kubernetes.io/component: {{ block "componentType" . }}{{ end }}
  app.kubernetes.io/managed-by: odoo-operator
  app.kubernetes.io/part-of: {{ .Instance.Name }}
  app.kubernetes.io/version: {{ .Instance.Spec.Version }}
{{ end }}