package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zexy-swami/SRP/SRP_web/internal/db"
	"github.com/zexy-swami/SRP/SRP_web/pkg/validation"
)

func rootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "root_page.html", gin.H{})
}

func signUpHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "signup_page.html", gin.H{})
}

func methodNotAllowedHandler(c *gin.Context) {
	c.HTML(http.StatusMethodNotAllowed, "error_page.html", gin.H{
		"Error": "Method not allowed",
	})
}

func verificationHandler(c *gin.Context) {
	if err := c.Request.ParseForm(); err != nil {
		c.HTML(http.StatusInternalServerError, "error_page.html", gin.H{
			"Error": "Error occurred while parsing user form",
		})
		return
	}

	userData := [3]string{
		strings.TrimSpace(c.Request.Form.Get("firstname")),
		strings.TrimSpace(c.Request.Form.Get("lastname")),
		c.Request.Form.Get("password"),
	}

	if isEmptyDataFound := validation.CheckDataForEmptiness(userData); isEmptyDataFound {
		c.HTML(http.StatusBadRequest, "error_page.html", gin.H{
			"Error": "Data from fields can't be empty",
		})
		return
	}

	if isXSSFound, detectedXSS := validation.CheckDataForXSS(userData); isXSSFound {
		c.HTML(http.StatusBadRequest, "error_page.html", gin.H{
			"Error": fmt.Sprintf("XSS detected in %q", detectedXSS),
		})
		return
	}

	if srpID, err := db.SignupUser(userData[0], userData[1], userData[2]); err != nil {
		c.HTML(http.StatusBadRequest, "error_page.html", gin.H{
			"Error": err.Error(),
		})
	} else {
		c.HTML(http.StatusOK, "verification.html", gin.H{
			"SRP_id": srpID,
		})
	}
}
