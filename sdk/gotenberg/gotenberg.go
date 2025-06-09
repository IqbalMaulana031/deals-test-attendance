package gotenberg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/pkg/errors"

	"starter-go-gin/config"
)

// Gotenberg is an struct for gotenberg SDK
type Gotenberg struct {
	cfg config.Config
}

// NewGotenberg initiate gotenberg SDK
func NewGotenberg(cfg config.Config) *Gotenberg {
	return &Gotenberg{
		cfg: cfg,
	}
}

// FromHTML is a function to create pdf from html
func (g *Gotenberg) FromHTML(httpWriter http.ResponseWriter, htmlData []byte, forms map[string]string) error {
	ctx := context.Background()
	endpoint := "/forms/chromium/convert/html"

	// init multipart
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// create file
	part, err := writer.CreateFormFile("file", "index.html")
	if err != nil {
		return errors.Wrap(err, "Error while add multipart file")
	}
	_, err = io.Copy(part, bytes.NewReader(htmlData))
	if err != nil {
		return errors.Wrap(err, "Error while write multipart file")
	}

	// add another form request
	for key, value := range forms {
		var part io.Writer
		if part, err = writer.CreateFormField(key); err != nil {
			return errors.Wrap(err, "Error while add form request")
		}

		if _, err = io.Copy(part, bytes.NewBufferString(value)); err != nil {
			return errors.Wrap(err, "Error while write form request")
		}
	}

	if err := writer.Close(); err != nil {
		return errors.Wrap(err, "Error while close response")
	}

	// hit gotenberg API
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", g.cfg.Gotenberg.Host, endpoint), buf)
	if err != nil {
		return errors.Wrap(err, "Error while init request gotenberg API")
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "Error while request gotenberg API")
	}
	if res.StatusCode != http.StatusOK {
		return errors.Wrap(err, "Error response gotenberg API")
	}

	// copy response body to http writer
	if _, err = io.Copy(httpWriter, res.Body); err != nil {
		return errors.Wrap(err, "Error while return response")
	}

	if err := res.Body.Close(); err != nil {
		return errors.Wrap(err, "Error while close response")
	}

	return nil
}
