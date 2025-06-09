package interfaces

import "net/http"

// GotenbergUseCase define interface for gotenberg
type GotenbergUseCase interface {
	FromHTML(httpWriter http.ResponseWriter, htmlData []byte, forms map[string]string) error
}
