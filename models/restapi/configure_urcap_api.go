// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"universal-robots.com/urservice/restapi/operations"
	"universal-robots.com/urservice/restapi/operations/artifact"
	"universal-robots.com/urservice/restapi/operations/container"
	"universal-robots.com/urservice/restapi/operations/urcap"
)

//go:generate swagger generate server --target ../../src --name UrcapAPI --spec ../../../tmp/go-swagger-1158340383/URCapAPI.yaml --principal interface{} --strict-responders

func configureFlags(api *operations.UrcapAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UrcapAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// urcap.AddUrcapMaxParseMemory = 32 << 20

	if api.UrcapAddUrcapHandler == nil {
		api.UrcapAddUrcapHandler = urcap.AddUrcapHandlerFunc(func(params urcap.AddUrcapParams) urcap.AddUrcapResponder {
			return urcap.AddUrcapNotImplemented()
		})
	}
	if api.UrcapDeleteUrcapByIDHandler == nil {
		api.UrcapDeleteUrcapByIDHandler = urcap.DeleteUrcapByIDHandlerFunc(func(params urcap.DeleteUrcapByIDParams) urcap.DeleteUrcapByIDResponder {
			return urcap.DeleteUrcapByIDNotImplemented()
		})
	}
	if api.ArtifactGetArtifactActualNameByPathHandler == nil {
		api.ArtifactGetArtifactActualNameByPathHandler = artifact.GetArtifactActualNameByPathHandlerFunc(func(params artifact.GetArtifactActualNameByPathParams) artifact.GetArtifactActualNameByPathResponder {
			return artifact.GetArtifactActualNameByPathNotImplemented()
		})
	}
	if api.ArtifactGetArtifactMetadataByPathHandler == nil {
		api.ArtifactGetArtifactMetadataByPathHandler = artifact.GetArtifactMetadataByPathHandlerFunc(func(params artifact.GetArtifactMetadataByPathParams) artifact.GetArtifactMetadataByPathResponder {
			return artifact.GetArtifactMetadataByPathNotImplemented()
		})
	}
	if api.ContainerGetContainerByIDHandler == nil {
		api.ContainerGetContainerByIDHandler = container.GetContainerByIDHandlerFunc(func(params container.GetContainerByIDParams) container.GetContainerByIDResponder {
			return container.GetContainerByIDNotImplemented()
		})
	}
	if api.ContainerGetContainerPortMappingByIDHandler == nil {
		api.ContainerGetContainerPortMappingByIDHandler = container.GetContainerPortMappingByIDHandlerFunc(func(params container.GetContainerPortMappingByIDParams) container.GetContainerPortMappingByIDResponder {
			return container.GetContainerPortMappingByIDNotImplemented()
		})
	}
	if api.UrcapGetUrcapByIDHandler == nil {
		api.UrcapGetUrcapByIDHandler = urcap.GetUrcapByIDHandlerFunc(func(params urcap.GetUrcapByIDParams) urcap.GetUrcapByIDResponder {
			return urcap.GetUrcapByIDNotImplemented()
		})
	}
	if api.UrcapGetUrcapsHandler == nil {
		api.UrcapGetUrcapsHandler = urcap.GetUrcapsHandlerFunc(func(params urcap.GetUrcapsParams) urcap.GetUrcapsResponder {
			return urcap.GetUrcapsNotImplemented()
		})
	}
	if api.ContainerStartContainerByIDHandler == nil {
		api.ContainerStartContainerByIDHandler = container.StartContainerByIDHandlerFunc(func(params container.StartContainerByIDParams) container.StartContainerByIDResponder {
			return container.StartContainerByIDNotImplemented()
		})
	}
	if api.ContainerStatusContainerByIDHandler == nil {
		api.ContainerStatusContainerByIDHandler = container.StatusContainerByIDHandlerFunc(func(params container.StatusContainerByIDParams) container.StatusContainerByIDResponder {
			return container.StatusContainerByIDNotImplemented()
		})
	}
	if api.ContainerStopAllContainersHandler == nil {
		api.ContainerStopAllContainersHandler = container.StopAllContainersHandlerFunc(func(params container.StopAllContainersParams) container.StopAllContainersResponder {
			return container.StopAllContainersNotImplemented()
		})
	}
	if api.ContainerStopContainerByIDHandler == nil {
		api.ContainerStopContainerByIDHandler = container.StopContainerByIDHandlerFunc(func(params container.StopContainerByIDParams) container.StopContainerByIDResponder {
			return container.StopContainerByIDNotImplemented()
		})
	}
	if api.UrcapUpdateUrcapByIDHandler == nil {
		api.UrcapUpdateUrcapByIDHandler = urcap.UpdateUrcapByIDHandlerFunc(func(params urcap.UpdateUrcapByIDParams) urcap.UpdateUrcapByIDResponder {
			return urcap.UpdateUrcapByIDNotImplemented()
		})
	}
	if api.ContainerValidateContainerByIDHandler == nil {
		api.ContainerValidateContainerByIDHandler = container.ValidateContainerByIDHandlerFunc(func(params container.ValidateContainerByIDParams) container.ValidateContainerByIDResponder {
			return container.ValidateContainerByIDNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
