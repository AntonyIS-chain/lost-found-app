package ports

import "github.com/AntonyIS-chain/lost-found-app/backend/gateway/internal/domain"


type ProxyService interface {
	ForwardRequest(req domain.ProxyRequest) (domain.ProxyResponse, error)
}