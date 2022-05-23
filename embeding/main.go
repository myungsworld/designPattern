package main

type crud interface {
	Create()
	Read()
	Update()
	Delete()
}

type blog struct {
}

func (b *blog) Create() {
	// 생성
}

func (b *blog) Read() {
	// 읽기
}

func (b *blog) Update() {
	// 수정
}

func (b *blog) Delete() {
	// 삭제
}

type tistoryBlog struct {
	blog
}

func main() {
	tblog := tistoryBlog{}
	tblog.Create()
	tblog.Read()
	tblog.Update()
	tblog.Delete()
}
