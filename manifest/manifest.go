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
func Update(oldManifest []byte) (string, error) {
	jsonManifest, err := loadJSONManifest(oldManifest)
	if err != nil {
		return "", err
	}
	newManifest, err := updateApplication(oldManifest)
	if err != nil {
		return "", err
	}
	if manifestApplications, ok := jsonManifest["applications"]; ok {
		newApplications, manifestErr := updateApplications(manifestApplications, jsonManifest)
		if manifestErr != nil {
			return "", manifestErr
		}
		newManifest["applications"], _ = json.Marshal(newApplications)
		delete(newManifest, "routes")
		return marshal(newManifest)
	}
	return marshal(newManifest)
}

func updateApplications(manifestApplications []byte, baseManifest manifest) ([]manifest, error) {
	var applicationsJSON []json.RawMessage
	if manifestErr := json.Unmarshal(manifestApplications, &applicationsJSON); manifestErr != nil {
		return nil, manifestErr
	}
	host, err := baseManifest.getHost()
	if err != nil {
		return nil, err
	}
	domain, err := baseManifest.getDomain()
	if err != nil {
		return nil, err
	}
	domains, err := baseManifest.getDomains()
	if err != nil {
		return nil, err
	}
	var newApplications []manifest
	for _, application := range applicationsJSON {
		applicationObj, err := loadJSONManifest(application)
		if err != nil {
			return nil, err
		}
		if _, ok := applicationObj["host"]; !ok {
			if host != "" {
				applicationObj["host"], err = json.Marshal(host)
				if err != nil {
					return nil, err
				}
			}
		}
		if _, ok := applicationObj["domain"]; !ok {
			if domain != "" {
				applicationObj["domain"], err = json.Marshal(domain)
				if err != nil {
					return nil, err
				}
			}
		}
		specificDomainsJSON, _ := applicationObj["domains"]
		var specificDomains []string
		if mashalErr := json.Unmarshal(specificDomainsJSON, &specificDomains); err != nil {
			return nil, mashalErr
		}
		specificDomains = append(specificDomains, domains...)
		if len(specificDomains) != 0 {
			applicationObj["domains"], err = json.Marshal(specificDomains)
			if err != nil {
				return nil, err
			}
		}
		marshalledApplication, err := json.Marshal(applicationObj)
		if err != nil {
			return nil, err
		}
		applicationManifest, appErr := updateApplication(marshalledApplication)
		if appErr != nil {
			return nil, appErr
		}
		newApplications = append(newApplications, applicationManifest)
	}
	return newApplications, nil
}

func updateApplication(oldManifest []byte) (manifest, error) {
	jsonManifest, err := loadJSONManifest(oldManifest)
	if err != nil {
		return nil, err
	}
	host, err := jsonManifest.getHost()
	if err != nil {
		return nil, err
	}
	if manifestErr := jsonManifest.addRoutes(host); manifestErr != nil {
		return nil, manifestErr
	}
	return jsonManifest, nil
}

func loadJSONManifest(oldManifest []byte) (manifest, error) {
	jsonManifestBytes, err := yaml.YAMLToJSON(oldManifest)
	if err != nil {
		return nil, err
	}
	jsonManifest := make(manifest)
	if jsonErr := json.Unmarshal(jsonManifestBytes, &jsonManifest); jsonErr != nil {
		return nil, jsonErr
	}
	return jsonManifest, nil
}

func marshal(jsonManifest interface{}) (string, error) {
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
	if manifestName, ok := jsonManifest["name"]; ok {
		if err := json.Unmarshal(manifestName, &host); err != nil {
			return "", err
		}
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
	if len(routes) != 0 {
		marshalledRoutes, err := json.Marshal(routes)
		if err != nil {
			return err
		}
		jsonManifest["routes"] = marshalledRoutes
	}
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

func (jsonManifest manifest) getDomain() (string, error) {
	var domain string
	if manifestDomain, ok := jsonManifest["domain"]; ok {
		if err := json.Unmarshal(manifestDomain, &domain); err != nil {
			return "", err
		}
		delete(jsonManifest, "domain")
	}
	return domain, nil
}

func (jsonManifest manifest) getDomains() ([]string, error) {
	var domains []string
	if manifestDomain, ok := jsonManifest["domains"]; ok {
		if err := json.Unmarshal(manifestDomain, &domains); err != nil {
			return nil, err
		}
		delete(jsonManifest, "domains")
	}
	return domains, nil
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
