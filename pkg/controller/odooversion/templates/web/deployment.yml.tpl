{{ define "componentName" }}web{{ end }}
{{ define "componentType" }}odoo{{ end }}

{{ define "deploymentArgs" }}["--config", "/run/configs/odoo/", "--db_maxconn=16", "--workers=0", "--max-cron-threads=0"]{{ end }}
{{ define "deploymentPorts" }}[{name: server-port, containerPort: 8069, protocol: TCP}]{{ end }}
{{ define "deploymentHealchecks" }}
livenessProbe:
  handler:
    exec:
      command: ["curl", "--connect-timeout", "5", "--max-time", "10", "-k", "-s", "http://localhost:8069"]
  initialDelaySeconds: 10
  timeoutSeconds: 10
  periodSeconds: 60
  failureThreshold: 3
  successThreshold: 1
readinessProbe:
  handler:
  	httpGet:
  	  port: 8069
  	  scheme: HTTP
  initialDelaySeconds: 10
  timeoutSeconds: 10
  periodSeconds: 60
  failureThreshold: 3
  successThreshold: 1
{{ end }}

{{ template "deployment" . }}