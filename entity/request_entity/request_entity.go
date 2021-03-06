package request_entity

type RelationActionRequest struct {
	UserId     int
	Token      string
	ToUserId   int
	ActionType string
}

type Video struct {
	UserId     int
	Token      string
	VideoId    int
	ActionType int
}

type Comment struct {
	UserId      int
	Token       string
	VideoId     int
	ActionType  int
	ContentText string
	ContentId   int
}
