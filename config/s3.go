package config

import (
    "bytes"
    "context"
    "fmt"
    "mime/multipart"
    "os"
    "time"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func ConnectS3() {
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(os.Getenv("AWS_REGION")),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
            os.Getenv("AWS_ACCESS_KEY"),
            os.Getenv("AWS_SECRET_KEY"),
            "",
        )),
    )
    if err != nil {
        panic("Failed to load AWS config")
    }
    S3Client = s3.NewFromConfig(cfg)
}

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
    bucket := os.Getenv("AWS_BUCKET_NAME")
    key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), fileHeader.Filename)

    buffer := bytes.NewBuffer(nil)
    _, err := buffer.ReadFrom(file)
    if err != nil {
        return "", err
    }

    _, err = S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: &bucket,
        Key:    &key,
        Body:   bytes.NewReader(buffer.Bytes()),
    })
    if err != nil {
        return "", err
    }

    url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key)
    return url, nil
}
