package src

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type Card struct {
	Number string
	Cvv    string
	Month  string
	Year   string
}

func (s srvSt) Validate(c *gin.Context) {
	var response Response
	crdn := c.Query("cardNumber")
	cvv := c.Query("cvv")
	month := c.Query("month")
	year := c.Query("year")

	boolValue := validNumber(crdn)
	if boolValue == false {
		response.Error = "invalid card number"
		response.Message = "invalid card number"
		c.JSON(http.StatusBadGateway, response)
		c.Abort()
		return
	}

	if err := ValidateExpiration(month, year); err != nil {
		response.Error = "invalid card"
		response.Message = "invalid input"
		c.JSON(http.StatusBadGateway, response)
		c.Abort()
		return
	}

	if len(cvv) < 3 || len(cvv) > 4 {
		response.Error = "invalid cvv"
		response.Message = "invalid card details"
		c.JSON(http.StatusBadGateway, response)
		c.Abort()
		return

	}
	company, err := s.r.CompanyDetails(crdn)
	if err != nil {
		response.Error = "invalid cvv"
		response.Message = "invalid card details"
		c.JSON(http.StatusBadGateway, response)
		c.Abort()
		return
	}
	response.Success = true
	response.Data = company
	response.Message = "valid card no"
	c.JSON(http.StatusOK, response)
}
