#!/bin/bash

mockgen -destination service_mock.go \
   -package mocks \
   github.com/coderbiq/pointsgo/base \
   AppServices,Infra,RegisterService,DepositService,ConsumeService