package models

type OAuthType int

const (
	GITHUB OAuthType = iota + 1
	GOOGLE
	TWITTER
	GENERIC
	SMARTEMS_COM
	GITLAB
)
