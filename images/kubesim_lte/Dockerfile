FROM scratch
ADD kubesim_lte/ca-certificates.crt /etc/ssl/certs/
ADD kubesim_lte/goclient /goclient
EXPOSE 8080
CMD ["/goclient"]
