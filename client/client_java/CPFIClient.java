// package org.apache.zookeeper.client_java;
public class CPFIClient
{
    // 载入 so 动态链接库
    static{
        //System.loadLibrary("HelloWorld");//dll方式
        System.load("/root/mvn_test/my-app/src/main/java/com/mycompany/client_java/CPFIClient.so");//so方式
    }

    // 声明 so 库中的方法
    public native int BeforeSendReq();
}