FROM ubuntu:20.10
RUN mkdir /app
COPY webappdemo /app/
WORKDIR /app
CMD ["/app/webappdemo"]
EXPOSE 80