package github

//CreateRepoRequest é uma classe
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Descricao   string `json:"descricao"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

//CreateRepoResponse  é uma classe
type CreateRepoResponse struct {
	Id           int64  `json:"id"`
	Nome         string `json:"nome"`
	NomeCompleto string `json:"nome_completo"`
	Owner 		 RepoOwner `json:"owner"`
	Permissions  RepoPermissions `json:"permissions"`
}

//RepoOwner  é uma classe
type RepoOwner struct {
	Id      int64  `json:"id"`
	Login   string `json:"login"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

//RepoPermissions  é uma classe
type RepoPermissions struct {
	IsAdmin      bool  `json:"admin"`
	HasPull   bool `json:"pull"`
	HasPush   bool `json:"push"`
}
