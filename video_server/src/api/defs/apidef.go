package defs

//request
type UserCredential struct {
	Username string 'json:"user_name"'
	Pwd string 'json:"pwd"'
}

//Data Model
type VideoInfo struct{
	Id string
	AuthorId int
	Name string
	DisplayCtimer string

}