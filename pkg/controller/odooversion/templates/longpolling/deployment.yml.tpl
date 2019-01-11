{{- define "componentName" }}longpolling{{ end }}
{{- define "componentType" }}app{{ end }}
{{- define "deploymentArgs" -}}
        - gevent
        - --config
        - /run/configs/odoo/
        - --db_maxconn=16
{{- end -}}
{{- define "deploymentPorts" }}[{name: lp-port, containerPort: 8072, protocol: TCP}]{{ end }}
{{- define "deploymentHealchecks" -}}
        livenessProbe:
          exec:
            command: ["curl", "--connect-timeout", "5", "--max-time", "10", "-k", "-s", "http://localhost:8072"]
          initialDelaySeconds: 10
          timeoutSeconds: 10
          periodSeconds: 60
          failureThreshold: 3
          successThreshold: 1
        readinessProbe:
          httpGet:
            port: 8072
            scheme: HTTP
          initialDelaySeconds: 10
          timeoutSeconds: 10
          periodSeconds: 60
          failureThreshold: 3
          successThreshold: 1
{{- end -}}
{{- template "deployment" . -}}