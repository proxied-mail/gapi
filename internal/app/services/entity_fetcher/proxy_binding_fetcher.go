package entity_fetcher

import (
	"github.com/abrouter/gapi/internal/app/models"
	"github.com/abrouter/gapi/internal/app/repository"
	"github.com/abrouter/gapi/pkg/entityId"
	"go.uber.org/fx"
)

type ProxyBindingFetcher struct {
	fx.In
	ProxyBindingRepository repository.ProxyBindingRepository
	encoder                entityId.Encoder
}

func (pbf ProxyBindingFetcher) GetModel(entityId string) (models.ProxyBinding, error) {
	decodedId, err := pbf.encoder.Decode(entityId, "proxy_bindings")
	if err != nil {
		return models.ProxyBinding{}, err
	}

	proxyBinding := pbf.ProxyBindingRepository.GetById(int(decodedId))

	return proxyBinding, nil
}
