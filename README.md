# Conan: Detect Consensus Issues in Distributed Databases using Multi-Feedback Fuzzing with Hybrid Fault Sequences
Conan is a testing framework for detecting consensus issues in distributed databases by fault injection. 

* We observe that solely using coarse-grained faults is not sufficient for bug detection. Hence, Conan combines fine-grained and coarse-grained faults to form fault sequence to uncover consensus issues. 
* We observe that runtime metrics improve the efficiency of generating fault sequences. Hence, Conan monitors runtime metrics as feedback to guide fuzzing for the effective generation of fault sequences. 

## Bug Detected by Conan
We appied Conan to 3 widely used distributed databases, including etcd, rqlite and openGauss. Conan successfully detects 8 consensus issues in these real-world databases, 6 of which are previous-unknown issues and 5 of them have been confirmed by developers. 
| Bug ID | Description | 
|-----|-----|
| [etcd-17332](https://github.com/etcd-io/etcd/issues/17332) | Inconsistent behaviors between server and client. | 
| [rqlite-1629](https://github.com/rqlite/rqlite/pull/1629) | Duplicate data insertion. | 
| [rqlite-1633](https://github.com/rqlite/rqlite/pull/1633) | Inappropriate error messages.  | 
| [rqlite-1712](https://github.com/rqlite/rqlite/issues/1712) | Unexpected election. | 
| [openGauss-I8I19W](https://gitee.com/opengauss/openGauss-server/issues/I8I19W) | Data inconsistency between nodes. | 
| [openGauss-I8H1YQ](https://gitee.com/opengauss/openGauss-server/issues/I8H1YQ) | No Leader. | 
| [openGauss-I8MGB4](https://gitee.com/opengauss/openGauss-server/issues/I8MGB4) | No Leader. |

## Require
```
Docker
docker-compose 2.6
go 1.21.0
python 3.6
```

## Getting Started
### Initial directory, pull docker images, complie conan server
```
./init.sh
```
### 
