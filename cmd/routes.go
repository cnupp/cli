package cmd
import (
	"fmt"
	"github.com/cde/client/config"
	"github.com/cde/apisdk/net"
	"github.com/cde/apisdk/api"
)

// RouteCreate creates an route.
func RoutesCreate(domain string, path string) error {
	configRepository := config.NewConfigRepository(func(error) {})
	routeRepository := api.NewRouteRepository(configRepository, net.NewCloudControllerGateway(configRepository))
	routeParams := api.RouteParams{
		Domain: domain,
		Path: path,
	}
	err := routeRepository.Create(routeParams)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func RoutesList() error {
	configRepository := config.NewConfigRepository(func(error) {})
	routeRepository := api.NewRouteRepository(configRepository, net.NewCloudControllerGateway(configRepository))
	routes, err := routeRepository.GetRoutes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("=== Routes: [%d]\n", len(routes.Items()))

	for _, route := range routes.Items() {
		fmt.Printf("id: %s path: %s domain: %s\n", route.ID(), route.Path(), route.Domain().Name)
	}
	return err
}

func RouteBindWithApp(route, appName string) error {
	configRepository := config.NewConfigRepository(func(error) {})
	appRepo := api.NewAppRepository(configRepository, net.NewCloudControllerGateway(configRepository))
	app, err := appRepo.GetApp(appName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	routeParams := api.AppRouteParams{
		Route: route,
	}
	err = app.AssociateRoute(routeParams)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}