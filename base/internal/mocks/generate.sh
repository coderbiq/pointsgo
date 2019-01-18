#!/bin/bash

mockgen -destination model_mock.go \
   -package mocks \
   github.com/coderbiq/pointsgo/base/internal/model \
   AccountRepository,AccountLogStorer

mockgen -destination service_mock.go \
   -package mocks \
   github.com/coderbiq/pointsgo/base/internal/service \
   Infra,AppServices,RegisterService,DepositService,ConsumeService,AccountFinder


mockgen -destination dgo_event.go \
   -package mocks \
   github.com/coderbiq/dgo/base/devent \
   Bus
