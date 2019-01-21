#!/bin/bash

mockgen -destination model_mock.go \
   -package mocks \
   github.com/coderbiq/pointsgo/base/internal/model \
   AccountRepository,AccountLogStorer,Infra,AppServices,RegisterService,DepositService,ConsumeService,AccountFinder

mockgen -destination dgo_event.go \
   -package mocks \
   github.com/coderbiq/dgo/base/devent \
   Bus
