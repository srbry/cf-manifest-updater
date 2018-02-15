package manifest_test

const oldManifest = `name: example
instances: 2
memory: 30M
disk_quota: 50M
domain: example.com
`

const newManifest = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: example.example.com
`

const oldManifest1 = `name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domain: example.com
`

const newManifest1 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example.com
`

const oldManifest2 = `name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domains:
- example1.com
- example2.com
`

const newManifest2 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example1.com
- route: test.example2.com
`

const oldManifest3 = `name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domain: example.com
domains:
- example1.com
- example2.com
`

const newManifest3 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example.com
- route: test.example1.com
- route: test.example2.com
`

const oldManifest4 = `name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domain: example1.com
domains:
- example1.com
- example2.com
`

const newManifest4 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example1.com
- route: test.example2.com
`

const oldManifest5 = `name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domains:
- example1.com
- example2.com
routes:
- route: tcp.example.com:8000
- route: test.example1.com
`

const newManifest5 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example1.com
- route: test.example2.com
- route: tcp.example.com:8000
`
