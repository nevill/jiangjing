
ELASTIC_PASSWORD = elastic823
KIBANA_PASSWORD = kibana823
ENCRYPTION_KEYS = 0df8f98f7c7e98accf4a4eae0c489fcba53467199165b7aab63981a57017dd6a
SESSION_KEY = 1662e70fa04b61b992045575ab9d6f0285c49450c5e78a3c8ed507843e343487f878891484d6c5d9543ec8f550846e20a2d1842705166fc063b12c23a12ae2ee

define set-password
  echo "Waiting for Elasticsearch availability";
  until docker exec -it elasticsearch curl --cacert /usr/share/elasticsearch/config/certs/http_ca.crt https://localhost:9200 | grep -q "missing authentication credentials"; do sleep 30; done;

  echo "Setting kibana_system password";
  until docker exec -it elasticsearch curl -X POST --cacert /usr/share/elasticsearch/config/certs/http_ca.crt -u elastic:${ELASTIC_PASSWORD} -H "Content-Type: application/json" https://localhost:9200/_security/user/kibana_system/_password -d "{\"password\":\"${KIBANA_PASSWORD}\"}" | grep -q "^{}"; do sleep 20; done;
  echo "All done!";
endef

docker-start:
	docker run --name elasticsearch --hostname elasticsearch -d --net elastic -p 9200:9200 \
		-v es-config:/usr/share/elasticsearch/config \
		--env "node.name=elasticsearch" \
		--env "ES_JAVA_OPTS=-Xms1g -Xmx1g" \
		--env "ELASTIC_PASSWORD=${ELASTIC_PASSWORD}" \
		"docker.elastic.co/elasticsearch/elasticsearch:8.2.3"

	/usr/bin/env bash -c ${set-password}

	docker run --name kibana --hostname kibana -d --net elastic -p 5601:5601 \
		--volumes-from elasticsearch \
		-v kibana-data:/usr/share/kibana/data \
		--env "ELASTICSEARCH_HOSTS=https://elasticsearch:9200" \
		--env "ELASTICSEARCH_USERNAME=kibana_system" \
		--env "ELASTICSEARCH_PASSWORD=${KIBANA_PASSWORD}" \
		--env "ELASTICSEARCH_SSL_CERTIFICATEAUTHORITIES=/usr/share/elasticsearch/config/certs/http_ca.crt" \
		--env "ENTERPRISESEARCH_HOST=http://enterprise-search:3002" \
		docker.elastic.co/kibana/kibana:8.2.3

	docker run --name enterprise-search --hostname enterprise-search -d --net elastic -p 3002:3002 \
		-v "es-config:/usr/share/enterprise-search/es-config:ro" \
		--env "secret_management.encryption_keys=[${ENCRYPTION_KEYS}]" \
		--env "allow_es_settings_modification=true" \
		--env "elasticsearch.host=https://elasticsearch:9200" \
		--env "elasticsearch.username=elastic" \
		--env "elasticsearch.password=${ELASTIC_PASSWORD}" \
		--env "elasticsearch.ssl.enabled=true" \
		--env "elasticsearch.ssl.certificate_authority=/usr/share/enterprise-search/es-config/certs/http_ca.crt" \
		--env "secret_session_key=${SESSION_KEY}" \
		"docker.elastic.co/enterprise-search/enterprise-search:8.2.3"

docker-stop:
	docker stop enterprise-search kibana  elasticsearch

docker-clean: docker-stop
	docker rm -f kibana enterprise-search elasticsearch
	docker volume rm -f es-config kibana-data

.PHONY: docker-start docker-stop docker-clean
