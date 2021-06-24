// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nssaiavailability

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/nssf/logger"
	"github.com/free5gc/nssf/producer"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

func HTTPNSSAIAvailabilityUnsubscribe(c *gin.Context) {
	// Due to conflict of route matching, 'subscriptions' in the route is replaced with the existing wildcard ':nfId'
	nfID := c.Param("nfId")
	if nfID != "subscriptions" {
		c.JSON(http.StatusNotFound, gin.H{})
		logger.Nssaiavailability.Infof("404 Not Found")
		return
	}

	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	rsp := producer.HandleNSSAIAvailabilityUnsubscribe(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.HandlerLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
