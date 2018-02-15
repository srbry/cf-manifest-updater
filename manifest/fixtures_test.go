package manifest_test

var oldManifest = []byte(`name: example
instances: 2
memory: 30M
disk_quota: 50M
domain: example.com
`)

const newManifest = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: example.example.com
`

var oldManifest1 = []byte(`name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domain: example.com
`)

const newManifest1 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example.com
`

var oldManifest2 = []byte(`name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domains:
- example1.com
- example2.com
`)

const newManifest2 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example1.com
- route: test.example2.com
`

var oldManifest3 = []byte(`name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domain: example.com
domains:
- example1.com
- example2.com
`)

const newManifest3 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example.com
- route: test.example1.com
- route: test.example2.com
`

var oldManifest4 = []byte(`name: example
host: test
instances: 2
memory: 30M
disk_quota: 50M
domain: example1.com
domains:
- example1.com
- example2.com
`)

const newManifest4 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: test.example1.com
- route: test.example2.com
`

var oldManifest5 = []byte(`name: example
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
services:
- service1
- service2
`)

const newManifest5 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
routes:
- route: tcp.example.com:8000
- route: test.example1.com
- route: test.example2.com
services:
- service1
- service2
`

var oldManifest6 = []byte(`applications:
- name: example
  instances: 2
  memory: 30M
  disk_quota: 50M
  domain: example.com
- name: example2
  instances: 2
  memory: 30M
  disk_quota: 50M
  domain: example.com
`)

const newManifest6 = `applications:
- disk_quota: 50M
  instances: 2
  memory: 30M
  name: example
  routes:
  - route: example.example.com
- disk_quota: 50M
  instances: 2
  memory: 30M
  name: example2
  routes:
  - route: example2.example.com
`

var oldManifest7 = []byte(`applications:
- name: example
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
  services:
  - service1
  - service2
- name: example
  host: test
  instances: 2
  memory: 30M
  disk_quota: 50M
  domain: example1.com
  domains:
  - example1.com
  - example2.com
`)

const newManifest7 = `applications:
- disk_quota: 50M
  instances: 2
  memory: 30M
  name: example
  routes:
  - route: tcp.example.com:8000
  - route: test.example1.com
  - route: test.example2.com
  services:
  - service1
  - service2
- disk_quota: 50M
  instances: 2
  memory: 30M
  name: example
  routes:
  - route: test.example1.com
  - route: test.example2.com
`

var oldManifest8 = []byte(`name: example
domain: example.com
memory: 256M
stack: cflinuxfs2
timeout: 180
applications:
- name: example
  env:
    ENV1: value
    ENV2: value
  services:
  - service1
  command: bundle exec rake
`)

const newManifest8 = `applications:
- command: bundle exec rake
  env:
    ENV1: value
    ENV2: value
  name: example
  routes:
  - route: example.example.com
  services:
  - service1
memory: 256M
name: example
stack: cflinuxfs2
timeout: 180
`

var oldManifest9 = []byte(`name: example
instances: 2
memory: 30M
disk_quota: 50M
`)

const newManifest9 = `disk_quota: 50M
instances: 2
memory: 30M
name: example
`

var oldManifest10 = []byte(`name: example
domains:
  - example.com
memory: 256M
stack: cflinuxfs2
timeout: 180
applications:
- name: example
  env:
    ENV1: value
    ENV2: value
  domains:
    - example2.com
  services:
  - service1
  command: bundle exec rake
`)

const newManifest10 = `applications:
- command: bundle exec rake
  env:
    ENV1: value
    ENV2: value
  name: example
  routes:
  - route: example.example.com
  - route: example.example2.com
  services:
  - service1
memory: 256M
name: example
stack: cflinuxfs2
timeout: 180
`
