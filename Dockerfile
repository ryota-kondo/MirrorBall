FROM golang:onbuild

COPY ./main.go ./main.go

CMD ["go", "run", "main.go"]
