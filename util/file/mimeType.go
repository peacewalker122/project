package file

type MIME string

var (
	MIME_JPG  MIME = "image/jpg"
	MIME_JPEG MIME = "image/jpeg"
	MIME_GIF  MIME = "image/gif"
	MIME_PNG  MIME = "image/png"
	MIME_WEBP MIME = "image/webp"
	MIME_MP4  MIME = "video/mp4"

	MIMEAUTH = []MIME{
		MIME_JPEG,
		MIME_JPG,
		MIME_GIF,
		MIME_PNG,
		MIME_WEBP,
		MIME_MP4,
	}
)

func (m MIME) String() string {
	return string(m)
}
