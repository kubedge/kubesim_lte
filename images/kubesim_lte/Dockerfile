FROM scratch
ADD kubesim_lte/ca-certificates.crt /etc/ssl/certs/
ADD kubesim_lte/goserv /goserv
EXPOSE 8080
CMD ["/goserv"]
