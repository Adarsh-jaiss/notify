package main

import "time"

type GithubNotification struct{
	Id string 
	Unread bool
	Subject Subject
	UpdatedAt time.Time `json:"updated_at"`
}

type Subject struct{
	Title string
	Type string
	URL string
	
}

type URL struct{
	HtmlURL string
}