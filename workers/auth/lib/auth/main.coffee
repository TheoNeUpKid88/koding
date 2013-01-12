{argv} = require 'optimist'

koding = require './bongo'
koding.connect()

AuthWorker = require './authworker'

{librato, authWorker} = require('koding-config-manager').load("main.#{argv.c}")

processMonitor = (require 'processes-monitor').start
  name : "Auth Worker #{process.pid}"
  stats_id: "worker.auth." + process.pid
  interval : 30000
  librato: librato

authWorker = new AuthWorker koding, authWorker.authResourceName
authWorker.connect()
