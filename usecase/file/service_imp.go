package file

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"hashing-file/constants"
	"hashing-file/domain/entity"
	"hashing-file/domain/repository"
	"hashing-file/helpers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type service struct {
	fileRepository repository.FileRepository
}

func NewService(file repository.FileRepository) FileService {
	return &service{fileRepository: file}
}

func (s *service) UploadFile(req *multipart.FileHeader) (res DefaultResponse, err error) {
	var (
		file multipart.File
	)
	err = helpers.ValidateFileHeader(req)
	if err != nil {
		return
	}
	fmt.Printf("Uploaded File: %+v\n", req.Filename)
	fmt.Printf("File Size: %+v\n", req.Size)
	fmt.Printf("MIME Header: %+v\n", req.Header)
	filenameSplit := strings.Split(req.Filename, ".")
	if filenameSplit[len(filenameSplit)-1] != "txt" {
		constants.Extensions = filenameSplit[len(filenameSplit)-1]
	}
	constants.Key = req.Filename
	file, err = req.Open()
	if err != nil {
		res = DefaultResponse{
			Status:  constants.STATUS_FAILED,
			Message: constants.MESSAGE_FAILED,
			Data:    struct{}{},
		}
		err = errors.New("Can't Open File")
		return
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	extension := strings.Split(constants.Key, ".")
	if extension[1] == "txt" {
		constants.Key = fmt.Sprintf("%s.%s", extension[0], constants.Extensions)
		f, errCreate := os.Create(constants.Key)
		if errCreate != nil {
			return res, errCreate
		}
		f.WriteString(string(fileBytes))
	} else {
		fUpload, errCreate := os.Create(constants.Key)
		if errCreate != nil {
			return res, errCreate
		}
		fUpload.Write(fileBytes)
	}
	constants.Data = fileBytes
	data := entity.File{
		UploadedFile: fmt.Sprint(fileBytes),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	upload, err := s.fileRepository.Upload(data)
	if err != nil {
		return
	}
	constants.EncryptID, constants.DecryptID = fmt.Sprint(upload.ID), fmt.Sprint(upload.ID)
	res = DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: struct {
			ID int64 `json:"id"`
		}{
			upload.ID,
		},
	}
	return
}

func (s *service) Encrypt(data []byte, id, passphrase string) (chiperText []byte, err error) {
	block, _ := aes.NewCipher([]byte(helpers.RC4Hash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		err = errors.New("Can't Create 128-bit chiper")
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		err = errors.New("Can't Read Full")
		return
	}
	io.ReadFull(rand.Reader, nonce)
	chiperText = gcm.Seal(nonce, nonce, data, nil)
	newId, err := strconv.Atoi(id)
	if err != nil {
		err = errors.New("Invalid ID")
		return
	}
	entity := entity.File{
		ID:            int64(newId),
		EncryptedFile: fmt.Sprint(chiperText),
	}
	_, err = s.fileRepository.Encrypt(entity)
	if err != nil {
		return
	}
	return
}

func (s *service) Decrypt(data []byte, id, passphrase string) (plainText []byte, err error) {
	key := []byte(helpers.RC4Hash(passphrase))
	block, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		err = errors.New("Can't Create 128-bit chiper")
		return
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plainText, _ = gcm.Open(nil, nonce, ciphertext, nil)
	newId, err := strconv.Atoi(id)
	if err != nil {
		err = errors.New("Invalid ID")
		return
	}
	entity := entity.File{
		ID:            int64(newId),
		DecryptedFile: fmt.Sprint(plainText),
	}
	_, err = s.fileRepository.Decrypt(entity)
	if err != nil {
		return
	}
	return
}

func (s *service) EncryptFile(id, filename, passphrase string, data []byte) (res DefaultResponse, err error) {
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	encrypt, err := s.Encrypt(data, id, passphrase)
	if err != nil {
		return
	}
	_, err = f.WriteString(string(encrypt))
	if err != nil {
		return
	}
	res = DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    struct{}{},
	}
	return
}

func (s *service) DecryptFile(id, filename, passphrase string) (res DefaultResponse, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	file, err := os.Create(fmt.Sprintf("decrypt-%s", filename))
	if err != nil {
		return
	}
	defer file.Close()
	decrypt, err := s.Decrypt(data, id, passphrase)
	if err != nil {
		return
	}
	n, err := file.WriteString(string(decrypt))
	if err != nil {
		return
	}
	if n == 0 {
		file.WriteString(string(data))
	}
	res = DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    struct{}{},
	}
	return
}

func (s *service) GetAllDocument() (res DefaultResponse, err error) {
	doc, err := s.fileRepository.GetAllDocument()
	if err != nil {
		return
	}
	res = DefaultResponse{
		Status:  constants.STATUS_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data:    doc,
	}
	return
}
