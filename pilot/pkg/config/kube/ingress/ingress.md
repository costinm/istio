# Ingress content

There are 3 main types of ingress.

## Host + paths:

```yaml
rules:
- http:
      paths:
      - path: /testpath
        backend:
          serviceName: test
          servicePort: 80
```

## Default + paths

```yaml
          
- host: ingress.v10.istio.webinf.info
  http:
      paths:
        - backend:
            serviceName: cm-acme-http-solver-mzpvw
            servicePort: 8089
          path: /.well-known/acme-challenge/zJX8KR6BSBCPufuCFh2_PSezbViS5YDVzbQcq7ioSfA
          
```

## Default for the ingress

```yaml
      
      backend:
          serviceName: test
          servicePort: 80

```

Equivalent with Host: '*', 'Path: /'

Must be used at the end, only one allowed. If VirtualService defines
one - drop Ingress.

# TLS settings

Gateway also defines TLS settings - since TLs setup involves secrets
and other manual operations, we will require that user configures a 
Gateway and ignore the TLS settings in Ingress.

# Troubleshoting

1. /debug/configz 
2. Look for 