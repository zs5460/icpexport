package main

type item struct {
	ID      int    `json:"id,omitempty"`
	CateID  int    `json:"cateid"`
	Title   string `json:"title,omitempty"`
	Source  string `json:"source,omitempty"`
	Editor  string `json:"editor,omitempty"`
	Addtime string `json:"addtime,omitempty"`
	Image   string `json:"image"`
	ImageSM string `json:"imagesm"`
	Resume  string `json:"resume"`
	Content string `json:"content,omitempty"`
}

type category struct {
	CateID   int    `json:"cateid,omitempty"`
	CateName string `json:"catename,omitempty"`
	Sort     int    `json:"sort,omitempty"`
}

type topic struct {
	CateID   int    `json:"cateid,omitempty"`
	CateName string `json:"catename,omitempty"`
	Sort     int    `json:"sort,omitempty"`
}

type result struct {
	Base  topic      `json:"base,omitempty"`
	Cates []category `json:"cates,omitempty"`
	Items []item     `json:"items,omitempty"`
}
