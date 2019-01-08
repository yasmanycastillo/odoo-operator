{{ define "service" }}
kind: Service
apiVersion: v1
{{ template "metadata" . }}
spec:
  selector:
    app.kubernetes.io/name: {{ block "componentName" . }}{{ end }}
    app.kubernetes.io/instance: {{ .Instance.Name }}-{{ block "componentName" . }}{{ end }}
  ports: {{ block "servicePorts" . }}{{ end }}
{{ end }}