package user

type Role string

const (
    Default Role = "default"
    Elevated Role = "elevated"
)

type User interface {
    Init() error
}

type SimpleUser struct {
    Username string
    Password string
    UserRole Role 
}

func (su *SimpleUser) Init() error {
   return nil 
}

