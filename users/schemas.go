package users

// ============================== Requests
type UserSummary struct{}

// ============================== Requests

// Register user request schema
type ReqRegister struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"min=8"`
}

// Update user request schema
type ReqUpdate struct {
}

// Change password request schema
type ReqChangePassword struct {
}

// Change email request schema
type ReqChangeEmail struct {
}

// ============================== Responses

// Register user response schema
type ResRegister struct {
	Uid string `json:"uid"`
}

type ResListUser = []UserSummary
