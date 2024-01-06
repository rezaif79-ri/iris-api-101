package domain

import "mime/multipart"

type File struct {
	InputFile multipart.FileHeader `form:"input_file"`
	FileName  string               `form:"file_name"`
}
