package soal

type requestJawaban struct {
	KeyJawaban  string `json:"key_jawaban" form:"key_jawaban"`
	Content     string `json:"content" form:"content"`
	ContentType string `json:"content_type" form:"content_type"`
}
type requestSoal struct {
	UserId       string           `json:"user_id" form:"user_id"`
	Title        string           `json:"title" form:"title"`
	TypeSoal     string           `json:"type_soal" form:"type_soal"`
	Image        string           `json:"image" form:"image"`
	KunciJawaban string           `json:"kunci_jawaban" form:"kunci_jawaban"`
	Answers      []requestJawaban `json:"answers" form:"answers"`
}
