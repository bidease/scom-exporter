# servers.com exporter

Config `/etc/scom_exporter.yml`:

```yaml
---
web_metrics_addr: localhost:9876    # optional, default: localhost:9999
metrics_endpoint: /metrics          # optional, default
scrape_nterval: 600                 # optional, seconds, default: 300
auth:
  email: zzz@xxx.com
  token: XXXXXXXXXXX
```
