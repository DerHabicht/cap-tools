package capa5

type Role int

const (
	Root Role = iota
	Command
	Admin
	Member
	Parent
)
