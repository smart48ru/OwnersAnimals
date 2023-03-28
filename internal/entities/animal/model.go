package animal

type Animal struct {
	ID       int    `json:"id,omitempty"`
	NickName string `json:"nick_name,omitempty"`
	Type     string `json:"type,omitempty"`
	Age      int    `json:"age,omitempty"`
	Color    string `json:"color,omitempty"`
}
