package router

import (
	"compare/controllers/alive"
	"compare/controllers/files"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Package classification Documentation of compare Service
//
// Documentation for compare Service
//
// Schemes: https
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// securityDefinitions:
//   BasicAuth:
//     type: basic
//     name: Authorization
//     in: header
//
// swagger:meta
func router(ginEngine *gin.Engine) {
	ginEngine.Use(cors.Default())

	ginEngine.GET("/alive", alive.Alive)
	ginEngine.GET("/files/health", files.Health)
}
