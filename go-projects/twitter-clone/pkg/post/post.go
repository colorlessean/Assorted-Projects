package post

type PostType int64

const (
	Default PostType = 0
	Repost  PostType = 1
	Reply   PostType = 2
)

func (pt PostType) String() string {
	if pt == 0 {
		return "default"
	} else if pt == 1 {
		return "repost"
	} else if pt == 2 {
		return "reply"
	} else {
		return ""
	}
}

type Post struct {
	postID   int64
	postType PostType
	username string
	content  string
}


