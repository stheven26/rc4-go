package file

import "mime/multipart"

type FileService interface {
	Encrypt(data []byte, id, passphrase string) ([]byte, error)
	Decrypt(data []byte, id, passphrase string) ([]byte, error)
	EncryptFile(id, filename, passphrase string, data []byte) (res DefaultResponse, err error)
	DecryptFile(id, filename, passphrase string) (res DefaultResponse, err error)
	UploadFile(req *multipart.FileHeader) (res DefaultResponse, err error)
	GetAllDocument() (res DefaultResponse, err error)
}
