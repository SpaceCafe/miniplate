package pkg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/spacecafe/miniplate/pkg/functions"
)

var ErrInvalidContext = errors.New("invalid context")

func LoadContexts(ctxList []string) (ctx map[string]any, err error) {
	ctx = make(map[string]any)

	for _, ctxItem := range ctxList {
		var (
			ctxURL *url.URL
			data   any
		)

		ctxName, ctxRawURL, ok := strings.Cut(ctxItem, "=")
		if !ok || (ctxName == "" && ctxRawURL == "") {
			return ctx, fmt.Errorf("%w: unsupported context format: %s", ErrInvalidContext, ctxItem)
		}

		ctxURL, err = url.Parse(ctxRawURL)
		if err != nil {
			return nil, err
		}

		switch ctxURL.Scheme {
		case "stdin":
			data, err = loadContextFromStdin(ctxURL)
		case "file", "":
			data, err = loadContextFromFile(ctxURL)
		case "http", "https":
			data, err = loadContextFromWeb(ctxURL)
		default:
			return nil, fmt.Errorf(
				"%w: unsupported context schema: %s",
				ErrInvalidContext,
				ctxURL.Scheme,
			)
		}

		if err != nil {
			return nil, err
		}

		ctx[ctxName] = data
	}

	return ctx, err
}

func loadContextFromStdin(ctxURL *url.URL) (data any, err error) {
	var buf []byte

	buf, err = io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	return parseContext(ctxURL, buf)
}

func loadContextFromFile(ctxURL *url.URL) (data any, err error) {
	var buf []byte

	filePath := path.Clean(ctxURL.Host + ctxURL.Path)

	buf, err = os.ReadFile(path.Clean(filePath))
	if err != nil {
		return
	}

	return parseContext(ctxURL, buf)
}

func loadContextFromWeb(ctxURL *url.URL) (data any, err error) {
	var buf []byte

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		ctxURL.String(),
		http.NoBody,
	)
	if err != nil {
		return data, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, err
	}

	defer func() { _ = resp.Body.Close() }()

	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"%w: failed to load context at %s: %s",
			ErrInvalidContext,
			ctxURL.String(),
			resp.Status,
		)
	}

	ctxURL.RawQuery = "type=" + resp.Header.Get("Content-Type")

	return parseContext(ctxURL, buf)
}

func parseContext(ctxURL *url.URL, buf []byte) (data any, err error) {
	dataFuncs := &functions.DataFuncs{}

	if len(buf) == 0 {
		return nil, fmt.Errorf("%w: empty context at %s", ErrInvalidContext, ctxURL.String())
	}

	switch path.Ext(ctxURL.Path) {
	case ".json":
		data, err = dataFuncs.JSON(buf)
	case ".toml":
		data, err = dataFuncs.TOML(buf)
	case ".yaml":
		data, err = dataFuncs.YAML(buf)
	default:
		mime := ctxURL.Query().Get("type")
		if mime == "" {
			mime, _, _ = strings.Cut(http.DetectContentType(buf), ";")
		}

		switch mime {
		case "application/json":
			data, err = dataFuncs.JSON(buf)
		case "application/toml":
			data, err = dataFuncs.TOML(buf)
		case "application/yaml", "text/plain":
			data, err = dataFuncs.YAML(buf)
		default:
			return nil, fmt.Errorf(
				"%w: unsupported context mime type at %s: %s",
				ErrInvalidContext,
				ctxURL.String(),
				mime,
			)
		}
	}

	if err != nil {
		return nil, fmt.Errorf(
			"%w: failed to parse context at %s: %w",
			ErrInvalidContext,
			ctxURL.String(),
			err,
		)
	}

	return data, err
}
