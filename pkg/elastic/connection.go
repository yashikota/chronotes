package elastic

import (
	"log/slog"

	"github.com/elastic/go-elasticsearch/v8"
)

var ElasticTypedClient *elasticsearch.TypedClient

func Connect() {
	var err error
	ElasticTypedClient, err = elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://elasticsearch:9200"},
	})
	if err != nil {
		slog.Error("Error creating the client: " + err.Error())
	}
	slog.Info("Connected to Elasticsearch")
}
