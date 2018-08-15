# gpmt2  [![Build Status](https://travis-ci.org/pivotal-gss/gpmt2.svg?branch=master)](https://travis-ci.org/pivotal-gss/gpmt2)
Greenplum Magic Tool 2 - open source rebuild of GPMT


## Goals

- The ultimate goal of the project is to make GPMT more easiliy manageable, extensible, and supportable.
- Initial development will focus on replicating functionality of the original GPMT tool, in the following order of tools (subject to adjustment):
  1. gp_log_collector
  2. packcore
  3. gpcheckcat
  4. gpcatalog_backup
  5. analyze_session
  6. gpstatscheck
  
  
- Other goals include:
  - Compatibility with gpdb6
  - Additional features
  

## Contributing
CI will be provided by Travis - if sufficient code is written, it is expected that appropriate go tests will be included in any PR as well

Please ensure all submitted code is formatted with gofmt. 

This repository uses [dep](https://github.com/golang/dep) so please ensure you run `dep ensure --update` if you add imports. 



Of course, github issue requests are the best way to start a request for new features or bugs.