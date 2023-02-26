# TODO

<p align="center"><b><i>
	A Small TODO list Project
</i></b></p>

## install golang
```
sudo apt install golang-go
```
or use https://linguinecode.com/post/install-golang-linux-terminal as guidance after downloading below package
```
wget https://dl.google.com/go/go1.20.1.linux-amd64.tar.gz
```


## run in docker
```
sudo docker build .
```

## localy
```
go build -o todo_list
```

## Abstract

This project is based on the following elements: Golang, Gin Framework, MongoDB,. We use Postman to test all the API functionality.

## Features

1. users/todo_list: Post method

    Sample Input

    ```JSON
    {
      "todo_list": ["do homework", "write diary"]
    }
    ```

    Sample Output

    ```JSON
    {
      "MatchedCount": 1,
      "ModifiedCount": 1,
      "UpsertedCount": 0,
      "UpsertedID": null
    }
    ```

2. users/todo_list: Get method

    Sample Output

    ```JSON
    {
      "gym",
      "diary"
    }
    ```

    > Note that this method does not have input.

    ```
3. users/todo_list: Put update
