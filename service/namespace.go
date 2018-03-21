package service

import "strings"

// NAMESPACE is the namespace used for the docker services
const NAMESPACE string = "MESG"

func (service *Service) namespace() string {
	return strings.Join([]string{
		NAMESPACE,
		strings.Replace(service.Name, " ", "-", -1),
	}, "-")
}