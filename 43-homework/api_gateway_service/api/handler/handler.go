package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// Create
func (h *handler) Create(g *gin.Context) {
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	client := http.Client{}
	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		fmt.Println(111111)
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(2222222)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(333333333)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"data": string(res)})
}

// Update
func (h *handler) Update(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"data": string(res)})
}

// Delete
func (h *handler) Delete(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path

	req, err := http.NewRequest(method, "http://localhost:8075"+url, nil)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	resq, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resq.Body.Close()

	g.JSON(http.StatusCreated, "SUCCESS")
}

// Get by id
func (h *handler) GetById(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resq, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resq.Body.Close()

	res, err := io.ReadAll(resq.Body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusCreated, res)
}

// Get all
func (h *handler) GetAll(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resq, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resq.Body.Close()

	res, err := io.ReadAll(resq.Body)
	if err != nil {

		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusCreated, string(res))
}


// ------------------------------------

// Create
func (h *handler) Create1(g *gin.Context) {
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	client := http.Client{}
	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		fmt.Println(111111)
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(2222222)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(333333333)
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"data": string(res)})
}

// Update
func (h *handler) Update1(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"data": string(res)})
}

// Delete
func (h *handler) Delete1(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path

	req, err := http.NewRequest(method, "http://localhost:8075"+url, nil)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	resq, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resq.Body.Close()

	g.JSON(http.StatusCreated, "SUCCESS")
}

// Get by id
func (h *handler) GetById1(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resq, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resq.Body.Close()

	res, err := io.ReadAll(resq.Body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusCreated, res)
}

// Get all
func (h *handler) GetAll1(g *gin.Context) {
	client := http.Client{}
	method := g.Request.Method
	url := g.Request.URL.Path
	body := g.Request.Body

	req, err := http.NewRequest(method, "http://localhost:8075"+url, body)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resq, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	defer resq.Body.Close()

	res, err := io.ReadAll(resq.Body)
	if err != nil {

		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("error:", err.Error())
		return
	}

	g.JSON(http.StatusCreated, string(res))
}
