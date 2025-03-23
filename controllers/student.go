package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/therusetiawan/lks2025cc-backend/config"
    "github.com/therusetiawan/lks2025cc-backend/models"
)

func CreateStudent(c *gin.Context) {
    name := c.PostForm("name")
    class := c.PostForm("class")

    file, fileHeader, err := c.Request.FormFile("photo")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Photo upload failed"})
        return
    }
    defer file.Close()

    photoURL, err := config.UploadFileToS3(file, fileHeader)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload to S3"})
        return
    }

    student := models.Student{Name: name, Class: class, Photo: photoURL}
    config.DB.Create(&student)

    studentJSON, _ := json.Marshal(student)
    redisKey := "student:" + string(student.ID)
    config.RedisClient.Set(config.Ctx, redisKey, studentJSON, 0)

    c.JSON(http.StatusCreated, student)
}
