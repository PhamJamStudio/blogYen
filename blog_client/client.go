package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/PhamJamStudio/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Blag client")
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	// createBlog(c)
	// readBlog(c)
	// updateBlog(c)
	// deleteBlog(c)
	listBlog(c)

}

func createBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Create blog request sent")
	blog := &blogpb.Blog{
		AuthorId: "Yen",
		Title:    "My first derp blog",
		Content:  "Content for first blog",
	}

	fmt.Println("Creating the blog")
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	fmt.Printf("Blog has been created %v\n", createBlogRes)
}

func readBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Read blog request sent")
	// ReadBlog attempt expected to fail
	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "derp"})
	if err != nil {
		fmt.Printf("ERROR occured while reading: %v\n", err)
		return
	}

	// ReadBlog request expected to pass
	readBlogReq := &blogpb.ReadBlogRequest{
		BlogId: "000000000000000000000000"}
	res, err = c.ReadBlog(context.Background(), readBlogReq)
	if err != nil {
		fmt.Printf("ERROR occured while reading: %v\n", err)
		return
	}
	fmt.Printf("Blog successfuly read: %v\n", res)
}

func updateBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Update blog request sent")

	blog := &blogpb.Blog{
		Id:       "000000000000000000000000",
		AuthorId: "Yenmo",
		Title:    "My first derp blog",
		Content:  "Content for first blog",
	}
	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		fmt.Printf("ERROR occured while updating: %v\n", err)
		return
	}

	fmt.Printf("Blog successfuly updated: %v\n", res)
}

func deleteBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Delete blog request sent")
	// ReadBlog attempt expected to fail
	res, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: "derp"})
	if err != nil {
		fmt.Printf("ERROR occured while trying to delete: %v\n", err)
	}

	// ReadBlog request expected to pass
	deleteBlogReq := &blogpb.DeleteBlogRequest{
		BlogId: "5f23a6e908aaeec706411926"}
	res, err = c.DeleteBlog(context.Background(), deleteBlogReq)
	if err != nil {
		fmt.Printf("ERROR occured while reading: %v\n", err)
		return
	}

	fmt.Printf("Blog successfuly deleted: %v\n", res)

}

func listBlog(c blogpb.BlogServiceClient) {
	fmt.Println("List blog request sent")
	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})

	if err != nil {
		log.Fatalf("ERROR occured during MongoDB list request: %v\n", err)
		return
	}
	fmt.Println("Printing out blog posts:")
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ERROR during list streaming: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
