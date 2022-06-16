package healthcheck

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func PingHandler(c *gin.Context) {
	//返回Pong
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
	})
}

func DockerOtherHandler(c *gin.Context) {
	ip := c.Query("ip")
	uri := c.Request.RequestURI
	url_get := fmt.Sprintf("http://%s:%s/%s", ip, "8001", strings.Split(uri, "router/")[1])
	resp, err := http.Get(url_get)
	if err != nil {
		fmt.Println("get request failed: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var jsondata interface{}
	err = json.Unmarshal(body, &jsondata)
	if err != nil {
		fmt.Println("json unmarshal failed: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, jsondata)
	return
}

func ConImageGetHandler(c *gin.Context) {
	ip := c.Query("ip")
	uri := c.Request.RequestURI
	url_get := fmt.Sprintf("http://%s:%s/%s", ip, "8001", strings.Split(uri, "router/")[1])
	resp, err := http.Get(url_get)
	if err != nil {
		fmt.Println("get request failed: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var jsondata interface{}
	err = json.Unmarshal(body, &jsondata)
	if err != nil {
		fmt.Println("json unmarshal failed: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, jsondata)
	return
}

func CACHandler(c *gin.Context) {
	ip := c.Query("ip")
	url_post := fmt.Sprintf("http://%s:%s/createandruncontainer", ip, "8001")
	resp, err := http.Post(url_post, c.ContentType(), c.Request.Body)
	if err != nil {
		fmt.Println("post request createandruncontainer failed: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var jsondata interface{}
	err = json.Unmarshal(body, &jsondata)
	if err != nil {
		fmt.Println("json unmarshal failed: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, jsondata)
	return
}

func RALHandler(c *gin.Context) {
	ip := c.Query("ip")
	url_post := fmt.Sprintf("http://%s:%s/receiveandloadimage", ip, "8001")
	fmt.Printf("contenttype:%s\n", c.Request.Header.Get("Content-Type"))
	resp, err := http.Post(url_post, c.Request.Header.Get("Content-Type"), c.Request.Body)
	if err != nil {
		fmt.Println("post request receiveandloadimage failed: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var jsondata interface{}
	err = json.Unmarshal(body, &jsondata)
	if err != nil {
		fmt.Println("json unmarshal failed: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(resp.StatusCode, jsondata)
	return
}

func DeviceDSPStatusHandler(c *gin.Context){


}

func DeviceFPGAStatusHandler(c *gin.Context){

	
}