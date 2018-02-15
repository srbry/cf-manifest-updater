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
