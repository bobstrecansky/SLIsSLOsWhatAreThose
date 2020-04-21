FROM golang
RUN go get -u github.com/rakyll/hey
ADD loadtest.sh .
RUN chmod +x ./loadtest.sh
ENTRYPOINT ["sh", "loadtest.sh"]
