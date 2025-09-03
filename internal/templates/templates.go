package templates
import "embed"

//go:embed all:common all:java_spring_boot all:python_fastapi all:php_laravel all:nodejs_nextjs
var TemplateFS embed.FS