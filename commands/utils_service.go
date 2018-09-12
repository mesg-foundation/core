package commands

import (
	"github.com/mesg-foundation/core/interface/grpc/core"
	"github.com/mesg-foundation/core/service"
)

// TODO(ilgooz) rm this when we stop using internal methods of service in cmd.
func toServices(ss []*core.Service) []*service.Service {
	services := make([]*service.Service, 0)
	for _, s := range ss {
		services = append(services, toService(s))
	}
	return services
}

// TODO(ilgooz) rm this when we stop using internal methods of service in cmd.
func toService(s *core.Service) *service.Service {
	return &service.Service{
		ID:            s.ID,
		Name:          s.Name,
		Description:   s.Description,
		Repository:    s.Repository,
		Tasks:         toTasks(s.Tasks),
		Events:        toEvents(s.Events),
		Dependencies:  toDependencies(s.Dependencies),
		Configuration: toDependency(s.Configuration),
	}
}

func toTasks(tasks map[string]*core.Task) map[string]*service.Task {
	ts := make(map[string]*service.Task, 0)
	for taskKey, task := range tasks {
		t := &service.Task{
			Key:         task.Key,
			Name:        task.Name,
			Description: task.Description,
			ServiceName: task.ServiceName,
			Inputs:      toParameters(task.Inputs),
			Outputs:     map[string]*service.Output{},
		}
		for outputKey, output := range task.Outputs {
			t.Outputs[outputKey] = &service.Output{
				Key:         output.Key,
				Name:        output.Name,
				Description: output.Description,
				TaskKey:     output.TaskKey,
				ServiceName: output.ServiceName,
				Data:        toParameters(output.Data),
			}
		}
		ts[taskKey] = t
	}
	return ts
}

func toEvents(events map[string]*core.Event) map[string]*service.Event {
	es := make(map[string]*service.Event, 0)
	for eventKey, event := range events {
		es[eventKey] = &service.Event{
			Key:         event.Key,
			Name:        event.Name,
			Description: event.Description,
			ServiceName: event.ServiceName,
			Data:        toParameters(event.Data),
		}
	}
	return es
}

// TODO(ilgooz) rm this when we stop using internal methods of service in cmd.
func toParameters(params map[string]*core.Parameter) map[string]*service.Parameter {
	gParams := make(map[string]*service.Parameter, 0)
	for key, param := range params {
		gParams[key] = &service.Parameter{
			Name:        param.Name,
			Description: param.Description,
			Type:        param.Type,
			Optional:    param.Optional,
		}
	}
	return gParams
}

func toDependency(dep *core.Dependency) *service.Dependency {
	if dep == nil {
		return nil
	}
	return &service.Dependency{
		Image:       dep.Image,
		Volumes:     dep.Volumes,
		VolumesFrom: dep.Volumesfrom,
		Ports:       dep.Ports,
		Command:     dep.Command,
	}
}

func toDependencies(deps map[string]*core.Dependency) map[string]*service.Dependency {
	ds := make(map[string]*service.Dependency, 0)
	for key, dep := range deps {
		ds[key] = toDependency(dep)
	}
	return ds
}