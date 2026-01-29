package pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/spacecafe/miniplate/pkg/functions"
)

func LoadContexts(ctxList []string) (ctx map[string]any, err error) {
	ctx = make(map[string]any)
	for _, ctxItem := range ctxList {
		var (
			ctxURL *url.URL
			data   any
		)

		ctxName, ctxRawURL, ok := strings.Cut(ctxItem, "=")
		if !ok || (ctxName == "" && ctxRawURL == "") {
			return ctx, fmt.Errorf("unsupported context format: %s", ctxItem)
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
			return nil, fmt.Errorf("unsupported context schema: %s", ctxURL.Scheme)
		}
		if err != nil {
			return nil, err
		}
		ctx[ctxName] = data
	}
	return
}

func loadContextFromStdin(ctxURL *url.URL) (data any, err error) {
	var (
		buf []byte
	)

	buf, err = io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	return parseContext(ctxURL, buf)
}

func loadContextFromFile(ctxURL *url.URL) (data any, err error) {
	var (
		buf []byte
	)

	filePath := path.Clean(ctxURL.Host + ctxURL.Path)
	buf, err = os.ReadFile(path.Clean(filePath))
	if err != nil {
		return
	}

	return parseContext(ctxURL, buf)
}

func loadContextFromWeb(ctxURL *url.URL) (data any, err error) {
	var (
		buf []byte
	)

	resp, err := http.Get(ctxURL.String())
	if err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()
	buf, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to load context at %s: %s", ctxURL.String(), resp.Status)
	}
	ctxURL.RawQuery = "type=" + resp.Header.Get("Content-Type")

	return parseContext(ctxURL, buf)
}

func parseContext(ctxURL *url.URL, buf []byte) (data any, err error) {
	var (
		dataFuncs = &functions.DataFuncs{}
	)

	if len(buf) == 0 {
		return nil, fmt.Errorf("empty context at %s", ctxURL.String())
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
			return nil, fmt.Errorf("unsupported context mime type at %s: %s", ctxURL.String(), mime)
		}
	}
	if err != nil {
		return nil, fmt.Errorf("failed to parse context at %s: %w", ctxURL.String(), err)
	}
	return
}
