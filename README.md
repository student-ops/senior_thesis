## 開発中

dev/slack ブランチ

docker run -d -p 3000:3000 grafana/grafana-oss:main

## Getting started

import grafana volume on

- ./grafana/grafana_volume:/var/lib/grafansa

let's compose up from local file

```
docker compose -f docker-compose.local.yml up -d

cd server_tests/sample_generator/

python3 sample_generator.py
```

then open http://localhost:3000
