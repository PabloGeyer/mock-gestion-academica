package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Especialidad struct {
	Id           int    `fake:"{number:1}" json:"id"`
	Especialidad string `fake:"{jobtitle}" json:"especialidad"`
	Facultad     string `fake:"{company}" json:"facultad"`
	Universidad  string `fake:"{company}" json:"universidad"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	especialidades := generateMockData()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Mock API is running",
		})
	})

	r.GET("/api/v1/especialidades", func(c *gin.Context) {
		c.JSON(http.StatusOK, especialidades)
	})

	r.GET("/api/v1/especialidades/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}

		// Buscar especialidad por ID
		for _, esp := range especialidades {
			if esp.Id == id {
				c.JSON(http.StatusOK, esp)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Especialidad not found"})

	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func generateMockData() []Especialidad {
	// Datos predefinidos de especialidades
	return []Especialidad{
		{
			Id:           1,
			Especialidad: "Ingeniería en Sistemas de Información",
			Facultad:     "Facultad Regional Buenos Aires",
			Universidad:  "Universidad Tecnológica Nacional",
		},
		{
			Id:           2,
			Especialidad: "Ingeniería Civil",
			Facultad:     "Facultad Regional Buenos Aires",
			Universidad:  "Universidad Tecnológica Nacional",
		},
		{
			Id:           3,
			Especialidad: "Ingeniería Electrónica",
			Facultad:     "Facultad Regional Buenos Aires",
			Universidad:  "Universidad Tecnológica Nacional",
		},
		{
			Id:           4,
			Especialidad: "Ingeniería Mecánica",
			Facultad:     "Facultad Regional Buenos Aires",
			Universidad:  "Universidad Tecnológica Nacional",
		},
		{
			Id:           5,
			Especialidad: "Ingeniería Industrial",
			Facultad:     "Facultad Regional Buenos Aires",
			Universidad:  "Universidad Tecnológica Nacional",
		},
		{
			Id:           6,
			Especialidad: "Licenciatura en Administración de Empresas",
			Facultad:     "Facultad de Ciencias Económicas",
			Universidad:  "Universidad de Buenos Aires",
		},
		{
			Id:           7,
			Especialidad: "Contador Público",
			Facultad:     "Facultad de Ciencias Económicas",
			Universidad:  "Universidad de Buenos Aires",
		},
		{
			Id:           8,
			Especialidad: "Medicina",
			Facultad:     "Facultad de Medicina",
			Universidad:  "Universidad de Buenos Aires",
		},
		{
			Id:           9,
			Especialidad: "Derecho",
			Facultad:     "Facultad de Derecho",
			Universidad:  "Universidad de Buenos Aires",
		},
		{
			Id:           10,
			Especialidad: "Arquitectura",
			Facultad:     "Facultad de Arquitectura, Diseño y Urbanismo",
			Universidad:  "Universidad de Buenos Aires",
		},
	}
}
