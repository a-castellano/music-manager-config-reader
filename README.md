# Config Reader

[Actual Repo](https://git.windmaker.net/musicmanager/Config-Reader)

 [![pipeline status](https://git.windmaker.net/musicmanager/Config-Reader/badges/master/pipeline.svg)](https://git.windmaker.net/musicmanager/Config-Reader/-/commits/master) [![coverage report](https://git.windmaker.net/musicmanager/Config-Reader/badges/master/coverage.svg)](https://git.windmaker.net/musicmanager/Config-Reader/-/commits/master) [![Quality Gate Status](https://sonarqube.windmaker.net/api/project_badges/measure?project=config-reader&metric=alert_status)](https://sonarqube.windmaker.net/dashboard?id=config-reader)

This service reads configs and creates config object that is used by multiple Music Manager Projects.

For the time being a single incoming and outgoing queue config is allowed, more config options will be allowed in future releases.

### Config example

This service will look for its config in **/etc/music-manager-service/config.toml**, parent folder can be changed setting the environment variable **MUSIC_MANAGER_SERVICE_CONFIG_FILE_LOCATION**.

Here is a config example:

```toml
[server]
host = "localhost"
port = 5672
user = "guest"
password = "pass"

[incoming]
name = "incoming"

[outgoing]
name = "outgoing"
```
