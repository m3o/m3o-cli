# Public APIs

M3O is a cloud platform for Micro services served as public APIs.

## Overview

Micro is an open source distributed cloud platform which enables the development of cloud services. 
Cloud services are primarily consumed API first. M3O hosts Micro and then makes publicly available 
a catalog of open source Micro services which can be consumed via a single account and token through 
one unified platform.

## Services

Micro services are built as single purpose APIs that do one thing well. They offer a uniform method 
of consumption and the ability to compose as building blocks for higher level services. We use this 
as our abstraction layer for third party infrastructure and services.

The source for services reside in https://github.com/micro/services.

## Publishing

Micro services are published to our catalog via a publicapi service. This service maintains all the 
information for the APIs we offer including examples, openapi spec and pricing details. Services 
include a standard formatted publicapi.json in the root folder along with a separate examples.json 
file. If these are found then they're included in the publishing process.

Find an example here https://github.com/micro/services/tree/master/helloworld.
 
