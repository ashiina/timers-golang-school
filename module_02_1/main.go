package main

import "fmt"

type MediaStorage interface {
	Store(path string)
	GetStoredURL() string
	GetBaseURL() string
}

type PhotoStorage struct {
	baseURL string
	path    string
}

func (p *PhotoStorage) Store(path string) {
	p.path = path + ".jpg"
}

func (p PhotoStorage) GetBaseURL() string {
	return p.baseURL
}

func (p PhotoStorage) GetStoredURL() string {
	return p.baseURL + p.path
}

type VideoStorage struct {
	baseURL     string
	path        string
	lengthLimit int
}

func (v *VideoStorage) Store(path string) {
	v.path = path + ".mp4"
}

func (v VideoStorage) GetBaseURL() string {
	return v.baseURL
}

func (v VideoStorage) GetStoredURL() string {
	return v.baseURL + v.path
}

func (v *VideoStorage) SetLengthLimit(l int) {
	v.lengthLimit = l
}

func main() {
	ps := &PhotoStorage{baseURL: "example.example/image"}
	vs := &VideoStorage{baseURL: "example.example/video", lengthLimit: 100}

	fmt.Println("Video length limit:\t", vs.lengthLimit)
	vs.SetLengthLimit(200)
	fmt.Println("Video length limit:\t", vs.lengthLimit)

	medias := []MediaStorage{ps, vs}
	for _, m := range medias {
		fmt.Println("Base URL:\t", m.GetBaseURL())
		m.Store("/somefile")
		fmt.Println("Stored URL:\t", m.GetStoredURL())
	}
}
