package models

import "time"

type Link struct {
	Id             string `json:"id"`
	Owner_id       string `json:"owner_id"`
	Title          string `json:"title"`
	Url            string `json:"url"`
	Link_execution string `json:"link_execution"`
}

type LinkResponses struct {
	Id                 string
	ResponseMessage    string
	ResponseStatusCode int
	RequestTime        int64
	VerifiedDate       time.Time
	LinkId             string
}

type Profile struct {
	Email    string
	Username string
}
