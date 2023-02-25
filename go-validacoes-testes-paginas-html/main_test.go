package main

import (
	"go-validacoes-testes-paginas-html/controllers"
	"go-validacoes-testes-paginas-html/database"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestVefica1(t *testing.T) {
	r := SetUpRoutes()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code, "Deveriam ser iguais")
	mockResp := `{"API diz":"E ai gui, Tudo beleza?"}`
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockResp, string(respBody))
}

func TestVerifica2(t *testing.T) {
	database.ConectaComBancoDeDados()
	r := SetUpRoutes()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("Get", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
