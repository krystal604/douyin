package request_entity

type RelationActionRequest struct {
	UserId     int
	Token      string
	ToUserId   int
	ActionType string
}
