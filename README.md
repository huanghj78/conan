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
### Run conan server
```
# ./run.sh system mode 
```
* Detect mode
```
./run.sh etcd Detect
```
> 2024/04/09 14:43:08 [INFO] /root/Conan/server/server.go:131 Start in Detect mode
2024/04/09 14:43:08 [INFO] /root/Conan/server/server.go:132 Target System is etcd
2024/04/09 14:43:08 [INFO] /root/Conan/server/monitor.go:35 Setup etcd cluster
2024/04/09 14:43:08 [INFO] /root/Conan/server/injector.go:19 Injector start
2024/04/09 14:43:08 [INFO] /root/Conan/server/fuzzer.go:122 Init seed seq [FTP Before Leader MsgApp][FIP EnumCPUHog 92 for 4s at node2][FIP EnumRestartNode Node2][FIP EnumNetworkLoss 97 at node1 for 3s][FIP EnumMessageFault  Modify term 0 ]
2024/04/09 14:43:08 [INFO] /root/Conan/server/fuzzer.go:122 Init seed seq [FTP Before Leader MsgAppResp][FIP EnumMessageFault  Delay for 813ms ]
2024/04/09 14:43:08 [INFO] /root/Conan/server/fuzzer.go:122 Init seed seq [FTP Before Follower MsgApp][FIP EnumMessageFault  Delay for 1.175s ]
2024/04/09 14:43:08 [INFO] /root/Conan/server/fuzzer.go:122 Init seed seq [FTP Before Follower MsgAppResp][FIP EnumNetworkLoss 80 at node1 for 4s][FIP EnumMessageFault  Delay for 794ms ]
2024/04/09 14:43:08 [INFO] /root/Conan/server/monitor.go:26 Monitor start
2024/04/09 14:43:08 [INFO] /root/Conan/server/checker.go:19 Checker Start
2024/04/09 14:43:24 [INFO] /root/Conan/server/monitor.go:60 Current leader num: 2
2024/04/09 14:43:24 [INFO] /root/Conan/server/monitor.go:61 Run etcd workload
2024/04/09 14:43:25 [INFO] /root/Conan/server/monitor.go:86 Current leader num: 2
2024/04/09 14:43:25 [INFO] /root/Conan/server/server.go:32 Start Server....
2024/04/09 14:43:33 [INFO] /root/Conan/server/fuzzer.go:163 select seed seq [FTP Before Follower MsgApp][FIP EnumMessageFault  Delay for 1.175s ]
2024/04/09 14:43:33 [INFO] /root/Conan/server/fuzzer.go:177 Select op *main.ModifyFIPOp
2024/04/09 14:43:33 [INFO] /root/Conan/server/operator.go:348 Modify args at 0 0
2024/04/09 14:43:33 [INFO] /root/Conan/server/monitor.go:35 Setup etcd cluster
2024/04/09 14:43:51 [INFO] /root/Conan/server/injector.go:30 Injector run, receive testSeq [FTP Before Follower MsgApp][FIP EnumMessageFault Omit]
2024/04/09 14:43:51 [INFO] /root/Conan/server/monitor.go:60 Current leader num: 2
2024/04/09 14:43:51 [INFO] /root/Conan/server/monitor.go:61 Run etcd workload
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:165 notification:  Before Leader MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:166 FaultPoint:  Before Follower MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:165 notification:  After Leader MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:166 FaultPoint:  Before Follower MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:165 notification:  Before Leader MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:166 FaultPoint:  Before Follower MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:165 notification:  After Leader MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:166 FaultPoint:  Before Follower MsgApp
2024/04/09 14:43:52 [INFO] /root/Conan/server/injector.go:170 Inject Fault!
2024/04/09 14:43:52 [INFO] /root/Conan/server/monitor.go:86 Current leader num: 2
2024/04/09 14:43:52 [INFO] /root/Conan/server/monitor.go:104 Fitness score: 0
2024/04/09 14:43:58 [INFO] /root/Conan/server/injector.go:42 Injector stop
2024/04/09 14:43:58 [INFO] /root/Conan/server/checker.go:27 Checker run
2024/04/09 14:43:59 [INFO] /root/Conan/server/checker.go:36 =====================CHECK PASS=====================

* Reproduce mode

