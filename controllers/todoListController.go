package controllers

import (
	"context"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"Todo_List/database"

	"Todo_List/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoListCollection *mongo.Collection = database.OpenCollection(database.Client, "todoList")

//TodoListPost is the api used to add a todo_list element
func TodoListPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var insert_todo_list models.TodoList
		
		if todo_list_array, exist := c.GetPostFormArray("todo_list"); exist && len(todo_list_array) != 0 {
			insert_todo_list.Todo_list = todo_list_array
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "todo list is empty"})
			return
		}

		insert_todo_list.ID = primitive.NewObjectID()

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "default todo list id not found"})
			return
		}
		_, err = TodoCollection.InsertOne(ctx, insert_todo_list)

		if err != nil {
            c.IndentedJSON(500, "Internal Server Error")
            fmt.Println("Your Todo has been listed")
            c.Abort()
        }
        c.IndentedJSON(200, todo)

    }
}

//TodoListGet is the api used to get a todo_list
func TodoListGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var insert_todo_list models.TodoList

		id := c.Query("id")
		fmt.Println(id)
		if id == "" {
            c.Header("Content-Type", "application/json")
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
            c.Abort()
 
        }
 
        fmt.Println("")
		ID1, err := primitive.ObjectIDFromHex(id)
        if err != nil {
            c.IndentedJSON(500, "Internal Server Error")
 
        }
        fmt.Println(ID1)
        err = TodoCollection.FindOne(ctx, bson.M{"_id": ID1}).Decode(&todo)

		if todo_list_array, exist := c.GetPostFormArray("todo_list"); exist && len(todo_list_array) != 0 {
			insert_todo_list.Todo_list = todo_list_array
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "todo list is empty"})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "default todo list id not found"})
			return
		}

		if err != nil {
            c.IndentedJSON(500, "Internal Server Error")
            fmt.Println("Here is your task")
            c.Abort()
        }
        c.IndentedJSON(http.StatusOK, insert_todo_list)
    }
}

//TodoListGet is the api used to get a todo_list
func TodoListUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
	var updateproduct models.TodoList

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := c.BindJSON(&updateproduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Query("id")
	if id == "" {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		c.Abort()

	}

	// fmt.Println(updateproduct.Title)

	ID1, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(500, "Unable to Hex encode")

	}

	// fmt.Println(ID1)
	// filter := bson.D{primitive.E{Key: "_id", Value: ID1}}
	filter := bson.M{"_id": ID1}
	update := bson.D{{"$set",
		bson.D{
			{"title", updateproduct.Title},
			{"completed", updateproduct.Completed},
		}}}

	_, err = TodoCollection.UpdateOne(ctx, filter, update)
	// fmt.Println(abc)
	if err != nil {
		c.IndentedJSON(500, err)
		// fmt.Println("errorhai")
		c.Abort()
	}

	c.IndentedJSON(200, "Successfully updated the todo list")
}
}

//TodoListDelete is the api used to delete a todo_list element
// func TodoListDelete() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		var delete_element string

// 		if element, exist := c.GetPostForm("delete_element"); exist && element != "" {
// 			delete_element = element
// 		} else {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "todo element to be deleted is empty"})
// 			return
// 		}

// 		find_result := todoListCollection.FindOne(ctx, bson.M{"_id": user_id})
// 		defer cancel()
// 		if find_result.Err() != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "user's todo list id not found"})
// 			return
// 		}

// 		resultInsertionNumber, insertErr := todoListCollection.UpdateOne(ctx, bson.M{"user_id": user_id}, bson.M{"$pull": bson.M{"todo_list": delete_element}})
// 		if insertErr != nil {
// 			msg := fmt.Sprintf("Todo list not found")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}
// 		c.JSON(http.StatusOK, )
// 	}
}
