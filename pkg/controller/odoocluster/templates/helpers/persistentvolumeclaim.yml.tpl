{{ define "pvc" }}
apiVersion: v1
kind: PersistentVolumeClaim
{{ template "metadata" . }}
spec: {{ .Extra.VolumeSpec }}
{{ end }}
