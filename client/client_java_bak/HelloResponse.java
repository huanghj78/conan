import java.io.Serializable;

public class HelloResponse implements Serializable {
    private static final long serialVersionUID = 1L;

    private String message;

    // Constructors, getters, setters

    public HelloResponse() {
    }

    public HelloResponse(String message) {
        this.message = message;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}
