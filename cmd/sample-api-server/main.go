package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"swagger/restapi"
	"swagger/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSampleAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Specify the paths to your TLS certificate and key files
	server.TLSHost = "localhost"
	server.TLSCertificate = flags.Filename("server.crt")
	server.TLSCertificateKey = flags.Filename("server.key")

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Sample API"
	parser.LongDescription = "API description in Markdown."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	// Serve the API with TLS
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
