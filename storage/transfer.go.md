# Transfer


```go
/*
 * https://docs.cloud.google.com/storage/docs/samples/storage-transfer-manager-download-chunks-concurrently
 */

// downloadChunksConcurrently downloads a single file in chunks, concurrently in a process pool.
func downloadChunksConcurrently(w io.Writer, bucketName, blobName, filename string) error {
	// bucketName := "your-bucket-name"
	// blobName := "target-file"
	// filename := "path/to/your/local/file.txt"

	// The chunkSize is the size of each chunk to be downloaded.
	// The performance impact of this value depends on the use case.
	// For example, for a slow network, using a smaller chunkSize may be better.
	// Providing this parameter is optional and the default value is 32 MiB.
	chunkSize := 16 * 1024 * 1024 // 16 MiB

	// The maximum number of workers to use for the operation.
	// Please note, providing this parameter is optional.
	// The performance impact of this value depends on the use case.
	// To download one large file, the default value: NumCPU / 2 is usually fine.
	workers := 8

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %w", err)
	}
	defer client.Close()

	d, err := transfermanager.NewDownloader(client, transfermanager.WithPartSize(int64(chunkSize)), transfermanager.WithWorkers(workers))
	if err != nil {
		return fmt.Errorf("transfermanager.NewDownloader: %w", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}
	defer f.Close()

	in := &transfermanager.DownloadObjectInput{
		Bucket:      bucketName,
		Object:      blobName,
		Destination: f,
	}

	if err := d.DownloadObject(ctx, in); err != nil {
		return fmt.Errorf("d.DownloadObject: %w", err)
	}

	// Wait for all downloads to complete and close the downloader.
	// This allows to synchronize the download processes.
	results, err := d.WaitAndClose()
	if err != nil {
		return fmt.Errorf("d.WaitAndClose: %w", err)
	}

	// Process the downloader result.
	if len(results) != 1 {
		return fmt.Errorf("expected 1 result, got %d", len(results))
	}
	result := results[0]
	if result.Err != nil {
		fmt.Fprintf(w, "download of %v failed with error %v\n", result.Object, result.Err)
		return result.Err
	}
	fmt.Fprintf(w, "Downloaded %v to %v.\n", blobName, filename)

	return nil
}

```
