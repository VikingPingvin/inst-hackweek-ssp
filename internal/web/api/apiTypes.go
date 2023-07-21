package api

type ApiData struct {
	ApiID   int
	ApiName string
	Message string
	ApiConfiguration
}

type ApiConfiguration struct {
	Enabled   bool
	TargetUrl string
	EndPoint  string
}
