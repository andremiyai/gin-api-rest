package main

import (
	"bytes"
	"encoding/json"
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) {

	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// }
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz:":"E ai gui, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "12312312312", RG: "1212312312"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12312312312", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Aluno Teste", alunoMock.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12312312312", alunoMock.CPF, "Os cpfs devem ser iguais")
	assert.Equal(t, "1212312312", alunoMock.RG, "Os Rgs devem ser iguais")
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	aluno := models.Aluno{Nome: "Aluno Teste Editado", CPF: "12312312300", RG: "121231230"}
	valorJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)

	var alunoAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoAtualizado)

	assert.Equal(t, "Aluno Teste Editado", alunoAtualizado.Nome, "Os nomes devem ser iguais")
	assert.Equal(t, "12312312300", alunoAtualizado.CPF, "Os cpfs devem ser iguais")
	assert.Equal(t, "121231230", alunoAtualizado.RG, "Os Rgs devem ser iguais")
}
