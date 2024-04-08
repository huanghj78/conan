# Conan: Detect Consensus Issues in Distributed Databases using Multi-Feedback Fuzzing with Hybrid Fault Sequences
Conan is a testing framework for detecting consensus issues in distributed databases by fault injection. 

* We observe that solely using coarse-grained faults is not sufficient for bug detection. Hence, Conan combines fine-grained and coarse-grained faults to form fault sequence to uncover consensus issues. 
* We observe that runtime metrics improve the efficiency of generating fault sequences. Hence, Conan monitors runtime metrics as feedback to guide fuzzing for the effective generation of fault sequences. 

## Bug Detected by Conan
We appied Conan to 3 widely used distributed databases, including etcd, rqlite and openGauss. Conan successfully detects 8 consensus issues in these real-world databases, 6 of which are previous-unknown issues and 5 of them have been confirmed by developers. 




## Getting Started
