{{ define "job" }}
apiVersion: batch/v1
kind: Job
metadata:
  name: {{ .Instance.Name }}-{{ block "componentName" . }}{{ end }}
  namespace: {{ .Instance.Namespace }}
  labels:
    cluster.odoo.io/name: {{ .Extra.ClusterName }}
    app.kubernetes.io/name: {{ block "componentName" . }}{{ end }}
    app.kubernetes.io/instance: {{ .Instance.Name }}-{{ block "componentName" . }}{{ end }}
    app.kubernetes.io/version: {{ .Instance.Spec.Version }}
    app.kubernetes.io/component: {{ block "componentType" . }}{{ end }}
    app.kubernetes.io/part-of: {{ .Instance.Name }}
    app.kubernetes.io/managed-by: odoo-operator
spec:
  completions: 1
  backoffLimit: 1
  activeDeadlineSeconds: 360
  template:
    metadata:
      labels:
        cluster.odoo.io/name: {{ .Extra.ClusterName }}
        app.kubernetes.io/name: {{ block "componentName" . }}{{ end }}
        app.kubernetes.io/instance: {{ .Instance.Name }}-{{ block "componentName" . }}{{ end }}
        app.kubernetes.io/version: {{ .Instance.Spec.Version }}
        app.kubernetes.io/component: {{ block "componentType" . }}{{ end }}
        app.kubernetes.io/part-of: {{ .Instance.Name }}
        app.kubernetes.io/managed-by: odoo-operator
    spec:
      restartPolicy: Never
      imagePullSecrets:
      - name: pull-secret
      securityContext:
        runAsUser: 9001
        runAsNonRoot: true
        supplementalGroups: 2000
        fsGroup: 9001
      containers:
      - name: default
        image: {{ .Extra.Image }}:devops-{{ .Instance.Spec.Version }}
        imagePullPolicy: Always
        command: {{ block "command" . }}{{ end }}
        resources:
          requests:
            memory: 512M
            cpu: 200m
          limits:
            memory: 1G
            cpu: 500m
        env:
         - name: PGHOST
           value: {{ .Extra.Database.Host }}
         - name: PGUSER
           value: {{ .Extra.Database.User }}
         - name: PGPORT
           value: {{ .Extra.Database.Port }}
         - name: PGPASSWORD
           value: {{ .Extra.Database.Password }}
         - name: ODOO_RC
           value: /run/configs/odoo/
         - name: ODOO_PASSFILE
           value: /run/secrets/odoo/adminpwd
        volumeMounts:
        - name: data-volume
          mountPath: /mnt/odoo/data/
        - name: config-volume
          mountPath: /run/configs/odoo/
          readonly: true
        - name: app-secrets
          mountPath: /run/secrets/odoo/
          readonly: true

      volumes:
        - name: data-volume
          configMap:
            name: {{ .Extra.ClusterName }}-data-volume
        - name: config-volume
          configMap:
            name: {{ .Extra.ClusterName }}-{{ .Instance.Version }}-config
        - name: app-secrets
          secret:
            secretName: {{ .Extra.ClusterName }}-{{ .Instance.Version }}-secret
{{ end }}