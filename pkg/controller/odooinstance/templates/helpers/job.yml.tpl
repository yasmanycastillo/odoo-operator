{{ define "job" }}
apiVersion: batch/v1
kind: Job
{{ template "metadata" . }}
{{ template "jobspec" . }}
{{ end }}