package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func initClient() (*minio.Client, error){

	endpoint := "localhost:9090"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return minioClient, nil
}

func makeBucket(ctx context.Context, minioClient *minio.Client, bucketName string) {
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: ""})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func uploadFile(ctx context.Context, minioClient *minio.Client, bucketName, objectName, filePath string) {
	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	fmt.Println(info)
}

func getFile(ctx context.Context, minioClient *minio.Client, bucketName, objectName, filePath string) {
	object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	localFile, err := os.Create("go.sum.u")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}

func statFile(ctx context.Context, minioClient *minio.Client, bucketName, objectName, filePath string) {
	objInfo, err := minioClient.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(objInfo)
}

func list(ctx context.Context, minioClient *minio.Client, bucketName, objectName, filePath string) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix: filePath,
		Recursive: false,
	})
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object.Key)
	}
}

func remove(ctx context.Context, minioClient *minio.Client, bucketName, objectName, filePath string) {
	opts := minio.RemoveObjectOptions {
		GovernanceBypass: true,
	}
	err := minioClient.RemoveObject(context.Background(), bucketName, objectName, opts)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func removeDir(ctx context.Context, minioClient *minio.Client, bucketName, objectName, filePath string) {
	opts := minio.RemoveObjectOptions {
		GovernanceBypass: true,
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix: filePath,
		Recursive: false,
	})
	rc := make(chan object)
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		if object.Key == objectName {
			break
		}
		fmt.Println(object)
	}
	err := minioClient.RemoveObjects(context.Background(), bucketName, objectName, opts)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	ctx := context.Background()
	minioClient, err := initClient()
	if err != nil {
		log.Fatalln(err)
	}
	// Make a new bucket called mymusic.
	bucketName := "testbucket"
	//location := "us-east-1" region

	// Upload the zip file
	objectName := "go.mod"
	filePath := "go.mod"
	//contentType := "application/zip"
	list(ctx, minioClient, bucketName, objectName, filePath)
}
