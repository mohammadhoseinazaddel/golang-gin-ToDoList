FROM golang:1.15.6

RUN mkdir /ToDoList

WORKDIR /ToDoList

COPY . /ToDoList/

RUN go build -o todo_list

EXPOSE 80

ENTRYPOINT ["./todo_list", "-d", "-p", "8080:80" ]
