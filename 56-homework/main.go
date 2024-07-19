package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	r := gin.Default()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	r.Use(func(c *gin.Context) {
		role := c.GetHeader("Role")
		c.Set("isAdmin", role == "admin")
		c.Next()
	})

	r.POST("/create/:key", func(c *gin.Context) {
		if !c.MustGet("isAdmin").(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
			return
		}

		key := c.Param("key")
		var json struct {
			Value string `json:"value"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := rdb.Set(ctx, key, json.Value, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Created"})
	})

	r.GET("/read/:key", func(c *gin.Context) {
		key := c.Param("key")

		value, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{key: value})
	})

	r.PUT("/update/:key", func(c *gin.Context) {
		if !c.MustGet("isAdmin").(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
			return
		}

		key := c.Param("key")
		var json struct {
			Value string `json:"value"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := rdb.Set(ctx, key, json.Value, 0).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Updated"})
	})

	r.DELETE("/delete/:key", func(c *gin.Context) {
		if !c.MustGet("isAdmin").(bool) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden"})
			return
		}

		key := c.Param("key")

		err := rdb.Del(ctx, key).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
	})

	r.Run(":8080")
}
