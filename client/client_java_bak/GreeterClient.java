import com.googlecode.jsonrpc4j.JsonRpcClient;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.*;
public class GreeterClient {
    public static void main(String[] args) {
        try {
                        System.out.println(System.getProperty("java.classpath"));
	    Socket socket = new Socket("127.0.0.1", 8080);
            JsonRpcClient client = new JsonRpcClient();

            InputStream ips = socket.getInputStream();
            OutputStream ops = socket.getOutputStream();

            int reply = client.invokeAndReadResponse("Server.Hello", new Object[]{"!!!"}, int.class, ops, ips);

            System.out.println("reply: " + reply);
        } catch (IOException e) {
            e.printStackTrace();
        } catch (Throwable throwable) {
            throwable.printStackTrace();
        }
    }
}
