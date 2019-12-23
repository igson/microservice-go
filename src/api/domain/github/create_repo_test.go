package github

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"encoding/json"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name: "introdução ao go lang",
		Descricao: "introdução a repositório em go",
		Homepage: "https://github.com",
		Private: true,
		HasIssues: true,
		HasWiki: false,
		HasProjects: true,
	}

	resultado, err := json.Marshal(request)

	var target CreateRepoRequest

	err = json.Unmarshal(resultado, &target)

	fmt.Println(string(resultado))
	
	assert.Nil(t, err)
	assert.NotNil(t,resultado)
	assert.Nil(t, err) 

	assert.EqualValues(t, target.Name, request.Name)
 
}