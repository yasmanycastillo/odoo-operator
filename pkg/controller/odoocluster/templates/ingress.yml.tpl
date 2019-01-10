{{ define "componentName" }}lb/l7router{{ end }}
{{ define "componentType" }}networking{{ end }}

apiVersion: extensions/v1beta1
kind: Ingress
{{ template "metadata" . }}
  annotations:
    kubernetes.io/ingress.class: nginx
    kubernetes.io/tls-acme: "true"  # Still support kube-lego in addition to certmanager
    certmanager.k8s.io/cluster-issuer: letsencrypt-prod
spec:
  backend:
    serviceName: default-backend
    servicePort: default-backend-port
  rules:
{{ range _, $instance := .Extra.InstanceList.Items }}
{{ template "ingressrule" $instance }}
{{ end }}
