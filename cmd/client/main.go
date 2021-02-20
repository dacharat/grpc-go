package client

func main() {
	// conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	// if err != nil {
	// 	panic(err)
	// }

	// client := proto.NewAddServiceClient(conn)
	// req := &proto.Request{A: int64(a), B: int64(b)}

	// if response, err := client.Add(ctx, req); err == nil {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"result": fmt.Sprint(response.Result),
	// 	})
	// } else {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }
}
