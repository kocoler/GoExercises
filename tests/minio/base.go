package main

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"os"
)

func initClient() (*minio.Client, error){

	endpoint := "localhost:9090"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := true

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
		fmt.Println("err", err)
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
		fmt.Println(object.ETag)
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
	opts := minio.RemoveObjectsOptions {
		GovernanceBypass: true,
	}
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	//objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
	//	Prefix: filePath,
	//	Recursive: false,
	//})
	//
	//keyName := ""
	//for object := range objectCh {
	//	if object.Err != nil {
	//		fmt.Println(object.Err)
	//		return
	//	}
	//	if path.Clean(object.Key) == objectName {
	//		keyName = object.Key
	//		break
	//	}
	//	fmt.Println(object)
	//}

	fmt.Println(filePath)
	findChan := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix: filePath,
		Recursive: true,
	})

	err := minioClient.RemoveObjects(ctx, bucketName, findChan, opts)
	if err != nil {
		fmt.Println((<- err).Err)
		return
	}
}

func put(ctx context.Context, minioClient *minio.Client, bucketName, objectName, filePath string) {
	file, err := os.Open("go.mod")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	uploadInfo, err := minioClient.PutObject(context.Background(), bucketName, objectName, file, fileStat.Size(), minio.PutObjectOptions{ContentType:"application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
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
	objectName := "www/tt"
	filePath := "www"
	//contentType := "application/zip"
	statFile(ctx, minioClient, bucketName, objectName, filePath)
}
