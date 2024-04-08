package org.apache.zookeeper.server;
// javac CPFIClient.java
// javah -jni CPFIClient
public class CPFIClient {
    static {
        System.load("/opt/zookeeper/src/java/main/org/apache/zookeeper/server/quorum/client_java/cpfi.so");
        // System.loadLibrary("cpfi");
    }

    // 声明本地方法
    public static native int Hello();
}