import java.io.Serializable;

public class HelloRequest implements Serializable {
    private static final long serialVersionUID = 1L;

    private String text;

    // Constructors, getters, setters

    public HelloRequest() {
    }

    public HelloRequest(String text) {
        this.text = text;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }
}
