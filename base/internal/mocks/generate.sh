#!/bin/bash

mockgen -destination service_mock.go \
   -package mocks \
   github.com/coderbiq/pointsgo/base \
   AppServices,Infra,RegisterService,DepositService,ConsumeService,AccountRepository

mockgen -destination dgo_mock.go \
   -package mocks \
   github.com/coderbiq/dgo/model \
   EventPublisher
