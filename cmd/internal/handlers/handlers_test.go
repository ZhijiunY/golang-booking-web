package handlers

type postData struct {
	key string
	value string
}

type theTest []struct {
	name  string
	url  string
	method string
	params []postData
} 