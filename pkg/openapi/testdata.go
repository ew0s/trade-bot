package openapi

const WrongSwaggerHostFormat = `
{
    "swagger": "2.0",
    "info": {
        "title": "Trade-bot API",
        "contact": {
            "email": "company@gmail.com"
        },
        "license": {},
		"version": "1.0"
    },
	"host": "/invalidHost"
}
`
