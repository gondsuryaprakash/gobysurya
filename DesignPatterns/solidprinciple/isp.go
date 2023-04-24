package solidprinciple

/*
The Interface Segregation Principle states that clients should not be forced to
depend on interfaces they do not use.
In other words, we should use multiple smaller, specialized
interfaces instead of a single, monolithic interface.
*/

type CloudStorage interface {
	UploadFile(file string) error
	DeleteFile(file string) error
	DownloadFile(file string) error
	GetFileMetadata(file string) (Metadata, error)
}

type Metadata struct {
	Size     int64
	Created  int64
	Modified int64
}

/*
If a client only needs to upload and download files,
they are forced to depend on methods they do not use (GetFileMetadata and DeleteFile).
To adhere to the ISP, we can refactor the code:
*/

type FileUploader interface {
	UploadFile(file string) error
}

type FileDownloader interface {
	DownloadFile(file string) error
}

type MetadataRetriever interface {
	GetFileMetadata(file string) (Metadata, error)
}

type FileDeleter interface {
	DeleteFile(file string) error
}
