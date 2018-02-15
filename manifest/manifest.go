package manifest

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/ghodss/yaml"
)

type manifest map[string]json.RawMessage

// Routes - a collection of Route
type Routes []Route

// Route - a cf route, typically host.domain, host.domain/path or tcp.domain:port
type Route struct {
	Route string `json:"route,omitempty"`
}

// Update - updates deperecated elements of a CF manifest
func Update(oldManifest string) (string, error) {
	jsonManifest, err := loadJSONManifest(oldManifest)
	if err != nil {
		return "", err
	}
	host, err := jsonManifest.getHost()
	if err != nil {
		return "", err
	}
	if manifestErr := jsonManifest.addRoutes(host); manifestErr != nil {
		return "", manifestErr
	}
	return jsonManifest.marshal()
}

func loadJSONManifest(oldManifest string) (manifest, error) {
	jsonManifestBytes, err := yaml.YAMLToJSON([]byte(oldManifest))
	if err != nil {
		return nil, err
	}
	jsonManifest := make(manifest)
	if jsonErr := json.Unmarshal(jsonManifestBytes, &jsonManifest); jsonErr != nil {
		return nil, jsonErr
	}
	return jsonManifest, nil
}

func (jsonManifest manifest) marshal() (string, error) {
	newJSONManifestBytes, err := json.Marshal(jsonManifest)
	if err != nil {
		return "", err
	}
	newManifestBytes, err := yaml.JSONToYAML(newJSONManifestBytes)
	if err != nil {
		return "", err
	}
	return string(newManifestBytes), nil
}

func (jsonManifest manifest) getHost() (string, error) {
	var host string
	if err := json.Unmarshal(jsonManifest["name"], &host); err != nil {
		return "", err
	}
	if manifestHost, ok := jsonManifest["host"]; ok {
		if err := json.Unmarshal(manifestHost, &host); err != nil {
			return "", err
		}
		delete(jsonManifest, "host")
	}
	return host, nil
}

func (jsonManifest manifest) addRoutes(host string) error {
	var routes Routes
	if manifestRoutes, ok := jsonManifest["routes"]; ok {
		if err := json.Unmarshal(manifestRoutes, &routes); err != nil {
			return err
		}
	}
	domainRoute, err := jsonManifest.processDomain(host)
	if err != nil {
		return err
	}
	domainsRoutes, err := jsonManifest.processDomains(host)
	if err != nil {
		return err
	}
	routes = append(routes, domainRoute...)
	routes = append(routes, domainsRoutes...)
	routes = routes.removeDuplicates()
	marshalledRoutes, err := json.Marshal(routes)
	if err != nil {
		return err
	}
	jsonManifest["routes"] = marshalledRoutes
	return nil
}

func (jsonManifest manifest) processDomain(host string) (Routes, error) {
	var routes Routes
	if manifestDomain, ok := jsonManifest["domain"]; ok {
		var domain string
		if err := json.Unmarshal(manifestDomain, &domain); err != nil {
			return nil, err
		}
		routes = append(routes, Route{Route: fmt.Sprintf("%s.%s", host, domain)})
		delete(jsonManifest, "domain")
	}
	return routes, nil
}

func (jsonManifest manifest) processDomains(host string) (Routes, error) {
	var routes Routes
	if manifestDomain, ok := jsonManifest["domains"]; ok {
		var domains []string
		if err := json.Unmarshal(manifestDomain, &domains); err != nil {
			return nil, err
		}
		delete(jsonManifest, "domains")
		for _, domain := range domains {
			routes = append(routes, Route{Route: fmt.Sprintf("%s.%s", host, domain)})
		}
	}
	return routes, nil
}

func (routes Routes) removeDuplicates() Routes {
	routesMap := make(map[string]string)
	for _, route := range routes {
		routesMap[route.Route] = ""
	}
	var dedupedRoutes Routes
	for route := range routesMap {
		dedupedRoutes = append(dedupedRoutes, Route{Route: route})
	}
	sort.Slice(dedupedRoutes, func(i int, j int) bool {
		return dedupedRoutes[i].Route < dedupedRoutes[j].Route
	})
	return dedupedRoutes
}
