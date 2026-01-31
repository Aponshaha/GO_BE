package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupSwagger configures Swagger/OpenAPI routes
func SetupSwagger(router *gin.Engine) {
	// Serve Swagger YAML spec
	router.GET("/api/v1/swagger.yaml", func(c *gin.Context) {
		c.File("docs/swagger.yaml")
	})

	// Serve Swagger UI
	router.GET("/swagger/index.html", func(c *gin.Context) {
		swaggerHTML := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
				<meta charset="UTF-8">
				<title>API Documentation</title>
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@3/swagger-ui.css">
				<style>
						html {
								box-sizing: border-box;
								overflow: -moz-scrollbars-vertical;
								overflow-y: scroll;
						}
						*, *:before, *:after {
								box-sizing: inherit;
						}
						body {
								margin: 0;
								padding: 0;
						}
				</style>
		</head>
		<body>
				<div id="swagger-ui"></div>
				<script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js" charset="UTF-8"></script>
				<script src="https://unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js" charset="UTF-8"></script>
				<script>
						window.onload = function() {
								const ui = SwaggerUIBundle({
										url: "/api/v1/swagger.yaml",
										dom_id: '#swagger-ui',
										deepLinking: true,
										presets: [
												SwaggerUIBundle.presets.apis,
												SwaggerUIStandalonePreset
										],
										plugins: [
												SwaggerUIBundle.plugins.DownloadUrl
										],
										layout: "StandaloneLayout"
								});
								window.ui = ui;
						}
				</script>
		</body>
		</html>`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(swaggerHTML))
	})

	// Redirect /swagger to /swagger/index.html
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Redirect /swagger/ to /swagger/index.html
	router.GET("/swagger/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
